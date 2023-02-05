/*
 *  Brown University, CS138, Spring 2018
 *
 *  Purpose: a LiteMiner mining pool.
 */

package liteminer

import (
	"encoding/gob"
	"io"
	"math"
	"net"
	"sync"
	"time"
)

const HEARTBEAT_TIMEOUT = 3 * HEARTBEAT_FREQ

// Represents a LiteMiner mining pool
type Pool struct {
	Addr          net.Addr                // Address of the pool
	Miners        map[net.Addr]MiningConn // Currently connected miners
	Client        MiningConn              // The current client
	busy          bool                    // True when processing a transaction
	mutex         sync.Mutex              // To manage concurrent access to these members
	Schedule      chan Interval           // Channel for job intervals for miners
	Nonce         uint64                  // Nonce to return to client
	NumProofsLeft int                     // Number of proofs of work remaining before end
	msg           Message                 // Saved transaction message
}

// CreatePool creates a new pool at the specified port.
func CreatePool(port string) (pp *Pool, err error) {
	var pool Pool

	pp = &pool

	pool.busy = false
	pool.Client.Conn = nil
	pool.Miners = make(map[net.Addr]MiningConn)
	// TODO: Students should (if necessary) initialize any additional members
	// to the Pool struct here.
	pool.Schedule = make(chan Interval, 10000)
	pool.NumProofsLeft = -1
	err = pp.startListener(port)

	return
}

// startListener starts listening for new connections.
func (p *Pool) startListener(port string) (err error) {
	listener, portId, err := OpenListener(port)
	if err != nil {
		return
	}

	p.Addr = listener.Addr()

	Out.Printf("Listening on port %v\n", portId)

	// Listen for and accept connections
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				Err.Printf("Received error %v when listening for connections\n", err)
				continue
			}

			go p.handleConnection(conn)
		}
	}()

	return
}

// handleConnection handles an incoming connection and delegates to
// handleMinerConnection or handleClientConnection.
func (p *Pool) handleConnection(nc net.Conn) {
	// Set up connection
	conn := MiningConn{}
	conn.Conn = nc
	conn.Enc = gob.NewEncoder(nc)
	conn.Dec = gob.NewDecoder(nc)
	//Debug.Printf("handling connection")
	// Wait for Hello message
	msg, err := RecvMsg(conn)
	//Debug.Printf("received msg")
	if err != nil {
		Err.Printf(
			"Received error %v when processing Hello message from %v\n",
			err,
			conn.Conn.RemoteAddr(),
		)
		conn.Conn.Close() // Close the connection
		return
	}
	//Debug.Printf("waiting to switch %d", msg.Type)
	switch msg.Type {
	case MinerHello:
		p.handleMinerConnection(conn)
	case ClientHello:
		p.handleClientConnection(conn)
	default:
		SendMsg(conn, ErrorMsg("Unexpected message type"))
	}
}

// handleClientConnection handles a connection from a client.
func (p *Pool) handleClientConnection(conn MiningConn) {
	Debug.Printf("Received client connection from %v", conn.Conn.RemoteAddr())

	p.mutex.Lock()
	if p.Client.Conn != nil {
		Debug.Printf(
			"Busy with client %v, sending BusyPool message to client %v",
			p.Client.Conn.RemoteAddr(),
			conn.Conn.RemoteAddr(),
		)
		SendMsg(conn, BusyPoolMsg())
		p.mutex.Unlock()
		return
	} else if p.busy {
		Debug.Printf(
			"Busy with previous transaction, sending BusyPool message to client %v",
			conn.Conn.RemoteAddr(),
		)
		SendMsg(conn, BusyPoolMsg())
		p.mutex.Unlock()
		return
	}
	p.Client = conn
	p.mutex.Unlock()

	// Listen for and handle incoming messages
	for {
		msg, err := RecvMsg(conn)
		if err != nil {
			if err == io.EOF {
				Out.Printf("Client %v disconnected\n", conn.Conn.RemoteAddr())

				conn.Conn.Close() // Close the connection

				p.mutex.Lock()
				p.Client.Conn = nil
				p.mutex.Unlock()

				return
			}
			Err.Printf(
				"Received error %v when processing message from client %v\n",
				err,
				conn.Conn.RemoteAddr(),
			)
			continue
		}

		if msg.Type != Transaction {
			SendMsg(conn, ErrorMsg("Expected Transaction message"))
			continue
		}

		Debug.Printf(
			"Received transaction from client %v with data %v and upper bound %v",
			conn.Conn.RemoteAddr(),
			msg.Data,
			msg.Upper,
		)

		p.mutex.Lock()
		if len(p.Miners) == 0 {
			SendMsg(conn, ErrorMsg("No miners connected"))
			p.mutex.Unlock()
			continue
		}
		p.mutex.Unlock()

		// TODO: Students should handle an incoming transaction from a client. A
		// pool may process one transaction at a time – thus, if you receive
		// another transaction while busy, you should send a BusyPool message.

		if msg.Type == Transaction {
			if p.busy {
				SendMsg(conn, BusyPoolMsg())
			} else {
				p.mutex.Lock()
				p.msg = msg
				p.busy = true
				p.mutex.Unlock()

				// handle transaction
				miners := p.GetMiners()

				p.mutex.Lock()
				num_intervals := int(math.Min(float64(5*len(miners)), float64(msg.Upper)))

				p.NumProofsLeft = num_intervals
				p.mutex.Unlock()

				intervals := GenerateIntervals(msg.Upper, num_intervals)
				p.mutex.Lock()
				for _, i := range intervals {
					p.Schedule <- i
				}
				p.mutex.Unlock()
			}
		}
	}
}

