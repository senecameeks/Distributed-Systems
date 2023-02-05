/*
 *  Brown University, CS138, Spring 2018
 *
 *  Purpose: a LiteMiner miner.
 */

package liteminer

import (
	"fmt"
	"io"
	"sync"
	"time"
)

const HEARTBEAT_FREQ = 1000 * time.Millisecond

// Represents a LiteMiner miner
type Miner struct {
	IsShutdown   bool
	Mining       bool
	NumProcessed uint64     // Number of values processed in the current mining range
	mutex        sync.Mutex // To manage concurrent access to these members
}

// CreateMiner creates a new miner connected to the pool at the specified address.
func CreateMiner(addr string) (mp *Miner, err error) {
	var miner Miner

	mp = &miner

	miner.Mining = false
	miner.NumProcessed = 0
	miner.IsShutdown = false

	err = miner.connect(addr)

	return
}

// connect connects the miner to the pool at the specified address.
func (m *Miner) connect(addr string) (err error) {
	conn, err := MinerConnect(addr)
	if err != nil {
		return fmt.Errorf("Received error %v when connecting to pool %v\n", err, addr)
	}

	go m.receiveFromPool(conn)
	go m.sendHeartBeats(conn)

	return
}

// receiveFromPool processes messages from the pool represented by conn.
func (m *Miner) receiveFromPool(conn MiningConn) {
	for {
		m.mutex.Lock()
		if m.IsShutdown {
			conn.Conn.Close() // Close the connection
			m.mutex.Unlock()
			return
		}
		m.mutex.Unlock()

		msg, err := RecvMsg(conn)
		if err != nil {
			if err == io.EOF {
				Err.Printf("Lost connection to pool %v\n", conn.Conn.RemoteAddr())
				conn.Conn.Close() // Close the connection
				return
			}

			Err.Printf(
				"Received error %v when processing pool %v\n",
				err,
				conn.Conn.RemoteAddr(),
			)
			continue
		}

		if msg.Type != MineRequest {
			Err.Printf(
				"Received unexpected message of type %v from pool %v\n",
				msg.Type,
				conn.Conn.RemoteAddr(),
			)
		}

		nonce := m.Mine(msg.Data, msg.Lower, msg.Upper) // Service the mine request

		// Send result
		res := ProofOfWorkMsg(msg.Data, nonce, Hash(msg.Data, nonce))
		SendMsg(conn, res)
	}
}

// sendHeartBeats periodically sends heartbeats to the pool represented by conn
// while mining. sendHeartBeats should NOT send heartbeats to the pool if the
// miner is not mining. It should close the connection and return if the miner
// is shutdown.
func (m *Miner) sendHeartBeats(conn MiningConn) {
	// TODO: Students should send a StatusUpdate message every HEARTBEAT_FREQ
	// while mining.
	tick := time.Tick(HEARTBEAT_FREQ)

	go func() {
		for range tick {
			if m.Mining {
				msg := StatusUpdateMsg(m.NumProcessed)
				SendMsg(conn, msg)
			}
			if m.IsShutdown {
				conn.Conn.Close()
				return
			}
		}
	}()

}

// Given a data string, a lower bound (inclusive), and an upper bound
// (exclusive), Mine returns the nonce in the range [lower, upper) that
// corresponds to the lowest hash value. With each value processed in the range,
// NumProcessed should increase.
func (m *Miner) Mine(data string, lower, upper uint64) (nonce uint64) {
	// TODO: Students should implement this. Make sure to use the Hash method
	// in hash.go
	m.mutex.Lock()
	min_hash := Hash(data, lower)

	m.NumProcessed = 0
	nonce = lower
	var index uint64
	var hash_temp uint64
	m.Mining = true
	m.mutex.Unlock()

	for index = lower + 1; index < upper; index++ {
		hash_temp = Hash(data, index)
		if hash_temp < min_hash {
			min_hash = hash_temp
			nonce = index
		}
		m.mutex.Lock()
		m.NumProcessed++
		m.mutex.Unlock()
	}

	m.mutex.Lock()
	m.Mining = false
	m.mutex.Unlock()
	return
}

// Shutdown marks the miner as shutdown and asynchronously disconnects it from
// its pool.
func (m *Miner) Shutdown() {
	m.mutex.Lock()
	Debug.Printf("Shutting down")
	m.IsShutdown = true
	m.Mining = false
	m.mutex.Unlock()
}
