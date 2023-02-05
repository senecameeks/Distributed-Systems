package raft

import (
	"time"
)

// doLeader implements the logic for a Raft node in the leader state.
func (r *RaftNode) doLeader() stateFunction {
	r.Out("Transitioning to LEADER_STATE")
	r.State = LEADER_STATE

	// TODO: Students should implement this method
	// Hint: perform any initial work, and then consider what a node in the
	// leader state should do when it receives an incoming message on every
	// possible channel.
	index := r.getLastLogIndex() + 1
	term := r.GetCurrentTerm()
	cmdType := CommandType_NOOP
	r.leaderMutex.Lock()
	r.appendLogEntry(LogEntry{Index: index, TermId: term, Type: cmdType})
	Out.Printf("Here is my entry: %v\n", r.getLogEntry(r.getLastLogIndex()))
	r.leaderMutex.Unlock()

	fallbackChan := make(chan bool)
	go r.startHeartbeats(fallbackChan)

	// fallback, _ := r.sendHeartbeats()
	// if fallback {
	// 	Out.Println("i am falling back u heard it here first !")
	// 	return r.doFollower
	// }

	for {
		select {
		case <-fallbackChan:
			return r.doFollower

		case heartbeat := <-r.appendEntries:
			Out.Printf("[leader node: %x] where am i -- getting appendEntries\n", r.Id)
			// should only receive if other potential leader check this for better leader
			_, fallback := r.handleAppendEntries(heartbeat)
			Out.Printf("would u look at that, it is %v that i am falling back\n", fallback)
			if fallback {
				return r.doFollower
			}

		case vote := <-r.requestVote:
			Out.Printf("\t[leader node: %x] requestVote: got a competing requestVote as a leader\n", r.GetRemoteSelf().GetId())
			fallback := r.handleCompetingRequestVote(vote)
			if fallback {
				return r.doFollower
			} else {
				// PUT IN PLACE
				// the problem is that just ignoring it will not resolve itself
				// right now, because other person still gets majority and you
				// you both think you're leader now.
				// send appendEntriesRPC to this fool
				//Out.Printf("telling node %x to STEP TF DOWN by real leader node %x", vote.request.GetCandidate().GetId(), r.GetRemoteSelf().GetId())
				if _, contains := r.nextIndex[vote.request.GetCandidate().GetId()]; !contains {
					r.leaderMutex.Lock()
					r.nextIndex[vote.request.GetCandidate().GetId()] = r.getLastLogIndex() + 1
					r.matchIndex[vote.request.GetCandidate().GetId()] = 0
					r.leaderMutex.Unlock()
				}

				request := &AppendEntriesRequest{Term: term, Leader: r.GetRemoteSelf(),
					PrevLogIndex: r.nextIndex[vote.request.GetCandidate().GetId()] - 1, PrevLogTerm: r.getLogEntry(r.nextIndex[vote.request.GetCandidate().GetId()] - 1).GetTermId(),
					Entries: r.getPartialLog(r.nextIndex[vote.request.GetCandidate().GetId()] - 1), LeaderCommit: r.commitIndex}

				//reply, err :=
				vote.request.GetCandidate().AppendEntriesRPC(r, request)
			}

		case client := <-r.registerClient:
			Out.Printf("[leader node %v], where am i -- register client why tf not\n", r.Id)
			// register the client
			clientId := r.getLastLogIndex() + 1
			r.leaderMutex.Lock()
			Out.Printf("[leader node: %x] appending to log in registerClient", r.Id)
			r.appendLogEntry(LogEntry{Index: clientId, TermId: r.GetCurrentTerm(), Type: CommandType_CLIENT_REGISTRATION})
			r.leaderMutex.Unlock()
			// need to send quick heartbeat to all nodes, and reply to client once received by a majority
			// DO WE CHECK FOR FALLBACK HERE why not
			fallback, sentToMajority := r.sendHeartbeats()
			if sentToMajority {
				Out.Printf("REPLYING TO CLIENT cuz sentToMajority is true")
				client.reply <- RegisterClientReply{Status: ClientStatus_OK, ClientId: clientId}
			} else {
				Out.Printf("REPLYING TO CLIENT but sentToMajority is false")
				client.reply <- RegisterClientReply{Status: ClientStatus_REQ_FAILED}
			}
			if fallback {
				return r.doFollower
			}

		case req := <-r.clientRequest:
			Out.Println("where am i -- client request sure")
			// actually do stuff
			// req.request contains a GetSequenceNum to say whether or not it is duplicate, so prob check for that

			clientReply, ok := r.GetCachedReply(*req.request)

			if ok { // reply exists in cache
				req.reply <- *clientReply
			} else {

				// check if in request cache
				// TODO: reset request channel if node fallbacks
				clientID := createCacheId(req.request.GetClientId(), req.request.GetSequenceNum())
				if _, contains := r.requestsByCacheId[clientID]; !contains {
					// Out.Printf("clientID is in cache...")
					// // GO FUNC!!!
					// go func() {
					// 	for {
					// 		select {
					// 		case clientReplyThing := <-clientReplyChan:
					// 			Out.Printf("we might be stuck in clientReplyThing")
					// 			req.reply <- clientReplyThing
					// 		}
					// 	}
					// }()
					r.requestsByCacheId[clientID] = make(chan ClientReply)

					go r.waitForClientReply(req, clientID)

					// should append new entry to log
					index := r.getLastLogIndex() + 1
					term := r.GetCurrentTerm()
					cmdType := CommandType_STATE_MACHINE_COMMAND
					cmd := req.request.GetStateMachineCmd()
					data := req.request.GetData()
					r.leaderMutex.Lock()
					Out.Printf("[leader node: %x] appending to the log in client request", r.Id)
					r.appendLogEntry(LogEntry{Index: index, TermId: term, Type: cmdType, Command: cmd, Data: data, CacheId: clientID})
					r.leaderMutex.Unlock()

					// then send heartbeats until holds across a majority
					fallback, sentToMajority := r.sendHeartbeats() // how to respond ?
					// am lazy someone ask our ta if this is ok
					if fallback {
						Out.Println("i am falling back u heard it here first !")
						return r.doFollower
					}

					// then only respond once entry is applied to state machine
					// ask our ta how to set up way to check if majority of servers ever have a certain log entry
					if sentToMajority {
						// cache request as soon as we get it.
						// after a majority have replicated once it is committed
						// then reply to client
						// then cache reply
						Out.Printf("sentToMajority is true replying to client request blah")
						r.updateCommitIndex()
						clientReply := r.processLogEntry(*r.getLogEntry(index))
						req.reply <- clientReply
						//r.CacheClientReply(clientID, clientReply)

					} else {
						Out.Println("majority was not reached putting client request back in channel")
						r.clientRequest <- req
					}
				} else {
					go r.waitForClientReply(req, clientID)
				}
			}

		case shutdown := <-r.gracefulExit:
			Out.Println("where am i -- SHUTDOWN")
			if shutdown {
				return nil
			}
		}
	}
}

