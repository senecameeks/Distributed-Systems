package tapestry

import (
	"testing"
)

// WHAT WE TEST
// 1) Empty table-- add something in all the places
// 2) if somethings there, does it put it into the array in correct spot
// 3) if something is already there-- does it kick one of them out if it needs to or // NOTE:
// 4) if the node is already there, does it not do anything

func TestAddOnEmptyTable(t *testing.T) {
	SetDebug(true)
  sharedPrefixLen := uint8(4)
  id := RandomID()
  Debug.Printf("%v", id)
  //remoteNode = new RemoteNode{Id: id, Address: "address"}
  //tapestry := newTapestryNode(RemoteNode{Id: id, Address: "address"})
  table := NewRoutingTable(RemoteNode{Id: id, Address: "address"})
  newNode := GenerateSharedPrefixNode(id, sharedPrefixLen)
  _b, prev := table.Add(newNode)
  if _b == false {
    t.Errorf("Added boolean should return true")
  }
  if prev != nil {
    t.Errorf("Table was empty. Previous node should be nil.")
  }

  //Since table is empty, newNode should be in all places up to prev
  for i := uint8(0); i < sharedPrefixLen; i++ {
    if (*table.rows[i][newNode.Id[i]])[1] != newNode {
      t.Errorf("newNode should have been added in this place behind the root at row: %v, col: %v", i, newNode.Id[i])
    }
  }

  if (*table.rows[sharedPrefixLen][newNode.Id[sharedPrefixLen]])[0] != newNode {
      t.Errorf("newNode should have been added in this place at row: %v, col: %v", sharedPrefixLen, newNode.Id[sharedPrefixLen])
    }
}

// WHAT WE TEST
// 1) Empty table-- does it remove nothing
// 2) if somethings there, does it remove it correctly
// 3) if multiple things are there-- does it remove the correct one

func TestRemoveOnEmptyTable(t *testing.T) {
  SetDebug(true)

  id := RandomID()
  node := RemoteNode{Id: id, Address: "address"}
  tab := NewRoutingTable(node)

  // check how it handles removing node that does not exist
  if tab.Remove(RemoteNode{Id: RandomID(), Address: "address"}) {
    t.Errorf("Remove claimed it removed node that never existed")
  }
}

func TestRemoveOnOneElemTable(t *testing.T) {
  SetDebug(true)

  id := RandomID()
  node := RemoteNode{Id: id, Address: "address"}
  tab := NewRoutingTable(node)

  sharedPrefixLen := uint8(5)
  newNode1 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  tab.Add(newNode1)

  // check how it handles removing node that does not exist
  if tab.Remove(RemoteNode{Id: RandomID(), Address: "address"}) {
    t.Errorf("Remove claimed it removed node that never existed")
  }

  // check removing single other node in table
  if !tab.Remove(newNode1) {
    t.Errorf("Remove claimed a present node does not exist")
  }

  // check if node removed
  for i := uint8(0); i < sharedPrefixLen; i++ {
    if len(*tab.rows[i][newNode1.Id[i]]) > 1 && (*tab.rows[i][newNode1.Id[i]])[1] == newNode1 {
      t.Errorf("newNode should have been removed in this place behind the root at row: %v, col: %v", i, newNode1.Id[i])
    }
  }

  if len(*tab.rows[sharedPrefixLen][newNode1.Id[sharedPrefixLen]]) > 1 {
    t.Errorf("newNode should have been removed in this place behind the root at row: %v, col: %v", sharedPrefixLen, newNode1.Id[sharedPrefixLen])
  }
}

func TestRemoveOnMultipleElementTable(t *testing.T) {
  SetDebug(true) 

  id := RandomID()
  node := RemoteNode{Id: id, Address: "address"}
  tab := NewRoutingTable(node)

  sharedPrefixLen := uint8(2)
  newNode2 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  tab.Add(newNode2)
  sharedPrefixLen = uint8(4)
  newNode3 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  tab.Add(newNode3)
  sharedPrefixLen = uint8(7)
  newNode4 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  tab.Add(newNode4)

  // check how it handles removing node that does not exist
  if tab.Remove(RemoteNode{Id: RandomID(), Address: "address"}) {
    t.Errorf("Remove claimed it removed node that never existed")
  }

  // check removing another node in table
  if !tab.Remove(newNode3) {
    t.Errorf("Remove claimed a present node does not exist")
  }

  // check if node actually removed
  for i := uint8(0); i < uint8(4); i++ {
    if len(*tab.rows[i][newNode3.Id[i]]) > 2 && (*tab.rows[i][newNode3.Id[i]])[2] == newNode3 {
      t.Errorf("newNode should have been removed in this place behind the root at row: %v, col: %v", i, newNode3.Id[i])
    }
  }

  if len(*tab.rows[uint8(4)][newNode3.Id[uint8(4)]]) > 1 {
    t.Errorf("newNode should have been removed in this place behind the root at row: %v, col: %v", 4, newNode3.Id[uint8(4)])
  }
}

// Helper method that generates a new RemoteId with the desired shared prefix len

func GenerateSharedPrefixNode(id ID, desiredSharedLen uint8) (node RemoteNode) {
  newId := RandomID()
  for i := uint8(0); i < desiredSharedLen; i++ {
    newId[i] = id[i]
  }

  if newId[desiredSharedLen] == id[desiredSharedLen] {
    newId[desiredSharedLen] += Digit(1)
    newId[desiredSharedLen] %= BASE
  }

  node = RemoteNode{Id: newId, Address:"address"}

  return
}
