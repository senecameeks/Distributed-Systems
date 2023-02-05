/*
 *  Brown University, CS138, Spring 2019
 *
 *  Purpose: Defines global constants and functions to create and join a new
 *  node into a Tapestry mesh, and functions for altering the routing table
 *  and backpointers of the local node that are invoked over RPC.
 */

package tapestry

import (
	"fmt"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	// Uncomment for xtrace
	// util "github.com/brown-csci1380/tracing-framework-go/xtrace/grpcutil"
	// xtr "github.com/brown-csci1380/tracing-framework-go/xtrace/client"
)

const BASE = 16    // The base of a digit of an ID.  By default, a digit is base-16.
const DIGITS = 40  // The number of digits in an ID.  By default, an ID has 40 digits.
const RETRIES = 3  // The number of retries on failure. By default we have 3 retries.
const K = 10       // During neighbor traversal, trim the neighborset to this size before fetching backpointers. By default this has a value of 10.
const SLOTSIZE = 3 // The each slot in the routing table should store this many nodes. By default this is 3.

const REPUBLISH = 10 * time.Second // Object republish interval for nodes advertising objects.
const TIMEOUT = 25 * time.Second   // Object timeout interval for nodes storing objects.

// The main struct for the local Tapestry node. Methods can be invoked locally on this struct.
type Node struct {
	node           RemoteNode    // The ID and address of this node
	table          *RoutingTable // The routing table
	backpointers   *Backpointers // Backpointers to keep track of other nodes that point to us
	locationsByKey *LocationMap  // Stores keys for which this node is the root
	blobstore      *BlobStore    // Stores blobs on the local node
	server         *grpc.Server
}

func (local *Node) String() string {
	return fmt.Sprintf("Tapestry Node %v at %v", local.node.Id, local.node.Address)
}

// Called in tapestry initialization to create a tapestry node struct
func newTapestryNode(node RemoteNode) *Node {
	serverOptions := []grpc.ServerOption{}
	// Uncomment for xtrace
	// serverOptions = append(serverOptions, grpc.UnaryInterceptor(util.XTraceServerInterceptor))
	n := new(Node)
	Debug.Printf("Starting new tapestry node")
	n.node = node
	n.table = NewRoutingTable(node)
	n.backpointers = NewBackpointers(node)
	n.locationsByKey = NewLocationMap()
	n.blobstore = NewBlobStore()
	n.server = grpc.NewServer(serverOptions...)

	return n
}

// Start a tapestry node on the specified port. Optionally, specify the address
// of an existing node in the tapestry mesh to connect to; otherwise set to "".
func Start(port int, connectTo string) (*Node, error) {
	return start(RandomID(), port, connectTo)
}

// Private method, useful for testing: start a node with the specified ID rather than a random ID.
func start(id ID, port int, connectTo string) (tapestry *Node, err error) {

	// Create the RPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, err
	}

	// Get the hostname of this machine
	name, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("Unable to get hostname of local machine to start Tapestry node. Reason: %v", err)
	}

	// Get the port we are bound to
	_, actualport, err := net.SplitHostPort(lis.Addr().String()) //fmt.Sprintf("%v:%v", name, port)
	if err != nil {
		return nil, err
	}

	// The actual address of this node
	address := fmt.Sprintf("%s:%s", name, actualport)

	// Uncomment for xtrace
	// xtr.NewTask("startup")
	// Trace.Print("Tapestry Starting up...")
	// xtr.SetProcessName(fmt.Sprintf("Tapestry %X... (%v)", id[:5], address))

	// Create the local node
	tapestry = newTapestryNode(RemoteNode{Id: id, Address: address})
	fmt.Printf("Created tapestry node %v\n", tapestry)
	Trace.Printf("Created tapestry node")

	RegisterTapestryRPCServer(tapestry.server, tapestry)
	fmt.Printf("Registered RPC Server\n")
	go tapestry.server.Serve(lis)

	// If specified, connect to the provided address
	if connectTo != "" {
		// Get the node we're joining
		node, err := SayHelloRPC(connectTo, tapestry.node)
		if err != nil {
			return nil, fmt.Errorf("Error joining existing tapestry node %v, reason: %v", address, err)
		}
		err = tapestry.Join(node)
		if err != nil {
			return nil, err
		}
	}

	return tapestry, nil
}

