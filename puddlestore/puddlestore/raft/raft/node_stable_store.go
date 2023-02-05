package raft

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
)

// StableState contains parts of a Raft node's state that needs to be persisted
// between sessions, and is thus written to disk. This file provides setters and
// getters to read and update the stable state.
type StableState struct {
	// Latest term the server has seen (initialized to 0 on start, increases monotonically)
	CurrentTerm uint64

	// The candidate Id that received our vote in the current term (or "" if none)
	VotedFor string

	// A remote node representation of the current node
	RemoteSelf RemoteNode

	// List of all nodes (including ourselves in the cluster)
	NodeList []RemoteNode

	// Client reply cache, maps a client request cacheId to the response that was sent to them
	ClientReplyCache map[string]ClientReply

	// New stuff for snapshot
	// Snapshot index offset?
	Snapshot      StateSnapshot
	SnapshotIndex uint64
}

// LogCache is the list of log entries on the current node. It is backed up to
// disk; this file provides setters and getters to read and update the log cache.
type LogCache []LogEntry

// SNAPSHOT STUFF START
//
// StateSnapshot contains all the relevant log information up to the given index and term
type StateSnapshot struct {
	// Snapshot metadata
	LastIncludedIndex uint64
	LastIncludedTerm  uint64

	// The state machine
	KVStore map[string]string
}

// effectively a zero-value constructor for a snapshot!
func getDefaultSnapshot() (snap *StateSnapshot) {
	snap = new(StateSnapshot)
	snap.LastIncludedIndex = 0
	snap.LastIncludedTerm = 0
	return
}

func (r *RaftNode) TakeSnapshot() {

	// Lock stable store and node mutexes to freeze node during snapshot process
	r.ssMutex.Lock()
	defer r.ssMutex.Unlock()

	r.nodeMutex.Lock()
	defer r.nodeMutex.Unlock()

	// TODO: do stuff to a working copy of a StateSnapshot in case snapshot doesnt work
	// workingSnapshot := new(StateSnapshot)

	// initialize snapshot struct
	r.stableState.Snapshot = *new(StateSnapshot)

	// get state from state machine
	//r.stableState.Snapshot.KVStore = make(map[string]string)

	// copy all mappings from node's stateMachine into snapshot's stateMachine
	// need to make a copy of the kvstore because maps are passed by reference
	for k, v := range (r.stateMachine.GetState()).(map[string]string) {
		r.stableState.Snapshot.KVStore[k] = v
	}

	// store snapshot metadata
	r.stableState.Snapshot.LastIncludedIndex = r.getLastLogIndex()
	r.stableState.Snapshot.LastIncludedTerm = r.GetCurrentTerm()

	// Write to disk
	err := WriteStableState(&r.raftMetaFd, r.stableState)

	// stablestate is recorded up to r.commitIndex
	if err != nil {
		Error.Printf("Unable to record snapshot to disk: %v\n", err)
		// panic(err) // help! panic! hehe!
	} else {
		// Once written to disk with no errors,
		// truncate log and record offset
		// r.stableState.SnapshotIndex += r.getLastLogIndex() // THIS WILL MESS THINGS UP
		r.stableState.SnapshotIndex += (uint64)(len(r.logCache))
		r.truncateLog(0)
	}

}

// This function initiates a snapshot
func (r *RaftNode) PropagateSnapshot(remote *RemoteNode) {
	// Can only be invoked by a leader node
	if r.State != LEADER_STATE {
		return
	}
	// Take a snapshot on this machine
	r.TakeSnapshot()
	// Encode snapshot as byte array
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(r.stableState.Snapshot)
	// Initialize a snapshot requests
	req := SnapshotRequest{
		Term:              r.GetCurrentTerm(),
		LeaderId:          r.Id,
		LastIncludedTerm:  r.getLogEntry(r.commitIndex).GetTermId(),
		LastIncludedIndex: r.commitIndex,
		Data:              buf.Bytes(),
	}
	// Send over data to the follwer remote node
	remote.InstallSnapshotRPC(r, &req)
	return
}

// // SNAPSHOT STUFF END

