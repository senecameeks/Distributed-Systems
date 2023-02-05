package tapestry

import (
	"testing"
)

func TestRemoveDuplicatesAndTrimToK(t *testing.T) {
	SetDebug(true)

  sharedPrefixLen := uint8(4)
  id := RandomID()
  //Debug.Printf("%v", id)
  remoteNode := RemoteNode{Id: id, Address: "address"}
  tapestry := newTapestryNode(remoteNode)

  neighbs := make([]RemoteNode, 0, K)

  neighbs = tapestry.RemoveDuplicatesAndTrimToK(neighbs)

  if len(neighbs) != 0 {
    t.Errorf("RemoveDuplicatesAndTrimToK claimed a nonexistent node exists")
  }

  node1 := GenerateSharedPrefixNode(id, sharedPrefixLen)

  neighbs = append(neighbs, node1)
  neighbs = tapestry.RemoveDuplicatesAndTrimToK(neighbs)

  if len(neighbs) != 1 {
    t.Errorf("RemoveDuplicatesAndTrimToK removed a single node")
  }

  neighbs = append(neighbs, node1)
  neighbs = tapestry.RemoveDuplicatesAndTrimToK(neighbs)

  if len(neighbs) != 1 {
    t.Errorf("RemoveDuplicatesAndTrimToK failed to remove a duplicate node, %v", neighbs)
  }

  sharedPrefixLen = uint8(6)
  node2 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  node3 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  node4 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  node5 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  node6 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  node7 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  node8 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  node9 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  node10 := GenerateSharedPrefixNode(id, sharedPrefixLen)
  node11 := GenerateSharedPrefixNode(id, sharedPrefixLen)

  neighbs = make([]RemoteNode, 0, K+1)
  neighbs = append(neighbs, node1)
  neighbs = append(neighbs, node2)
  neighbs = append(neighbs, node3)
  neighbs = append(neighbs, node4)
  neighbs = append(neighbs, node5)
  neighbs = append(neighbs, node6)
  neighbs = append(neighbs, node7)
  neighbs = append(neighbs, node8)
  neighbs = append(neighbs, node9)
  neighbs = append(neighbs, node10)
  neighbs = append(neighbs, node11)

  neighbs = tapestry.RemoveDuplicatesAndTrimToK(neighbs)

  if len(neighbs) != 10 {
    t.Errorf("RemoveDuplicatesAndTrimToK failed to trim to K nodes, has len %v", len(neighbs))
  }

  for _, n := range neighbs {
      if n == node1 {
        t.Errorf("RemoveDuplicatesAndTrimToK failed to trim to the furthest node")
      }
  }

  neighbs = make([]RemoteNode, 0, K+5)
  neighbs = append(neighbs, node1)
  neighbs = append(neighbs, node2)
  neighbs = append(neighbs, node3)
  neighbs = append(neighbs, node4)
  neighbs = append(neighbs, node5)
  neighbs = append(neighbs, node6)
  neighbs = append(neighbs, node7)
  neighbs = append(neighbs, node8)
  neighbs = append(neighbs, node9)
  neighbs = append(neighbs, node10)
  neighbs = append(neighbs, node11)
  neighbs = append(neighbs, node6)
  neighbs = append(neighbs, node7)
  neighbs = append(neighbs, node8)
  neighbs = append(neighbs, node9)

  neighbs = tapestry.RemoveDuplicatesAndTrimToK(neighbs)

  if len(neighbs) != 10 {
    t.Errorf("RemoveDuplicatesAndTrimToK failed to trim to K nodes, has len %v", len(neighbs))
  }

}