// Invoked when starting the local node, if we are connecting to an existing Tapestry.
//
// - Find the root for our node's ID
// - Call AddNode on our root to initiate the multicast and receive our initial neighbor set
// - Iteratively get backpointers from the neighbor set and populate routing table
func (local *Node) Join(otherNode RemoteNode) (err error) {
	Debug.Println("Joining", otherNode)

	// Route to our root
	root, err := local.findRoot(otherNode, local.node.Id)
	if err != nil {
		return fmt.Errorf("Error joining existing tapestry node %v, reason: %v", otherNode, err)
	}

	// Add ourselves to our root by invoking AddNode on the remote node
	neighbors, err := root.AddNodeRPC(local.node)
	if err != nil { //
		return fmt.Errorf("Error adding ourselves to root node %v, reason: %v", root, err)
	}

	// Add the neighbors to our local routing table.
	for _, n := range neighbors {
		local.addRoute(n)
	}

	// TODO: students should implement the backpointer traversal portion of Join
	local.TraverseBackpointers(neighbors, SharedPrefixLength(local.node.Id, otherNode.Id))

	return nil
}

// Helper function used to recursively get backpointers from the neighbor set and populate routing table
//
// - Implements pseudocode for TraverseBackpointers from handout
func (local *Node) TraverseBackpointers(neighbors []RemoteNode, level int) {
	if level >= 0 {
		nextNeighbors := neighbors

		for _, neighbor := range neighbors {
			bps, err := neighbor.GetBackpointersRPC(local.node, level)
			if err != nil {
				// haha i do not trust this
				local.table.Remove(neighbor)
			}

			nextNeighbors = append(nextNeighbors, bps...)
		}

		for _, nextNeighbor := range nextNeighbors {
			local.addRoute(nextNeighbor)
		}

		next := local.RemoveDuplicatesAndTrimToK(nextNeighbors)
		local.TraverseBackpointers(next, level-1)
	}

	return
}

// Helper function used to remove duplicate neighbors and trim the total number to K based on closeness
//
func (local *Node) RemoveDuplicatesAndTrimToK(neighbors []RemoteNode) (next []RemoteNode) {
	next = make([]RemoteNode, 0, K)
	var nextNeighbor bool

	for _, neighbor := range neighbors {
		nextNeighbor = false
		if len(next) == 0 {
			next = append(next, neighbor)
			continue
		}

		for i := 0; i < len(next); i++ {
			//Debug.Printf("id of next: %v \nid of neighb: %v \nequality: %v",next[i].Id, neighbor.Id, next[i].Id == neighbor.Id)
			if next[i].Id == neighbor.Id {
				nextNeighbor = true
				break
			}
			if local.node.Id.Closer(neighbor.Id, next[i].Id) {
				if len(next) == K {
					copy(next[i+1:], next[i:K])
				} else {
					next = append(next, neighbor)
					copy(next[i+1:], next[i:])
				}
				next[i] = neighbor
				nextNeighbor = true
				break
			}
		}

		if !nextNeighbor && len(next) < K {
			next = append(next, neighbor)
		}
	}

	return
}

// We are the root node for some new node joining the tapestry.
//
// - Begin the acknowledged multicast
// - Return the neighborset from the multicast
func (local *Node) AddNode(node RemoteNode) (neighborset []RemoteNode, err error) {
	return local.AddNodeMulticast(node, SharedPrefixLength(node.Id, local.node.Id))
}

