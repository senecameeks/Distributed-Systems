package raft

import (
	//"fmt"
	"github.com/brown-csci1380/mkohn-smeeks-s19/raft/hashmachine"
	"golang.org/x/net/context"
	"testing"
	"time"
)

// Example test making sure leaders can register the client and process the request from clients properly
func TestClientInteraction_Leader(t *testing.T) {
	suppressLoggers()
	config := DefaultConfig()
	cluster, _ := CreateLocalCluster(config)
	defer cleanupCluster(cluster)

	time.Sleep(2 * time.Second)

	leader, err := findLeader(cluster)
	if err != nil {
		t.Fatal(err)
	}

	// First make sure we can register a client correctly
	reply, _ := leader.RegisterClientCaller(context.Background(), &RegisterClientRequest{})

	if reply.Status != ClientStatus_OK {
		t.Fatal("Could not register client")
	}

	clientid := reply.ClientId

	// Hash initialization request
	initReq := ClientRequest{
		ClientId:        clientid,
		SequenceNum:     1,
		StateMachineCmd: hashmachine.HASH_CHAIN_INIT,
		Data:            []byte("hello"),
	}
	clientResult, _ := leader.ClientRequestCaller(context.Background(), &initReq)
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
}

// Example test making sure the follower would reject the registration and requests from clients with correct messages
// The test on candidates can be similar with these sample tests
func TestClientInteraction_Follower(t *testing.T) {
	suppressLoggers()
	config := DefaultConfig()
	// set the ElectionTimeout long enough to keep nodes in the state of follower
	config.ElectionTimeout = 60 * time.Second
	config.ClusterSize = 3
	config.HeartbeatTimeout = 500 * time.Millisecond
	cluster, err := CreateLocalCluster(config)
	defer cleanupCluster(cluster)

	time.Sleep(2 * time.Second)

	if err != nil {
		t.Fatal(err)
	}

	// make sure the client get the correct response while registering itself with a follower
	reply, _ := cluster[0].RegisterClientCaller(context.Background(), &RegisterClientRequest{})
	if reply.Status != ClientStatus_NOT_LEADER && reply.Status != ClientStatus_ELECTION_IN_PROGRESS {
		t.Error(reply.Status)
		t.Fatal("Wrong response when registering a client to a follower")
	}

	// make sure the client get the correct response while sending a request to a follower
	req := ClientRequest{
		ClientId:        1,
		SequenceNum:     1,
		StateMachineCmd: hashmachine.HASH_CHAIN_INIT,
		Data:            []byte("hello"),
	}
	clientResult, _ := cluster[0].ClientRequestCaller(context.Background(), &req)
	if clientResult.Status != ClientStatus_NOT_LEADER && clientResult.Status != ClientStatus_ELECTION_IN_PROGRESS {
		t.Fatal("Wrong response when sending a client request to a follower")
	}
}

// Example test making sure the candidate would reject the registration and requests from clients with correct messages
// The test on candidates can be similar with these sample tests
func TestClientInteraction_Candidate(t *testing.T) {
	suppressLoggers()
	config := DefaultConfig()

	// set the ElectionTimeout short enough to keep nodes in the state of candidate
	err := CheckConfig(config)
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
	oldLeader.NetworkPolicy.PauseWorld(true)

	// // wait for new leader to be elected
	// time.Sleep(time.Second * WAIT_PERIOD)
	//
	// newLeader, err := findLeader(cluster)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// newLeader.NetworkPolicy.PauseWorld(true)

	candidateIndex := 0
	for i := 0; i < config.ClusterSize; i++ {
		if cluster[i].State == CANDIDATE_STATE {
			Out.Println("found candidate saving index")
			candidateIndex = i
		} else if cluster[i].State != LEADER_STATE {
			Out.Println("found candidate saving index")
			candidateIndex = i
		}
	}

	// make sure the client get the correct response while registering itself with a follower
	reply, _ := cluster[candidateIndex].RegisterClientCaller(context.Background(), &RegisterClientRequest{})
	if reply.Status != ClientStatus_NOT_LEADER && reply.Status != ClientStatus_ELECTION_IN_PROGRESS {
		t.Error(reply.Status)
		t.Fatal("Wrong response when registering a client to a candidate")
	} else {
		return
	}

	// // unpause leader
	oldLeader.NetworkPolicy.PauseWorld(true)
	// newLeader.NetworkPolicy.PauseWorld(false)
	// newLeader.RegisterClientCaller(context.Background(), &RegisterClientRequest{})
	//
	// make sure the client get the correct response while sending a request to a follower
	req := ClientRequest{
		ClientId:        1,
		SequenceNum:     1,
		StateMachineCmd: hashmachine.HASH_CHAIN_INIT,
		Data:            []byte("hello"),
	}
	clientResult, _ := cluster[candidateIndex].ClientRequestCaller(context.Background(), &req)
	if clientResult.Status != ClientStatus_NOT_LEADER && clientResult.Status != ClientStatus_ELECTION_IN_PROGRESS {
		t.Fatal("Wrong response when sending a client request to a follower")
	}

	// gracefully exit all NodeState
	for j := 0; j < config.ClusterSize; j++ {
		cluster[j].gracefulExit <- true
	}
}
