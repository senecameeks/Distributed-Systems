// Brown University, CS138, Spring 2019
//
// Purpose: Implements functions that are invoked by clients and other nodes
// over RPC.

package raft

import "golang.org/x/net/context"

// JoinCaller is called through GRPC to execute a join request.
func (local *RaftNode) JoinCaller(ctx context.Context, r *RemoteNode) (*Ok, error) {
	// Check if the network policy prevents incoming requests from the requesting node
	if local.NetworkPolicy.IsDenied(*r, *local.GetRemoteSelf()) {
		return nil, ErrorNetworkPolicyDenied
	}

	// Defer to the local Join implementation, and marshall the results to
	// respond to the GRPC request.
	err := local.Join(r)
	return &Ok{Ok: err == nil}, err
}

// StartNodeCaller is called through GRPC to execute a start node request.
func (local *RaftNode) StartNodeCaller(ctx context.Context, req *StartNodeRequest) (*Ok, error) {
	// TODO: Students should implement this method
	// Make sure to check the provided Raft node's network policy to make sure
	// we are allowed to respond to this request. Respond with
	// ErrorNetworkPolicyDenied if not.

	if local.NetworkPolicy.IsDenied(*req.GetFromNode(), *local.GetRemoteSelf()) {
		return nil, ErrorNetworkPolicyDenied
	}

	err := local.StartNode(req)
	return &Ok{Ok: err == nil}, err
}

// AppendEntriesCaller is called through GRPC to respond to an append entries request.
func (local *RaftNode) AppendEntriesCaller(ctx context.Context, req *AppendEntriesRequest) (*AppendEntriesReply, error) {
	// TODO: Students should implement this method
	// Make sure to check the provided Raft node's network policy to make sure
	// we are allowed to respond to this request. Respond with
	// ErrorNetworkPolicyDenied if not.
	if local.NetworkPolicy.IsDenied(*req.GetLeader(), *local.GetRemoteSelf()) {
		return nil, ErrorNetworkPolicyDenied
	}

	reply, err := local.AppendEntries(req)
	if err != nil {
		return &reply, err
	}
	return &reply, err
}

// RequestVoteCaller is called through GRPC to respond to a vote request.
func (local *RaftNode) RequestVoteCaller(ctx context.Context, req *RequestVoteRequest) (*RequestVoteReply, error) {
	// TODO: Students should implement this method
	// Make sure to check the provided Raft node's network policy to make sure
	// we are allowed to respond to this request. Respond with
	// ErrorNetworkPolicyDenied if not.
	if local.NetworkPolicy.IsDenied(*req.GetCandidate(), *local.GetRemoteSelf()) {
		return nil, ErrorNetworkPolicyDenied
	}

	reply, err := local.RequestVote(req)
	if err != nil {
		return &reply, err
	}
	return &reply, err
}

// RegisterClientCaller is called through GRPC to respond to a client
// registration request.
func (local *RaftNode) RegisterClientCaller(ctx context.Context, req *RegisterClientRequest) (*RegisterClientReply, error) {
	// TODO: Students should implement this method
	reply, err := local.RegisterClient(req)
	if err != nil {
		return &reply, err
	}
	return &reply, err
}

// ClientRequestCaller is called through GRPC to respond to a client request.
func (local *RaftNode) ClientRequestCaller(ctx context.Context, req *ClientRequest) (*ClientReply, error) {
	// TODO: Students should implement this method
	reply, err := local.ClientRequest(req)
	if err != nil {
		return &reply, err
	}
	return &reply, err
}
