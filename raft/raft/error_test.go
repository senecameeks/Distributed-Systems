package raft

import (
	//"fmt"
	"testing"
	"time"
)

func TestNetworkPolicyDenied(t *testing.T) {
	suppressLoggers()
	config := DefaultConfig()
	config.ClusterSize = 5

	cluster, err := CreateLocalCluster(config)
	defer cleanupCluster(cluster)
	if err != nil {
		t.Fatal("error with createTestCluster")
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
		t.Fatal("leader should remain leader even when partitioned")
	}

	// now try and send messages between two nodes across a partition. Should fail

	appendRequest := &AppendEntriesRequest{Term: leader.GetCurrentTerm(), Leader: leader.GetRemoteSelf()}
	_, err = newLeader.GetRemoteSelf().AppendEntriesRPC(leader, appendRequest)
	if err != ErrorNetworkPolicyDenied {
		t.Fatal("leaders are in different clusters appendEntriesRPC should return error")
	}

	voteRequest := &RequestVoteRequest{Term: leader.GetCurrentTerm(), Candidate: leader.GetRemoteSelf()}
	_, err = newLeader.GetRemoteSelf().RequestVoteRPC(leader, voteRequest)
	if err != ErrorNetworkPolicyDenied {
		t.Fatal("leaders are in different clusters RequestVoteRPC should return error")
	}

}

// func TestNilRequests(t *testing.T) {
// 	suppressLoggers()
//
// 	cluster, err := createTestCluster([]int{5001, 5002, 5003, 5004, 5005})
// 	defer cleanupCluster(cluster)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// wait for a leader to be elected
// 	time.Sleep(time.Second * WAIT_PERIOD)
// 	leader, err := findLeader(cluster)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// send nil requests
// 	_, err = leader.GetRemoteSelf().AppendEntriesRPC(leader, nil)
// 	if err != nil {
// 		t.Fatal("Should throw error because request is nil")
// 	}
//
// 	_, err = leader.GetRemoteSelf().RequestVoteRPC(leader, nil)
// 	if err != nil {
// 		t.Fatal("Should throw error because request is nil")
// 	}
//
// }
