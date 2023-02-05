package raft

import (
	"testing"
	"time"
)

func TestRandomTimeout(t *testing.T) {
	hbt := time.Millisecond * 50
	heartbeatTimeout := randomTimeout(hbt)

	time.Sleep(hbt)
	time.Sleep(hbt)

	select {
	case _, ok := <-heartbeatTimeout:
		if !ok {
			t.Fatal("heartbeat timeout is closed")
		}
	default:
		t.Fatal("no heartbeat timeout sent")
	}

	et := time.Millisecond * 150
	electionTimeout := randomTimeout(et)

	time.Sleep(et)
	time.Sleep(et)

	select {
	case _, ok := <-electionTimeout:
		if !ok {
			t.Fatal("election timeout is closed")
		}
	default:
		t.Fatal("no election timeout sent")
	}
}
