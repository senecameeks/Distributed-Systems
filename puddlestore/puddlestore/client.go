/*
 *  Brown University, CS138, Spring 2019
 *
 *  Purpose: Allows third-party clients to connect to a Tapestry node (such as
 *  a web app, mobile app, or CLI that you write), and put and get objects.
 */

package puddlestore

import (
	//"fmt"
	//p "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore"
	raftCli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/client"
	kv "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/kvstatemachine"
	"github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/raft"
	tapCli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/client"
	"github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/tapestry"
	//"github.com/samuel/go-zookeeper/zk"
)

// checks if root directory exists, creates one if not
func CreateRootDir(tc *tapCli.Client, rc *raftCli.Client) string {
	// hash "/" to get AGUID
	rootAguid := tapestry.Hash("/").String()

	// PUT THE AGUID THROUGH THE MAP MACHINE TO GET THE VERSION
	rootVguidReply, err := rc.SendRequestReceiveReply(kv.KV_STATE_GET, []byte(rootAguid))
	handleError(err)

	rootVguid := rootVguidReply.GetResponse()

	var rootInode *Inode

	if rootVguid == "" {
		// if does not exist
		rootInode = CreateDirectoryInode("/", "/")
		// construct the indirect block for the root
		rootIndirect := CreateIndirectBlock()
		rootInode.Indirect = rootIndirect.AGUID
		putIndirectIntoTap(rootIndirect, tc, rc)
		// assign the initial vguid
		rootVguid = rootAguid //does not need to be versioned? p.makeVguidForAguid(rootAguid)
		updateVguidForAguid(rootVguid, rootAguid, rc)
		// put it into map machine
		_, err := rc.SendRequestReceiveReply(kv.KV_STATE_ADD, []byte(rootVguid+"."+rootAguid))
		handleError(err)
		// put it into tapestry
		data, err := GetInodeBytes(rootInode)
		handleError(err)

		err = tc.Store(rootVguid, data)
		handleError(err)

	} else {
		// if root directory does exist
		// retrieve from tapestry
		bytes, err := tc.Get(rootVguid)
		handleError(err)

		// convert from bytes to rootInode
		rootInode, err = GetInodeFromBytes(bytes)
		handleError(err)
	}

	return rootInode.AGUID
}

// starts a defined local raft cluster and tapestry network
func StartRaftAndTapestry(raftSize int, tapSize int) {
	// launch a raft cluster of size raftSize
	ports := make([]int, raftSize)
	for i := 0; i < raftSize; i++ {
		ports[i] = 3001 + i
	}

	rc, err := raft.CreateDefinedLocalCluster(nil, ports)
	if err != nil {
		handleError(err)
	}

	setRaftCluster(rc)

	_, err = tapestry.MakeRandomTapestries(int64(1), tapSize)
	if err != nil {
		handleError(err)
	}
}
