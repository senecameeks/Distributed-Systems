package raft

import (
	"errors"
	//"fmt"
	"testing"
	"time"
)

func TestPartition(t *testing.T) {
	suppressLoggers()

	config := DefaultConfig()
	config.ClusterSize = 5

	cluster, err := CreateLocalCluster(config)
	defer cleanupCluster(cluster)
	if err != nil {
		t.Fatal(err)
	}

	// wait for a leader to be elected
	time.Sleep(time.Second * WAIT_PERIOD)
	leader, err := findLeader(cluster)
	if err != nil {
		t.Fatal(err)
	}

	followers := make([]*RaftNode, 0)
	for _, node := range cluster {
		if node != leader {
			followers = append(followers, node)
		}
	}

	// partition into 2 clusters: one with leader and first follower; other with remaining 3 followers
	ff := followers[0]
	for _, follower := range followers[1:] {
		follower.NetworkPolicy.RegisterPolicy(*follower.GetRemoteSelf(), *ff.GetRemoteSelf(), false)
		follower.NetworkPolicy.RegisterPolicy(*follower.GetRemoteSelf(), *leader.GetRemoteSelf(), false)

		ff.NetworkPolicy.RegisterPolicy(*ff.GetRemoteSelf(), *follower.GetRemoteSelf(), false)
		leader.NetworkPolicy.RegisterPolicy(*leader.GetRemoteSelf(), *follower.GetRemoteSelf(), false)
	}
	//fmt.Printf("JUST PARTITIONED\n\n\n\n\n\n\n\n\n\n")
	// allow a new leader to be elected in partition of 3 nodes
	time.Sleep(time.Second * WAIT_PERIOD)
	newLeader, err := findLeader(followers)
	if err != nil {
		t.Fatal(err)
	}

	// check that old leader, which is cut off from new leader, still thinks it's leader
	if leader.State != LEADER_STATE {
		t.Fatal(errors.New("leader should remain leader even when partitioned"))
	}

	if leader.GetCurrentTerm() >= newLeader.GetCurrentTerm() {
		t.Fatal(errors.New("new leader should have higher term"))
	}

	// add a new log entry to the old leader; should NOT be replicated
	leader.leaderMutex.Lock()
	logEntry := LogEntry{
		Index:  leader.getLastLogIndex() + 1,
		TermId: leader.GetCurrentTerm(),
		Type:   CommandType_NOOP,
		Data:   []byte{1, 2, 3, 4},
	}
	leader.appendLogEntry(logEntry)
	leader.leaderMutex.Unlock()

	// add a new log entry to the new leader; SHOULD be replicated
	newLeader.leaderMutex.Lock()
	logEntry = LogEntry{
		Index:  newLeader.getLastLogIndex() + 1,
		TermId: newLeader.GetCurrentTerm(),
		Type:   CommandType_NOOP,
		Data:   []byte{5, 6, 7, 8},
	}
	newLeader.appendLogEntry(logEntry)
	newLeader.leaderMutex.Unlock()

	time.Sleep(time.Second * WAIT_PERIOD)

	// rejoin the cluster
	for _, follower := range followers[1:] {
		follower.NetworkPolicy.RegisterPolicy(*follower.GetRemoteSelf(), *ff.GetRemoteSelf(), true)
		follower.NetworkPolicy.RegisterPolicy(*follower.GetRemoteSelf(), *leader.GetRemoteSelf(), true)

		ff.NetworkPolicy.RegisterPolicy(*ff.GetRemoteSelf(), *follower.GetRemoteSelf(), true)
		leader.NetworkPolicy.RegisterPolicy(*leader.GetRemoteSelf(), *follower.GetRemoteSelf(), true)
	}
	//Out.Printf("ABOUT TO STABILIZE\n\n\n\n\n\n\n\n\n\n\n")
	// wait for larger cluster to stabilize
	time.Sleep(time.Second * WAIT_PERIOD)

	if newLeader.State != LEADER_STATE {
		t.Errorf("New leader should still be leader when old leader rejoins, but in %v state", newLeader.State)
	}

	if leader.State != FOLLOWER_STATE {
		t.Errorf("Old leader should fall back to the follower state after rejoining (was in %v state)", leader.State)
	}

	if !logsMatch(newLeader, cluster) {
		t.Errorf("logs incorrect")
	}
}
