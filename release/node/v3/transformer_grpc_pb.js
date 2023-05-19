// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var v3_transformer_pb = require('../v3/transformer_pb.js');
var v3_types_pb = require('../v3/types_pb.js');

function serialize_v3_TransformManyRequest(arg) {
  if (!(arg instanceof v3_transformer_pb.TransformManyRequest)) {
    throw new Error('Expected argument of type v3.TransformManyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_TransformManyRequest(buffer_arg) {
  return v3_transformer_pb.TransformManyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_TransformManyResponse(arg) {
  if (!(arg instanceof v3_transformer_pb.TransformManyResponse)) {
    throw new Error('Expected argument of type v3.TransformManyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_TransformManyResponse(buffer_arg) {
  return v3_transformer_pb.TransformManyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var TransformerServiceService = exports.TransformerServiceService = {
  transformMany: {
    path: '/v3.TransformerService/TransformMany',
    requestStream: false,
    responseStream: false,
    requestType: v3_transformer_pb.TransformManyRequest,
    responseType: v3_transformer_pb.TransformManyResponse,
    requestSerialize: serialize_v3_TransformManyRequest,
    requestDeserialize: deserialize_v3_TransformManyRequest,
    responseSerialize: serialize_v3_TransformManyResponse,
    responseDeserialize: deserialize_v3_TransformManyResponse,
  },
};

exports.TransformerServiceClient = grpc.makeGenericClientConstructor(TransformerServiceService);
// NOTES:
// rpc TransformOne
// rpc TransformStream
