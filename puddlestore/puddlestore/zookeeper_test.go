package puddlestore

import (
	"bytes"
	"fmt"
	"github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/raft"
	tc "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/client"
	"github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/tapestry"
	"github.com/samuel/go-zookeeper/zk"
	"os"
	"strings"
	"testing"
	"time"
)

// tests if can connect to zookeeper server, then start up a raft cluster,
//  store the raft nodes appropriately, then access them again
func TestRaftConnection(t *testing.T) {
	t.Log("testing raft connection")
	// check server state w ruok
	ok := zk.FLWRuok([]string{"127.0.0.1:2181"}, time.Second*2)
	if !ok[0] {
		t.Fatal("ZooKeeper server not ok")
	}

	// establish a connection to zookeeper server w address 127.0.0.1 that times out after 3s
	c, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second*3)
	if err != nil {
		t.Fatal(err)
	}

	// remove any existing raftlogs
	removeRaftLogs()

	// launch a raft cluster of size 3
	raftCluster, err := raft.CreateDefinedLocalCluster(nil, []int{3001, 3002, 3003})
	if err != nil {
		t.Fatal(err)
	}

	addrList := make([]string, 3)
	for i := 0; i < len(raftCluster); i++ {
		addrList[i] = raftCluster[i].GetRemoteSelf().GetAddr()
	}

	// start up a tapestry network of size 5
	tapNode, _ := tapestry.Start(0, "")
	indexOfAddr := strings.Index(tapNode.String(), "at ")
	addr := tapNode.String()[indexOfAddr+3:]
	tapestry.Start(0, addr)
	tapestry.Start(0, addr)
	tapestry.Start(0, addr)
	tapestry.Start(0, addr)

	// wait 2s for things to start
	time.Sleep(time.Second * 2)

	path := "/Raft"
	// check if Raft folder already exists
	exists, _, err := c.Exists(path)
	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		// create the Raft folder inside of zookeeper
		_, err = c.Create(path, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
		if err != nil {
			t.Fatal(err)
		}

	}

	// now try and retrieve the raft nodes from zookeeper
	//  children contains paths which should be in form "/Raft/{addr}"
	children, _, err := c.Children(path)
	if err != nil {
		t.Fatal(err)
	}

	// check format of paths
	printChildren(children)

	//retrievedNodes := make([]*raft.RaftNode, len(children))
	versions := make([]int32, len(children))

	for i, childPath := range children {
		// data should be address of node
		data, stat, err := c.Get("/Raft/" + childPath)
		t.Log(childPath)
		if err != nil {
			t.Fatal(err)
		}

		contains := false
		for _, addr := range addrList {
			if string(data) == addr {
				contains = true
			}
		}

		// compare retrieved node w original node
		if !contains {
			fmt.Printf("Retrieved node %v nonexistent in original nodes\n", childPath)
			t.Fatal("Node mismatch failure")
		}

		versions[i] = stat.Version
	}

	// post-test cleanup
	cleanupCluster(raftCluster)
}

// tests if can connect to zookeeper server, then start up a tapestry network,
//  store the tapestry nodes appropriately, then access them again
func TestTapestryInit(t *testing.T) {
	t.Log("testing tapestry initialization")
	// check server state w ruok
	ok := zk.FLWRuok([]string{"127.0.0.1:2181"}, time.Second*2)
	if !ok[0] {
		t.Fatal("ZooKeeper server not ok")
	}

	// establish a connection to zookeeper server w address 127.0.0.1 that times out after 3s
	c, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second*3)
	if err != nil {
		t.Fatal(err)
	}

	// remove any existing raftlogs
	removeRaftLogs()

	// launch a raft cluster of size 3
	raftCluster, err := raft.CreateDefinedLocalCluster(nil, []int{3001, 3002, 3003})
	if err != nil {
		t.Fatal(err)
	}

	// start up a tapestry network of size 5
	/*
		tapNode, _ := tapestry.Start(0, "")
		indexOfAddr := strings.Index(tapNode.String(), "at ")
		addr := tapNode.String()[indexOfAddr+3:]
		tapestry.Start(0, addr)
		tapestry.Start(0, addr)
		tapestry.Start(0, addr)
		tapestry.Start(0, addr)*/
	// start up a tapestry network of size 5
	ts, err := tapestry.MakeRandomTapestries(int64(1), 5)
	if err != nil {
		t.Fatal(err)
	}

	ts[1].Store("look at this lad", []byte("an absolute unit"))

	// wait 2s for things to start
	time.Sleep(time.Second * 2)

	path := "/Tapestry"
	// check if Tapestry folder already exists
	exists, _, err := c.Exists(path)
	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		// create the Tapestry folder inside of zookeeper
		_, err = c.Create(path, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
		if err != nil {
			t.Fatal(err)
		}

	}

	// now try and retrieve the raft nodes from zookeeper
	//  children contains paths which should be in form "/Tapestry/{addr}"
	children, _, err := c.Children(path)
	if err != nil {
		t.Fatal(err)
	}

	// check format of paths
	printChildren(children)

	//retrievedNodes := make([]*raft.RaftNode, len(children))
	versions := make([]int32, len(children))

	for i, childPath := range children {
		data, stat, err := c.Get("/Tapestry/" + childPath)
		if err != nil {
			t.Fatal(err)
		}

		// test whether shit is stored ??
		tClient := tc.Connect(string(data))
		result, _ := tClient.Get("look at this lad")          //Store a KV pair and try to fetch it
		if !bytes.Equal(result, []byte("an absolute unit")) { //Ensure we correctly get our KV
			fmt.Printf("failure node address: %v \n", data)
			t.Fatal("Get failed")
		}

		versions[i] = stat.Version
	}

	// post-test cleanup
	cleanupCluster(raftCluster)

}

// write another tapestry test in which you actually store and then retrieve stuff

/*func cleanupCluster(nodes []*raft.RaftNode) {
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		//node.server.Stop()
		go func(node *raft.RaftNode) {
			node.GracefulExit()
			node.RemoveLogs()
		}(node)
	}
	time.Sleep(5 * time.Second)
}*/

func removeRaftLogs() {
	os.RemoveAll("raftlogs/")
}

/*
func printChildren(children []string) {
	for _, path := range children {
		fmt.Println(path)
	}
} */