// A new node is joining the tapestry, and we are a need-to-know node participating in the multicast.
//
// - Propagate the multicast to the specified row in our routing table
// - Await multicast response and return the neighborset
// - Add the route for the new node
// - Begin transfer of appropriate replica info to the new node
func (local *Node) AddNodeMulticast(newnode RemoteNode, level int) (neighbors []RemoteNode, err error) {
	Debug.Printf("Add node multicast %v at level %v\n", newnode, level)
	// TODO: students should implement this
	if level == DIGITS {
		neighbors = nil
		err = nil
		return
	}

	targets := local.table.GetLevel(level) // includes local node
	//Debug.Printf("hey there are %v targets for level %v and the targets are %v", len(targets), level, targets)
	var results []RemoteNode

	for _, target := range targets {
		newRes, err := target.AddNodeMulticastRPC(newnode, level+1)
		if err != nil {
			// check status of node probably dead and to be removed, do not necessarily return right away, by cases
			// remove target
			local.table.Remove(target)
		}

		results = append(results, newRes...)
	}

	local.addRoute(newnode)

	ret := local.locationsByKey.GetTransferRegistrations(local.node, newnode)
	err = newnode.TransferRPC(local.node, ret)
	// consider handling HANDLE ME
	if err != nil {
		local.table.Remove(newnode)
	}
	neighbors = append(targets, results...)

	return
}

// Add the from node to our backpointers, and possibly add the node to our
// routing table, if appropriate
func (local *Node) AddBackpointer(from RemoteNode) (err error) {
	if local.backpointers.Add(from) {
		Debug.Printf("Added backpointer %v\n", from)
	}
	local.addRoute(from)
	return
}

// Remove the from node from our backpointers
func (local *Node) RemoveBackpointer(from RemoteNode) (err error) {
	if local.backpointers.Remove(from) {
		Debug.Printf("Removed backpointer %v\n", from)
	}
	return
}

// Get all backpointers at the level specified, and possibly add the node to our
// routing table, if appropriate
func (local *Node) GetBackpointers(from RemoteNode, level int) (backpointers []RemoteNode, err error) {
	Debug.Printf("Sending level %v backpointers to %v\n", level, from)
	backpointers = local.backpointers.Get(level)
	local.addRoute(from)
	return
}

// The provided nodes are bad and we should discard them
// - Remove each node from our routing table
// - Remove each node from our set of backpointers
func (local *Node) RemoveBadNodes(badnodes []RemoteNode) (err error) {
	for _, badnode := range badnodes {
		if local.table.Remove(badnode) {
			Debug.Printf("Removed bad node %v\n", badnode)
		}
		if local.backpointers.Remove(badnode) {
			Debug.Printf("Removed bad node backpointer %v\n", badnode)
		}
	}
	return
}

// Utility function that adds a node to our routing table.
//
// - Adds the provided node to the routing table, if appropriate.
// - If the node was added to the routing table, notify the node of a backpointer
// - If an old node was removed from the routing table, notify the old node of a removed backpointer
func (local *Node) addRoute(node RemoteNode) (err error) {
	// TODO: students should implement this
	var succ bool
	var prev *RemoteNode
	if local.node.Id != node.Id {
		succ, prev = local.table.Add(node)
	} else {
		//Debug.Printf("Adding route to itself- returning")
		return
	}
	if !succ {
		return
	} else {
		err = node.AddBackpointerRPC(local.node)
		if err != nil {
			//delete node from my (local) routing table
			Debug.Printf("AddBackpointerRPC failed, remove node from routing table")
			Debug.Printf("removing node: %v from local:%v", node, local.node)
			local.table.Remove(node)
		}
		if prev != nil {
			err = prev.RemoveBackpointerRPC(local.node)
			if err != nil {
				Debug.Printf("RemoveBackpointerRPC failed, remove node from routing table")
				Debug.Printf("removing prev: %v from local:%v", prev, local.node)
				local.table.Remove(*prev)
			}
		}
	}

	return
}
