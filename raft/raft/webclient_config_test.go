package raft

import (
	"fmt"
	"testing"
	"time"
)

//TestClusterForWeb
func TestClusterForWeb(t *testing.T) {
	ports := []int{6000, 9091, 9092, 9093, 9094}
	fmt.Print(ports)
	t.Log(ports)
	// suppressLoggers()
	cluster, err := createTestCluster(ports)
	defer cleanupCluster(cluster)
	if err != nil {
		t.Fatal(err)
	}

	// wait for a leader to be elected
	time.Sleep(time.Second * WAIT_PERIOD)

	followers, candidates, leaders := 0, 0, 0
	for i := 0; i < 5; i++ {
		node := cluster[i]
		fmt.Printf("NODE %v\n", node.port)
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
	time.Sleep(time.Second * 600)
}