// handleMinerConnection handles a connection from a miner.
// assign miners here
func (p *Pool) handleMinerConnection(conn MiningConn) {
	Debug.Printf("Received miner connection from %v", conn.Conn.RemoteAddr())

	p.mutex.Lock()
	p.Miners[conn.Conn.RemoteAddr()] = conn

	p.mutex.Unlock()

	msgChan := make(chan Message, 10)
	go p.receiveFromMiner(conn, msgChan)

	// TODO: Students should handle a miner connection. If a miner does not
	// send a StatusUpdate message every HEARTBEAT_TIMEOUT while mining,
	// any work assigned to them should be redistributed and they should be
	// disconnected and removed from p.Miners.

	// assign miners an intervals

	for new_interval := range p.Schedule {

		SendMsg(conn, MineRequestMsg(p.msg.Data, new_interval.Lower, new_interval.Upper))
		timeout := time.Tick(time.Duration(HEARTBEAT_TIMEOUT))
	loop:
		for {
			select {
			case msg := <-msgChan:
				switch msg.Type {
				case StatusUpdate:
					timeout = time.Tick(time.Duration(HEARTBEAT_TIMEOUT))
					continue
				case ProofOfWork:
					p.mutex.Lock()
					p.NumProofsLeft--

					p.mutex.Unlock()

					p.mutex.Lock()
					if msg.Hash < Hash(msg.Data, p.Nonce) {
						p.Nonce = msg.Nonce
					}
					p.mutex.Unlock()

					// check if all the intervals have been mined and send POW msg to client
					p.mutex.Lock()
					if p.NumProofsLeft == 0 {
						SendMsg(p.Client, ProofOfWorkMsg(p.msg.Data, uint64(p.Nonce), Hash(p.msg.Data, uint64(p.Nonce))))
						p.busy = false
					}
					p.mutex.Unlock()

					break loop
				}

			case <-timeout:
				if p.busy {
					Debug.Printf("Killing miner")
					p.mutex.Lock()

					p.Schedule <- new_interval

					delete(p.Miners, conn.Conn.RemoteAddr()) //delete miner

					p.mutex.Unlock()
					conn.Conn.Close() // Close the connection
					return
				}
			}
		}
	}

	//Debug.Printf("Left forever loop in Miner Connection")
}

// receiveFromMiner waits for messages from the miner specified by conn and
// forwards them over msgChan.
func (p *Pool) receiveFromMiner(conn MiningConn, msgChan chan Message) {
	for {
		msg, err := RecvMsg(conn)
		if err != nil {
			if _, ok := err.(*net.OpError); ok || err == io.EOF {
				Out.Printf("Miner %v disconnected\n", conn.Conn.RemoteAddr())

				p.mutex.Lock()
				delete(p.Miners, conn.Conn.RemoteAddr())
				p.mutex.Unlock()

				conn.Conn.Close() // Close the connection

				return
			}
			Err.Printf(
				"Received error %v when processing message from miner %v\n",
				err,
				conn.Conn.RemoteAddr(),
			)
			continue
		}

		msgChan <- msg
	}
}

// GetMiners returns the addresses of any connected miners.
func (p *Pool) GetMiners() []net.Addr {
	miners := []net.Addr{}
	p.mutex.Lock()
	for _, m := range p.Miners {
		miners = append(miners, m.Conn.RemoteAddr())
	}
	p.mutex.Unlock()
	return miners
}

// GetClient returns the address of the current client or nil if there is no
// current client.
func (p *Pool) GetClient() net.Addr {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.Client.Conn == nil {
		return nil
	}
	return p.Client.Conn.RemoteAddr()
}
