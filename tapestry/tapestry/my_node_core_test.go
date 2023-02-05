package tapestry

import (
	"testing"
)

func TestFindRoot(t *testing.T) {
	SetDebug(true)
  tap, _:= MakeTapestries(true,"1", "2", "3","432")
  id := MakeID("47")
  root, _ := tap[0].findRoot(tap[1].node,id)
  if root != tap[3].node {
    t.Errorf("Returned wrong root for ID:%v. Returned %v", id, root)
  }

}
