/**
 * @fileoverview gRPC-Web generated client stub for raft
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.raft = require('./raft_rpc_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.raft.RaftRPCClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.raft.RaftRPCPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.raft.RemoteNode,
 *   !proto.raft.Ok>}
 */
const methodInfo_RaftRPC_JoinCaller = new grpc.web.AbstractClientBase.MethodInfo(
  proto.raft.Ok,
  /** @param {!proto.raft.RemoteNode} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.raft.Ok.deserializeBinary
);


/**
 * @param {!proto.raft.RemoteNode} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.raft.Ok)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.raft.Ok>|undefined}
 *     The XHR Node Readable Stream
 */
proto.raft.RaftRPCClient.prototype.joinCaller =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/raft.RaftRPC/JoinCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_JoinCaller,
      callback);
};


/**
 * @param {!proto.raft.RemoteNode} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.raft.Ok>}
 *     A native promise that resolves to the response
 */
proto.raft.RaftRPCPromiseClient.prototype.joinCaller =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/raft.RaftRPC/JoinCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_JoinCaller);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.raft.StartNodeRequest,
 *   !proto.raft.Ok>}
 */
const methodInfo_RaftRPC_StartNodeCaller = new grpc.web.AbstractClientBase.MethodInfo(
  proto.raft.Ok,
  /** @param {!proto.raft.StartNodeRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.raft.Ok.deserializeBinary
);


/**
 * @param {!proto.raft.StartNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.raft.Ok)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.raft.Ok>|undefined}
 *     The XHR Node Readable Stream
 */
proto.raft.RaftRPCClient.prototype.startNodeCaller =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/raft.RaftRPC/StartNodeCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_StartNodeCaller,
      callback);
};


/**
 * @param {!proto.raft.StartNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.raft.Ok>}
 *     A native promise that resolves to the response
 */
proto.raft.RaftRPCPromiseClient.prototype.startNodeCaller =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/raft.RaftRPC/StartNodeCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_StartNodeCaller);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.raft.AppendEntriesRequest,
 *   !proto.raft.AppendEntriesReply>}
 */
const methodInfo_RaftRPC_AppendEntriesCaller = new grpc.web.AbstractClientBase.MethodInfo(
  proto.raft.AppendEntriesReply,
  /** @param {!proto.raft.AppendEntriesRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.raft.AppendEntriesReply.deserializeBinary
);


/**
 * @param {!proto.raft.AppendEntriesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.raft.AppendEntriesReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.raft.AppendEntriesReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.raft.RaftRPCClient.prototype.appendEntriesCaller =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/raft.RaftRPC/AppendEntriesCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_AppendEntriesCaller,
      callback);
};


/**
 * @param {!proto.raft.AppendEntriesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.raft.AppendEntriesReply>}
 *     A native promise that resolves to the response
 */
proto.raft.RaftRPCPromiseClient.prototype.appendEntriesCaller =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/raft.RaftRPC/AppendEntriesCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_AppendEntriesCaller);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.raft.RequestVoteRequest,
 *   !proto.raft.RequestVoteReply>}
 */
const methodInfo_RaftRPC_RequestVoteCaller = new grpc.web.AbstractClientBase.MethodInfo(
  proto.raft.RequestVoteReply,
  /** @param {!proto.raft.RequestVoteRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.raft.RequestVoteReply.deserializeBinary
);


/**
 * @param {!proto.raft.RequestVoteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.raft.RequestVoteReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.raft.RequestVoteReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.raft.RaftRPCClient.prototype.requestVoteCaller =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/raft.RaftRPC/RequestVoteCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_RequestVoteCaller,
      callback);
};


/**
 * @param {!proto.raft.RequestVoteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.raft.RequestVoteReply>}
 *     A native promise that resolves to the response
 */
proto.raft.RaftRPCPromiseClient.prototype.requestVoteCaller =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/raft.RaftRPC/RequestVoteCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_RequestVoteCaller);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.raft.RegisterClientRequest,
 *   !proto.raft.RegisterClientReply>}
 */
const methodInfo_RaftRPC_RegisterClientCaller = new grpc.web.AbstractClientBase.MethodInfo(
  proto.raft.RegisterClientReply,
  /** @param {!proto.raft.RegisterClientRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.raft.RegisterClientReply.deserializeBinary
);


/**
 * @param {!proto.raft.RegisterClientRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.raft.RegisterClientReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.raft.RegisterClientReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.raft.RaftRPCClient.prototype.registerClientCaller =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/raft.RaftRPC/RegisterClientCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_RegisterClientCaller,
      callback);
};


/**
 * @param {!proto.raft.RegisterClientRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.raft.RegisterClientReply>}
 *     A native promise that resolves to the response
 */
proto.raft.RaftRPCPromiseClient.prototype.registerClientCaller =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/raft.RaftRPC/RegisterClientCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_RegisterClientCaller);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.raft.ClientRequest,
 *   !proto.raft.ClientReply>}
 */
const methodInfo_RaftRPC_ClientRequestCaller = new grpc.web.AbstractClientBase.MethodInfo(
  proto.raft.ClientReply,
  /** @param {!proto.raft.ClientRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.raft.ClientReply.deserializeBinary
);


/**
 * @param {!proto.raft.ClientRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.raft.ClientReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.raft.ClientReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.raft.RaftRPCClient.prototype.clientRequestCaller =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/raft.RaftRPC/ClientRequestCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_ClientRequestCaller,
      callback);
};


/**
 * @param {!proto.raft.ClientRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.raft.ClientReply>}
 *     A native promise that resolves to the response
 */
proto.raft.RaftRPCPromiseClient.prototype.clientRequestCaller =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/raft.RaftRPC/ClientRequestCaller',
      request,
      metadata || {},
      methodInfo_RaftRPC_ClientRequestCaller);
};


module.exports = proto.raft;

