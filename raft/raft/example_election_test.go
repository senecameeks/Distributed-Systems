package raft

import (
	//"fmt"
	"testing"
	"time"
)

// Tests that nodes can successfully join a cluster and elect a leader.
func TestInit(t *testing.T) {
	suppressLoggers()
	config := DefaultConfig()
	cluster, err := CreateLocalCluster(config)
	defer cleanupCluster(cluster)
	if err != nil {
		t.Fatal(err)
	}

	// wait for a leader to be elected
	time.Sleep(time.Second * WAIT_PERIOD)

	followers, candidates, leaders := 0, 0, 0
	for i := 0; i < config.ClusterSize; i++ {
		node := cluster[i]
		switch node.State {
		case FOLLOWER_STATE:
			followers++
		case CANDIDATE_STATE:
			candidates++
		case LEADER_STATE:
			Out.Printf("found leader %x \n", node.GetRemoteSelf().GetId())
			leaders++
		}
	}

	if followers != config.ClusterSize-1 {
		t.Errorf("follower count mismatch, expected %v, got %v", config.ClusterSize-1, followers)
	}

	if candidates != 0 {
		t.Errorf("candidate count mismatch, expected %v, got %v", 0, candidates)
	}

	if leaders != 1 {
		t.Errorf("leader count mismatch, expected %v, got %v", 1, leaders)
	}
}

// Tests that if a leader is partitioned from its followers, a
// new leader is elected.
func TestNewElection(t *testing.T) {
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
	oldLeader, err := findLeader(cluster)
	if err != nil {
		t.Fatal(err)
	}

	// partition leader, triggering election
	oldTerm := oldLeader.GetCurrentTerm()
	oldLeader.NetworkPolicy.PauseWorld(true)

	// wait for new leader to be elected
	time.Sleep(time.Second * WAIT_PERIOD)

	// unpause old leader and wait for it to become a follower
	oldLeader.NetworkPolicy.PauseWorld(false)
	time.Sleep(time.Second * WAIT_PERIOD)

	newLeader, err := findLeader(cluster)
	if err != nil {
		t.Fatal(err)
	}

	if oldLeader.Id == newLeader.Id {
		t.Errorf("leader did not change")
	}

	if newLeader.GetCurrentTerm() == oldTerm {
		t.Errorf("term did not change")
	}
}

// Send requestVote to node with Higher term. Candidate should
func TestReqVoteToHigherTerm(t *testing.T) {
	suppressLoggers()
	config := DefaultConfig()
	config.ClusterSize = 5

	cluster, err := CreateLocalCluster(config)
	cluster[2].setCurrentTerm(3)
	defer cleanupCluster(cluster)
	if err != nil {
		t.Fatal(err)
	}
	// election will begin and try to do reqVote to cluster[2] with higher term
	time.Sleep(time.Second * WAIT_PERIOD)

	// check if cluster[0] is still a follower
	if cluster[0].State != FOLLOWER_STATE {
		t.Fatal("node that sent request should still be follower")
	}

}