func (r *RaftNode) waitForClientReply(request ClientRequestMsg, cacheID string) {
	for {
		select {
		case clientReply := <-r.requestsByCacheId[cacheID]:
			request.reply <- clientReply
		}
	}
}

// sendHeartbeats is used by the leader to send out heartbeats to each of
// the other nodes. It returns true if the leader should fall back to the
// follower state. (This happens if we discover that we are in an old term.)
//
// If another node isn't up-to-date, then the leader should attempt to
// update them, and, if an index has made it to a quorum of nodes, commit
// up to that index. Once committed to that index, the replicated state
// machine should be given the new log entries via processLogEntry.
func (r *RaftNode) sendHeartbeats() (fallback, sentToMajority bool) {
	// TODO: Students should implement this method
	//r.updateCommitIndex()
	nodeList := r.GetNodeList()
	term := r.GetCurrentTerm()
	// Out.Printf("in send heartbeats \n")
	//fallback = false
	//sentToMajority = false
	replies := make(chan AppendEntriesReply)
	//timer := time.Tick(r.config.HeartbeatTimeout)

	for i := 0; i < len(nodeList); i++ {
		//for _, node := range nodeList {
		if nodeList[i].GetId() != r.GetRemoteSelf().GetId() {
			Out.Printf("[leader node: %x] sending heartbeat to node %x \n", r.GetRemoteSelf().GetId(), nodeList[i].GetId())
			// check if nextIndex and matchIndex exists for this node
			if _, contains := r.nextIndex[nodeList[i].Id]; !contains {
				r.leaderMutex.Lock()
				r.nextIndex[nodeList[i].Id] = r.getLastLogIndex() + 1
				r.matchIndex[nodeList[i].Id] = 0
				r.leaderMutex.Unlock()
			}

			go r.sendHeartbeatsHelper(replies, &nodeList[i], term)
		}

	}
	//Out.Printf("sent heart beats to all nodes in nodelist glub")
	succ_count := 1
	repCount := 1
outer:
	for {
		select {
		// case <-timer:
		// 	break

		case rep := <-replies:
			repCount++
			Out.Printf("[leader node: %x] reply received with term %x\n\t versus my currTerm %x\n", r.GetRemoteSelf().GetId(), rep.GetTerm(), r.GetCurrentTerm())
			// Out.Printf("versus my term %x \n", )
			if rep.GetSuccess() {
				Out.Printf("\t[leader node: %x] reply successful \n", r.GetRemoteSelf().GetId())
				succ_count++
			} else {
				Out.Printf("\t[leader node: %x] reply failed \n", r.GetRemoteSelf().GetId())
			}
			//got replies from everyone
			if succ_count > len(nodeList)/2 {
				Out.Printf("majority replied should commit now with succ count: %x \n", succ_count)
				sentToMajority = true

				r.updateCommitIndex()
				if r.lastApplied < r.commitIndex {
					Out.Println("sentToMajority is true replying to client request")
					Out.Printf("[leader node: %x] gonna process log entries from %v to %v\n",
						r.GetRemoteSelf().GetId(), r.lastApplied+1, r.commitIndex)
					// dont have index of exact client request details, but we know that successfully
					//  replicated log to majority of servers => loop from lastApplied to updateCommitIndex
					// prob updateCommitIndex right before tbh, loop through indices and process all these logentries
					// update lastApplied as u do so, able to pull cacheID from logentry using logEntry.GetCacheId()
					// use this cache id to send clientReply from processlogentry through requestsByCacheId channel
					// directly to client, then cache the reply. if does not work, something wrong w go func in request
					// logic above
					//clientReply := r.processLogEntry(*r.getLogEntry(index))
					//req.reply <- clientReply
					//r.requestsByCacheId[clientID] <- clientReply
					//r.CacheClientReply(clientID, clientReply)
					var logEntry LogEntry
					for ; r.lastApplied < r.commitIndex; r.lastApplied++ {
						logEntry = *r.getLogEntry(r.lastApplied + 1)
						Out.Printf("[leader node: %x] i am out here %v, processing these ENTRIES man %v\n",
							r.GetRemoteSelf().GetId(), r.lastApplied+1, logEntry)
						r.processLogEntry(logEntry)
					}
				}

				break outer
			} else {
				sentToMajority = false
			}
			Out.Printf("not yet majority: %x \n", succ_count)
			if rep.GetTerm() > term {
				Out.Printf("\t[leader node: %x] should fallback \n", r.GetRemoteSelf().GetId())
				fallback = true
				break outer
			} else {
				fallback = false
			}

			if repCount == r.config.ClusterSize-1 {
				break outer
			}
		}
	}
	Out.Printf("broke out of outer should commit now if sentToMajority is true \n")
	return fallback, sentToMajority // not unreachable code, vet is confuseds
}

