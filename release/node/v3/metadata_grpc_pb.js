// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var v3_metadata_pb = require('../v3/metadata_pb.js');
var v3_types_pb = require('../v3/types_pb.js');

function serialize_v3_GetDatastreamRequest(arg) {
  if (!(arg instanceof v3_metadata_pb.GetDatastreamRequest)) {
    throw new Error('Expected argument of type v3.GetDatastreamRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_GetDatastreamRequest(buffer_arg) {
  return v3_metadata_pb.GetDatastreamRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_GetDatastreamResponse(arg) {
  if (!(arg instanceof v3_metadata_pb.GetDatastreamResponse)) {
    throw new Error('Expected argument of type v3.GetDatastreamResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_GetDatastreamResponse(buffer_arg) {
  return v3_metadata_pb.GetDatastreamResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var MetadataServiceService = exports.MetadataServiceService = {
  getDatastream: {
    path: '/v3.MetadataService/GetDatastream',
    requestStream: false,
    responseStream: false,
    requestType: v3_metadata_pb.GetDatastreamRequest,
    responseType: v3_metadata_pb.GetDatastreamResponse,
    requestSerialize: serialize_v3_GetDatastreamRequest,
    requestDeserialize: deserialize_v3_GetDatastreamRequest,
    responseSerialize: serialize_v3_GetDatastreamResponse,
    responseDeserialize: deserialize_v3_GetDatastreamResponse,
  },
};

exports.MetadataServiceClient = grpc.makeGenericClientConstructor(MetadataServiceService);
