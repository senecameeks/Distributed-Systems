/*
 *  Brown University, CS138, Spring 2019
 *
 *  Purpose: Provides wrappers around the client interface of GRPC to invoke
 *  functions on remote tapestry nodes.
 */

package tapestry

import (
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	// Uncomment for xtrace
	// util "github.com/brown-csci1380/tracing-framework-go/xtrace/grpcutil"
)

var dialOptions []grpc.DialOption

func init() {
	dialOptions = []grpc.DialOption{grpc.WithInsecure(), grpc.FailOnNonTempDialError(true)}
	// Uncomment for xtrace
	// dialOptions = append(dialOptions, grpc.WithUnaryInterceptor(util.XTraceClientInterceptor))
}

type RemoteNode struct {
	Id      ID
	Address string
}

// Turns a NodeMsg into a RemoteNode
func (n *NodeMsg) toRemoteNode() RemoteNode {
	if n == nil {
		return RemoteNode{}
	}
	idVal, err := ParseID(n.Id)
	if err != nil {
		return RemoteNode{}
	}
	return RemoteNode{
		Id:      idVal,
		Address: n.Address,
	}
}

// Turns a RemoteNode into a NodeMsg
func (n *RemoteNode) toNodeMsg() *NodeMsg {
	if n == nil {
		return nil
	}
	return &NodeMsg{
		Id:      n.Id.String(),
		Address: n.Address,
	}
}

/**
 *  RPC invocation functions
 */

var connMap = make(map[string]*grpc.ClientConn)
var connMapLock = &sync.RWMutex{}

func closeAllConnections() {
	connMapLock.Lock()
	defer connMapLock.Unlock()
	for k, conn := range connMap {
		conn.Close()
		delete(connMap, k)
	}
}

// Creates a new client connection to the given remote node
func makeClientConn(remote *RemoteNode) (*grpc.ClientConn, error) {
	return grpc.Dial(remote.Address, dialOptions...)
}

// Creates or returns a cached RPC client for the given remote node
func (remote *RemoteNode) ClientConn() (TapestryRPCClient, error) {
	connMapLock.RLock()
	if cc, ok := connMap[remote.Address]; ok {
		connMapLock.RUnlock()
		return NewTapestryRPCClient(cc), nil
	}
	connMapLock.RUnlock()

	cc, err := makeClientConn(remote)
	if err != nil {
		return nil, err
	}
	connMapLock.Lock()
	connMap[remote.Address] = cc
	connMapLock.Unlock()

	return NewTapestryRPCClient(cc), err
}

// Remove the client connection to the given node, if present
func (remote *RemoteNode) RemoveClientConn() {
	connMapLock.Lock()
	defer connMapLock.Unlock()
	if cc, ok := connMap[remote.Address]; ok {
		cc.Close()
		delete(connMap, remote.Address)
	}
}

// Check the error and remove the client connection if necessary
func (remote *RemoteNode) connCheck(err error) error {
	if err != nil {
		remote.RemoveClientConn()
	}
	return err
}

// Say hello to a remote address, and get the tapestry node there
func SayHelloRPC(addr string, joiner RemoteNode) (RemoteNode, error) {
	remote := &RemoteNode{Address: addr}
	cc, err := remote.ClientConn()
	if err != nil {
		return RemoteNode{}, err
	}
	node, err := cc.HelloCaller(context.Background(), joiner.toNodeMsg())
	return node.toRemoteNode(), remote.connCheck(err)
}

func (remote *RemoteNode) GetNextHopRPC(id ID, level int32) (RemoteNode, error, *NodeSet) {
	// TODO: students should implement this
	//Debug.Printf("what")
	cc, err := remote.ClientConn()
	if err != nil {
		return RemoteNode{}, err, nil
	}
	//Debug.Printf("is")
	nextHop, err := cc.GetNextHopCaller(context.Background(), &IdMsg{Id: id.String(), Level: level})
	//Debug.Printf("the")
	nodeSet := NewNodeSet()
	nodeSet.AddAll(nodeMsgsToRemoteNodes(nextHop.GetToRemove()))
	//Debug.Printf("problem ocifer")
	return nextHop.GetNext().toRemoteNode(), remote.connCheck(err), nodeSet
}

func (remote *RemoteNode) RegisterRPC(key string, replica RemoteNode) (bool, error) {
	cc, err := remote.ClientConn()
	if err != nil {
		return false, err
	}
	rsp, err := cc.RegisterCaller(context.Background(), &Registration{
		FromNode: replica.toNodeMsg(),
		Key:      key,
	})
	return rsp.GetOk(), remote.connCheck(err)
}

func (remote *RemoteNode) FetchRPC(key string) (bool, []RemoteNode, error) {
	// TODO: students should implement this
	cc, err := remote.ClientConn()
	if err != nil {
		return false, nil, err
	}
	fetchLocations, err := cc.FetchCaller(context.Background(), &Key{Key: key})
	nodes := nodeMsgsToRemoteNodes(fetchLocations.GetValues())
	return fetchLocations.GetIsRoot(), nodes, remote.connCheck(err)
}

