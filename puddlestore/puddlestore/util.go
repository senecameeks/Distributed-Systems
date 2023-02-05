package puddlestore

import (
	"errors"
	"fmt"
	raftCli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/client"
	kv "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/kvstatemachine"
	"github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/raft"
	tapCli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/client"
	"github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/tapestry"
	"github.com/samuel/go-zookeeper/zk"
	"math/rand"
	"strings"
	"time"
)

var myRaftCluster []*raft.RaftNode

func setRaftCluster(rc []*raft.RaftNode) {
	myRaftCluster = rc
}

func Cleanup() {
	cleanupCluster(myRaftCluster)
}

// contains helpful functions
// acquires address of random node in children at path
func GetNodeAddress(conn *zk.Conn, path string) (addr string) {
	seed := int64(time.Now().Nanosecond())
	rand.Seed(seed)

	children, _, err := conn.Children(path)
	handleError(err)

	if len(children) == 0 {
		err = errors.New("No children in path" + path)
		handleError(err)
	}

	index := rand.Intn(len(children))
	startOfAddr := strings.Index(children[index], "/")

	return children[index][startOfAddr+1:]
}

// traversing incrementally the given path and checks that everything is okay
func traversePath(absPath string, tapClient *tapCli.Client, raftClient *raftCli.Client) (inode *Inode, found bool) {
	if absPath == "/" {
		//fmt.Printf("abs Path is /")
		toReturnInode := getInodeFromVguid(tapestry.Hash("/").String(), tapClient)
		return toReturnInode, true
	}
	//fmt.Printf("ABSOLUTE PATH: %v\n", absPath)
	names := strings.Split(absPath, "/")
	//printChildren(names)
	//fmt.Printf("traversing path and abs path is " + absPath)
	//fmt.Printf("\n\n\n\n")
	destInode := getInodeFromPath(absPath, raftClient, tapClient)
	if destInode == nil {
		//fmt.Println("returning nil up in this bih")
		return nil, false
	}
	if len(names) == 2 {
		return destInode, true
	}
	destVguid := getVguidFromAguid(destInode.AGUID, raftClient)
	currInode := getInodeFromPath("/", raftClient, tapClient)
	currVguid := getVguidFromAguid(currInode.AGUID, raftClient)

	invalid_path := false
	full_curr_path := "/"
	count := 1
	for currVguid != destVguid && !invalid_path {
		//fmt.Printf("name[0]" + names[0])
		//fmt.Printf("length of names %d", len(names))
		//fmt.Printf("\n\n\n\n")

		if len(names)-1 == count {
			full_curr_path = full_curr_path + names[count]
		} else {
			full_curr_path = full_curr_path + names[count] + "/"
		}
		count++
		//fmt.Printf("full cur path " + full_curr_path + "\n")

		nextInode := getInodeFromPath(full_curr_path, raftClient, tapClient)
		//fmt.Printf("nextInode: %v\n", nextInode)
		if nextInode == nil || nextInode.Type == FILE {
			return nil, false
		}
		nextVguid := getVguidFromAguid(nextInode.AGUID, raftClient)
		//fmt.Printf("full cur path vguid " + nextVguid)
		indirectVguid := getVguidFromAguid(currInode.Indirect, raftClient)
		indirectBlock := getIndirectBlockFromVguid(indirectVguid, tapClient)
		//fmt.Printf("\n\n\n\n")
		found = false
		for _, aguid := range indirectBlock.DataBlockList {
			//fmt.Println("i work")
			if getVguidFromAguid(aguid, raftClient) == nextVguid {
				found = true
				break
			}
			//fmt.Println("i still work")
		}

		currVguid = nextVguid
		//fmt.Printf("currVguid: %v\n destVguid: %v\n", currVguid, destVguid)
		if currVguid == destVguid {
			found = true
		}

		if !found {
			invalid_path = true
		}

	}
	//fmt.Printf("out of loop, fuck w me %v\n", invalid_path)
	if invalid_path {
		return nil, false
	} else {
		return destInode, true
	}
}

// get Inode to begin traversal
func getInodeFromPath(path string, raftClient *raftCli.Client, tapClient *tapCli.Client) (inode *Inode) {
	aguid := tapestry.Hash(path).String()
	reply, _ := raftClient.SendRequestReceiveReply(kv.KV_STATE_GET, []byte(aguid))
	inodeVguid := reply.GetResponse()
	data, err := tapClient.Get(inodeVguid)
	//fmt.Printf("end my life u wont %v\n", path)
	handleError(err)
	if data == nil || len(data) == 0 {
		//fmt.Println("i will pay u to end my life")
		return nil
	}

	i, err := GetInodeFromBytes(data)
	handleError(err)

	return i
}

