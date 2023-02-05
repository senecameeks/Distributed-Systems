package puddlestore

import (
	//"fmt"
	"bytes"
	raftCli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/client"
	kv "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/kvstatemachine"
	"github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/raft"
	tapCli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/client"
	"github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/tapestry"
	//cli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/client"
	"errors"
	zk "github.com/samuel/go-zookeeper/zk"
	"testing"
	"time"
)

func TestCommands(t *testing.T) {
	t.Log("testing commands")
	// check server state w ruok
	ok := zk.FLWRuok([]string{"127.0.0.1:2181"}, time.Second*2)

	if !ok[0] {
		t.Errorf("ZooKeeper server not ok")
		t.Fatal(errors.New("ZooKeeper server not ok"))
	}

	// establish a connection to zookeeper server w address 127.0.0.1 at port 2181
	//  that times out after 3s
	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second*3)
	if err != nil {
		t.Errorf("didn't connect to zk")
		t.Fatal(errors.New("didn't connect to zk"))
		handleError(err)
	}

	// remove any existing raftlogs
	removeRaftLogs()

	// launch a raft cluster of size 3
	raftCluster, err := raft.CreateDefinedLocalCluster(nil, []int{3001, 3002, 3003})
	if err != nil {
		t.Fatal(err)
	}

	_, err = tapestry.MakeRandomTapestries(int64(1), 5)
	if err != nil {
		t.Fatal(err)
	}

	tapClient := tapCli.Connect(GetNodeAddress(conn, "/Tapestry"))

	raftClient, err := raftCli.Connect(GetNodeAddress(conn, "/Raft"))
	if err != nil {
		t.Errorf("didn't connect to raft")
		t.Fatal(errors.New("didn't connect to raft"))
		handleError(err)
	}

	// initialize kv state machine
	err = raftClient.SendRequest(kv.KV_STATE_INIT, []byte{})
	handleError(err)

	// create root directory
	rootAguid := CreateRootDir(tapClient, raftClient)
	t.Log("root aguid below")
	t.Log(rootAguid)

	// CREATE FILE IN ROOT
	//absPath := "/hello.txt"
	parentPath := "/"
	fileName := "hello.txt"
	t.Log("about to make file")
	_, err = Mkfileordir(parentPath, fileName, tapClient, raftClient, true)
	if err != nil {
		t.Log("failed to create file")
		t.Errorf("failed to create file %v", fileName)
	}

	items, err := Ls(parentPath, tapClient, raftClient)
	if err != nil {
		t.Errorf("error with ls command")
	}
	printChildren(items)
	if !contains(items, fileName, tapClient, raftClient) {
		t.Errorf("filename not in ls items")
	}

	// CREATE DIR IN ROOT
	//absPath := "/world"
	parentPath = "/"
	dirName := "world"
	_, err = Mkfileordir(parentPath, dirName, tapClient, raftClient, false)
	if err != nil {
		t.Errorf("failed to create directory %v", dirName)
	}

	items, err = Ls(parentPath, tapClient, raftClient)
	handleError(err)

	if !contains(items, dirName, tapClient, raftClient) {
		t.Errorf("Fail to create a file name %v", dirName)
	}

	// remove file foo.txt PROB NEED TO CHANGE TO MEET NEW STANDARDS
	err = Rm(parentPath, fileName, tapClient, raftClient)
	if err != nil {
		t.Errorf("error with rm file %v", fileName)
	}

	items, err = Ls(parentPath, tapClient, raftClient)
	if err != nil {
		t.Errorf("error with ls command")
	}

	//printChildren(items)
	// file/dir should no longer be in items
	if contains(items, fileName, tapClient, raftClient) || len(items) != 1 {
		t.Errorf("filename not in ls items")
	}

	// ########## TESTING WRITE AND READ ############### //
	parentPath = "/"
	fileName1 := "test.txt"
	_, err = Mkfileordir(parentPath, fileName1, tapClient, raftClient, true)

	// ########## 1.READING FROM AN EMPTY FILE ##########################
	rbuffer := Read("/test.txt", "test.txt", 0, make([]byte, 0), raftClient, tapClient)
	expectedBuffer := []byte("")
	if !bytes.Equal(expectedBuffer, rbuffer) {
		t.Errorf("Reading from an empty file expect to be %v but return %v", string(expectedBuffer), string(rbuffer))
	}

	// ########### 2.WRITING TO AN EMPTY FILE ############################
	err = Write("/", "test.txt", 0, []byte("abcdefg"), tapClient, raftClient)
	rbuffer = Read("/test.txt", "test.txt", 0, make([]byte, 0), raftClient, tapClient)
	expectedBuffer = []byte("abcdefg")
	if !bytes.Equal(expectedBuffer, rbuffer) {
		t.Errorf("Reading from an empty file expect to be %v but return %v", string(expectedBuffer), string(rbuffer))
	}

	// ########### 3.APPENDING TO A FILE ############################
	err = Write("/", "test.txt", 0, []byte("123"), tapClient, raftClient)
	rbuffer = Read("/test.txt", "test.txt", 0, make([]byte, 0), raftClient, tapClient)
	expectedBuffer = []byte("123abcdefg")
	if !bytes.Equal(expectedBuffer, rbuffer) {
		t.Errorf("Reading from an empty file expect to be %v but return %v", string(expectedBuffer), string(rbuffer))
	}

	cleanupCluster(raftCluster)
	removeRaftLogs()
}
