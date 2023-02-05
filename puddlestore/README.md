# Puddlestore :cloud::zap:
---

lmao so we meet again.
Sometimes when we run our CLI it throws an error and crashes. If this occurs just try running it again. It'll work eventually.
Also when closing the CLI, make sure to run the "quit" command, as it makes sure that the Raft nodes properly shut down.

# Extra Features:
---
## Raft Log Compaction
We implemented Log Compaction in Raft through Snapshotting. A snapshot contains a node's stable state (representing committed entries in the log) up to the moment the snapshot is taken, the last comitted entry, and last comitted entry's term. After a snapshot is taken and is confirmed to be flushed to disk storage, the log is truncated to include only the non-committed entries. 

The relevant parts of the code for snapshotting are as follows:

In `raft_rpc.proto`, `rpc InstallSnapshotCaller()`, `message SnapshotRequest`, and `message SnapshotReply` have been defined to enable the sending of snapshots from a node to another node in the cluster. Accordingly, a new version of `raft_rpc.pb.go` has been compiled.

In `node_stable_store.go`:
- `struct StateSnapshot` struct contains relevant data for a snapshot
- `getDefaultSnapshot()` returns a zero-value snapshot used upon initializing the stable store
- `TakeSnapshot()` saves the node's stable state and truncates the log to only uncomitted entries. Can be invoked by any node regardless of leader or follower state.
- `PropagateSnapshot(remoteNode)` can only be invoked by a leader node. This function calls `TakeSnapshot` on the leader, and calls `InstallSnapshotRPC` on the follower remote node.
- `getLogEntry()` and `getLastLogIndex()` are modified to account for the log index offset caused by a snapshot. Otherwise, they work the same as before.

In `node.go`:
- `InstallSnapshot()` installs a snapshot from a leader into the local state machine if the snapshot's term is higher than the current term.

In `raft_rpc_client.go`, `InstallSnapshotRPC()` can be called to send a snapshot to a remote node, while in `raft_rpc_server.go`, `InstallSnapshotCaller()` receives the RPC call and calls `InstallSnapshot()` as described above.

## ZooKeeper
We utilized ZooKeeper for our membership server, as can be evidenced in our code.

## Tapestry Salting
We implemented Tapestry Salting by appending 3 unique salts to a key and storing 4 replicas
of data into Tapestry. And then during Lookup we just lookup using the key and the key followed by each
our salts.

## Tapestry Hotspot Caching
For Hotspot Caching, we added a `objectFreqs` map to each Tapestry node which keeps a count of how many
times that specific node looks up a specific object, and if the frequency surpasses 3, we move the object to
be stored onto that node.

As for our Puddlestore implementation, we do not support the `cd` command, so each command should use the absolute path.
