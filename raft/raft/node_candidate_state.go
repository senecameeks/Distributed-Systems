package raft

// import "fmt"

// doCandidate implements the logic for a Raft node in the candidate state.
func (r *RaftNode) doCandidate() stateFunction {
	Out.Printf("[node %x] Transitioning to CANDIDATE_STATE", r.GetRemoteSelf().GetId())
	r.State = CANDIDATE_STATE
	// TODO: Students should implement this method
	// Hint: perform any initial work, and then consider what a node in the
	// candidate state should do when it receives an incoming message on every
	// possible channel.
	term := r.GetCurrentTerm() + 1
	r.setCurrentTerm(term)
	r.setVotedFor(r.GetRemoteSelf().GetId())
	electionTimeout := randomTimeout(r.config.ElectionTimeout)
	electionResults := make(chan bool)
	fallbackChan := make(chan bool)

	r.requestVotes(electionResults, fallbackChan, term)

	voteCount := 1

	for {
		select {
		case <-electionTimeout:
			Out.Printf("where am i--- candidate node %x electionTimeout \n", r.GetRemoteSelf().GetId())
			// start an election
			return r.doCandidate

		case result := <-electionResults:

			if result {
				Out.Printf("\t[node: %x] requestVote: got a vote++\n", r.GetRemoteSelf().GetId())

				voteCount++
			}
			if voteCount > r.config.ClusterSize/2 {
				Out.Printf("\t[node: %x]: elected, term %v \n", r.GetRemoteSelf().GetId(), r.GetCurrentTerm())
				return r.doLeader
			}
			//electionTimeout = randomTimeout(r.config.ElectionTimeout)

		case fb := <-fallbackChan:
			Out.Printf("[node: %x] where am i -- candidate fallback", r.Id)
			if fb {
				Out.Printf("\t[node: %x] requestVote: fallback, becoming follower\n", r.GetRemoteSelf().GetId())
				return r.doFollower
			}
			//electionTimeout = randomTimeout(r.config.ElectionTimeout)

		case heartbeat := <-r.appendEntries:
			Out.Printf("[node: %x] where am i -- candidate getting appendEntry", r.Id)
			_, fallback := r.handleAppendEntries(heartbeat)
			if fallback {
				return r.doFollower
			}
			//electionTimeout = randomTimeout(r.config.ElectionTimeout)

		case vote := <-r.requestVote:
			Out.Printf("\t[node: %x] requestVote: got a competing requestVote\n", r.GetRemoteSelf().GetId())

			fallback := r.handleCompetingRequestVote(vote)
			if fallback {
				Out.Printf("[node: %x] candidate falling back due to other votes", r.Id)
				return r.doFollower
			}
			//electionTimeout = randomTimeout(r.config.ElectionTimeout)

		case client := <-r.registerClient:
			client.reply <- RegisterClientReply{Status: ClientStatus_ELECTION_IN_PROGRESS, LeaderHint: r.GetRemoteSelf()}

		case req := <-r.clientRequest:
			req.reply <- ClientReply{Status: ClientStatus_ELECTION_IN_PROGRESS, LeaderHint: r.GetRemoteSelf()}

		case shutdown := <-r.gracefulExit:
			if shutdown {
				return nil
			}
		}
	}
}

// requestVotes is called to request votes from all other nodes. It takes in a
// channel on which the result of the vote should be sent over: true for a
// successful election, false otherwise.
func (r *RaftNode) requestVotes(electionResults chan bool, fallback chan bool, currTerm uint64) {
	// TODO: Students should implement this method
	Out.Printf("[node: %x] requesting votes!\n", r.GetRemoteSelf().GetId())
	nodeList := r.GetNodeList()
	Out.Printf("requesting votes. Length of nodelist %x \n", len(nodeList))
	for j := 0; j < len(nodeList); j++ {
		//for i, node := range nodeList {
		//inner := nodeList[i]

		// check for self
		inner := nodeList[j]
		term := r.GetCurrentTerm()
		if inner.GetId() != r.GetRemoteSelf().GetId() {
			Out.Printf("requesting vote from node %x\n", nodeList[j].GetId())
			go func() {
				request := RequestVoteRequest{Term: term, Candidate: r.GetRemoteSelf(),
					LastLogIndex: r.getLastLogIndex(), LastLogTerm: r.getLogEntry(r.getLastLogIndex()).GetTermId()}
				reply, _ := inner.RequestVoteRPC(r, &request) // HANDLE ERROR LATER

				fallback <- (reply.GetTerm() > currTerm)
				electionResults <- reply.GetVoteGranted()
			}()
		}
	}

	return
}

// handleCompetingRequestVote handles an incoming vote request when the current
// node is in the candidate or leader state. It returns true if the caller
// should fall back to the follower state, false otherwise.
func (r *RaftNode) handleCompetingRequestVote(msg RequestVoteMsg) (fallback bool) {
	// TODO: Students should implement this method
	term := r.GetCurrentTerm()
	if msg.request.GetTerm() > term {
		Out.Printf("[node: %x] handleCompetingRequestVote: falling back to other node %x\n", r.GetRemoteSelf().GetId(), msg.request.GetCandidate().GetId())

		return true
	}
	Out.Printf("[node: %x] handleCompetingRequestVote: NOT falling back to other node %x\n", r.GetRemoteSelf().GetId(), msg.request.GetCandidate().GetId())

	return false
}
