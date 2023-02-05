# Raft

Add your README here!

Things To Know:
- Our test run locally but we may not run in ssh due to git erros where we apparently don't have specific googleapi packages. 
- The example tests like TestClientInteraction_Leader, TestClientInteraction_Follower and TestPartition will sometimes fail if they are run concurrently (i.e go test), but will pass if run individually.
- Sometimes tests will fail if they preload the raft metalogs, in that case we usually run it again and it passes. We've also ran each test individually and they pass.
- Our test cover changes dramatically depending on the run. Sometimes it is over 70% and others times it will be as low as 64%. Unclear why that is but we suspect that it's because of the random components of our code like randomTimeout which cause different parts of our code to be put in use.

Extra Tests We Wrote:
- TestClientInteraction_Candidate
  This test, similar to the others in this file tries to register a client with a node in a candidate state and expects that the candidate will reject this request. We also send it a client request with the same expectations.

- TestReqVoteToHigherTerm
  This test targets the part of the code where a node with a higher term gets a requestVote and tells the other node to fallback.

- TestMultipleClientInteraction_Leader
  This function tests multiple clients interacting with a leader, all of which should end successfully with the client requests accurately appended to the log and committed.

- TestSpecialClientRequests
  We send a clientRequests and registration request to followers with Leaders and thus should be sent back a LeaderHint. It also sends duplicate client requests to leaders with the expectation that the reply is cached and that the request is no longer cached.

- TestPartitionPossibilities:
  We start out with five nodes and partition out a follower and becomes a candidate with ever increasing terms. Then we partition away the leader and it's by itself. This forces the three node cluster remaining to elect a leader. THis means that this cluster has a higher term than the previous leader. THen we reintroduce the old leader back into the cluster and we ensure that the old leader fallsback to a follower. At the end we reintroduce the partitioned follower that became a candidate in its own cluster and make sure that even though that node has a high term, it does not become leader because it's log is not up-to-date.

It's important to note that sometimes partition tests will fail if it is run in tandem.

- TestRandomTimeout
  This is a unit test that tests that our use of the RandomTimeout channel is correct. This was integral during development because we had a lot of issues with elections timing out to early and we wanted to make sure that our use of RandomTimeout was correct.

- TestNetworkPolicyDenied
  This tests that if a node calls an RPC function on another node that it does not have a connection to because of a partition that it gets back TestNetworkPolicyDenied.



Extra Features:

We have implemented a basic web client to interact with the raft nodes.

grpc-web and protoc were used to compile raft_rpc_pb.js and raft_rpc_grpc_web_pb.js which are implementations of the raft procol in javascript.
- The client (raft_rpc_client.js) works analogously to the go client.

- Docker with Envoy is used to connect grpc-web calls from the front end to grpc calls in the backend.
  The web client sends web-grpc requests to port 8080, which is then converted into a regular grpc call and
  sent to port 9090. Therefore, a node needs to be receiving on port 9090 in order to interface with the web client.

- Detailed information on how to build the docker/envoy container and run the web client is found in docker_cheatsheet.txt.
