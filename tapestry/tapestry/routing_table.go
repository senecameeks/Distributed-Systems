/*
 *  Brown University, CS138, Spring 2019
 *
 *  Purpose: Defines the RoutingTable type and provides methods for interacting
 *  with it.
 */

package tapestry

import (
	"sync"
)

// A routing table has a number of levels equal to the number of digits in an ID
// (default 40). Each level has a number of slots equal to the digit base
// (default 16). A node that exists on level n thereby shares a prefix of length
// n with the local node. Access to the routing table protected by a mutex.
type RoutingTable struct {
	local RemoteNode                  // The local tapestry node
	rows  [DIGITS][BASE]*[]RemoteNode // The rows of the routing table
	mutex sync.Mutex                  // To manage concurrent access to the routing table (could also have a per-level mutex)
}

// Creates and returns a new routing table, placing the local node at the
// appropriate slot in each level of the table.
func NewRoutingTable(me RemoteNode) *RoutingTable {
	t := new(RoutingTable)
	t.local = me

	// Create the node lists with capacity of SLOTSIZE
	for i := 0; i < DIGITS; i++ {
		for j := 0; j < BASE; j++ {
			slot := make([]RemoteNode, 0, SLOTSIZE)
			t.rows[i][j] = &slot
		}
	}

	// Make sure each row has at least our node in it
	for i := 0; i < DIGITS; i++ {
		slot := t.rows[i][t.local.Id[i]]
		*slot = append(*slot, t.local)
	}

	return t
}

// Adds the given node to the routing table.
//
// Returns true if the node did not previously exist in the table and was subsequently added.
// Returns the previous node in the table, if one was overwritten.
func (t *RoutingTable) Add(node RemoteNode) (added bool, previous *RemoteNode) {
	//Debug.Printf("lock add %v to %v", node.Id, t.local.Id)
	t.mutex.Lock()

	// TODO: students should implement this
	// temp variable for potential swapping
	var temp RemoteNode
	// iterate through the levels
	for i := 0; i <= SharedPrefixLength(t.local.Id, node.Id); i++ {
		if i == DIGITS {
			break
		}

		slot := t.rows[i][node.Id[i]]
		// if the slot is not already full
		size := len(*slot)
		// check if already there
		if size > 0 {
			for i := 0; i < size; i++ {
				if node.Id == (*slot)[i].Id {
					added = false
					previous = nil
					t.mutex.Unlock()
					//Debug.Printf("unlock add %v to %v", node.Id, t.local.Id)
					return
				}
			}
		}

		if size < SLOTSIZE {
			*slot = append(*slot, node)
			if size > 1 {
				// deal with closeness
				if size == 2 {
					if t.local.Id.Closer((*slot)[1].Id, (*slot)[0].Id) {
						temp = (*slot)[0]
						(*slot)[0] = (*slot)[1]
						(*slot)[1] = temp
					}
				} else {
					if t.local.Id.Closer((*slot)[2].Id, (*slot)[1].Id) {
						temp = (*slot)[1]
						(*slot)[1] = (*slot)[2]
						(*slot)[2] = temp
						if t.local.Id.Closer((*slot)[1].Id, (*slot)[0].Id) {
							temp = (*slot)[0]
							(*slot)[0] = (*slot)[1]
							(*slot)[1] = temp
						}
					}
				}

			}
			added = true
			previous = nil
		} else { // if the slot is full
			if t.local.Id.Closer(node.Id, (*slot)[2].Id) {
				previous = &(*slot)[2]
				(*slot)[2] = node
				if t.local.Id.Closer((*slot)[2].Id, (*slot)[1].Id) {
					temp = (*slot)[1]
					(*slot)[1] = (*slot)[2]
					(*slot)[2] = temp
					if t.local.Id.Closer((*slot)[1].Id, (*slot)[0].Id) {
						temp = (*slot)[0]
						(*slot)[0] = (*slot)[1]
						(*slot)[1] = temp
					}
				}
				added = true
			} else {
				added = false
				previous = nil
			}
		}
	}

	t.mutex.Unlock()
	//Debug.Printf("unlock add %v to %v", node.Id, t.local.Id)
	return
}