// initStableStore initializes the StableState and LogCache. If previous logs
// exist, it reads from them, otherwise it creates log files and populates
// StableState and LogCache with default values.
func (r *RaftNode) initStableStore() (bool, error) {
	freshNode := false

	// Create log path directory if it doesn't already exist
	_, err := os.Stat(r.config.LogPath)
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(r.config.LogPath, 0777)

		/*if err == nil {
			Out.Printf("Created log directory: %v\n", r.config.LogPath)
		}*/

		if err != nil && !os.IsExist(err) {
			Error.Printf("error creating dir %v\n", err)
			return freshNode, err
		}
	}

	// Initialize logCache and metadata files
	r.raftLogFd = FileData{
		fd:       nil,
		size:     0,
		filename: fmt.Sprintf("%v/%d_raftlog.dat", r.config.LogPath, r.port),
	}
	r.raftMetaFd = FileData{
		fd:       nil,
		size:     0,
		filename: fmt.Sprintf("%v/%d_raftmeta.dat", r.config.LogPath, r.port),
	}

	raftLogSize, raftLogExists := getFileInfo(r.raftLogFd.filename)
	r.raftLogFd.size = raftLogSize

	raftMetaSize, raftMetaExists := getFileInfo(r.raftMetaFd.filename)
	r.raftMetaFd.size = raftMetaSize

	if raftLogExists && raftMetaExists {
		// If previous state exists, re-populate everything...
		//fmt.Printf("Reloading previous raftlog (%v) and raftmeta (%v)\n",
		//r.raftLogFd.filename, r.raftMetaFd.filename)

		// Read in previous log and populate index mappings
		entries, _ := ReadRaftLog(&r.raftLogFd)
		if err != nil {
			Error.Printf("Error reading in raft log: %v\n", err)
			return freshNode, err
		}
		r.logCache = entries

		// Create append-only file descriptor for later writing out of log entries.
		err = openRaftLogForWrite(&r.raftLogFd)
		if err != nil {
			Error.Printf("Error opening raftlog for write: %v\n", err)
			return freshNode, err
		}

		// Read in previous metalog and set cache
		ss, _ := ReadStableState(&r.raftMetaFd)
		if err != nil {
			Error.Printf("Error reading stable state: %v\n", err)
			return freshNode, err
		}
		r.stableState = *ss

	} else if (!raftLogExists && raftMetaExists) || (raftLogExists && !raftMetaExists) {
		// If only one of the two logs exists, throw error
		Error.Println("Both raftlog and raftmeta files must exist to proceed!")
		err = errors.New("Both raftlog and raftmeta files must exist to start this node")
		return freshNode, err

	} else {
		// If neither of the two logs exist, create new files
		freshNode = true
		//Out.Printf("Creating new raftlog and raftmeta files")

		err := CreateRaftLog(&r.raftLogFd)
		if err != nil {
			Error.Printf("Error creating new raftlog: %v\n", err)
			return freshNode, err
		}

		err = CreateStableState(&r.raftMetaFd)
		if err != nil {
			Error.Printf("Error creating new stable state: %v\n", err)
			return freshNode, err
		}

		// Init other nodes to zero, this will become populated
		r.stableState.NodeList = make([]RemoteNode, 0)

		// Init client reply cache
		r.stableState.ClientReplyCache = make(map[string]ClientReply)

		// No previous log cache exists, so a fresh one must be created.
		r.logCache = make(LogCache, 0)

		// If the log is empty we need to bootstrap it by adding the first committed entry.
		initEntry := LogEntry{
			Index:  0,
			TermId: r.GetCurrentTerm(),
			Type:   CommandType_INIT,
			Data:   []byte{0},
		}
		r.appendLogEntry(initEntry)
		r.setCurrentTerm(0)

		// Snapshot stuff start
		// initialize StateSnapshot to nil?
		r.stableState.Snapshot = *getDefaultSnapshot()
		r.stableState.SnapshotIndex = 0
		// Snapshot stuff end
	}

	return freshNode, nil
}

////////////////////////////////////////////////////////////////////////////////
// Setters and Getters for StableState                                        //
////////////////////////////////////////////////////////////////////////////////

// setCurrentTerm sets the current node's term and writes log to disk
func (r *RaftNode) setCurrentTerm(newTerm uint64) {
	r.ssMutex.Lock()
	defer r.ssMutex.Unlock()

	// Out.Printf("CURRENT STABLE STATE: %v", r.stableState)
	// Out.Printf("CURRENT STABLE CURRENT TERM: %v", r.stableState.CurrentTerm)

	/*if newTerm != r.stableState.CurrentTerm {
		Out.Printf("(%v) Setting current term from %v -> %v", r.Id, r.stableState.CurrentTerm, newTerm)
	}*/
	// Out.Printf("CURRRENT STABLESTORE SNAPOFFSET: %v", r.stableState.SnapshotIndex)

	r.stableState.CurrentTerm = newTerm

	err := WriteStableState(&r.raftMetaFd, r.stableState)
	if err != nil {
		Error.Printf("Unable to flush new term to disk: %v\n", err)
		panic(err)
	}

}

// GetCurrentTerm returns the current node's term
func (r *RaftNode) GetCurrentTerm() uint64 {
	return r.stableState.CurrentTerm
}

// setVotedFor sets the candidateId for which the current node voted for, and writes log to disk
func (r *RaftNode) setVotedFor(candidateId string) {
	r.ssMutex.Lock()
	defer r.ssMutex.Unlock()

	r.stableState.VotedFor = candidateId

	err := WriteStableState(&r.raftMetaFd, r.stableState)
	if err != nil {
		Error.Printf("Unable to flush new votedFor to disk: %v\n", err)
		panic(err)
	}
}

// GetVotedFor returns the Id of the candidate that the current node voted for
func (r *RaftNode) GetVotedFor() string {
	return r.stableState.VotedFor
}