func (r *RaftNode) sendHeartbeatsHelper(replies chan AppendEntriesReply, node *RemoteNode, term uint64) {
	//Out.Printf("node that actually got to sendHeartHelper: %x \n", node.GetId())
	//Out.Printf("HOLLA setting PrevLogIndex+1 to be %v for node %x\n", r.nextIndex[node.GetId()], node.GetId())
	request := AppendEntriesRequest{Term: term, Leader: r.GetRemoteSelf(),
		PrevLogIndex: r.nextIndex[node.GetId()] - 1, PrevLogTerm: r.getLogEntry(r.nextIndex[node.GetId()] - 1).GetTermId(),
		Entries: r.getPartialLog(r.nextIndex[node.GetId()]), LeaderCommit: r.commitIndex}

	reply, err := node.AppendEntriesRPC(r, &request) // HANDLE ERROR LATER

	if reply != nil {
		Out.Printf("[leader node: %x] got append entry reply from node %x with success: %v \n", r.GetRemoteSelf().GetId(), node.GetId(), reply.GetSuccess())
		if !reply.GetSuccess() {
			Out.Printf("[node: %x] append entries failed, decrementing nextIndex\n", node.GetId())
			r.leaderMutex.Lock()
			if r.nextIndex[node.GetId()] > 0 {
				r.nextIndex[node.GetId()]--
			}
			r.leaderMutex.Unlock()
			Out.Printf("nextIndex for node %x is now %v\n", node.GetId(), r.nextIndex[node.GetId()])
		} else {
			// set matchIndex here
			//Out.Printf("should be updating matchIndex and nextIndex for node %x", node.GetId())
			r.leaderMutex.Lock()
			r.nextIndex[node.GetId()] += uint64(len(request.Entries))
			r.matchIndex[node.GetId()] = r.nextIndex[node.GetId()] - 1 //request.GetPrevLogIndex()
			r.leaderMutex.Unlock()
		}
		replies <- *reply
	} else {
		Out.Printf("AppendEntries error: %v\n", err)
		replies <- AppendEntriesReply{Term: 0, Success: false}
	}

	return
}

