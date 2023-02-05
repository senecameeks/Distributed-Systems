package tapestry

import "testing"

func TestSimpleRoutingTable(t *testing.T) {
	ida, _ := ParseID("0000000000000000000000000000000000000000")
	node := RemoteNode{ida, "testroutingtable"}
	table := NewRoutingTable(node)

	idb, _ := ParseID("0000000000000000000000000000000000000001")
	idc, _ := ParseID("0000000000000000000000000000000000000002")
	idd, _ := ParseID("0000000000000000000000000000000000000003")

	nodeb := RemoteNode{idb, ""}
	nodec := RemoteNode{idc, ""}
	noded := RemoteNode{idd, ""}

	added, _ := table.Add(nodeb)
	if !added {
		t.Errorf("Failed to add %v", nodeb)
	}
	added, _ = table.Add(nodec)
	if !added {
		t.Errorf("Failed to add %v", nodec)
	}
	added, _ = table.Add(noded)
	if !added {
		t.Errorf("Failed to add %v", noded)
	}
}

func TestSimpleRoutingTable2(t *testing.T) {
	ida, _ := ParseID("0000000000000000000000000000000000000010")
	node := RemoteNode{ida, "testroutingtable"}
	table := NewRoutingTable(node)

	idb, _ := ParseID("0000000000000000000000000000000000000021")
	idc, _ := ParseID("0000000000000000000000000000000000000022")
	idd, _ := ParseID("0000000000000000000000000000000000000023")
	ide, _ := ParseID("0000000000000000000000000000000000000024")

	nodeb := RemoteNode{idb, ""}
	nodec := RemoteNode{idc, ""}
	noded := RemoteNode{idd, ""}
	nodee := RemoteNode{ide, ""}

	added, _ := table.Add(nodeb)
	if !added {
		t.Errorf("Failed to add %v", nodeb)
	}
	added, _ = table.Add(nodec)
	if !added {
		t.Errorf("Failed to add %v", nodec)
	}
	added, _ = table.Add(noded)
	if !added {
		t.Errorf("Failed to add %v", noded)
	}
	added, _ = table.Add(nodee)
	if added {
		t.Errorf("Unexpectedly added %v", nodee)
	}
}

func TestSimpleRoutingTable3(t *testing.T) {
	ida, _ := ParseID("0000000000000000000000000000000000000010")
	node := RemoteNode{ida, "testroutingtable"}
	table := NewRoutingTable(node)

	idb, _ := ParseID("0000000000000000000000000000000000000001")
	idc, _ := ParseID("0000000000000000000000000000000000000002")
	idd, _ := ParseID("0000000000000000000000000000000000000003")
	ide, _ := ParseID("0000000000000000000000000000000000000004")

	nodeb := RemoteNode{idb, ""}
	nodec := RemoteNode{idc, ""}
	noded := RemoteNode{idd, ""}
	nodee := RemoteNode{ide, ""}

	added, _ := table.Add(nodeb)
	if !added {
		t.Errorf("Failed to add %v", nodeb)
	}
	added, _ = table.Add(nodec)
	if !added {
		t.Errorf("Failed to add %v", nodec)
	}
	added, _ = table.Add(noded)
	if !added {
		t.Errorf("Failed to add %v", noded)
	}
	added, prev := table.Add(nodee)
	if !added {
		t.Errorf("Failed to add %v", nodee)
	}
	if prev == nil {
		t.Errorf("Prev is nil")
	} else if *prev != nodeb {
		t.Errorf("Prev was %v, expected %v", prev, nodeb)
	}

}

func TestSimpleRoutingTable4(t *testing.T) {
	ida, _ := ParseID("0000000000000000000000000000000000000010")
	node := RemoteNode{ida, "testroutingtable"}
	table := NewRoutingTable(node)

	idb, _ := ParseID("0000000000000000000000000000000000000001")
	idc, _ := ParseID("0000000000000000000000000000000000000002")
	idd, _ := ParseID("0000000000000000000000000000000000000023")
	ide, _ := ParseID("0000000000000000000000000000000000000024")

	nodeb := RemoteNode{idb, ""}
	nodec := RemoteNode{idc, ""}
	noded := RemoteNode{idd, ""}
	nodee := RemoteNode{ide, ""}

	added, _ := table.Add(nodeb)
	if !added {
		t.Errorf("Failed to add %v", nodeb)
	}
	added, _ = table.Add(nodec)
	if !added {
		t.Errorf("Failed to add %v", nodec)
	}
	added, _ = table.Add(noded)
	if !added {
		t.Errorf("Failed to add %v", noded)
	}
	added, _ = table.Add(nodee)
	if !added {
		t.Errorf("Failed to add %v", nodee)
	}

}