func (remote *RemoteNode) RemoveBadNodesRPC(badnodes []RemoteNode) error {
	cc, err := remote.ClientConn()
	if err != nil {
		return err
	}
	_, err = cc.RemoveBadNodesCaller(context.Background(), &Neighbors{remoteNodesToNodeMsgs(badnodes)})
	// TODO deal with ok
	return remote.connCheck(err)
}

func (remote *RemoteNode) AddNodeRPC(toAdd RemoteNode) ([]RemoteNode, error) {
	// TODO: students should implement this
	cc, err := remote.ClientConn()
	if err != nil {
		return nil, err
	}
	neighbors, err := cc.AddNodeCaller(context.Background(), toAdd.toNodeMsg())
	nodes := nodeMsgsToRemoteNodes(neighbors.GetNeighbors())
	return nodes, remote.connCheck(err)
}

func (remote *RemoteNode) AddNodeMulticastRPC(newNode RemoteNode, level int) ([]RemoteNode, error) {
	cc, err := remote.ClientConn()
	if err != nil {
		return nil, err
	}
	rsp, err := cc.AddNodeMulticastCaller(context.Background(), &MulticastRequest{
		NewNode: newNode.toNodeMsg(),
		Level:   int32(level),
	})
	if err != nil {
		return nil, remote.connCheck(err)
	}
	return nodeMsgsToRemoteNodes(rsp.Neighbors), remote.connCheck(err)
}

func (remote *RemoteNode) TransferRPC(from RemoteNode, data map[string][]RemoteNode) error {
	// TODO: students should implement this
	cc, err := remote.ClientConn()
	if err != nil {
		return err
	}
	transferData := make(map[string]*Neighbors)
	for k, v := range data {
		transferData[k] = &Neighbors{remoteNodesToNodeMsgs(v)}
	}
	_, err = cc.TransferCaller(context.Background(), &TransferData{From: from.toNodeMsg(), Data: transferData})
	// deal w Ok
	return remote.connCheck(err)
}

func (remote *RemoteNode) AddBackpointerRPC(bp RemoteNode) error {
	cc, err := remote.ClientConn()
	if err != nil {
		return err
	}
	_, err = cc.AddBackpointerCaller(context.Background(), bp.toNodeMsg())
	// TODO deal with Ok
	return remote.connCheck(err)
}

func (remote *RemoteNode) RemoveBackpointerRPC(bp RemoteNode) error {
	// TODO: students should implement this
	cc, err := remote.ClientConn()
	if err != nil {
		return err
	}
	_, err = cc.RemoveBackpointerCaller(context.Background(), bp.toNodeMsg())
	// deal w Ok
	return remote.connCheck(err)
}

func (remote *RemoteNode) GetBackpointersRPC(from RemoteNode, level int) ([]RemoteNode, error) {
	cc, err := remote.ClientConn()
	if err != nil {
		return nil, err
	}
	rsp, err := cc.GetBackpointersCaller(context.Background(), &BackpointerRequest{from.toNodeMsg(), int32(level)})
	if err != nil {
		return nil, remote.connCheck(err)
	}
	return nodeMsgsToRemoteNodes(rsp.Neighbors), remote.connCheck(err)
}

func (remote *RemoteNode) NotifyLeaveRPC(from RemoteNode, replacement *RemoteNode) error {
	// TODO: students should implement this
	cc, err := remote.ClientConn()
	if err != nil {
		return err
	}
	_, err = cc.NotifyLeaveCaller(context.Background(), &LeaveNotification{From: from.toNodeMsg(), Replacement: replacement.toNodeMsg()})
	// deal w Ok
	return remote.connCheck(err)
}

func (remote *RemoteNode) BlobStoreFetchRPC(key string) (*[]byte, error) {
	cc, err := remote.ClientConn()
	if err != nil {
		return nil, err
	}
	rsp, err := cc.BlobStoreFetchCaller(context.Background(), &Key{key})
	if err != nil {
		return nil, remote.connCheck(err)
	}
	return &rsp.Data, remote.connCheck(err)
}

func (remote *RemoteNode) TapestryLookupRPC(key string) ([]RemoteNode, error) {
	cc, err := remote.ClientConn()
	if err != nil {
		return nil, err
	}
	rsp, err := cc.TapestryLookupCaller(context.Background(), &Key{key})
	if err != nil {
		return nil, remote.connCheck(err)
	}
	return nodeMsgsToRemoteNodes(rsp.Neighbors), remote.connCheck(err)
}

func (remote *RemoteNode) TapestryStoreRPC(key string, value []byte) error {
	cc, err := remote.ClientConn()
	if err != nil {
		return err
	}
	_, err = cc.TapestryStoreCaller(context.Background(), &DataBlob{
		Key:  key,
		Data: value,
	})
	return remote.connCheck(err)
}
