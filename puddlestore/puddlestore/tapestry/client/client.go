/*
 *  Brown University, CS138, Spring 2018
 *
 *  Purpose: Allows third-party clients to connect to a Tapestry node (such as
 *  a web app, mobile app, or CLI that you write), and put and get objects.
 */

package client

import (
	"fmt"

	t "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/tapestry"
)

type Client struct {
	Id   string
	node *t.RemoteNode
}

var salts []string

// Connect to a Tapestry node
func Connect(addr string) *Client {
	initializeSalts()
	node, err := t.SayHelloRPC(addr, t.RemoteNode{})
	if err != nil {
		t.Error.Printf("Failed to make clientection to Tapestry node\n")
	}
	return &Client{node.Id.String(), &node}
}

// initialize salts
func initializeSalts() {
	salts = make([]string, 3)
	salts[0] = "f1nd1ngn3m0"
	salts[1] = "pwn3d"
	salts[2] = "f1nd1ngd0ry"
}

// Invoke tapestry.Store on the remote Tapestry node
func (client *Client) Store(key string, value []byte) error {
	t.Debug.Printf("Making remote TapestryStore call\n")
	errs := make([]error, 4)
	errs[0] = client.node.TapestryStoreRPC(key, value)
	errs[1] = client.node.TapestryStoreRPC(key+salts[0], value)
	errs[2] = client.node.TapestryStoreRPC(key+salts[1], value)
	errs[3] = client.node.TapestryStoreRPC(key+salts[2], value)

	for _, e := range errs {
		if e != nil {
			return e
		}
	}

	return nil
}

// Invoke tapestry.Lookup on a remote Tapestry node
func (client *Client) Lookup(key string) ([]*Client, error) {
	t.Debug.Printf("Making remote TapestryLookup call\n")
	var clients []*Client
	errs := make([]error, 4)
	nodes, err := client.node.TapestryLookupRPC(key)
	errs[0] = err
	//clients := make([]*Client, len(nodes))
	for _, n := range nodes {
		clients = append(clients, &Client{n.Id.String(), &n})
	}

	for i, salt := range salts {
		nodes, err := client.node.TapestryLookupRPC(key + salt)
		errs[i+1] = err
		//clients := make([]*Client, len(nodes))
		for _, n := range nodes {
			clients = append(clients, &Client{n.Id.String(), &n})
		}
	}

	for _, e := range errs {
		if e != nil {
			err = e
		}
	}

	return clients, err
}

// Get data from a Tapestry node. Looks up key then fetches directly.
func (client *Client) Get(key string) ([]byte, error) {
	t.Debug.Printf("Making remote TapestryGet call\n")
	// Lookup the key
	replicas, err := client.node.TapestryLookupRPC(key)
	if err != nil {
		return nil, err
	}
	if len(replicas) == 0 {
		return nil, fmt.Errorf("No replicas returned for key %v", key)
	}

	// Contact replicas
	var errs []error
	for _, replica := range replicas {
		blob, err := replica.BlobStoreFetchRPC(key)
		if err != nil {
			errs = append(errs, err)
		}
		if blob != nil {
			return *blob, nil
		}
	}

	return nil, fmt.Errorf("Error contacting replicas, %v: %v", replicas, errs)
}
