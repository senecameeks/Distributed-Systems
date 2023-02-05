/*
 *  Brown University, CS138, Spring 2019
 *
 *  Purpose: Defines functions to publish and lookup objects in a Tapestry mesh
 */

package tapestry

import (
	"fmt"
	"time"
	// Uncomment for xtrace
	// xtr "github.com/brown-csci1380/tracing-framework-go/xtrace/client"
)

// Store a blob on the local node and publish the key to the tapestry.
func (local *Node) Store(key string, value []byte) (err error) {
	done, err := local.Publish(key)
	if err != nil {
		return err
	}
	local.blobstore.Put(key, value, done)
	return nil
}

// Lookup a key in the tapestry then fetch the corresponding blob from the
// remote blob store.
func (local *Node) Get(key string) ([]byte, error) {
	// Lookup the key
	replicas, err := local.Lookup(key)

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

// Remove the blob from the local blob store and stop advertising
func (local *Node) Remove(key string) bool {
	return local.blobstore.Delete(key)
}

// Publishes the key in tapestry.
//
// - Route to the root node for the key
// - Register our ln node on the root
// - Start periodically republishing the key
// - Return a channel for cancelling the publish
func (local *Node) Publish(key string) (done chan bool, err error) { // jacked UP
	// TODO: students should implement this
	root, err := local.findRoot(local.node, Hash(key))

	done = make(chan bool, 10)
	_, err = root.RegisterRPC(key, local.node)

	if err != nil {
		Debug.Printf("error: %v", err)
	}

	// if <-done then return, stop republishing

	ticker := time.NewTicker(REPUBLISH)
	// periodically republish key
	go func() { // most likely maybe correct !
		for t := range ticker.C {
			select {
				case <-done:
					return
				default:
					// root might fail, so refind each time you republish
					root, err = local.findRoot(local.node, Hash(key))
					_, err = root.RegisterRPC(key, local.node)
					if err != nil { // error means loss connection between local and root
						// whenever any RPC fails it means I should retry, Register should find new Root for me
						Debug.Printf("error: %v", err)
					}
					Debug.Printf("Republishing at time: %v", t)
			}
		}
	}()

	return
}

// Look up the Tapestry nodes that are storing the blob for the specified key.
//
// - Find the root node for the key
// - Fetch the replicas (nodes storing the blob) from the root's location map
// - Attempt up to RETRIES times
func (local *Node) Lookup(key string) (nodes []RemoteNode, err error) {
	// TODO: students should implement this
	var root RemoteNode
	var isRoot bool
	retries := 0
	id := Hash(key)
	//Debug.Printf("id for key: %v is %v", key, id)

	root, err = local.findRoot(local.node, id)
	for {
		if err != nil {
			if retries >= RETRIES {
				fmt.Errorf("failed after RETRIES times")
				return
			}
			retries++
			root, err = local.findRoot(local.node, id)
		} else {
			retries = 0
			break
		}
	}
	isRoot, nodes, err = root.FetchRPC(key)
	Debug.Printf("isRoot: %v", isRoot)
	for {
		if err != nil {
			if retries >= RETRIES {
				fmt.Errorf("failed after RETRIES times")
				return
			}
			retries++
			isRoot, nodes, err = root.FetchRPC(key)
			Debug.Printf("isRoot: %v", isRoot)
		} else {
			retries = 0
			break
		}
	}

	return
}

// Returns the best candidate from our routing table for routing to the provided ID.
// That node should recursively get the best candidate from its routing table, so the result of calling this should be
// the best candidate in the network for the given id.
func (local *Node) GetNextHop(id ID, level int32) (nexthop RemoteNode, err error, toRemove *NodeSet) {
	// TODO: students should implement this
	nexthop, err, toRemove = local.table.GetNextHop(id, level)
	return
}

// Register the specified node as an advertiser of the specified key.
// - Check that we are the root node for the key
// - Add the node to the location map
// - Kick off a timer to remove the node if it's not advertised again after a set amount of time
func (local *Node) Register(key string, replica RemoteNode) (isRoot bool, err error) {
	// TODO: students should implement this
	id := Hash(key)
	root, err := local.findRoot(local.node, id)

	if err != nil || root.Id != local.node.Id {
		// no idea honestly
		isRoot = false
		return
	}

	isRoot = true
	added := local.locationsByKey.Register(key, replica, TIMEOUT)
	//Debug.Printf("REGISTERS supposedly registered %v at %v", key, replica)
	if !added {
		// handle me
		Debug.Printf("Already registered")
	}


	return
}

// - Check that we are the root node for the requested key
// - Return all nodes that are registered in the local location map for this key
func (local *Node) Fetch(key string) (isRoot bool, replicas []RemoteNode, err error) {
	// TODO: students should implement this
	id := Hash(key)
	root, err := local.findRoot(local.node, id)
	if err != nil || root != local.node {
		// no idea honestly
		isRoot = false
		replicas = nil
		return
	}

	isRoot = true
	//Debug.Printf("FETCH the root is %v, but also i am %v", root, local.node)
	replicas = local.locationsByKey.Get(key)

	return
}

// - Register all of the provided objects in the local location map
// - If appropriate, add the from node to our local routing table
func (local *Node) Transfer(from RemoteNode, replicamap map[string][]RemoteNode) (err error) {
	// TODO: students should implement this
	local.locationsByKey.RegisterAll(replicamap, TIMEOUT)
	local.table.Add(from)
	return nil
}

// Utility function for recursively contacting nodes to get the root node for the provided ID.
//
// -  Starting from the specified node, call your recursive GetNextHop. This should
//    continually call GetNextHop until the root node is returned.
// -  Also keep track of any bad nodes that errored during lookup, and
//    as you come back from your recursive calls, remove the bad nodes from local
func (local *Node) findRoot(start RemoteNode, id ID) (RemoteNode, error) {
	Debug.Printf("Routing to %v\n", id)
	// TODO: students should implement this
	root, err, toRemove := start.GetNextHopRPC(id, 0)
	Debug.Printf("hey maybe the root for %v is this: %v", id, root)

	local.RemoveBadNodes(toRemove.Nodes())
	return root, err
}
