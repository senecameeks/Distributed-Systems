package raft

import (
	//"fmt"
	"fmt"
	"github.com/brown-csci1380/mkohn-smeeks-s19/raft/hashmachine"
	"golang.org/x/net/context"
	"testing"
	"time"
)

// Example test making sure clusters can handle multiple clients
func TestMultipleClientInteraction_Leader(t *testing.T) {
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

	//// Second Client

	reply2, _ := leader.RegisterClientCaller(context.Background(), &RegisterClientRequest{})

	if reply2.Status != ClientStatus_OK {
		t.Fatal("Could not register client 2")
	}

	client2id := reply2.ClientId

	// Make sure further request is correct processed
	Client2Req := ClientRequest{
		ClientId:        client2id,
		SequenceNum:     1,
		StateMachineCmd: hashmachine.HASH_CHAIN_ADD,
		Data:            []byte{},
	}
	client2Result, _ := leader.ClientRequestCaller(context.Background(), &Client2Req)
	if client2Result.Status != ClientStatus_OK {
		fmt.Printf("%v", client2Result.Status)
		t.Fatal("Leader failed to commit a client 2 request")
	}

}
