/*
 *  Brown University, CS138, Spring 2018
 *
 *  Purpose: a LiteMiner client.
 */

package liteminer

import (
	"fmt"
	"io"
	"net"
	"sync"
)

// Represents a LiteMiner client
type Client struct {
	PoolConns map[net.Addr]MiningConn // Pool(s) that the client is currently connected to
	Nonces    map[net.Addr]int64      // Nonce(s) received by the various pool(s) for the current transaction
	TxResults chan map[net.Addr]int64 // Used to send results of transaction
	mutex     sync.Mutex              // To manage concurrent access to these members
}

// CreateClient creates a new client connected to the given pool addresses.
func CreateClient(addrs []string) (cp *Client) {
	var client Client

	cp = &client

	client.PoolConns = make(map[net.Addr]MiningConn)
	client.Nonces = make(map[net.Addr]int64)
	client.TxResults = make(chan map[net.Addr]int64)

	cp.Connect(addrs)

	return
}

// Connect connects the client to the specified pool addresses.
func (c *Client) Connect(addrs []string) {
	for _, addr := range addrs {
		conn, err := ClientConnect(addr)
		if err != nil {
			Err.Printf("Received error %v when connecting to pool %v\n", err, addr)
			continue
		}

		c.mutex.Lock()
		c.PoolConns[conn.Conn.RemoteAddr()] = conn
		c.mutex.Unlock()

		go c.processPool(conn)
	}
}

// processPool handles incoming messages from the pool represented by conn.
func (c *Client) processPool(conn MiningConn) {
	for {
		msg, err := RecvMsg(conn)
		if err != nil {
			if err == io.EOF {
				Err.Printf("Lost connection to pool %v\n", conn.Conn.RemoteAddr())

				c.mutex.Lock()
				delete(c.PoolConns, conn.Conn.RemoteAddr())
				if len(c.Nonces) == len(c.PoolConns) {
					c.TxResults <- c.Nonces
				}
				c.mutex.Unlock()

				conn.Conn.Close() // Close the connection

				return
			}

			Err.Printf(
				"Received error %v when processing pool %v\n",
				err,
				conn.Conn.RemoteAddr(),
			)

			c.mutex.Lock()
			c.Nonces[conn.Conn.RemoteAddr()] = -1 // -1 used to indicate error
			if len(c.Nonces) == len(c.PoolConns) {
				c.TxResults <- c.Nonces
			}
			c.mutex.Unlock()

			continue
		}

		switch msg.Type {
		case BusyPool:
			Out.Printf("Pool %v is currently busy, disconnecting\n", conn.Conn.RemoteAddr())

			c.mutex.Lock()
			delete(c.PoolConns, conn.Conn.RemoteAddr())
			c.mutex.Unlock()

			conn.Conn.Close() // Close the connection

			return
		case ProofOfWork:
			Debug.Printf("Pool %v found nonce %v\n", conn.Conn.RemoteAddr(), msg.Nonce)

			c.mutex.Lock()
			c.Nonces[conn.Conn.RemoteAddr()] = int64(msg.Nonce)
			if len(c.Nonces) == len(c.PoolConns) {
				c.TxResults <- c.Nonces
			}
			c.mutex.Unlock()
		default:
			Err.Printf(
				"Received unexpected message of type %v from pool %v\n",
				msg.Type,
				conn.Conn.RemoteAddr(),
			)

			c.mutex.Lock()
			c.Nonces[conn.Conn.RemoteAddr()] = -1 // -1 used to indicate error
			if len(c.Nonces) == len(c.PoolConns) {
				c.TxResults <- c.Nonces
			}
			c.mutex.Unlock()
		}
	}
}

// Given a transaction encoded as a string and an unsigned integer, Mine returns
// the nonce(s) calculated by any connected pool(s). This method should NOT be
// executed concurrently by the same miner.
func (c *Client) Mine(data string, upperBound uint64) (map[net.Addr]int64, error) {
	c.mutex.Lock()

	if len(c.PoolConns) == 0 {
		c.mutex.Unlock()
		return nil, fmt.Errorf("Not connected to any pools")
	}

	c.Nonces = make(map[net.Addr]int64)

	// Send transaction to connected pool(s)
	tx := TransactionMsg(data, upperBound)
	for _, conn := range c.PoolConns {
		SendMsg(conn, tx)
	}
	c.mutex.Unlock()

	nonces := <-c.TxResults

	return nonces, nil
}