// Removes the specified node from the routing table, if it exists.
// Returns true if the node was in the table and was successfully removed.
func (t *RoutingTable) Remove(node RemoteNode) (wasRemoved bool) {
	//Debug.Printf("lock remove %v from %v", node.Id, t.local.Id)
	t.mutex.Lock()

	// TODO: students should implement this
	sharedPrefixLen := SharedPrefixLength(node.Id, t.local.Id)
	wasRemoved = false
	for i := 0; i <= sharedPrefixLen; i++ {
		slot := t.rows[i][node.Id[i]]
		for j := 0; j < len(*slot); j++ {
			Debug.Printf("%v %v %v", j, len(*slot), (*slot))
			if (*slot)[j].Id == node.Id {
				//Debug.Printf("%v %v", len(*slot), (*slot)[j])
				if len(*slot) == 1 {
					(*slot) = make([]RemoteNode, 0, SLOTSIZE)
				} else {
					if j < SLOTSIZE-1 {  //bound check
						*slot = append((*slot)[:j],(*slot)[j+1:]...) //remove jth element
					} else {
						*slot = (*slot)[:j]
					}
					wasRemoved = true
				}
			}
		}
	}
	//t.PrintTable()
	t.mutex.Unlock()
	//Debug.Printf("unlock remove %v", t.local.Id)
	return
}

// Get all nodes on the specified level of the routing table, EXCLUDING the local node.
func (t *RoutingTable) GetLevel(level int) (nodes []RemoteNode) {
	//Debug.Printf("lock getlevel %v", t.local.Id)
	t.mutex.Lock()

	// TODO: students should implement this
	nodes = make([]RemoteNode, 0, DIGITS * BASE - 1)
	for i := 0; i < BASE; i++ {
		for j := 0; j < len(*t.rows[level][i]); j++ {
			if t.local.Id != (*t.rows[level][i])[j].Id {
				nodes = append(nodes, (*t.rows[level][i])[j])
			}
		}
	}

	t.mutex.Unlock()
	//Debug.Printf("unlock getlevel %v", t.local.Id)
	return
}

// Search the table for the closest next-hop node for the provided ID.
// Recur on that node with GetNextHop, until you arrive at the root.
// keep track of bad nodes in "toRemove", and remove all known bad nodes
// locally as you unwind from your recursion.
func (t *RoutingTable) GetNextHop(id ID, level int32) (node RemoteNode, err error, toRemove *NodeSet) {
	t.mutex.Lock()
	// TODO: students should implement this
	Debug.Printf("table id: %v\n looking for root id of: %v", t.local.Id, id)
	//t.PrintTable()

	row := SharedPrefixLength(id, t.local.Id)
	//Debug.Printf("Shared prefix length: %v", row)
	if row == DIGITS {
		return t.local, nil, NewNodeSet()
	}

	slot := t.rows[Digit(row)][id[row]]
	//t.PrintTable()
	if len(*slot) == 0 {
		// the id is not a node in the routing table
		col := id[row]
		for len(*t.rows[Digit(row)][col]) == 0 {
			col ++
			col %= BASE
		}
		// if this works then check for closest in slot
		slot = t.rows[row][col]
		// also see whatever tf is up w errors and toRemove prob from errors
		node, err, toRemove = (*slot)[0], nil, NewNodeSet()
		//Debug.Printf("len of slot = 0, im choosing %v to be my root ", node)
		t.mutex.Unlock()
		return
	}
	level++
	node, err, toRemove = (*slot)[0].GetNextHopRPC(id, level)

	//Debug.Printf("im choosing %v to be my root", node)

	t.mutex.Unlock()

	return
}

func (t * RoutingTable) PrintTable() {
	Debug.Printf("PRINTING TABLE for %v\n\n", t.local.Id)
	for i := 0; i < DIGITS; i++ {
		for j := 0; j < BASE; j++ {
			if len(*t.rows[i][j]) != 0 {
				Debug.Printf("======== : slotsize: %v", len(*t.rows[i][j]))
				for k := 0; k < len(*t.rows[i][j]); k++ {
					Debug.Printf("%v\n", (*t.rows[i][j])[k])
				}

			}
		}
	}
}
