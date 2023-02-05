/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.raft = (function() {

    /**
     * Namespace raft.
     * @exports raft
     * @namespace
     */
    var raft = {};

    raft.RaftRPC = (function() {

        /**
         * Constructs a new RaftRPC service.
         * @memberof raft
         * @classdesc Represents a RaftRPC
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function RaftRPC(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (RaftRPC.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = RaftRPC;

        /**
         * Creates new RaftRPC service using the specified rpc implementation.
         * @function create
         * @memberof raft.RaftRPC
         * @static
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {RaftRPC} RPC service. Useful where requests and/or responses are streamed.
         */
        RaftRPC.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link raft.RaftRPC#joinCaller}.
         * @memberof raft.RaftRPC
         * @typedef JoinCallerCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {raft.Ok} [response] Ok
         */

        /**
         * Calls JoinCaller.
         * @function joinCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IRemoteNode} request RemoteNode message or plain object
         * @param {raft.RaftRPC.JoinCallerCallback} callback Node-style callback called with the error, if any, and Ok
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(RaftRPC.prototype.joinCaller = function joinCaller(request, callback) {
            return this.rpcCall(joinCaller, $root.raft.RemoteNode, $root.raft.Ok, request, callback);
        }, "name", { value: "JoinCaller" });

        /**
         * Calls JoinCaller.
         * @function joinCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IRemoteNode} request RemoteNode message or plain object
         * @returns {Promise<raft.Ok>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link raft.RaftRPC#startNodeCaller}.
         * @memberof raft.RaftRPC
         * @typedef StartNodeCallerCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {raft.Ok} [response] Ok
         */

        /**
         * Calls StartNodeCaller.
         * @function startNodeCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IStartNodeRequest} request StartNodeRequest message or plain object
         * @param {raft.RaftRPC.StartNodeCallerCallback} callback Node-style callback called with the error, if any, and Ok
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(RaftRPC.prototype.startNodeCaller = function startNodeCaller(request, callback) {
            return this.rpcCall(startNodeCaller, $root.raft.StartNodeRequest, $root.raft.Ok, request, callback);
        }, "name", { value: "StartNodeCaller" });

        /**
         * Calls StartNodeCaller.
         * @function startNodeCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IStartNodeRequest} request StartNodeRequest message or plain object
         * @returns {Promise<raft.Ok>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link raft.RaftRPC#appendEntriesCaller}.
         * @memberof raft.RaftRPC
         * @typedef AppendEntriesCallerCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {raft.AppendEntriesReply} [response] AppendEntriesReply
         */

        /**
         * Calls AppendEntriesCaller.
         * @function appendEntriesCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IAppendEntriesRequest} request AppendEntriesRequest message or plain object
         * @param {raft.RaftRPC.AppendEntriesCallerCallback} callback Node-style callback called with the error, if any, and AppendEntriesReply
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(RaftRPC.prototype.appendEntriesCaller = function appendEntriesCaller(request, callback) {
            return this.rpcCall(appendEntriesCaller, $root.raft.AppendEntriesRequest, $root.raft.AppendEntriesReply, request, callback);
        }, "name", { value: "AppendEntriesCaller" });

        /**
         * Calls AppendEntriesCaller.
         * @function appendEntriesCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IAppendEntriesRequest} request AppendEntriesRequest message or plain object
         * @returns {Promise<raft.AppendEntriesReply>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link raft.RaftRPC#requestVoteCaller}.
         * @memberof raft.RaftRPC
         * @typedef RequestVoteCallerCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {raft.RequestVoteReply} [response] RequestVoteReply
         */

        /**
         * Calls RequestVoteCaller.
         * @function requestVoteCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IRequestVoteRequest} request RequestVoteRequest message or plain object
         * @param {raft.RaftRPC.RequestVoteCallerCallback} callback Node-style callback called with the error, if any, and RequestVoteReply
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(RaftRPC.prototype.requestVoteCaller = function requestVoteCaller(request, callback) {
            return this.rpcCall(requestVoteCaller, $root.raft.RequestVoteRequest, $root.raft.RequestVoteReply, request, callback);
        }, "name", { value: "RequestVoteCaller" });

        /**
         * Calls RequestVoteCaller.
         * @function requestVoteCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IRequestVoteRequest} request RequestVoteRequest message or plain object
         * @returns {Promise<raft.RequestVoteReply>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link raft.RaftRPC#registerClientCaller}.
         * @memberof raft.RaftRPC
         * @typedef RegisterClientCallerCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {raft.RegisterClientReply} [response] RegisterClientReply
         */

        /**
         * Calls RegisterClientCaller.
         * @function registerClientCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IRegisterClientRequest} request RegisterClientRequest message or plain object
         * @param {raft.RaftRPC.RegisterClientCallerCallback} callback Node-style callback called with the error, if any, and RegisterClientReply
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(RaftRPC.prototype.registerClientCaller = function registerClientCaller(request, callback) {
            return this.rpcCall(registerClientCaller, $root.raft.RegisterClientRequest, $root.raft.RegisterClientReply, request, callback);
        }, "name", { value: "RegisterClientCaller" });

        /**
         * Calls RegisterClientCaller.
         * @function registerClientCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IRegisterClientRequest} request RegisterClientRequest message or plain object
         * @returns {Promise<raft.RegisterClientReply>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link raft.RaftRPC#clientRequestCaller}.
         * @memberof raft.RaftRPC
         * @typedef ClientRequestCallerCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {raft.ClientReply} [response] ClientReply
         */

        /**
         * Calls ClientRequestCaller.
         * @function clientRequestCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IClientRequest} request ClientRequest message or plain object
         * @param {raft.RaftRPC.ClientRequestCallerCallback} callback Node-style callback called with the error, if any, and ClientReply
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(RaftRPC.prototype.clientRequestCaller = function clientRequestCaller(request, callback) {
            return this.rpcCall(clientRequestCaller, $root.raft.ClientRequest, $root.raft.ClientReply, request, callback);
        }, "name", { value: "ClientRequestCaller" });

        /**
         * Calls ClientRequestCaller.
         * @function clientRequestCaller
         * @memberof raft.RaftRPC
         * @instance
         * @param {raft.IClientRequest} request ClientRequest message or plain object
         * @returns {Promise<raft.ClientReply>} Promise
         * @variation 2
         */

        return RaftRPC;
    })();

    raft.Ok = (function() {

        /**
         * Properties of an Ok.
         * @memberof raft
         * @interface IOk
         * @property {boolean|null} [ok] Ok ok
         * @property {string|null} [reason] Ok reason
         */

        /**
         * Constructs a new Ok.
         * @memberof raft
         * @classdesc Represents an Ok.
         * @implements IOk
         * @constructor
         * @param {raft.IOk=} [properties] Properties to set
         */
        function Ok(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Ok ok.
         * @member {boolean} ok
         * @memberof raft.Ok
         * @instance
         */
        Ok.prototype.ok = false;

        /**
         * Ok reason.
         * @member {string} reason
         * @memberof raft.Ok
         * @instance
         */
        Ok.prototype.reason = "";

        /**
         * Creates a new Ok instance using the specified properties.
         * @function create
         * @memberof raft.Ok
         * @static
         * @param {raft.IOk=} [properties] Properties to set
         * @returns {raft.Ok} Ok instance
         */
        Ok.create = function create(properties) {
            return new Ok(properties);
        };

        /**
         * Encodes the specified Ok message. Does not implicitly {@link raft.Ok.verify|verify} messages.
         * @function encode
         * @memberof raft.Ok
         * @static
         * @param {raft.IOk} message Ok message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Ok.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ok != null && message.hasOwnProperty("ok"))
                writer.uint32(/* id 1, wireType 0 =*/8).bool(message.ok);
            if (message.reason != null && message.hasOwnProperty("reason"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.reason);
            return writer;
        };

        /**
         * Encodes the specified Ok message, length delimited. Does not implicitly {@link raft.Ok.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.Ok
         * @static
         * @param {raft.IOk} message Ok message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Ok.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an Ok message from the specified reader or buffer.
         * @function decode
         * @memberof raft.Ok
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.Ok} Ok
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Ok.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.Ok();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ok = reader.bool();
                    break;
                case 2:
                    message.reason = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an Ok message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.Ok
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.Ok} Ok
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Ok.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an Ok message.
         * @function verify
         * @memberof raft.Ok
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Ok.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ok != null && message.hasOwnProperty("ok"))
                if (typeof message.ok !== "boolean")
                    return "ok: boolean expected";
            if (message.reason != null && message.hasOwnProperty("reason"))
                if (!$util.isString(message.reason))
                    return "reason: string expected";
            return null;
        };

        /**
         * Creates an Ok message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.Ok
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.Ok} Ok
         */
        Ok.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.Ok)
                return object;
            var message = new $root.raft.Ok();
            if (object.ok != null)
                message.ok = Boolean(object.ok);
            if (object.reason != null)
                message.reason = String(object.reason);
            return message;
        };

        /**
         * Creates a plain object from an Ok message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.Ok
         * @static
         * @param {raft.Ok} message Ok
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Ok.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.ok = false;
                object.reason = "";
            }
            if (message.ok != null && message.hasOwnProperty("ok"))
                object.ok = message.ok;
            if (message.reason != null && message.hasOwnProperty("reason"))
                object.reason = message.reason;
            return object;
        };

        /**
         * Converts this Ok to JSON.
         * @function toJSON
         * @memberof raft.Ok
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Ok.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Ok;
    })();

    raft.RemoteNode = (function() {

        /**
         * Properties of a RemoteNode.
         * @memberof raft
         * @interface IRemoteNode
         * @property {string|null} [addr] RemoteNode addr
         * @property {string|null} [id] RemoteNode id
         */

        /**
         * Constructs a new RemoteNode.
         * @memberof raft
         * @classdesc Represents a RemoteNode.
         * @implements IRemoteNode
         * @constructor
         * @param {raft.IRemoteNode=} [properties] Properties to set
         */
        function RemoteNode(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RemoteNode addr.
         * @member {string} addr
         * @memberof raft.RemoteNode
         * @instance
         */
        RemoteNode.prototype.addr = "";

        /**
         * RemoteNode id.
         * @member {string} id
         * @memberof raft.RemoteNode
         * @instance
         */
        RemoteNode.prototype.id = "";

        /**
         * Creates a new RemoteNode instance using the specified properties.
         * @function create
         * @memberof raft.RemoteNode
         * @static
         * @param {raft.IRemoteNode=} [properties] Properties to set
         * @returns {raft.RemoteNode} RemoteNode instance
         */
        RemoteNode.create = function create(properties) {
            return new RemoteNode(properties);
        };

        /**
         * Encodes the specified RemoteNode message. Does not implicitly {@link raft.RemoteNode.verify|verify} messages.
         * @function encode
         * @memberof raft.RemoteNode
         * @static
         * @param {raft.IRemoteNode} message RemoteNode message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RemoteNode.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.addr != null && message.hasOwnProperty("addr"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.addr);
            if (message.id != null && message.hasOwnProperty("id"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.id);
            return writer;
        };

        /**
         * Encodes the specified RemoteNode message, length delimited. Does not implicitly {@link raft.RemoteNode.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.RemoteNode
         * @static
         * @param {raft.IRemoteNode} message RemoteNode message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RemoteNode.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RemoteNode message from the specified reader or buffer.
         * @function decode
         * @memberof raft.RemoteNode
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.RemoteNode} RemoteNode
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RemoteNode.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.RemoteNode();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.addr = reader.string();
                    break;
                case 2:
                    message.id = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RemoteNode message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.RemoteNode
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.RemoteNode} RemoteNode
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RemoteNode.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RemoteNode message.
         * @function verify
         * @memberof raft.RemoteNode
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RemoteNode.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.addr != null && message.hasOwnProperty("addr"))
                if (!$util.isString(message.addr))
                    return "addr: string expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            return null;
        };

        /**
         * Creates a RemoteNode message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.RemoteNode
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.RemoteNode} RemoteNode
         */
        RemoteNode.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.RemoteNode)
                return object;
            var message = new $root.raft.RemoteNode();
            if (object.addr != null)
                message.addr = String(object.addr);
            if (object.id != null)
                message.id = String(object.id);
            return message;
        };

        /**
         * Creates a plain object from a RemoteNode message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.RemoteNode
         * @static
         * @param {raft.RemoteNode} message RemoteNode
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RemoteNode.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.addr = "";
                object.id = "";
            }
            if (message.addr != null && message.hasOwnProperty("addr"))
                object.addr = message.addr;
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            return object;
        };

        /**
         * Converts this RemoteNode to JSON.
         * @function toJSON
         * @memberof raft.RemoteNode
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RemoteNode.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RemoteNode;
    })();

    /**
     * CommandType enum.
     * @name raft.CommandType
     * @enum {string}
     * @property {number} CLIENT_REGISTRATION=0 CLIENT_REGISTRATION value
     * @property {number} INIT=1 INIT value
     * @property {number} NOOP=2 NOOP value
     * @property {number} STATE_MACHINE_COMMAND=3 STATE_MACHINE_COMMAND value
     */
    raft.CommandType = (function() {
        var valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "CLIENT_REGISTRATION"] = 0;
        values[valuesById[1] = "INIT"] = 1;
        values[valuesById[2] = "NOOP"] = 2;
        values[valuesById[3] = "STATE_MACHINE_COMMAND"] = 3;
        return values;
    })();

    raft.StartNodeRequest = (function() {

        /**
         * Properties of a StartNodeRequest.
         * @memberof raft
         * @interface IStartNodeRequest
         * @property {raft.IRemoteNode|null} [fromNode] StartNodeRequest fromNode
         * @property {Array.<raft.IRemoteNode>|null} [nodeList] StartNodeRequest nodeList
         */

        /**
         * Constructs a new StartNodeRequest.
         * @memberof raft
         * @classdesc Represents a StartNodeRequest.
         * @implements IStartNodeRequest
         * @constructor
         * @param {raft.IStartNodeRequest=} [properties] Properties to set
         */
        function StartNodeRequest(properties) {
            this.nodeList = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * StartNodeRequest fromNode.
         * @member {raft.IRemoteNode|null|undefined} fromNode
         * @memberof raft.StartNodeRequest
         * @instance
         */
        StartNodeRequest.prototype.fromNode = null;

        /**
         * StartNodeRequest nodeList.
         * @member {Array.<raft.IRemoteNode>} nodeList
         * @memberof raft.StartNodeRequest
         * @instance
         */
        StartNodeRequest.prototype.nodeList = $util.emptyArray;

        /**
         * Creates a new StartNodeRequest instance using the specified properties.
         * @function create
         * @memberof raft.StartNodeRequest
         * @static
         * @param {raft.IStartNodeRequest=} [properties] Properties to set
         * @returns {raft.StartNodeRequest} StartNodeRequest instance
         */
        StartNodeRequest.create = function create(properties) {
            return new StartNodeRequest(properties);
        };

        /**
         * Encodes the specified StartNodeRequest message. Does not implicitly {@link raft.StartNodeRequest.verify|verify} messages.
         * @function encode
         * @memberof raft.StartNodeRequest
         * @static
         * @param {raft.IStartNodeRequest} message StartNodeRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StartNodeRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.fromNode != null && message.hasOwnProperty("fromNode"))
                $root.raft.RemoteNode.encode(message.fromNode, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.nodeList != null && message.nodeList.length)
                for (var i = 0; i < message.nodeList.length; ++i)
                    $root.raft.RemoteNode.encode(message.nodeList[i], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified StartNodeRequest message, length delimited. Does not implicitly {@link raft.StartNodeRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.StartNodeRequest
         * @static
         * @param {raft.IStartNodeRequest} message StartNodeRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StartNodeRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a StartNodeRequest message from the specified reader or buffer.
         * @function decode
         * @memberof raft.StartNodeRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.StartNodeRequest} StartNodeRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        StartNodeRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.StartNodeRequest();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.fromNode = $root.raft.RemoteNode.decode(reader, reader.uint32());
                    break;
                case 2:
                    if (!(message.nodeList && message.nodeList.length))
                        message.nodeList = [];
                    message.nodeList.push($root.raft.RemoteNode.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a StartNodeRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.StartNodeRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.StartNodeRequest} StartNodeRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        StartNodeRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a StartNodeRequest message.
         * @function verify
         * @memberof raft.StartNodeRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        StartNodeRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.fromNode != null && message.hasOwnProperty("fromNode")) {
                var error = $root.raft.RemoteNode.verify(message.fromNode);
                if (error)
                    return "fromNode." + error;
            }
            if (message.nodeList != null && message.hasOwnProperty("nodeList")) {
                if (!Array.isArray(message.nodeList))
                    return "nodeList: array expected";
                for (var i = 0; i < message.nodeList.length; ++i) {
                    var error = $root.raft.RemoteNode.verify(message.nodeList[i]);
                    if (error)
                        return "nodeList." + error;
                }
            }
            return null;
        };

        /**
         * Creates a StartNodeRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.StartNodeRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.StartNodeRequest} StartNodeRequest
         */
        StartNodeRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.StartNodeRequest)
                return object;
            var message = new $root.raft.StartNodeRequest();
            if (object.fromNode != null) {
                if (typeof object.fromNode !== "object")
                    throw TypeError(".raft.StartNodeRequest.fromNode: object expected");
                message.fromNode = $root.raft.RemoteNode.fromObject(object.fromNode);
            }
            if (object.nodeList) {
                if (!Array.isArray(object.nodeList))
                    throw TypeError(".raft.StartNodeRequest.nodeList: array expected");
                message.nodeList = [];
                for (var i = 0; i < object.nodeList.length; ++i) {
                    if (typeof object.nodeList[i] !== "object")
                        throw TypeError(".raft.StartNodeRequest.nodeList: object expected");
                    message.nodeList[i] = $root.raft.RemoteNode.fromObject(object.nodeList[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a StartNodeRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.StartNodeRequest
         * @static
         * @param {raft.StartNodeRequest} message StartNodeRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StartNodeRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.nodeList = [];
            if (options.defaults)
                object.fromNode = null;
            if (message.fromNode != null && message.hasOwnProperty("fromNode"))
                object.fromNode = $root.raft.RemoteNode.toObject(message.fromNode, options);
            if (message.nodeList && message.nodeList.length) {
                object.nodeList = [];
                for (var j = 0; j < message.nodeList.length; ++j)
                    object.nodeList[j] = $root.raft.RemoteNode.toObject(message.nodeList[j], options);
            }
            return object;
        };

        /**
         * Converts this StartNodeRequest to JSON.
         * @function toJSON
         * @memberof raft.StartNodeRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        StartNodeRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return StartNodeRequest;
    })();

    raft.LogEntry = (function() {

        /**
         * Properties of a LogEntry.
         * @memberof raft
         * @interface ILogEntry
         * @property {number|Long|null} [index] LogEntry index
         * @property {number|Long|null} [termId] LogEntry termId
         * @property {raft.CommandType|null} [type] LogEntry type
         * @property {number|Long|null} [command] LogEntry command
         * @property {Uint8Array|null} [data] LogEntry data
         * @property {string|null} [cacheId] LogEntry cacheId
         */

        /**
         * Constructs a new LogEntry.
         * @memberof raft
         * @classdesc Represents a LogEntry.
         * @implements ILogEntry
         * @constructor
         * @param {raft.ILogEntry=} [properties] Properties to set
         */
        function LogEntry(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * LogEntry index.
         * @member {number|Long} index
         * @memberof raft.LogEntry
         * @instance
         */
        LogEntry.prototype.index = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LogEntry termId.
         * @member {number|Long} termId
         * @memberof raft.LogEntry
         * @instance
         */
        LogEntry.prototype.termId = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LogEntry type.
         * @member {raft.CommandType} type
         * @memberof raft.LogEntry
         * @instance
         */
        LogEntry.prototype.type = 0;

        /**
         * LogEntry command.
         * @member {number|Long} command
         * @memberof raft.LogEntry
         * @instance
         */
        LogEntry.prototype.command = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LogEntry data.
         * @member {Uint8Array} data
         * @memberof raft.LogEntry
         * @instance
         */
        LogEntry.prototype.data = $util.newBuffer([]);

        /**
         * LogEntry cacheId.
         * @member {string} cacheId
         * @memberof raft.LogEntry
         * @instance
         */
        LogEntry.prototype.cacheId = "";

        /**
         * Creates a new LogEntry instance using the specified properties.
         * @function create
         * @memberof raft.LogEntry
         * @static
         * @param {raft.ILogEntry=} [properties] Properties to set
         * @returns {raft.LogEntry} LogEntry instance
         */
        LogEntry.create = function create(properties) {
            return new LogEntry(properties);
        };

        /**
         * Encodes the specified LogEntry message. Does not implicitly {@link raft.LogEntry.verify|verify} messages.
         * @function encode
         * @memberof raft.LogEntry
         * @static
         * @param {raft.ILogEntry} message LogEntry message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        LogEntry.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.index != null && message.hasOwnProperty("index"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.index);
            if (message.termId != null && message.hasOwnProperty("termId"))
                writer.uint32(/* id 2, wireType 0 =*/16).uint64(message.termId);
            if (message.type != null && message.hasOwnProperty("type"))
                writer.uint32(/* id 3, wireType 0 =*/24).int32(message.type);
            if (message.command != null && message.hasOwnProperty("command"))
                writer.uint32(/* id 4, wireType 0 =*/32).uint64(message.command);
            if (message.data != null && message.hasOwnProperty("data"))
                writer.uint32(/* id 5, wireType 2 =*/42).bytes(message.data);
            if (message.cacheId != null && message.hasOwnProperty("cacheId"))
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.cacheId);
            return writer;
        };

        /**
         * Encodes the specified LogEntry message, length delimited. Does not implicitly {@link raft.LogEntry.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.LogEntry
         * @static
         * @param {raft.ILogEntry} message LogEntry message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        LogEntry.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a LogEntry message from the specified reader or buffer.
         * @function decode
         * @memberof raft.LogEntry
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.LogEntry} LogEntry
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        LogEntry.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.LogEntry();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.index = reader.uint64();
                    break;
                case 2:
                    message.termId = reader.uint64();
                    break;
                case 3:
                    message.type = reader.int32();
                    break;
                case 4:
                    message.command = reader.uint64();
                    break;
                case 5:
                    message.data = reader.bytes();
                    break;
                case 6:
                    message.cacheId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a LogEntry message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.LogEntry
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.LogEntry} LogEntry
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        LogEntry.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a LogEntry message.
         * @function verify
         * @memberof raft.LogEntry
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        LogEntry.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.index != null && message.hasOwnProperty("index"))
                if (!$util.isInteger(message.index) && !(message.index && $util.isInteger(message.index.low) && $util.isInteger(message.index.high)))
                    return "index: integer|Long expected";
            if (message.termId != null && message.hasOwnProperty("termId"))
                if (!$util.isInteger(message.termId) && !(message.termId && $util.isInteger(message.termId.low) && $util.isInteger(message.termId.high)))
                    return "termId: integer|Long expected";
            if (message.type != null && message.hasOwnProperty("type"))
                switch (message.type) {
                default:
                    return "type: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                    break;
                }
            if (message.command != null && message.hasOwnProperty("command"))
                if (!$util.isInteger(message.command) && !(message.command && $util.isInteger(message.command.low) && $util.isInteger(message.command.high)))
                    return "command: integer|Long expected";
            if (message.data != null && message.hasOwnProperty("data"))
                if (!(message.data && typeof message.data.length === "number" || $util.isString(message.data)))
                    return "data: buffer expected";
            if (message.cacheId != null && message.hasOwnProperty("cacheId"))
                if (!$util.isString(message.cacheId))
                    return "cacheId: string expected";
            return null;
        };

        /**
         * Creates a LogEntry message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.LogEntry
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.LogEntry} LogEntry
         */
        LogEntry.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.LogEntry)
                return object;
            var message = new $root.raft.LogEntry();
            if (object.index != null)
                if ($util.Long)
                    (message.index = $util.Long.fromValue(object.index)).unsigned = true;
                else if (typeof object.index === "string")
                    message.index = parseInt(object.index, 10);
                else if (typeof object.index === "number")
                    message.index = object.index;
                else if (typeof object.index === "object")
                    message.index = new $util.LongBits(object.index.low >>> 0, object.index.high >>> 0).toNumber(true);
            if (object.termId != null)
                if ($util.Long)
                    (message.termId = $util.Long.fromValue(object.termId)).unsigned = true;
                else if (typeof object.termId === "string")
                    message.termId = parseInt(object.termId, 10);
                else if (typeof object.termId === "number")
                    message.termId = object.termId;
                else if (typeof object.termId === "object")
                    message.termId = new $util.LongBits(object.termId.low >>> 0, object.termId.high >>> 0).toNumber(true);
            switch (object.type) {
            case "CLIENT_REGISTRATION":
            case 0:
                message.type = 0;
                break;
            case "INIT":
            case 1:
                message.type = 1;
                break;
            case "NOOP":
            case 2:
                message.type = 2;
                break;
            case "STATE_MACHINE_COMMAND":
            case 3:
                message.type = 3;
                break;
            }
            if (object.command != null)
                if ($util.Long)
                    (message.command = $util.Long.fromValue(object.command)).unsigned = true;
                else if (typeof object.command === "string")
                    message.command = parseInt(object.command, 10);
                else if (typeof object.command === "number")
                    message.command = object.command;
                else if (typeof object.command === "object")
                    message.command = new $util.LongBits(object.command.low >>> 0, object.command.high >>> 0).toNumber(true);
            if (object.data != null)
                if (typeof object.data === "string")
                    $util.base64.decode(object.data, message.data = $util.newBuffer($util.base64.length(object.data)), 0);
                else if (object.data.length)
                    message.data = object.data;
            if (object.cacheId != null)
                message.cacheId = String(object.cacheId);
            return message;
        };

        /**
         * Creates a plain object from a LogEntry message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.LogEntry
         * @static
         * @param {raft.LogEntry} message LogEntry
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        LogEntry.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.index = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.index = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.termId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.termId = options.longs === String ? "0" : 0;
                object.type = options.enums === String ? "CLIENT_REGISTRATION" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.command = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.command = options.longs === String ? "0" : 0;
                if (options.bytes === String)
                    object.data = "";
                else {
                    object.data = [];
                    if (options.bytes !== Array)
                        object.data = $util.newBuffer(object.data);
                }
                object.cacheId = "";
            }
            if (message.index != null && message.hasOwnProperty("index"))
                if (typeof message.index === "number")
                    object.index = options.longs === String ? String(message.index) : message.index;
                else
                    object.index = options.longs === String ? $util.Long.prototype.toString.call(message.index) : options.longs === Number ? new $util.LongBits(message.index.low >>> 0, message.index.high >>> 0).toNumber(true) : message.index;
            if (message.termId != null && message.hasOwnProperty("termId"))
                if (typeof message.termId === "number")
                    object.termId = options.longs === String ? String(message.termId) : message.termId;
                else
                    object.termId = options.longs === String ? $util.Long.prototype.toString.call(message.termId) : options.longs === Number ? new $util.LongBits(message.termId.low >>> 0, message.termId.high >>> 0).toNumber(true) : message.termId;
            if (message.type != null && message.hasOwnProperty("type"))
                object.type = options.enums === String ? $root.raft.CommandType[message.type] : message.type;
            if (message.command != null && message.hasOwnProperty("command"))
                if (typeof message.command === "number")
                    object.command = options.longs === String ? String(message.command) : message.command;
                else
                    object.command = options.longs === String ? $util.Long.prototype.toString.call(message.command) : options.longs === Number ? new $util.LongBits(message.command.low >>> 0, message.command.high >>> 0).toNumber(true) : message.command;
            if (message.data != null && message.hasOwnProperty("data"))
                object.data = options.bytes === String ? $util.base64.encode(message.data, 0, message.data.length) : options.bytes === Array ? Array.prototype.slice.call(message.data) : message.data;
            if (message.cacheId != null && message.hasOwnProperty("cacheId"))
                object.cacheId = message.cacheId;
            return object;
        };

        /**
         * Converts this LogEntry to JSON.
         * @function toJSON
         * @memberof raft.LogEntry
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        LogEntry.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return LogEntry;
    })();

    raft.AppendEntriesRequest = (function() {

        /**
         * Properties of an AppendEntriesRequest.
         * @memberof raft
         * @interface IAppendEntriesRequest
         * @property {number|Long|null} [term] AppendEntriesRequest term
         * @property {raft.IRemoteNode|null} [leader] AppendEntriesRequest leader
         * @property {number|Long|null} [prevLogIndex] AppendEntriesRequest prevLogIndex
         * @property {number|Long|null} [prevLogTerm] AppendEntriesRequest prevLogTerm
         * @property {Array.<raft.ILogEntry>|null} [entries] AppendEntriesRequest entries
         * @property {number|Long|null} [leaderCommit] AppendEntriesRequest leaderCommit
         */

        /**
         * Constructs a new AppendEntriesRequest.
         * @memberof raft
         * @classdesc Represents an AppendEntriesRequest.
         * @implements IAppendEntriesRequest
         * @constructor
         * @param {raft.IAppendEntriesRequest=} [properties] Properties to set
         */
        function AppendEntriesRequest(properties) {
            this.entries = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * AppendEntriesRequest term.
         * @member {number|Long} term
         * @memberof raft.AppendEntriesRequest
         * @instance
         */
        AppendEntriesRequest.prototype.term = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * AppendEntriesRequest leader.
         * @member {raft.IRemoteNode|null|undefined} leader
         * @memberof raft.AppendEntriesRequest
         * @instance
         */
        AppendEntriesRequest.prototype.leader = null;

        /**
         * AppendEntriesRequest prevLogIndex.
         * @member {number|Long} prevLogIndex
         * @memberof raft.AppendEntriesRequest
         * @instance
         */
        AppendEntriesRequest.prototype.prevLogIndex = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * AppendEntriesRequest prevLogTerm.
         * @member {number|Long} prevLogTerm
         * @memberof raft.AppendEntriesRequest
         * @instance
         */
        AppendEntriesRequest.prototype.prevLogTerm = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * AppendEntriesRequest entries.
         * @member {Array.<raft.ILogEntry>} entries
         * @memberof raft.AppendEntriesRequest
         * @instance
         */
        AppendEntriesRequest.prototype.entries = $util.emptyArray;

        /**
         * AppendEntriesRequest leaderCommit.
         * @member {number|Long} leaderCommit
         * @memberof raft.AppendEntriesRequest
         * @instance
         */
        AppendEntriesRequest.prototype.leaderCommit = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Creates a new AppendEntriesRequest instance using the specified properties.
         * @function create
         * @memberof raft.AppendEntriesRequest
         * @static
         * @param {raft.IAppendEntriesRequest=} [properties] Properties to set
         * @returns {raft.AppendEntriesRequest} AppendEntriesRequest instance
         */
        AppendEntriesRequest.create = function create(properties) {
            return new AppendEntriesRequest(properties);
        };

        /**
         * Encodes the specified AppendEntriesRequest message. Does not implicitly {@link raft.AppendEntriesRequest.verify|verify} messages.
         * @function encode
         * @memberof raft.AppendEntriesRequest
         * @static
         * @param {raft.IAppendEntriesRequest} message AppendEntriesRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AppendEntriesRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.term != null && message.hasOwnProperty("term"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.term);
            if (message.leader != null && message.hasOwnProperty("leader"))
                $root.raft.RemoteNode.encode(message.leader, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            if (message.prevLogIndex != null && message.hasOwnProperty("prevLogIndex"))
                writer.uint32(/* id 3, wireType 0 =*/24).uint64(message.prevLogIndex);
            if (message.prevLogTerm != null && message.hasOwnProperty("prevLogTerm"))
                writer.uint32(/* id 4, wireType 0 =*/32).uint64(message.prevLogTerm);
            if (message.entries != null && message.entries.length)
                for (var i = 0; i < message.entries.length; ++i)
                    $root.raft.LogEntry.encode(message.entries[i], writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
            if (message.leaderCommit != null && message.hasOwnProperty("leaderCommit"))
                writer.uint32(/* id 6, wireType 0 =*/48).uint64(message.leaderCommit);
            return writer;
        };

        /**
         * Encodes the specified AppendEntriesRequest message, length delimited. Does not implicitly {@link raft.AppendEntriesRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.AppendEntriesRequest
         * @static
         * @param {raft.IAppendEntriesRequest} message AppendEntriesRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AppendEntriesRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an AppendEntriesRequest message from the specified reader or buffer.
         * @function decode
         * @memberof raft.AppendEntriesRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.AppendEntriesRequest} AppendEntriesRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AppendEntriesRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.AppendEntriesRequest();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.term = reader.uint64();
                    break;
                case 2:
                    message.leader = $root.raft.RemoteNode.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.prevLogIndex = reader.uint64();
                    break;
                case 4:
                    message.prevLogTerm = reader.uint64();
                    break;
                case 5:
                    if (!(message.entries && message.entries.length))
                        message.entries = [];
                    message.entries.push($root.raft.LogEntry.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.leaderCommit = reader.uint64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an AppendEntriesRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.AppendEntriesRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.AppendEntriesRequest} AppendEntriesRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AppendEntriesRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an AppendEntriesRequest message.
         * @function verify
         * @memberof raft.AppendEntriesRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        AppendEntriesRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.term != null && message.hasOwnProperty("term"))
                if (!$util.isInteger(message.term) && !(message.term && $util.isInteger(message.term.low) && $util.isInteger(message.term.high)))
                    return "term: integer|Long expected";
            if (message.leader != null && message.hasOwnProperty("leader")) {
                var error = $root.raft.RemoteNode.verify(message.leader);
                if (error)
                    return "leader." + error;
            }
            if (message.prevLogIndex != null && message.hasOwnProperty("prevLogIndex"))
                if (!$util.isInteger(message.prevLogIndex) && !(message.prevLogIndex && $util.isInteger(message.prevLogIndex.low) && $util.isInteger(message.prevLogIndex.high)))
                    return "prevLogIndex: integer|Long expected";
            if (message.prevLogTerm != null && message.hasOwnProperty("prevLogTerm"))
                if (!$util.isInteger(message.prevLogTerm) && !(message.prevLogTerm && $util.isInteger(message.prevLogTerm.low) && $util.isInteger(message.prevLogTerm.high)))
                    return "prevLogTerm: integer|Long expected";
            if (message.entries != null && message.hasOwnProperty("entries")) {
                if (!Array.isArray(message.entries))
                    return "entries: array expected";
                for (var i = 0; i < message.entries.length; ++i) {
                    var error = $root.raft.LogEntry.verify(message.entries[i]);
                    if (error)
                        return "entries." + error;
                }
            }
            if (message.leaderCommit != null && message.hasOwnProperty("leaderCommit"))
                if (!$util.isInteger(message.leaderCommit) && !(message.leaderCommit && $util.isInteger(message.leaderCommit.low) && $util.isInteger(message.leaderCommit.high)))
                    return "leaderCommit: integer|Long expected";
            return null;
        };

        /**
         * Creates an AppendEntriesRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.AppendEntriesRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.AppendEntriesRequest} AppendEntriesRequest
         */
        AppendEntriesRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.AppendEntriesRequest)
                return object;
            var message = new $root.raft.AppendEntriesRequest();
            if (object.term != null)
                if ($util.Long)
                    (message.term = $util.Long.fromValue(object.term)).unsigned = true;
                else if (typeof object.term === "string")
                    message.term = parseInt(object.term, 10);
                else if (typeof object.term === "number")
                    message.term = object.term;
                else if (typeof object.term === "object")
                    message.term = new $util.LongBits(object.term.low >>> 0, object.term.high >>> 0).toNumber(true);
            if (object.leader != null) {
                if (typeof object.leader !== "object")
                    throw TypeError(".raft.AppendEntriesRequest.leader: object expected");
                message.leader = $root.raft.RemoteNode.fromObject(object.leader);
            }
            if (object.prevLogIndex != null)
                if ($util.Long)
                    (message.prevLogIndex = $util.Long.fromValue(object.prevLogIndex)).unsigned = true;
                else if (typeof object.prevLogIndex === "string")
                    message.prevLogIndex = parseInt(object.prevLogIndex, 10);
                else if (typeof object.prevLogIndex === "number")
                    message.prevLogIndex = object.prevLogIndex;
                else if (typeof object.prevLogIndex === "object")
                    message.prevLogIndex = new $util.LongBits(object.prevLogIndex.low >>> 0, object.prevLogIndex.high >>> 0).toNumber(true);
            if (object.prevLogTerm != null)
                if ($util.Long)
                    (message.prevLogTerm = $util.Long.fromValue(object.prevLogTerm)).unsigned = true;
                else if (typeof object.prevLogTerm === "string")
                    message.prevLogTerm = parseInt(object.prevLogTerm, 10);
                else if (typeof object.prevLogTerm === "number")
                    message.prevLogTerm = object.prevLogTerm;
                else if (typeof object.prevLogTerm === "object")
                    message.prevLogTerm = new $util.LongBits(object.prevLogTerm.low >>> 0, object.prevLogTerm.high >>> 0).toNumber(true);
            if (object.entries) {
                if (!Array.isArray(object.entries))
                    throw TypeError(".raft.AppendEntriesRequest.entries: array expected");
                message.entries = [];
                for (var i = 0; i < object.entries.length; ++i) {
                    if (typeof object.entries[i] !== "object")
                        throw TypeError(".raft.AppendEntriesRequest.entries: object expected");
                    message.entries[i] = $root.raft.LogEntry.fromObject(object.entries[i]);
                }
            }
            if (object.leaderCommit != null)
                if ($util.Long)
                    (message.leaderCommit = $util.Long.fromValue(object.leaderCommit)).unsigned = true;
                else if (typeof object.leaderCommit === "string")
                    message.leaderCommit = parseInt(object.leaderCommit, 10);
                else if (typeof object.leaderCommit === "number")
                    message.leaderCommit = object.leaderCommit;
                else if (typeof object.leaderCommit === "object")
                    message.leaderCommit = new $util.LongBits(object.leaderCommit.low >>> 0, object.leaderCommit.high >>> 0).toNumber(true);
            return message;
        };

        /**
         * Creates a plain object from an AppendEntriesRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.AppendEntriesRequest
         * @static
         * @param {raft.AppendEntriesRequest} message AppendEntriesRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AppendEntriesRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.entries = [];
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.term = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.term = options.longs === String ? "0" : 0;
                object.leader = null;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.prevLogIndex = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.prevLogIndex = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.prevLogTerm = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.prevLogTerm = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.leaderCommit = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.leaderCommit = options.longs === String ? "0" : 0;
            }
            if (message.term != null && message.hasOwnProperty("term"))
                if (typeof message.term === "number")
                    object.term = options.longs === String ? String(message.term) : message.term;
                else
                    object.term = options.longs === String ? $util.Long.prototype.toString.call(message.term) : options.longs === Number ? new $util.LongBits(message.term.low >>> 0, message.term.high >>> 0).toNumber(true) : message.term;
            if (message.leader != null && message.hasOwnProperty("leader"))
                object.leader = $root.raft.RemoteNode.toObject(message.leader, options);
            if (message.prevLogIndex != null && message.hasOwnProperty("prevLogIndex"))
                if (typeof message.prevLogIndex === "number")
                    object.prevLogIndex = options.longs === String ? String(message.prevLogIndex) : message.prevLogIndex;
                else
                    object.prevLogIndex = options.longs === String ? $util.Long.prototype.toString.call(message.prevLogIndex) : options.longs === Number ? new $util.LongBits(message.prevLogIndex.low >>> 0, message.prevLogIndex.high >>> 0).toNumber(true) : message.prevLogIndex;
            if (message.prevLogTerm != null && message.hasOwnProperty("prevLogTerm"))
                if (typeof message.prevLogTerm === "number")
                    object.prevLogTerm = options.longs === String ? String(message.prevLogTerm) : message.prevLogTerm;
                else
                    object.prevLogTerm = options.longs === String ? $util.Long.prototype.toString.call(message.prevLogTerm) : options.longs === Number ? new $util.LongBits(message.prevLogTerm.low >>> 0, message.prevLogTerm.high >>> 0).toNumber(true) : message.prevLogTerm;
            if (message.entries && message.entries.length) {
                object.entries = [];
                for (var j = 0; j < message.entries.length; ++j)
                    object.entries[j] = $root.raft.LogEntry.toObject(message.entries[j], options);
            }
            if (message.leaderCommit != null && message.hasOwnProperty("leaderCommit"))
                if (typeof message.leaderCommit === "number")
                    object.leaderCommit = options.longs === String ? String(message.leaderCommit) : message.leaderCommit;
                else
                    object.leaderCommit = options.longs === String ? $util.Long.prototype.toString.call(message.leaderCommit) : options.longs === Number ? new $util.LongBits(message.leaderCommit.low >>> 0, message.leaderCommit.high >>> 0).toNumber(true) : message.leaderCommit;
            return object;
        };

        /**
         * Converts this AppendEntriesRequest to JSON.
         * @function toJSON
         * @memberof raft.AppendEntriesRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        AppendEntriesRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return AppendEntriesRequest;
    })();

    raft.AppendEntriesReply = (function() {

        /**
         * Properties of an AppendEntriesReply.
         * @memberof raft
         * @interface IAppendEntriesReply
         * @property {number|Long|null} [term] AppendEntriesReply term
         * @property {boolean|null} [success] AppendEntriesReply success
         */

        /**
         * Constructs a new AppendEntriesReply.
         * @memberof raft
         * @classdesc Represents an AppendEntriesReply.
         * @implements IAppendEntriesReply
         * @constructor
         * @param {raft.IAppendEntriesReply=} [properties] Properties to set
         */
        function AppendEntriesReply(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * AppendEntriesReply term.
         * @member {number|Long} term
         * @memberof raft.AppendEntriesReply
         * @instance
         */
        AppendEntriesReply.prototype.term = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * AppendEntriesReply success.
         * @member {boolean} success
         * @memberof raft.AppendEntriesReply
         * @instance
         */
        AppendEntriesReply.prototype.success = false;

        /**
         * Creates a new AppendEntriesReply instance using the specified properties.
         * @function create
         * @memberof raft.AppendEntriesReply
         * @static
         * @param {raft.IAppendEntriesReply=} [properties] Properties to set
         * @returns {raft.AppendEntriesReply} AppendEntriesReply instance
         */
        AppendEntriesReply.create = function create(properties) {
            return new AppendEntriesReply(properties);
        };

        /**
         * Encodes the specified AppendEntriesReply message. Does not implicitly {@link raft.AppendEntriesReply.verify|verify} messages.
         * @function encode
         * @memberof raft.AppendEntriesReply
         * @static
         * @param {raft.IAppendEntriesReply} message AppendEntriesReply message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AppendEntriesReply.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.term != null && message.hasOwnProperty("term"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.term);
            if (message.success != null && message.hasOwnProperty("success"))
                writer.uint32(/* id 2, wireType 0 =*/16).bool(message.success);
            return writer;
        };

        /**
         * Encodes the specified AppendEntriesReply message, length delimited. Does not implicitly {@link raft.AppendEntriesReply.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.AppendEntriesReply
         * @static
         * @param {raft.IAppendEntriesReply} message AppendEntriesReply message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AppendEntriesReply.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an AppendEntriesReply message from the specified reader or buffer.
         * @function decode
         * @memberof raft.AppendEntriesReply
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.AppendEntriesReply} AppendEntriesReply
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AppendEntriesReply.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.AppendEntriesReply();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.term = reader.uint64();
                    break;
                case 2:
                    message.success = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an AppendEntriesReply message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.AppendEntriesReply
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.AppendEntriesReply} AppendEntriesReply
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AppendEntriesReply.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an AppendEntriesReply message.
         * @function verify
         * @memberof raft.AppendEntriesReply
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        AppendEntriesReply.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.term != null && message.hasOwnProperty("term"))
                if (!$util.isInteger(message.term) && !(message.term && $util.isInteger(message.term.low) && $util.isInteger(message.term.high)))
                    return "term: integer|Long expected";
            if (message.success != null && message.hasOwnProperty("success"))
                if (typeof message.success !== "boolean")
                    return "success: boolean expected";
            return null;
        };

        /**
         * Creates an AppendEntriesReply message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.AppendEntriesReply
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.AppendEntriesReply} AppendEntriesReply
         */
        AppendEntriesReply.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.AppendEntriesReply)
                return object;
            var message = new $root.raft.AppendEntriesReply();
            if (object.term != null)
                if ($util.Long)
                    (message.term = $util.Long.fromValue(object.term)).unsigned = true;
                else if (typeof object.term === "string")
                    message.term = parseInt(object.term, 10);
                else if (typeof object.term === "number")
                    message.term = object.term;
                else if (typeof object.term === "object")
                    message.term = new $util.LongBits(object.term.low >>> 0, object.term.high >>> 0).toNumber(true);
            if (object.success != null)
                message.success = Boolean(object.success);
            return message;
        };

        /**
         * Creates a plain object from an AppendEntriesReply message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.AppendEntriesReply
         * @static
         * @param {raft.AppendEntriesReply} message AppendEntriesReply
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AppendEntriesReply.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.term = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.term = options.longs === String ? "0" : 0;
                object.success = false;
            }
            if (message.term != null && message.hasOwnProperty("term"))
                if (typeof message.term === "number")
                    object.term = options.longs === String ? String(message.term) : message.term;
                else
                    object.term = options.longs === String ? $util.Long.prototype.toString.call(message.term) : options.longs === Number ? new $util.LongBits(message.term.low >>> 0, message.term.high >>> 0).toNumber(true) : message.term;
            if (message.success != null && message.hasOwnProperty("success"))
                object.success = message.success;
            return object;
        };

        /**
         * Converts this AppendEntriesReply to JSON.
         * @function toJSON
         * @memberof raft.AppendEntriesReply
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        AppendEntriesReply.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return AppendEntriesReply;
    })();

    raft.RequestVoteRequest = (function() {

        /**
         * Properties of a RequestVoteRequest.
         * @memberof raft
         * @interface IRequestVoteRequest
         * @property {number|Long|null} [term] RequestVoteRequest term
         * @property {raft.IRemoteNode|null} [candidate] RequestVoteRequest candidate
         * @property {number|Long|null} [lastLogIndex] RequestVoteRequest lastLogIndex
         * @property {number|Long|null} [lastLogTerm] RequestVoteRequest lastLogTerm
         */

        /**
         * Constructs a new RequestVoteRequest.
         * @memberof raft
         * @classdesc Represents a RequestVoteRequest.
         * @implements IRequestVoteRequest
         * @constructor
         * @param {raft.IRequestVoteRequest=} [properties] Properties to set
         */
        function RequestVoteRequest(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RequestVoteRequest term.
         * @member {number|Long} term
         * @memberof raft.RequestVoteRequest
         * @instance
         */
        RequestVoteRequest.prototype.term = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * RequestVoteRequest candidate.
         * @member {raft.IRemoteNode|null|undefined} candidate
         * @memberof raft.RequestVoteRequest
         * @instance
         */
        RequestVoteRequest.prototype.candidate = null;

        /**
         * RequestVoteRequest lastLogIndex.
         * @member {number|Long} lastLogIndex
         * @memberof raft.RequestVoteRequest
         * @instance
         */
        RequestVoteRequest.prototype.lastLogIndex = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * RequestVoteRequest lastLogTerm.
         * @member {number|Long} lastLogTerm
         * @memberof raft.RequestVoteRequest
         * @instance
         */
        RequestVoteRequest.prototype.lastLogTerm = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Creates a new RequestVoteRequest instance using the specified properties.
         * @function create
         * @memberof raft.RequestVoteRequest
         * @static
         * @param {raft.IRequestVoteRequest=} [properties] Properties to set
         * @returns {raft.RequestVoteRequest} RequestVoteRequest instance
         */
        RequestVoteRequest.create = function create(properties) {
            return new RequestVoteRequest(properties);
        };

        /**
         * Encodes the specified RequestVoteRequest message. Does not implicitly {@link raft.RequestVoteRequest.verify|verify} messages.
         * @function encode
         * @memberof raft.RequestVoteRequest
         * @static
         * @param {raft.IRequestVoteRequest} message RequestVoteRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RequestVoteRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.term != null && message.hasOwnProperty("term"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.term);
            if (message.candidate != null && message.hasOwnProperty("candidate"))
                $root.raft.RemoteNode.encode(message.candidate, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            if (message.lastLogIndex != null && message.hasOwnProperty("lastLogIndex"))
                writer.uint32(/* id 3, wireType 0 =*/24).uint64(message.lastLogIndex);
            if (message.lastLogTerm != null && message.hasOwnProperty("lastLogTerm"))
                writer.uint32(/* id 4, wireType 0 =*/32).uint64(message.lastLogTerm);
            return writer;
        };

        /**
         * Encodes the specified RequestVoteRequest message, length delimited. Does not implicitly {@link raft.RequestVoteRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.RequestVoteRequest
         * @static
         * @param {raft.IRequestVoteRequest} message RequestVoteRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RequestVoteRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RequestVoteRequest message from the specified reader or buffer.
         * @function decode
         * @memberof raft.RequestVoteRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.RequestVoteRequest} RequestVoteRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RequestVoteRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.RequestVoteRequest();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.term = reader.uint64();
                    break;
                case 2:
                    message.candidate = $root.raft.RemoteNode.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.lastLogIndex = reader.uint64();
                    break;
                case 4:
                    message.lastLogTerm = reader.uint64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RequestVoteRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.RequestVoteRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.RequestVoteRequest} RequestVoteRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RequestVoteRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RequestVoteRequest message.
         * @function verify
         * @memberof raft.RequestVoteRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RequestVoteRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.term != null && message.hasOwnProperty("term"))
                if (!$util.isInteger(message.term) && !(message.term && $util.isInteger(message.term.low) && $util.isInteger(message.term.high)))
                    return "term: integer|Long expected";
            if (message.candidate != null && message.hasOwnProperty("candidate")) {
                var error = $root.raft.RemoteNode.verify(message.candidate);
                if (error)
                    return "candidate." + error;
            }
            if (message.lastLogIndex != null && message.hasOwnProperty("lastLogIndex"))
                if (!$util.isInteger(message.lastLogIndex) && !(message.lastLogIndex && $util.isInteger(message.lastLogIndex.low) && $util.isInteger(message.lastLogIndex.high)))
                    return "lastLogIndex: integer|Long expected";
            if (message.lastLogTerm != null && message.hasOwnProperty("lastLogTerm"))
                if (!$util.isInteger(message.lastLogTerm) && !(message.lastLogTerm && $util.isInteger(message.lastLogTerm.low) && $util.isInteger(message.lastLogTerm.high)))
                    return "lastLogTerm: integer|Long expected";
            return null;
        };

        /**
         * Creates a RequestVoteRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.RequestVoteRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.RequestVoteRequest} RequestVoteRequest
         */
        RequestVoteRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.RequestVoteRequest)
                return object;
            var message = new $root.raft.RequestVoteRequest();
            if (object.term != null)
                if ($util.Long)
                    (message.term = $util.Long.fromValue(object.term)).unsigned = true;
                else if (typeof object.term === "string")
                    message.term = parseInt(object.term, 10);
                else if (typeof object.term === "number")
                    message.term = object.term;
                else if (typeof object.term === "object")
                    message.term = new $util.LongBits(object.term.low >>> 0, object.term.high >>> 0).toNumber(true);
            if (object.candidate != null) {
                if (typeof object.candidate !== "object")
                    throw TypeError(".raft.RequestVoteRequest.candidate: object expected");
                message.candidate = $root.raft.RemoteNode.fromObject(object.candidate);
            }
            if (object.lastLogIndex != null)
                if ($util.Long)
                    (message.lastLogIndex = $util.Long.fromValue(object.lastLogIndex)).unsigned = true;
                else if (typeof object.lastLogIndex === "string")
                    message.lastLogIndex = parseInt(object.lastLogIndex, 10);
                else if (typeof object.lastLogIndex === "number")
                    message.lastLogIndex = object.lastLogIndex;
                else if (typeof object.lastLogIndex === "object")
                    message.lastLogIndex = new $util.LongBits(object.lastLogIndex.low >>> 0, object.lastLogIndex.high >>> 0).toNumber(true);
            if (object.lastLogTerm != null)
                if ($util.Long)
                    (message.lastLogTerm = $util.Long.fromValue(object.lastLogTerm)).unsigned = true;
                else if (typeof object.lastLogTerm === "string")
                    message.lastLogTerm = parseInt(object.lastLogTerm, 10);
                else if (typeof object.lastLogTerm === "number")
                    message.lastLogTerm = object.lastLogTerm;
                else if (typeof object.lastLogTerm === "object")
                    message.lastLogTerm = new $util.LongBits(object.lastLogTerm.low >>> 0, object.lastLogTerm.high >>> 0).toNumber(true);
            return message;
        };

        /**
         * Creates a plain object from a RequestVoteRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.RequestVoteRequest
         * @static
         * @param {raft.RequestVoteRequest} message RequestVoteRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RequestVoteRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.term = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.term = options.longs === String ? "0" : 0;
                object.candidate = null;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.lastLogIndex = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.lastLogIndex = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.lastLogTerm = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.lastLogTerm = options.longs === String ? "0" : 0;
            }
            if (message.term != null && message.hasOwnProperty("term"))
                if (typeof message.term === "number")
                    object.term = options.longs === String ? String(message.term) : message.term;
                else
                    object.term = options.longs === String ? $util.Long.prototype.toString.call(message.term) : options.longs === Number ? new $util.LongBits(message.term.low >>> 0, message.term.high >>> 0).toNumber(true) : message.term;
            if (message.candidate != null && message.hasOwnProperty("candidate"))
                object.candidate = $root.raft.RemoteNode.toObject(message.candidate, options);
            if (message.lastLogIndex != null && message.hasOwnProperty("lastLogIndex"))
                if (typeof message.lastLogIndex === "number")
                    object.lastLogIndex = options.longs === String ? String(message.lastLogIndex) : message.lastLogIndex;
                else
                    object.lastLogIndex = options.longs === String ? $util.Long.prototype.toString.call(message.lastLogIndex) : options.longs === Number ? new $util.LongBits(message.lastLogIndex.low >>> 0, message.lastLogIndex.high >>> 0).toNumber(true) : message.lastLogIndex;
            if (message.lastLogTerm != null && message.hasOwnProperty("lastLogTerm"))
                if (typeof message.lastLogTerm === "number")
                    object.lastLogTerm = options.longs === String ? String(message.lastLogTerm) : message.lastLogTerm;
                else
                    object.lastLogTerm = options.longs === String ? $util.Long.prototype.toString.call(message.lastLogTerm) : options.longs === Number ? new $util.LongBits(message.lastLogTerm.low >>> 0, message.lastLogTerm.high >>> 0).toNumber(true) : message.lastLogTerm;
            return object;
        };

        /**
         * Converts this RequestVoteRequest to JSON.
         * @function toJSON
         * @memberof raft.RequestVoteRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RequestVoteRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RequestVoteRequest;
    })();

    raft.RequestVoteReply = (function() {

        /**
         * Properties of a RequestVoteReply.
         * @memberof raft
         * @interface IRequestVoteReply
         * @property {number|Long|null} [term] RequestVoteReply term
         * @property {boolean|null} [voteGranted] RequestVoteReply voteGranted
         */

        /**
         * Constructs a new RequestVoteReply.
         * @memberof raft
         * @classdesc Represents a RequestVoteReply.
         * @implements IRequestVoteReply
         * @constructor
         * @param {raft.IRequestVoteReply=} [properties] Properties to set
         */
        function RequestVoteReply(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RequestVoteReply term.
         * @member {number|Long} term
         * @memberof raft.RequestVoteReply
         * @instance
         */
        RequestVoteReply.prototype.term = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * RequestVoteReply voteGranted.
         * @member {boolean} voteGranted
         * @memberof raft.RequestVoteReply
         * @instance
         */
        RequestVoteReply.prototype.voteGranted = false;

        /**
         * Creates a new RequestVoteReply instance using the specified properties.
         * @function create
         * @memberof raft.RequestVoteReply
         * @static
         * @param {raft.IRequestVoteReply=} [properties] Properties to set
         * @returns {raft.RequestVoteReply} RequestVoteReply instance
         */
        RequestVoteReply.create = function create(properties) {
            return new RequestVoteReply(properties);
        };

        /**
         * Encodes the specified RequestVoteReply message. Does not implicitly {@link raft.RequestVoteReply.verify|verify} messages.
         * @function encode
         * @memberof raft.RequestVoteReply
         * @static
         * @param {raft.IRequestVoteReply} message RequestVoteReply message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RequestVoteReply.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.term != null && message.hasOwnProperty("term"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.term);
            if (message.voteGranted != null && message.hasOwnProperty("voteGranted"))
                writer.uint32(/* id 2, wireType 0 =*/16).bool(message.voteGranted);
            return writer;
        };

        /**
         * Encodes the specified RequestVoteReply message, length delimited. Does not implicitly {@link raft.RequestVoteReply.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.RequestVoteReply
         * @static
         * @param {raft.IRequestVoteReply} message RequestVoteReply message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RequestVoteReply.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RequestVoteReply message from the specified reader or buffer.
         * @function decode
         * @memberof raft.RequestVoteReply
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.RequestVoteReply} RequestVoteReply
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RequestVoteReply.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.RequestVoteReply();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.term = reader.uint64();
                    break;
                case 2:
                    message.voteGranted = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RequestVoteReply message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.RequestVoteReply
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.RequestVoteReply} RequestVoteReply
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RequestVoteReply.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RequestVoteReply message.
         * @function verify
         * @memberof raft.RequestVoteReply
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RequestVoteReply.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.term != null && message.hasOwnProperty("term"))
                if (!$util.isInteger(message.term) && !(message.term && $util.isInteger(message.term.low) && $util.isInteger(message.term.high)))
                    return "term: integer|Long expected";
            if (message.voteGranted != null && message.hasOwnProperty("voteGranted"))
                if (typeof message.voteGranted !== "boolean")
                    return "voteGranted: boolean expected";
            return null;
        };

        /**
         * Creates a RequestVoteReply message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.RequestVoteReply
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.RequestVoteReply} RequestVoteReply
         */
        RequestVoteReply.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.RequestVoteReply)
                return object;
            var message = new $root.raft.RequestVoteReply();
            if (object.term != null)
                if ($util.Long)
                    (message.term = $util.Long.fromValue(object.term)).unsigned = true;
                else if (typeof object.term === "string")
                    message.term = parseInt(object.term, 10);
                else if (typeof object.term === "number")
                    message.term = object.term;
                else if (typeof object.term === "object")
                    message.term = new $util.LongBits(object.term.low >>> 0, object.term.high >>> 0).toNumber(true);
            if (object.voteGranted != null)
                message.voteGranted = Boolean(object.voteGranted);
            return message;
        };

        /**
         * Creates a plain object from a RequestVoteReply message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.RequestVoteReply
         * @static
         * @param {raft.RequestVoteReply} message RequestVoteReply
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RequestVoteReply.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.term = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.term = options.longs === String ? "0" : 0;
                object.voteGranted = false;
            }
            if (message.term != null && message.hasOwnProperty("term"))
                if (typeof message.term === "number")
                    object.term = options.longs === String ? String(message.term) : message.term;
                else
                    object.term = options.longs === String ? $util.Long.prototype.toString.call(message.term) : options.longs === Number ? new $util.LongBits(message.term.low >>> 0, message.term.high >>> 0).toNumber(true) : message.term;
            if (message.voteGranted != null && message.hasOwnProperty("voteGranted"))
                object.voteGranted = message.voteGranted;
            return object;
        };

        /**
         * Converts this RequestVoteReply to JSON.
         * @function toJSON
         * @memberof raft.RequestVoteReply
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RequestVoteReply.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RequestVoteReply;
    })();

    /**
     * ClientStatus enum.
     * @name raft.ClientStatus
     * @enum {string}
     * @property {number} OK=0 OK value
     * @property {number} NOT_LEADER=1 NOT_LEADER value
     * @property {number} ELECTION_IN_PROGRESS=2 ELECTION_IN_PROGRESS value
     * @property {number} CLUSTER_NOT_STARTED=3 CLUSTER_NOT_STARTED value
     * @property {number} REQ_FAILED=4 REQ_FAILED value
     */
    raft.ClientStatus = (function() {
        var valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "OK"] = 0;
        values[valuesById[1] = "NOT_LEADER"] = 1;
        values[valuesById[2] = "ELECTION_IN_PROGRESS"] = 2;
        values[valuesById[3] = "CLUSTER_NOT_STARTED"] = 3;
        values[valuesById[4] = "REQ_FAILED"] = 4;
        return values;
    })();

    raft.RegisterClientRequest = (function() {

        /**
         * Properties of a RegisterClientRequest.
         * @memberof raft
         * @interface IRegisterClientRequest
         */

        /**
         * Constructs a new RegisterClientRequest.
         * @memberof raft
         * @classdesc Represents a RegisterClientRequest.
         * @implements IRegisterClientRequest
         * @constructor
         * @param {raft.IRegisterClientRequest=} [properties] Properties to set
         */
        function RegisterClientRequest(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new RegisterClientRequest instance using the specified properties.
         * @function create
         * @memberof raft.RegisterClientRequest
         * @static
         * @param {raft.IRegisterClientRequest=} [properties] Properties to set
         * @returns {raft.RegisterClientRequest} RegisterClientRequest instance
         */
        RegisterClientRequest.create = function create(properties) {
            return new RegisterClientRequest(properties);
        };

        /**
         * Encodes the specified RegisterClientRequest message. Does not implicitly {@link raft.RegisterClientRequest.verify|verify} messages.
         * @function encode
         * @memberof raft.RegisterClientRequest
         * @static
         * @param {raft.IRegisterClientRequest} message RegisterClientRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegisterClientRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified RegisterClientRequest message, length delimited. Does not implicitly {@link raft.RegisterClientRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.RegisterClientRequest
         * @static
         * @param {raft.IRegisterClientRequest} message RegisterClientRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegisterClientRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegisterClientRequest message from the specified reader or buffer.
         * @function decode
         * @memberof raft.RegisterClientRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.RegisterClientRequest} RegisterClientRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegisterClientRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.RegisterClientRequest();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegisterClientRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.RegisterClientRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.RegisterClientRequest} RegisterClientRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegisterClientRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegisterClientRequest message.
         * @function verify
         * @memberof raft.RegisterClientRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegisterClientRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a RegisterClientRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.RegisterClientRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.RegisterClientRequest} RegisterClientRequest
         */
        RegisterClientRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.RegisterClientRequest)
                return object;
            return new $root.raft.RegisterClientRequest();
        };

        /**
         * Creates a plain object from a RegisterClientRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.RegisterClientRequest
         * @static
         * @param {raft.RegisterClientRequest} message RegisterClientRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegisterClientRequest.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this RegisterClientRequest to JSON.
         * @function toJSON
         * @memberof raft.RegisterClientRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegisterClientRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegisterClientRequest;
    })();

    raft.RegisterClientReply = (function() {

        /**
         * Properties of a RegisterClientReply.
         * @memberof raft
         * @interface IRegisterClientReply
         * @property {raft.ClientStatus|null} [status] RegisterClientReply status
         * @property {number|Long|null} [clientId] RegisterClientReply clientId
         * @property {raft.IRemoteNode|null} [leaderHint] RegisterClientReply leaderHint
         */

        /**
         * Constructs a new RegisterClientReply.
         * @memberof raft
         * @classdesc Represents a RegisterClientReply.
         * @implements IRegisterClientReply
         * @constructor
         * @param {raft.IRegisterClientReply=} [properties] Properties to set
         */
        function RegisterClientReply(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegisterClientReply status.
         * @member {raft.ClientStatus} status
         * @memberof raft.RegisterClientReply
         * @instance
         */
        RegisterClientReply.prototype.status = 0;

        /**
         * RegisterClientReply clientId.
         * @member {number|Long} clientId
         * @memberof raft.RegisterClientReply
         * @instance
         */
        RegisterClientReply.prototype.clientId = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * RegisterClientReply leaderHint.
         * @member {raft.IRemoteNode|null|undefined} leaderHint
         * @memberof raft.RegisterClientReply
         * @instance
         */
        RegisterClientReply.prototype.leaderHint = null;

        /**
         * Creates a new RegisterClientReply instance using the specified properties.
         * @function create
         * @memberof raft.RegisterClientReply
         * @static
         * @param {raft.IRegisterClientReply=} [properties] Properties to set
         * @returns {raft.RegisterClientReply} RegisterClientReply instance
         */
        RegisterClientReply.create = function create(properties) {
            return new RegisterClientReply(properties);
        };

        /**
         * Encodes the specified RegisterClientReply message. Does not implicitly {@link raft.RegisterClientReply.verify|verify} messages.
         * @function encode
         * @memberof raft.RegisterClientReply
         * @static
         * @param {raft.IRegisterClientReply} message RegisterClientReply message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegisterClientReply.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.status != null && message.hasOwnProperty("status"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.status);
            if (message.clientId != null && message.hasOwnProperty("clientId"))
                writer.uint32(/* id 2, wireType 0 =*/16).uint64(message.clientId);
            if (message.leaderHint != null && message.hasOwnProperty("leaderHint"))
                $root.raft.RemoteNode.encode(message.leaderHint, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified RegisterClientReply message, length delimited. Does not implicitly {@link raft.RegisterClientReply.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.RegisterClientReply
         * @static
         * @param {raft.IRegisterClientReply} message RegisterClientReply message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegisterClientReply.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegisterClientReply message from the specified reader or buffer.
         * @function decode
         * @memberof raft.RegisterClientReply
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.RegisterClientReply} RegisterClientReply
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegisterClientReply.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.RegisterClientReply();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.status = reader.int32();
                    break;
                case 2:
                    message.clientId = reader.uint64();
                    break;
                case 3:
                    message.leaderHint = $root.raft.RemoteNode.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegisterClientReply message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.RegisterClientReply
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.RegisterClientReply} RegisterClientReply
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegisterClientReply.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegisterClientReply message.
         * @function verify
         * @memberof raft.RegisterClientReply
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegisterClientReply.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.status != null && message.hasOwnProperty("status"))
                switch (message.status) {
                default:
                    return "status: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                    break;
                }
            if (message.clientId != null && message.hasOwnProperty("clientId"))
                if (!$util.isInteger(message.clientId) && !(message.clientId && $util.isInteger(message.clientId.low) && $util.isInteger(message.clientId.high)))
                    return "clientId: integer|Long expected";
            if (message.leaderHint != null && message.hasOwnProperty("leaderHint")) {
                var error = $root.raft.RemoteNode.verify(message.leaderHint);
                if (error)
                    return "leaderHint." + error;
            }
            return null;
        };

        /**
         * Creates a RegisterClientReply message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.RegisterClientReply
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.RegisterClientReply} RegisterClientReply
         */
        RegisterClientReply.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.RegisterClientReply)
                return object;
            var message = new $root.raft.RegisterClientReply();
            switch (object.status) {
            case "OK":
            case 0:
                message.status = 0;
                break;
            case "NOT_LEADER":
            case 1:
                message.status = 1;
                break;
            case "ELECTION_IN_PROGRESS":
            case 2:
                message.status = 2;
                break;
            case "CLUSTER_NOT_STARTED":
            case 3:
                message.status = 3;
                break;
            case "REQ_FAILED":
            case 4:
                message.status = 4;
                break;
            }
            if (object.clientId != null)
                if ($util.Long)
                    (message.clientId = $util.Long.fromValue(object.clientId)).unsigned = true;
                else if (typeof object.clientId === "string")
                    message.clientId = parseInt(object.clientId, 10);
                else if (typeof object.clientId === "number")
                    message.clientId = object.clientId;
                else if (typeof object.clientId === "object")
                    message.clientId = new $util.LongBits(object.clientId.low >>> 0, object.clientId.high >>> 0).toNumber(true);
            if (object.leaderHint != null) {
                if (typeof object.leaderHint !== "object")
                    throw TypeError(".raft.RegisterClientReply.leaderHint: object expected");
                message.leaderHint = $root.raft.RemoteNode.fromObject(object.leaderHint);
            }
            return message;
        };

        /**
         * Creates a plain object from a RegisterClientReply message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.RegisterClientReply
         * @static
         * @param {raft.RegisterClientReply} message RegisterClientReply
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegisterClientReply.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.status = options.enums === String ? "OK" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.clientId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.clientId = options.longs === String ? "0" : 0;
                object.leaderHint = null;
            }
            if (message.status != null && message.hasOwnProperty("status"))
                object.status = options.enums === String ? $root.raft.ClientStatus[message.status] : message.status;
            if (message.clientId != null && message.hasOwnProperty("clientId"))
                if (typeof message.clientId === "number")
                    object.clientId = options.longs === String ? String(message.clientId) : message.clientId;
                else
                    object.clientId = options.longs === String ? $util.Long.prototype.toString.call(message.clientId) : options.longs === Number ? new $util.LongBits(message.clientId.low >>> 0, message.clientId.high >>> 0).toNumber(true) : message.clientId;
            if (message.leaderHint != null && message.hasOwnProperty("leaderHint"))
                object.leaderHint = $root.raft.RemoteNode.toObject(message.leaderHint, options);
            return object;
        };

        /**
         * Converts this RegisterClientReply to JSON.
         * @function toJSON
         * @memberof raft.RegisterClientReply
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegisterClientReply.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegisterClientReply;
    })();

    raft.ClientRequest = (function() {

        /**
         * Properties of a ClientRequest.
         * @memberof raft
         * @interface IClientRequest
         * @property {number|Long|null} [clientId] ClientRequest clientId
         * @property {number|Long|null} [sequenceNum] ClientRequest sequenceNum
         * @property {number|Long|null} [stateMachineCmd] ClientRequest stateMachineCmd
         * @property {Uint8Array|null} [data] ClientRequest data
         */

        /**
         * Constructs a new ClientRequest.
         * @memberof raft
         * @classdesc Represents a ClientRequest.
         * @implements IClientRequest
         * @constructor
         * @param {raft.IClientRequest=} [properties] Properties to set
         */
        function ClientRequest(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * ClientRequest clientId.
         * @member {number|Long} clientId
         * @memberof raft.ClientRequest
         * @instance
         */
        ClientRequest.prototype.clientId = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * ClientRequest sequenceNum.
         * @member {number|Long} sequenceNum
         * @memberof raft.ClientRequest
         * @instance
         */
        ClientRequest.prototype.sequenceNum = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * ClientRequest stateMachineCmd.
         * @member {number|Long} stateMachineCmd
         * @memberof raft.ClientRequest
         * @instance
         */
        ClientRequest.prototype.stateMachineCmd = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * ClientRequest data.
         * @member {Uint8Array} data
         * @memberof raft.ClientRequest
         * @instance
         */
        ClientRequest.prototype.data = $util.newBuffer([]);

        /**
         * Creates a new ClientRequest instance using the specified properties.
         * @function create
         * @memberof raft.ClientRequest
         * @static
         * @param {raft.IClientRequest=} [properties] Properties to set
         * @returns {raft.ClientRequest} ClientRequest instance
         */
        ClientRequest.create = function create(properties) {
            return new ClientRequest(properties);
        };

        /**
         * Encodes the specified ClientRequest message. Does not implicitly {@link raft.ClientRequest.verify|verify} messages.
         * @function encode
         * @memberof raft.ClientRequest
         * @static
         * @param {raft.IClientRequest} message ClientRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ClientRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.clientId != null && message.hasOwnProperty("clientId"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.clientId);
            if (message.sequenceNum != null && message.hasOwnProperty("sequenceNum"))
                writer.uint32(/* id 2, wireType 0 =*/16).uint64(message.sequenceNum);
            if (message.stateMachineCmd != null && message.hasOwnProperty("stateMachineCmd"))
                writer.uint32(/* id 4, wireType 0 =*/32).uint64(message.stateMachineCmd);
            if (message.data != null && message.hasOwnProperty("data"))
                writer.uint32(/* id 5, wireType 2 =*/42).bytes(message.data);
            return writer;
        };

        /**
         * Encodes the specified ClientRequest message, length delimited. Does not implicitly {@link raft.ClientRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.ClientRequest
         * @static
         * @param {raft.IClientRequest} message ClientRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ClientRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ClientRequest message from the specified reader or buffer.
         * @function decode
         * @memberof raft.ClientRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.ClientRequest} ClientRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ClientRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.ClientRequest();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.clientId = reader.uint64();
                    break;
                case 2:
                    message.sequenceNum = reader.uint64();
                    break;
                case 4:
                    message.stateMachineCmd = reader.uint64();
                    break;
                case 5:
                    message.data = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ClientRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.ClientRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.ClientRequest} ClientRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ClientRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ClientRequest message.
         * @function verify
         * @memberof raft.ClientRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        ClientRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.clientId != null && message.hasOwnProperty("clientId"))
                if (!$util.isInteger(message.clientId) && !(message.clientId && $util.isInteger(message.clientId.low) && $util.isInteger(message.clientId.high)))
                    return "clientId: integer|Long expected";
            if (message.sequenceNum != null && message.hasOwnProperty("sequenceNum"))
                if (!$util.isInteger(message.sequenceNum) && !(message.sequenceNum && $util.isInteger(message.sequenceNum.low) && $util.isInteger(message.sequenceNum.high)))
                    return "sequenceNum: integer|Long expected";
            if (message.stateMachineCmd != null && message.hasOwnProperty("stateMachineCmd"))
                if (!$util.isInteger(message.stateMachineCmd) && !(message.stateMachineCmd && $util.isInteger(message.stateMachineCmd.low) && $util.isInteger(message.stateMachineCmd.high)))
                    return "stateMachineCmd: integer|Long expected";
            if (message.data != null && message.hasOwnProperty("data"))
                if (!(message.data && typeof message.data.length === "number" || $util.isString(message.data)))
                    return "data: buffer expected";
            return null;
        };

        /**
         * Creates a ClientRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.ClientRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.ClientRequest} ClientRequest
         */
        ClientRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.ClientRequest)
                return object;
            var message = new $root.raft.ClientRequest();
            if (object.clientId != null)
                if ($util.Long)
                    (message.clientId = $util.Long.fromValue(object.clientId)).unsigned = true;
                else if (typeof object.clientId === "string")
                    message.clientId = parseInt(object.clientId, 10);
                else if (typeof object.clientId === "number")
                    message.clientId = object.clientId;
                else if (typeof object.clientId === "object")
                    message.clientId = new $util.LongBits(object.clientId.low >>> 0, object.clientId.high >>> 0).toNumber(true);
            if (object.sequenceNum != null)
                if ($util.Long)
                    (message.sequenceNum = $util.Long.fromValue(object.sequenceNum)).unsigned = true;
                else if (typeof object.sequenceNum === "string")
                    message.sequenceNum = parseInt(object.sequenceNum, 10);
                else if (typeof object.sequenceNum === "number")
                    message.sequenceNum = object.sequenceNum;
                else if (typeof object.sequenceNum === "object")
                    message.sequenceNum = new $util.LongBits(object.sequenceNum.low >>> 0, object.sequenceNum.high >>> 0).toNumber(true);
            if (object.stateMachineCmd != null)
                if ($util.Long)
                    (message.stateMachineCmd = $util.Long.fromValue(object.stateMachineCmd)).unsigned = true;
                else if (typeof object.stateMachineCmd === "string")
                    message.stateMachineCmd = parseInt(object.stateMachineCmd, 10);
                else if (typeof object.stateMachineCmd === "number")
                    message.stateMachineCmd = object.stateMachineCmd;
                else if (typeof object.stateMachineCmd === "object")
                    message.stateMachineCmd = new $util.LongBits(object.stateMachineCmd.low >>> 0, object.stateMachineCmd.high >>> 0).toNumber(true);
            if (object.data != null)
                if (typeof object.data === "string")
                    $util.base64.decode(object.data, message.data = $util.newBuffer($util.base64.length(object.data)), 0);
                else if (object.data.length)
                    message.data = object.data;
            return message;
        };

        /**
         * Creates a plain object from a ClientRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.ClientRequest
         * @static
         * @param {raft.ClientRequest} message ClientRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ClientRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.clientId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.clientId = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.sequenceNum = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.sequenceNum = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.stateMachineCmd = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.stateMachineCmd = options.longs === String ? "0" : 0;
                if (options.bytes === String)
                    object.data = "";
                else {
                    object.data = [];
                    if (options.bytes !== Array)
                        object.data = $util.newBuffer(object.data);
                }
            }
            if (message.clientId != null && message.hasOwnProperty("clientId"))
                if (typeof message.clientId === "number")
                    object.clientId = options.longs === String ? String(message.clientId) : message.clientId;
                else
                    object.clientId = options.longs === String ? $util.Long.prototype.toString.call(message.clientId) : options.longs === Number ? new $util.LongBits(message.clientId.low >>> 0, message.clientId.high >>> 0).toNumber(true) : message.clientId;
            if (message.sequenceNum != null && message.hasOwnProperty("sequenceNum"))
                if (typeof message.sequenceNum === "number")
                    object.sequenceNum = options.longs === String ? String(message.sequenceNum) : message.sequenceNum;
                else
                    object.sequenceNum = options.longs === String ? $util.Long.prototype.toString.call(message.sequenceNum) : options.longs === Number ? new $util.LongBits(message.sequenceNum.low >>> 0, message.sequenceNum.high >>> 0).toNumber(true) : message.sequenceNum;
            if (message.stateMachineCmd != null && message.hasOwnProperty("stateMachineCmd"))
                if (typeof message.stateMachineCmd === "number")
                    object.stateMachineCmd = options.longs === String ? String(message.stateMachineCmd) : message.stateMachineCmd;
                else
                    object.stateMachineCmd = options.longs === String ? $util.Long.prototype.toString.call(message.stateMachineCmd) : options.longs === Number ? new $util.LongBits(message.stateMachineCmd.low >>> 0, message.stateMachineCmd.high >>> 0).toNumber(true) : message.stateMachineCmd;
            if (message.data != null && message.hasOwnProperty("data"))
                object.data = options.bytes === String ? $util.base64.encode(message.data, 0, message.data.length) : options.bytes === Array ? Array.prototype.slice.call(message.data) : message.data;
            return object;
        };

        /**
         * Converts this ClientRequest to JSON.
         * @function toJSON
         * @memberof raft.ClientRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        ClientRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ClientRequest;
    })();

    raft.ClientReply = (function() {

        /**
         * Properties of a ClientReply.
         * @memberof raft
         * @interface IClientReply
         * @property {raft.ClientStatus|null} [status] ClientReply status
         * @property {string|null} [response] ClientReply response
         * @property {raft.IRemoteNode|null} [leaderHint] ClientReply leaderHint
         */

        /**
         * Constructs a new ClientReply.
         * @memberof raft
         * @classdesc Represents a ClientReply.
         * @implements IClientReply
         * @constructor
         * @param {raft.IClientReply=} [properties] Properties to set
         */
        function ClientReply(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * ClientReply status.
         * @member {raft.ClientStatus} status
         * @memberof raft.ClientReply
         * @instance
         */
        ClientReply.prototype.status = 0;

        /**
         * ClientReply response.
         * @member {string} response
         * @memberof raft.ClientReply
         * @instance
         */
        ClientReply.prototype.response = "";

        /**
         * ClientReply leaderHint.
         * @member {raft.IRemoteNode|null|undefined} leaderHint
         * @memberof raft.ClientReply
         * @instance
         */
        ClientReply.prototype.leaderHint = null;

        /**
         * Creates a new ClientReply instance using the specified properties.
         * @function create
         * @memberof raft.ClientReply
         * @static
         * @param {raft.IClientReply=} [properties] Properties to set
         * @returns {raft.ClientReply} ClientReply instance
         */
        ClientReply.create = function create(properties) {
            return new ClientReply(properties);
        };

        /**
         * Encodes the specified ClientReply message. Does not implicitly {@link raft.ClientReply.verify|verify} messages.
         * @function encode
         * @memberof raft.ClientReply
         * @static
         * @param {raft.IClientReply} message ClientReply message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ClientReply.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.status != null && message.hasOwnProperty("status"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.status);
            if (message.response != null && message.hasOwnProperty("response"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.response);
            if (message.leaderHint != null && message.hasOwnProperty("leaderHint"))
                $root.raft.RemoteNode.encode(message.leaderHint, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified ClientReply message, length delimited. Does not implicitly {@link raft.ClientReply.verify|verify} messages.
         * @function encodeDelimited
         * @memberof raft.ClientReply
         * @static
         * @param {raft.IClientReply} message ClientReply message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ClientReply.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ClientReply message from the specified reader or buffer.
         * @function decode
         * @memberof raft.ClientReply
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {raft.ClientReply} ClientReply
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ClientReply.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.raft.ClientReply();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.status = reader.int32();
                    break;
                case 2:
                    message.response = reader.string();
                    break;
                case 3:
                    message.leaderHint = $root.raft.RemoteNode.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ClientReply message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof raft.ClientReply
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {raft.ClientReply} ClientReply
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ClientReply.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ClientReply message.
         * @function verify
         * @memberof raft.ClientReply
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        ClientReply.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.status != null && message.hasOwnProperty("status"))
                switch (message.status) {
                default:
                    return "status: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                    break;
                }
            if (message.response != null && message.hasOwnProperty("response"))
                if (!$util.isString(message.response))
                    return "response: string expected";
            if (message.leaderHint != null && message.hasOwnProperty("leaderHint")) {
                var error = $root.raft.RemoteNode.verify(message.leaderHint);
                if (error)
                    return "leaderHint." + error;
            }
            return null;
        };

        /**
         * Creates a ClientReply message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof raft.ClientReply
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {raft.ClientReply} ClientReply
         */
        ClientReply.fromObject = function fromObject(object) {
            if (object instanceof $root.raft.ClientReply)
                return object;
            var message = new $root.raft.ClientReply();
            switch (object.status) {
            case "OK":
            case 0:
                message.status = 0;
                break;
            case "NOT_LEADER":
            case 1:
                message.status = 1;
                break;
            case "ELECTION_IN_PROGRESS":
            case 2:
                message.status = 2;
                break;
            case "CLUSTER_NOT_STARTED":
            case 3:
                message.status = 3;
                break;
            case "REQ_FAILED":
            case 4:
                message.status = 4;
                break;
            }
            if (object.response != null)
                message.response = String(object.response);
            if (object.leaderHint != null) {
                if (typeof object.leaderHint !== "object")
                    throw TypeError(".raft.ClientReply.leaderHint: object expected");
                message.leaderHint = $root.raft.RemoteNode.fromObject(object.leaderHint);
            }
            return message;
        };

        /**
         * Creates a plain object from a ClientReply message. Also converts values to other types if specified.
         * @function toObject
         * @memberof raft.ClientReply
         * @static
         * @param {raft.ClientReply} message ClientReply
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ClientReply.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.status = options.enums === String ? "OK" : 0;
                object.response = "";
                object.leaderHint = null;
            }
            if (message.status != null && message.hasOwnProperty("status"))
                object.status = options.enums === String ? $root.raft.ClientStatus[message.status] : message.status;
            if (message.response != null && message.hasOwnProperty("response"))
                object.response = message.response;
            if (message.leaderHint != null && message.hasOwnProperty("leaderHint"))
                object.leaderHint = $root.raft.RemoteNode.toObject(message.leaderHint, options);
            return object;
        };

        /**
         * Converts this ClientReply to JSON.
         * @function toJSON
         * @memberof raft.ClientReply
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        ClientReply.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ClientReply;
    })();

    return raft;
})();

module.exports = $root;
