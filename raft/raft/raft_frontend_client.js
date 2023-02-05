// import grpc-web protobuf
const {RaftRPCClient} = require('./raft_rpc_grpc_web_pb.js');

// import js protobufs
const {RegisterClientRequest, RegisterClientReply} = require('./raft_rpc_grpc_web_pb.js');
const {ClientRequest, ClientReply} = require('./raft_rpc_grpc_web_pb.js');

var raftClient = new RaftRPCClient('localhost:8080');
var registerClientReq = new RegisterClientRequest();
var clientId;

// first register client
raftClient.registerClientCaller(registerClientReq, {}, (err, response)=>{
    if (err != nil) {
        console.log(err);
    } else {
        // if successful, print out relevant info
        console.log(response.getStatus());
        console.log(response.getResponse());
        console.log(response.getLeaderhint());
    }
});


// interact with client
var clientReq = new ClientRequest();
// set clientId to id from registration:
clientReq.setClientid(clientId);
clientReq.setSequencenum(2);
clientReq.setStatemachinecmd(5);

raftClient.clientRequestCaller(clientReq, {}, (err, response)=>{

    if (err != nil){
        console.log(err);
    } else {
        console.log(response.getStatus());
        console.log(response.getResponse());
        console.log(response.getLeaderhint());
    }
});
