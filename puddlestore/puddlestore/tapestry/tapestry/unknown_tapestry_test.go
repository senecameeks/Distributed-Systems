package tapestry

import (
	"testing"
)

/*
	TODO: test cases

	general strategy: start multiple tapestry nodes then function test

*/

func checkSelf(tapestry *Node, t *testing.T) {
	node := tapestry.node
	id := tapestry.node.Id
	table := tapestry.table
	for i := 0; i < DIGITS; i++ {
		if (*table.rows[i][id[i]])[0] != node {
			t.Errorf("Routing table entry [%v][%v] was %v instead of %v", i, id[i], (*table.rows[i][id[i]])[0], node)
		}
	}
}

func checkEntry(entry *[]RemoteNode, id ID) bool {
	for i := 0; i < len(*entry); i++ {
		if (*entry)[i].Id == id {
			return true
		}
	}
	return false
}

func checkTable(table *RoutingTable, id ID, t *testing.T) {
	shared_length := SharedPrefixLength(table.local.Id, id)
	for i := 0; i < DIGITS; i++ {
		for j := 0; j < BASE; j++ {
			if i == shared_length && Digit(j) == id[i] {
				if !checkEntry(table.rows[i][j], id) {
					t.Errorf("Expected (%v, %v) slot to contain %v, but slot was %v", i, j, id, table.rows[i][j])
				}
			} else {
				if checkEntry(table.rows[i][j], id) {
					t.Errorf("Slot (%v, %v) unexpectedly contained %v. Slot: %v", i, j, id, table.rows[i][j])
				}
			}
		}
	}
}

func TestTapestry(t *testing.T) {
	ida, _ := ParseID("0000000000000000000000000000000000000000")
	idb, _ := ParseID("0000000000000000000000000000000000000001")

	str := "127.0.0.1:6608"

	nodea, err := start(ida, 6608, "")
	if err != nil {
		t.Error(err)
		return
	}

	nodeb, err := start(idb, 6609, str)
	if err != nil {
		t.Error(err)
		return
	}

	checkSelf(nodea, t)
	checkSelf(nodeb, t)

	checkTable(nodea.table, idb, t)
	checkTable(nodeb.table, ida, t)

	//    nodea.Kill()
	//    nodeb.Kill()
}