// setRemoteSelf sets the current node's RemoteSelf and writes log to disk
func (r *RaftNode) setRemoteSelf(remoteSelf *RemoteNode) {
	r.ssMutex.Lock()
	defer r.ssMutex.Unlock()

	r.stableState.RemoteSelf = *remoteSelf

	err := WriteStableState(&r.raftMetaFd, r.stableState)
	if err != nil {
		Error.Printf("Unable to flush new RemoteSelf to disk: %v\n", err)
		panic(err)
	}
}

// GetRemoteSelf returns the current node's representation of itself as a remote node
func (r *RaftNode) GetRemoteSelf() *RemoteNode {
	return &r.stableState.RemoteSelf
}

// SetNodeList sets the current node's conception of all nodes in the cluster
func (r *RaftNode) SetNodeList(nodePointers []*RemoteNode) {
	// Get nodes from node pointers
	nodes := make([]RemoteNode, len(nodePointers))
	for i, np := range nodePointers {
		nodes[i] = *np
	}

	r.ssMutex.Lock()
	defer r.ssMutex.Unlock()

	r.stableState.NodeList = nodes

	err := WriteStableState(&r.raftMetaFd, r.stableState)
	if err != nil {
		Error.Printf("Unable to flush new other nodes to disk: %v\n", err)
		panic(err)
	}
}

// AppendToNodeList adds nodes to the current node's conception of all nodes in the cluster
func (r *RaftNode) AppendToNodeList(other RemoteNode) {
	r.ssMutex.Lock()
	defer r.ssMutex.Unlock()

	r.stableState.NodeList = append(r.stableState.NodeList, other)

	err := WriteStableState(&r.raftMetaFd, r.stableState)
	if err != nil {
		Error.Printf("Unable to flush new other nodes to disk: %v\n", err)
		panic(err)
	}
}

// GetNodeList returns the list of nodes in the Raft cluster as known by the current node
func (r *RaftNode) GetNodeList() []RemoteNode {
	return r.stableState.NodeList
}

// CacheClientReply caches the given client response with the provided cache ID.
func (r *RaftNode) CacheClientReply(cacheId string, reply ClientReply) error {
	r.ssMutex.Lock()
	defer r.ssMutex.Unlock()

	// Check if same cacheId already exists in cache
	_, ok := r.stableState.ClientReplyCache[cacheId]
	if ok {
		return errors.New("request with the same clientId and seqNum already exists")
	}

	r.stableState.ClientReplyCache[cacheId] = reply

	err := WriteStableState(&r.raftMetaFd, r.stableState)
	if err != nil {
		Error.Printf("Unable to flush new client request to disk: %v\n", err)
		panic(err)
	}

	return nil
}

// GetCachedReply checks if the given client request has a cached response.
// It returns the cached response (or nil) and a boolean indicating whether or not
// a cached response existed.
func (r *RaftNode) GetCachedReply(clientReq ClientRequest) (*ClientReply, bool) {
	cacheId := createCacheId(clientReq.ClientId, clientReq.SequenceNum)

	val, ok := r.stableState.ClientReplyCache[cacheId]

	if ok {
		return &val, ok
	} else {
		return nil, ok
	}
}

////////////////////////////////////////////////////////////////////////////////
// Setters and Getters for LogCache                                           //
////////////////////////////////////////////////////////////////////////////////

// appendLogEntry adds a log entry to the current node's cache, and writes it to disk
func (r *RaftNode) appendLogEntry(entry LogEntry) error {
	// Write entry to disk
	err := AppendLogEntry(&r.raftLogFd, &entry)
	if err != nil {
		return err
	}
	// Update entry in cache
	r.logCache = append(r.logCache, entry)
	return nil
}

// truncateLog removes all log entries at index and after it (an inclusive truncation!)
func (r *RaftNode) truncateLog(index uint64) error {
	// Truncate log on disk
	err := TruncateLog(&r.raftLogFd, index)
	if err != nil {
		return err
	}

	// Remove entries from cache
	r.logCache = r.logCache[:index]
	return nil
}

// getLogEntry returns the log entry at the given index
func (r *RaftNode) getLogEntry(index uint64) *LogEntry {

	//original
	// if index < uint64(len(r.logCache)) {
	// 	return &r.logCache[index]
	// }

	// Out.Printf("Hit getLogEntry\n")
	// Out.Printf("LOG CACHE: %v\n", r.logCache)

	// special case: everything is 0

	//make sure log index is valid
	if (index-r.stableState.SnapshotIndex) < uint64(len(r.logCache)) && (index >= r.stableState.SnapshotIndex) {
		return &r.logCache[index-r.stableState.SnapshotIndex]
	}

	return nil
}

// getLastLogIndex returns the index of the last log entry on the current node
func (r *RaftNode) getLastLogIndex() uint64 {
	return uint64(len(r.logCache)-1) + r.stableState.SnapshotIndex
}
