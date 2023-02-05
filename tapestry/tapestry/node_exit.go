/*
 *  Brown University, CS138, Spring 2019
 *
 *  Purpose: Defines functions for a node leaving the Tapestry mesh, and
 *  transferring its stored locations to a new node.
 */

package tapestry

import (
	// Uncomment for xtrace
	// xtr "github.com/brown-csci1380/tracing-framework-go/xtrace/client"
)

// Kill this node without gracefully leaving the tapestry.
func (local *Node) Kill() {
	local.blobstore.DeleteAll()
	local.server.Stop()
}

// Function to gracefully exit the Tapestry mesh.
//
// - Notify the nodes in our backpointers that we are leaving by calling NotifyLeave
// - If possible, give each backpointer a suitable alternative node from our routing table
func (local *Node) Leave() (err error) {
	// TODO: students should implement this
	for lvl := 0; lvl < DIGITS; lvl++ {
		bps := local.backpointers.Get(lvl)
		for i := 0; i < len(bps); i++ {
			neighbors := local.table.GetLevel(lvl)
			if len(neighbors) > 0 {
				closestNode := neighbors[0]
				for i := range neighbors {
					n := neighbors[i]
					if !local.node.Id.Closer(closestNode.Id, n.Id) {
						closestNode = n
					}
				}
			bps[i].NotifyLeaveRPC(local.node, &closestNode) // who tf is our replacement node ??
			}
			// find replacement potential?
		}
	}

	local.blobstore.DeleteAll()
	go local.server.GracefulStop()
	return
}

// Another node is informing us of a graceful exit.
// - Remove references to the from node from our routing table and backpointers
// - If replacement is not nil, add replacement to our routing table
func (local *Node) NotifyLeave(from RemoteNode, replacement *RemoteNode) (err error) {
	Debug.Printf("Received leave notification from %v with replacement node %v\n", from, replacement)

	// TODO: students should implement this
	wasRemoved := local.table.Remove(from)
	backPointerBool := local.backpointers.Remove(from)
	Debug.Printf("wasRemoved: %v and backPointerBool: %v", wasRemoved, backPointerBool)
	if replacement != nil {
		added, previous := local.table.Add(*replacement)
		Debug.Printf("added: %v, previous: %v", added, previous)
	}

	return
}