// return IndirectBlock given guid
func getIndirectBlock(guid string, raftClient *raftCli.Client, tapClient *tapCli.Client) (indirectBlock *IndirectBlock) {
	vguid := getVguidFromAguid(guid, raftClient)
	indirectBlock = getIndirectBlockFromVguid(vguid, tapClient)
	return indirectBlock
}

func getDataBlock(guid string, raftClient *raftCli.Client, tapClient *tapCli.Client) (dataBlock *DataBlock) {
	vguid := getVguidFromAguid(guid, raftClient)
	dataBlock = getDataBlockFromVguid(vguid, tapClient)
	return dataBlock
}

func updateVguidForAguid(vguid string, aguid string, raftClient *raftCli.Client) {
	err := raftClient.SendRequest(kv.KV_STATE_ADD, []byte(aguid+"."+vguid))
	handleError(err)
}

// stores the inode in tapestry w vguid as key
func putInodeIntoTap(inode *Inode, tapClient *tapCli.Client, raftClient *raftCli.Client) {
	vguid := makeVguidForAguid(inode.AGUID)
	updateVguidForAguid(vguid, inode.AGUID, raftClient)
	bytes, err := GetInodeBytes(inode)
	handleError(err)
	tapClient.Store(vguid, bytes)
}

// stores the indirect block in tapestry w vguid as key
func putIndirectIntoTap(ib *IndirectBlock, tapClient *tapCli.Client, raftClient *raftCli.Client) {
	vguid := makeVguidForAguid(ib.AGUID)
	updateVguidForAguid(vguid, ib.AGUID, raftClient)
	bytes, err := GetIndirectBytes(ib)
	handleError(err)
	tapClient.Store(vguid, bytes)
}

// stores the data block in tapestry w vguid as key
func putDataIntoTap(db *DataBlock, tapClient *tapCli.Client, raftClient *raftCli.Client) {
	vguid := makeVguidForAguid(db.AGUID)
	updateVguidForAguid(vguid, db.AGUID, raftClient)
	bytes, err := GetDataBytes(db)
	handleError(err)
	tapClient.Store(vguid, bytes)
}

// runs aguid through raft cluster to get vguid
func getVguidFromAguid(aguid string, raftClient *raftCli.Client) string {
	reply, err := raftClient.SendRequestReceiveReply(kv.KV_STATE_GET, []byte(aguid))
	handleError(err)

	return reply.GetResponse()
}

// runs vguid through tapestry to get inode reference
func getInodeFromVguid(vguid string, tapestryClient *tapCli.Client) *Inode {
	bytes, err := tapestryClient.Get(vguid)
	handleError(err)

	inode, err := GetInodeFromBytes(bytes)
	handleError(err)

	return inode
}

// runs vguid through tapestry to get indirect block reference
func getIndirectBlockFromVguid(vguid string, tapestryClient *tapCli.Client) *IndirectBlock {
	bytes, err := tapestryClient.Get(vguid)
	handleError(err)

	ib, err := GetIndirectFromBytes(bytes)
	handleError(err)

	return ib
}

// runs vguid through tapestry to get data block reference
func getDataBlockFromVguid(vguid string, tapestryClient *tapCli.Client) *DataBlock {
	bytes, err := tapestryClient.Get(vguid)
	handleError(err)

	db, err := GetDataFromBytes(bytes)
	handleError(err)

	return db
}

// makes a vguid using aguid and a random sequence of 32 characters
func makeVguidForAguid(aguid string) string {
	vguid := aguid

	// add 32 random characters
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().Unix())
	// 32 placeholder X's
	randSeq := ""

	for i := 0; i < 32; i++ {
		randSeq += string(chars[rand.Intn(len(chars))])
	}

	vguid += randSeq
	return vguid
}

func getParentPath(absPath string) string {
	fmt.Printf("abs path %v", absPath)
	names := strings.Split(absPath, "/")
	parentPath := "/"
	for i := 0; i < len(names)-1; i++ {
		parentPath += names[i] + "/"
	}

	return parentPath
}

func contains(arr []string, name string, tapClient *tapCli.Client, raftClient *raftCli.Client) bool {
	found := false
	for _, item_name := range arr {
		if item_name == name {
			found = true
		}
	}
	return found
}

// handles errors, panics if found
func handleError(err error) {
	if err != nil {
		//cleanupCluster(myRaftCluster)
		fmt.Println(err)
	}
}

func cleanupCluster(nodes []*raft.RaftNode) {
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		//node.server.Stop()
		go func(node *raft.RaftNode) {
			node.GracefulExit()
			node.RemoveLogs()
		}(node)
	}
	time.Sleep(5 * time.Second)
}

func printChildren(children []string) {
	for _, path := range children {
		fmt.Println(path)
	}
}
