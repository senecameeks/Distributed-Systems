package liteminer

import "testing"

// testing one miner with upper bound of 100
func TestOneMiner(t *testing.T) {
	SetDebug(true)

	p, err := CreatePool("")
	if err != nil {
		t.Errorf("Received error %v when creating pool", err)
	}

	addr := p.Addr.String()

	numMiners := 1
	miners := make([]*Miner, numMiners)
	for i := 0; i < numMiners; i++ {
		m, err := CreateMiner(addr)
		if err != nil {
			t.Errorf("Received error %v when creating miner", err)
		}
		miners[i] = m
	}

	client := CreateClient([]string{addr})

	data := "data"
	upperbound := uint64(100)
	nonces, err := client.Mine(data, upperbound)

	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(97)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}
}

// testing two miners with upper bound of 100
func TestTwoMiners(t *testing.T) {
	SetDebug(true)

	p, err := CreatePool("")
	if err != nil {
		t.Errorf("Received error %v when creating pool", err)
	}

	addr := p.Addr.String()

	numMiners := 2
	miners := make([]*Miner, numMiners)
	for i := 0; i < numMiners; i++ {
		m, err := CreateMiner(addr)
		if err != nil {
			t.Errorf("Received error %v when creating miner", err)
		}
		miners[i] = m
	}

	client := CreateClient([]string{addr})

	data := "data"
	upperbound := uint64(100)
	nonces, err := client.Mine(data, upperbound)

	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(97)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}
}

// testing two miners with upper bound of 100
func TestFiveMiners(t *testing.T) {
	SetDebug(true)

	p, err := CreatePool("")
	if err != nil {
		t.Errorf("Received error %v when creating pool", err)
	}

	addr := p.Addr.String()

	numMiners := 5
	miners := make([]*Miner, numMiners)
	for i := 0; i < numMiners; i++ {
		m, err := CreateMiner(addr)
		if err != nil {
			t.Errorf("Received error %v when creating miner", err)
		}
		miners[i] = m
	}

	client := CreateClient([]string{addr})

	data := "data"
	upperbound := uint64(100)
	nonces, err := client.Mine(data, upperbound)

	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(97)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}
}

func TestIntervals(t *testing.T) {
	SetDebug(true)

	interval1 := GenerateIntervals(100, 4)
	for i := 0; i < len(interval1); i++ {
		if int(interval1[i].Lower) != 25*i {
			t.Errorf("Lower not working")
		}
		if i != (len(interval1)-1) && int(interval1[i].Upper) != 25*(i+1) {
			t.Errorf("Upper not working")
		}
	}

	interval2 := GenerateIntervals(257, 7)
	for i := 0; i < len(interval2); i++ {
		if (i < 5 || i == len(interval2)-1) && int(interval2[i].Upper-interval2[i].Lower) != 37 {
			t.Errorf("Partition not working %v", interval2[i].Upper-interval2[i].Lower)
		} else if i == 5 && int(interval2[i].Upper-interval2[i].Lower) != 36 {
			t.Errorf("Partition not working %v", interval2[i].Upper-interval2[i].Lower)
		}
	}

	interval3 := GenerateIntervals(1, 2)
	if interval3[0].Lower != 0 || interval3[1].Lower != 1 {
		t.Errorf("Lower not working on interval of size 1")
	}
	if interval3[0].Upper != 1 || interval3[1].Upper != 2 {
		t.Errorf("Upper not working on interval of size 1")
	}
}

func TestMultipleTransactions(t *testing.T) {
	SetDebug(true)

	p, err := CreatePool("")
	if err != nil {
		t.Errorf("Received error %v when creating pool", err)
	}

	addr := p.Addr.String()

	numMiners := 2
	miners := make([]*Miner, numMiners)
	for i := 0; i < numMiners; i++ {
		m, err := CreateMiner(addr)
		if err != nil {
			t.Errorf("Received error %v when creating miner", err)
		}
		miners[i] = m
	}

	client := CreateClient([]string{addr})

	data := "josh"
	upperbound := uint64(5000000)
	nonces, err := client.Mine(data, upperbound)
	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(3586653)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}

	data = "sam"
	nonces, err = client.Mine(data, upperbound)
	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(3948011)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}

	data = "will"
	nonces, err = client.Mine(data, upperbound)
	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(3848253)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}

	data = "jim"
	nonces, err = client.Mine(data, upperbound)
	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(3420565)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}

	data = "tom"
	nonces, err = client.Mine(data, upperbound)
	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(782614)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}

}

func TestMultTransactionsWithFailures(t *testing.T) {
	SetDebug(true)

	p, err := CreatePool("")
	if err != nil {
		t.Errorf("Received error %v when creating pool", err)
	}

	addr := p.Addr.String()

	numMiners := 6
	miners := make([]*Miner, numMiners)
	for i := 0; i < numMiners-2; i++ {
		m, err := CreateMiner(addr)
		if err != nil {
			t.Errorf("Received error %v when creating miner", err)
		}
		miners[i] = m
	}

	client := CreateClient([]string{addr})

	data := "josh"
	upperbound := uint64(5000000)
	nonces, err := client.Mine(data, upperbound)
	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(3586653)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}

	miners[2].Shutdown()

	data = "sam"
	nonces, err = client.Mine(data, upperbound)
	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(3948011)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}

	miners[0].Shutdown()
	m, err := CreateMiner(addr)
	if err != nil {
		t.Errorf("Received error %v when creating miner", err)
	}
	miners[2] = m
	m, err = CreateMiner(addr)
	if err != nil {
		t.Errorf("Received error %v when creating miner", err)
	}
	miners[4] = m
	m, err = CreateMiner(addr)
	if err != nil {
		t.Errorf("Received error %v when creating miner", err)
	}
	miners[5] = m

	data = "will"
	nonces, err = client.Mine(data, upperbound)
	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(3848253)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}

	miners[3].Shutdown()

	data = "jim"
	nonces, err = client.Mine(data, upperbound)
	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(3420565)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}

	m, err = CreateMiner(addr)
	if err != nil {
		t.Errorf("Received error %v when creating miner", err)
	}
	miners[0] = m
	m, err = CreateMiner(addr)
	if err != nil {
		t.Errorf("Received error %v when creating miner", err)
	}
	miners[3] = m

	data = "tom"
	nonces, err = client.Mine(data, upperbound)
	if err != nil {
		t.Errorf("Received error %v when mining", err)
	} else {
		for _, nonce := range nonces {
			expected := int64(782614)
			if nonce != expected {
				t.Errorf("Expected nonce %d, but received %d", expected, nonce)
			}
		}
	}

}
