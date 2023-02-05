package tapestry

import (
	"testing"
	"bytes"
	//"time"
)

func TestSampleTapestrySetup(t *testing.T){
	SetDebug(true)
	tap, _:= MakeTapestries(true, "1", "3", "5", "7") //Make a tapestry with these ids
	//KillTapestries(tap[1], tap[2]) //Kill off two of them.
	tap[1].Leave()
	tap[2].Leave()
	next, _, toRemove := tap[0].table.GetNextHop(MakeID("2"), 0) //After killing 3 and 5, this should route to 7

	Debug.Printf("toRemove: %v", toRemove)
	Debug.Printf("next: %v\n desired: %v", next, tap[3].node)
	if next != tap[3].node {
		t.Errorf("Failed to kill successfully")
	}

}

func TestSampleTapestrySearch (t *testing.T){
	SetDebug(true)
	tap, _ := MakeTapestries(true, "100", "456", "1234") //make a sample tap
	tap[1].Store("look at this lad", []byte("an absolute unit"))

	result, _ := tap[0].Get("look at this lad") //Store a KV pair and try to fetch it
	if !bytes.Equal(result, []byte("an absolute unit")){ //Ensure we correctly get our KV
		t.Errorf("Get failed")
	}
}

func TestNewTapestry(t *testing.T) {
	SetDebug(true)
	tap, _ := MakeTapestries(true, "100", "456", "1234")
	tap[1].Store("key", []byte("value"))
	t1, err := start(MakeID("789"), 0, tap[1].node.Address)
	err = tap[1].Join(t1.node)  //Add another node
	if err != nil {
		t.Errorf("Error when joining new node")
	}
}

func TestAddOneNode(t *testing.T) {
	SetDebug(true)
	tap, _ := MakeTapestries(true, "100", "456", "1234")
	tap[1].Store("key", []byte("value"))
	//Debug.Printf("Neighbors: %v", neighbors)
	newNode, newTap, _ := AddOne("789", tap[1].node.Address, tap)
	if len(newTap) != len(tap) + 1 {
		t.Errorf("AddNode failed to add one more node: %v", newNode.node.Id)
	}

	//Add existing node
	//_, newTap, _ := AddOne("456", tap[0].node.Address, updatedTap)
	//if len(updatedTap) != len(newTap) {
	//	t.Errorf("New node: %v should not have been added because it already exists", newNode.node.Id)
	//}
}

func TestSampleTapestryAddNodes(t *testing.T){
	tap, _ := MakeTapestries(true, "1", "5", "9")
	node8, tap, _ := AddOne("8", tap[0].node.Address, tap) //Add some tap nodes after the initial construction
	_, tap, _ = AddOne("12", tap[0].node.Address, tap)


	next, _, _ := tap[1].table.GetNextHop(MakeID("7"), 0)
	if node8.node != next {
		t.Errorf("Addition of node failed")
	}
}

func TestRandomTapestry(t *testing.T) {
	SetDebug(true)
	tap, err := MakeRandomTapestries(47, 100)
	if err != nil {
		t.Errorf("MakeRandomTapestries failed with error: %v", err)
	}

	Debug.Printf("sup\n\n")
	tap[1].Store("look at this lad", []byte("an absolute unit"))
	Debug.Printf("able to store one thing")
	tap[79].Store("look at this lad", []byte("an absolute unit"))
	tap[83].Store("look at this lad", []byte("an absolute unit"))
	tap[42].Store("who shot ya", []byte("rest in peace biggie"))
	tap[39].Store("who shot ya", []byte("rest in peace biggie"))
	tap[57].Store("who shot ya", []byte("rest in peace biggie"))
	tap[7].Store("who shot ya", []byte("rest in peace biggie"))
	tap[42].Store("are you my dad?", []byte("never talk to me or my son ever again"))
	tap[1].Store("are you my dad?", []byte("never talk to me or my son ever again"))
	tap[45].Store("are you my dad?", []byte("never talk to me or my son ever again"))

	tap[81].Leave()
	tap[79].Leave()
	tap[62].Leave()
	tap[72].Leave()

	Debug.Printf("done storing stuff\n\n")
	time.Sleep(REPUBLISH)
	Debug.Printf("back at it\n\n")
	result, _ := tap[33].Get("look at this lad")
	if !bytes.Equal(result, []byte("an absolute unit")){ //Ensure we correctly get our KV
		t.Errorf("Get failed")
	}

	result, _ = tap[30].Get("who shot ya")
	if !bytes.Equal(result, []byte("rest in peace biggie")){ //Ensure we correctly get our KV
		t.Errorf("Get failed: returned")
	}

	result, _ = tap[70].Get("are you my dad?")
	if !bytes.Equal(result, []byte("never talk to me or my son ever again")){ //Ensure we correctly get our KV
		t.Errorf("Get failed")
	}
	/*
	for i := 0; i < len(tap)/2; i++ {
		tap[i].Kill()
	}


	KillTapestries(tap[1], tap[10], tap[42], tap[35], tap[24])
	Debug.Printf("\n\n\n\n\n\n\n\n\n")
	result, _ = tap[63].Get("look at this lad")
	if !bytes.Equal(result, []byte("an absolute unit")){ //Ensure we correctly get our KV
		t.Errorf("Get failed 2")
	}

	result, _ = tap[90].Get("who shot ya")
	if !bytes.Equal(result, []byte("rest in peace biggie")){ //Ensure we correctly get our KV
		t.Errorf("Get failed 3")
	}
	Debug.Printf("\n\n\n\n\n\n\n\n\n")
	result, _ = tap[70].Get("are you my dad?")
	if !bytes.Equal(result, []byte("never talk to me or my son ever again")){ //Ensure we correctly get our KV
		t.Errorf("Get failed")
	} */
}
