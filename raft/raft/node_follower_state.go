package raft

import (
	"math"
)

// doFollower implements the logic for a Raft node in the follower state.
func (r *RaftNode) doFollower() stateFunction {
	Out.Printf("[node %x] Transitioning to FOLLOWER_STATE", r.GetRemoteSelf().GetId())
	r.State = FOLLOWER_STATE

	// TODO: Students should implement this method
	// Hint: perform any initial work, and then consider what a node in the
	// follower state should do when it receives an incoming message on every
	// possible channel.

	// election timeout setup
	electionTimeout := randomTimeout(r.config.ElectionTimeout)
	r.nextIndex = make(map[string]uint64)
	r.matchIndex = make(map[string]uint64)
	r.requestsByCacheId = make(map[string]chan ClientReply)
	Out.Printf("following is my PASSION %x \n", r.GetRemoteSelf().GetId())

	for {
		select {
		case <-electionTimeout:
			Out.Printf("where am i--- follower node %x electionTimeout \n", r.GetRemoteSelf().GetId())

			// start an election
			return r.doCandidate

		case heartbeat := <-r.appendEntries:
			resetTimeout, _ := r.handleAppendEntries(heartbeat)
			if resetTimeout {
				Out.Printf("[node %x] RESETTING election timeout in follower ------- \n",
					r.GetRemoteSelf().GetId())
				electionTimeout = randomTimeout(r.config.ElectionTimeout)
			}

		case vote := <-r.requestVote:
			// check stuff and act accordingly

			electionTimeout = randomTimeout(r.config.ElectionTimeout)
			Out.Printf("node %x received vote request \n", r.GetRemoteSelf().GetId())

			term := r.GetCurrentTerm()
			if vote.request.GetTerm() < term {
				Out.Printf("rejecting vote to candidate %x", vote.request.GetCandidate().GetId())
				vote.reply <- RequestVoteReply{Term: term, VoteGranted: false}
			} else {
				//Out.Printf("[node %x] sending vote to candidate %x", r.GetRemoteSelf().GetId(), vote.request.GetCandidate().GetId())
				if vote.request.GetTerm() > term {
					r.setCurrentTerm(vote.request.GetTerm())
					term = vote.request.GetTerm()
					r.setVotedFor("")
				}

				if r.GetVotedFor() == "" || r.GetVotedFor() == vote.request.GetCandidate().GetId() {
					if vote.request.GetLastLogIndex() >= r.getLastLogIndex() {
						r.setVotedFor(vote.request.GetCandidate().GetId())
						Out.Printf("[node %x] sending vote to candidate %x", r.GetRemoteSelf().GetId(),
							vote.request.GetCandidate().GetId())
						vote.reply <- RequestVoteReply{Term: term, VoteGranted: true}
					} else {
						vote.reply <- RequestVoteReply{Term: term, VoteGranted: false}
					}
				} else {
					vote.reply <- RequestVoteReply{Term: term, VoteGranted: false}
				}
			}
			electionTimeout = randomTimeout(r.config.ElectionTimeout)

		case client := <-r.registerClient:
			// forward to leader
			if r.Leader == nil {
				client.reply <- RegisterClientReply{Status: ClientStatus_NOT_LEADER,
					LeaderHint: r.GetRemoteSelf()}
			} else {
				client.reply <- RegisterClientReply{Status: ClientStatus_NOT_LEADER, LeaderHint: r.Leader}
			}

		case req := <-r.clientRequest:
			// forward to leader
			if r.Leader == nil {
				req.reply <- ClientReply{Status: ClientStatus_NOT_LEADER, LeaderHint: r.GetRemoteSelf()}
			} else {
				req.reply <- ClientReply{Status: ClientStatus_NOT_LEADER, LeaderHint: r.Leader}
			}

		case shutdown := <-r.gracefulExit:
			if shutdown {
				return nil
			}
		}
	}
}

// handleAppendEntries handles an incoming AppendEntriesMsg. It is called by a
// node in a follower, candidate, or leader state. It returns two booleans:
// - resetTimeout is true if the follower node should reset the election timeout
// - fallback is true if the node should become a follower again
func (r *RaftNode) handleAppendEntries(msg AppendEntriesMsg) (resetTimeout, fallback bool) {
	// TODO: Students should implement this method
	Out.Printf("[node: %x] handleAppendEntry from [leader: %x] \n", r.GetRemoteSelf().GetId(),
		msg.request.GetLeader().GetId())
	req := msg.request
	resetTimeout = true
	term := r.GetCurrentTerm()
	resetTimeout = true
	fallback = true
	// set current leader to whoever this person is
	if req.GetTerm() < term {
		resetTimeout = false
		fallback = true
		msg.reply <- AppendEntriesReply{Term: term, Success: false}
	} else {
		logEntry := r.getLogEntry(req.GetPrevLogIndex())
		if logEntry == nil || logEntry.GetTermId() != req.GetPrevLogTerm() {
			Out.Printf("[node %x] length of my log: %v also looka that term tho %v",
				r.GetRemoteSelf().GetId(), len(r.logCache), logEntry.GetTermId())
			Out.Printf("logEntry: %v\nPrevLogIndex: %v\nPrevLogTerm: %v\n", logEntry,
				req.GetPrevLogIndex(), req.GetPrevLogTerm())
			Out.Printf("u have failed me, leader\n")
			msg.reply <- AppendEntriesReply{Term: term, Success: false}
		} else {
			//r.printLog()
			r.Leader = req.GetLeader()

			if r.getLogEntry(req.GetPrevLogIndex()+1) != nil {
				r.truncateLog(req.GetPrevLogIndex() + 1)
			}

			for i := 0; i < len(req.GetEntries()); i++ {
				r.appendLogEntry(*req.GetEntries()[i])
			}

			if req.GetLeaderCommit() > r.commitIndex {
				r.commitIndex = uint64(math.Min(float64(req.GetLeaderCommit()),
					float64(req.GetPrevLogIndex())+float64(len(req.GetEntries()))))
			}

			r.setCurrentTerm(req.GetTerm())
			msg.reply <- AppendEntriesReply{Term: req.GetTerm(), Success: true}
		}
	}

	for ; r.commitIndex > r.lastApplied; r.lastApplied++ {
		r.processLogEntry(*r.getLogEntry(r.lastApplied + 1))
	}

	return
}

// print the entire log
func (r *RaftNode) printLog() {
	Out.Printf("[node: %x] IM PRINTING THE LOG\n\n\n\n\n\n\n", r.GetRemoteSelf().GetId())
	for i := 0; i < len(r.logCache); i++ {
		Out.Printf("[node: %x] here is entry %v of mah Log: %v", r.GetRemoteSelf().GetId(), i, r.getLogEntry(uint64(i)))
	}
	Out.Println()
}
