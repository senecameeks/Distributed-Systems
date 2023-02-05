package raft

import (
	"errors"
	"github.com/brown-csci1380/mkohn-smeeks-s19/raft/hashmachine"
	"golang.org/x/net/context"
	"testing"
	"time"
)

// test for correctness of sending client requests and registrations to followers
//  with leaders, as well as making sure cached replies for duplicate requests work
func TestSpecialClientRequests(t *testing.T) {
	suppressLoggers()
	config := DefaultConfig()
	cluster, _ := CreateLocalCluster(config)
	defer cleanupCluster(cluster)

	time.Sleep(2 * time.Second)

	leader, err := findLeader(cluster)
	if err != nil {
		t.Fatal(err)
	}

	follower, err := findAFollower(cluster)
	for err != nil {
		follower, err = findAFollower(cluster)
	}

	// First make sure we can reroute registering a client correctly
	reply, _ := follower.RegisterClientCaller(context.Background(), &RegisterClientRequest{})

	if reply.Status != ClientStatus_NOT_LEADER {
		t.Fatal("Did not send correct register client reply status")
	}

	if reply.LeaderHint.Id != leader.GetRemoteSelf().GetId() {
		t.Fatal("Leader hint is not correct")
	}

	reply, _ = leader.RegisterClientCaller(context.Background(), &RegisterClientRequest{})

	clientid := reply.ClientId

	// Hash initialization request
	initReq := ClientRequest{
		ClientId:        clientid,
		SequenceNum:     1,
		StateMachineCmd: hashmachine.HASH_CHAIN_INIT,
		Data:            []byte("hello"),
	}

	clientResult, _ := follower.ClientRequestCaller(context.Background(), &initReq)
	if clientResult.Status != ClientStatus_NOT_LEADER {
		t.Fatal("Did not send correct client request reply status")
	}

	if clientResult.LeaderHint.Id != leader.GetRemoteSelf().GetId() {
		t.Fatal("Leader hint is not correct")
	}

	clientResult, _ = leader.ClientRequestCaller(context.Background(), &initReq)
	if clientResult.Status != ClientStatus_OK {
		t.Fatal("Leader failed to commit a client request")
	}

	// Make sure further request is correct processed
	ClientReq := ClientRequest{
		ClientId:        clientid,
		SequenceNum:     2,
		StateMachineCmd: hashmachine.HASH_CHAIN_ADD,
		Data:            []byte{},
	}
	clientResult, _ = leader.ClientRequestCaller(context.Background(), &ClientReq)
	if clientResult.Status != ClientStatus_OK {
		t.Fatal("Leader failed to commit a client request")
	}

	clientResult, _ = leader.ClientRequestCaller(context.Background(), &initReq)
	if clientResult.Status != ClientStatus_OK {
		t.Fatal("Leader failed to retrieve a (should be) cached client request")
	}
}

// test for
func TestPartitionPossibilities(t *testing.T) {
	suppressLoggers()
	config := DefaultConfig()
	config.ClusterSize = 5

	cluster, err := CreateLocalCluster(config)
	defer cleanupCluster(cluster)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(WAIT_PERIOD * time.Second)

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

	// partition into 2 clusters: one with first follower; other with remaining 3 followers and leader
	ff := followers[0]

	leader.NetworkPolicy.RegisterPolicy(*leader.GetRemoteSelf(), *ff.GetRemoteSelf(), false)
	ff.NetworkPolicy.RegisterPolicy(*ff.GetRemoteSelf(), *leader.GetRemoteSelf(), false)

	for _, follower := range followers[1:] {
		follower.NetworkPolicy.RegisterPolicy(*follower.GetRemoteSelf(), *ff.GetRemoteSelf(), false)
		ff.NetworkPolicy.RegisterPolicy(*ff.GetRemoteSelf(), *follower.GetRemoteSelf(), false)
	}

	// allow the lonely follower to become a candidate
	time.Sleep(time.Second * WAIT_PERIOD)

	if ff.State != CANDIDATE_STATE {
		t.Fatal(errors.New("follower not receiving heartbeats has not become candidate"))
	}

	// now partition away the leader
	for _, follower := range followers[1:] {
		follower.NetworkPolicy.RegisterPolicy(*follower.GetRemoteSelf(), *leader.GetRemoteSelf(), false)
		leader.NetworkPolicy.RegisterPolicy(*leader.GetRemoteSelf(), *follower.GetRemoteSelf(), false)
	}

	time.Sleep(time.Second * WAIT_PERIOD)

	// check that old leader, which is cut off from new leader, still thinks it's leader
	if leader.State != LEADER_STATE {
		t.Fatal(errors.New("leader should remain leader even when partitioned"))
	}

	// check if followers elected a new leader
	newLeader, err := findLeader(followers)
	if err != nil {
		t.Fatal(err)
	}

	// now put the leader back in with the 3 node cluster
	for _, follower := range followers[1:] {
		follower.NetworkPolicy.RegisterPolicy(*follower.GetRemoteSelf(), *leader.GetRemoteSelf(), true)
		leader.NetworkPolicy.RegisterPolicy(*leader.GetRemoteSelf(), *follower.GetRemoteSelf(), true)
	}

	time.Sleep(time.Second * WAIT_PERIOD)

	newLeader, err = findLeader(followers)
	if err != nil {
		t.Fatal(err)
	}

	if newLeader.State != LEADER_STATE {
		t.Errorf("New leader should still be leader when old leader rejoins, but in %v", newLeader.State)
	}

	if leader.State != FOLLOWER_STATE {
		t.Errorf("Old leader should fall back to the follower state after rejoining (was in %v)", leader.State)
	}

	// recombine all nodes
	leader.NetworkPolicy.RegisterPolicy(*leader.GetRemoteSelf(), *ff.GetRemoteSelf(), true)
	ff.NetworkPolicy.RegisterPolicy(*ff.GetRemoteSelf(), *leader.GetRemoteSelf(), true)

	for _, follower := range followers[1:] {
		follower.NetworkPolicy.RegisterPolicy(*follower.GetRemoteSelf(), *ff.GetRemoteSelf(), true)
		ff.NetworkPolicy.RegisterPolicy(*ff.GetRemoteSelf(), *follower.GetRemoteSelf(), true)
	}

	time.Sleep(time.Second * WAIT_PERIOD)

	if ff.State != FOLLOWER_STATE {
		t.Errorf("Old candidate should fall back to the follower state due to not having complete log (was in %v)", leader.State)
	}

}
