package tapestry

import "testing"

func TestNodeSetsSemantics(t *testing.T) {
	set := NewNodeSet()
	for numadds := 1; numadds <= 1024; numadds *= 2 {
		node := RemoteNode{RandomID(), ""}
		if set.Contains(node) {
			t.Errorf("Set should not contain node %v yet. Set: %v", node, set)
		}
		for i := 0; i < numadds; i++ {
			set.Add(node)
		}
		if !set.Contains(node) {
			t.Errorf("Set should now contain node %v. Set: %v", node, set)
		}
		set.Remove(node)
		if set.Contains(node) {
			t.Errorf("Set should no longer contain %v. Set: %v", node, set)
		}
	}
}