// starting a new go routine to periodically send heartbeats
func (r *RaftNode) startHeartbeats(fallbackChan chan bool) {
	tick := time.Tick(r.config.HeartbeatTimeout)
	count := 1
	//go func() {
	for {

		select {
		case <-tick:
			if r.State == LEADER_STATE {
				Out.Printf("where am i -- sending heartbeat number %v", count)
				count++

				fallback, _ := r.sendHeartbeats()
				if fallback {
					Out.Println("i am falling back u heard it here first !")
					fallbackChan <- true
					return
				}
			} else {
				return
			}
		}
	}
	//}()
}

func (r *RaftNode) updateCommitIndex() {
	// TODO: does cluster size keep track of nodes that have dropped out?
	// do we even need to care?
	majorityNum := r.config.ClusterSize / 2 // TODO: casting
	numMatchIndexOverN := 0

	// check if log[n] == current term
	// check if majority of matchIndex[i] >= n
	// check if n > commitIndex
	// if all above true, commitIndex = n

	for n := r.getLastLogIndex(); n > r.commitIndex; n-- {

		if r.getLogEntry(n).GetTermId() == r.GetCurrentTerm() {

			for _, mIndex := range r.matchIndex {
				//Out.Printf("[leader node: %x] matchIndex %x \n", r.GetRemoteSelf().GetId(), mIndex)
				if mIndex >= n {
					numMatchIndexOverN++
				}
			}
			Out.Println()
			if numMatchIndexOverN > majorityNum {
				r.leaderMutex.Lock()
				r.commitIndex = n
				r.leaderMutex.Unlock()
				break
			}
		}
	}
}

// getPartialLog returns the desired Entries for a leader to send in an
//  AppendEntriesRequest (all entries in its log starting from r.nextIndex[node.Id])
func (r *RaftNode) getPartialLog(startIndex uint64) (entries []*LogEntry) {
	// Out.Printf("NODE %x LAST LOG INDEX: %x\n", r.GetRemoteSelf().GetId(), r.getLastLogIndex())
	// Out.Printf("NODE %x START LOG INDEX: %x\n", r.GetRemoteSelf().GetId(), startIndex)
	entries = make([]*LogEntry, 0)

	//Out.Printf("startIndex: %v for partial log\n", startIndex)
	r.leaderMutex.Lock()
	for i := startIndex; i <= r.getLastLogIndex(); i++ {
		entries = append(entries, r.getLogEntry(i))
		Out.Printf("this LOG is looking like %v at index: %v\n", r.getLogEntry(i), i)
	}
	r.leaderMutex.Unlock()

	return
}

// processLogEntry applies a single log entry to the finite state machine. It is
// called once a log entry has been replicated to a majority and committed by
// the leader. Once the entry has been applied, the leader responds to the client
// with the result, and also caches the response.
func (r *RaftNode) processLogEntry(entry LogEntry) ClientReply {
	Out.Printf("Processing log entry: %v\n", entry)

	status := ClientStatus_OK
	response := ""
	var err error

	// Apply command on state machine
	if entry.Type == CommandType_STATE_MACHINE_COMMAND {
		response, err = r.stateMachine.ApplyCommand(entry.Command, entry.Data)
		if err != nil {
			status = ClientStatus_REQ_FAILED
			response = err.Error()
		}
	}

	// Construct reply
	reply := ClientReply{
		Status:     status,
		Response:   response,
		LeaderHint: r.GetRemoteSelf(),
	}

	// Add reply to cache
	if entry.CacheId != "" {
		r.CacheClientReply(entry.CacheId, reply)
	}

	// Send reply to client
	r.requestsMutex.Lock()
	replyChan, exists := r.requestsByCacheId[entry.CacheId]
	if exists {
		replyChan <- reply
		delete(r.requestsByCacheId, entry.CacheId)
	}
	r.requestsMutex.Unlock()

	return reply
}
