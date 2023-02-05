package liteminer

import (
	"testing"
	"time"
)

// testing if one miner fails
func TestOneMinerFails(t *testing.T) {
	SetDebug(true)

	p, err := CreatePool("")
	if err != nil {
		t.Errorf("Received error %v when creating pool", err)
	}

	addr := p.Addr.String()

	numMiners := 4
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
	// sleep after mining
	wait_time := 500 * time.Millisecond
	nextTime := time.Now()
	nextTime = nextTime.Add(wait_time)
	time.Sleep(time.Until(nextTime))

	//Shut down miner
	miners[1].Shutdown()
	if !miners[1].IsShutdown {
		t.Errorf("Miner should be shutdown")
	}
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
