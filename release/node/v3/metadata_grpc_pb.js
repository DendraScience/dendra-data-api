// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var v3_metadata_pb = require('../v3/metadata_pb.js');
var v3_types_pb = require('../v3/types_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

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

function serialize_v3_ListUomsResponse(arg) {
  if (!(arg instanceof v3_metadata_pb.ListUomsResponse)) {
    throw new Error('Expected argument of type v3.ListUomsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_ListUomsResponse(buffer_arg) {
  return v3_metadata_pb.ListUomsResponse.deserializeBinary(new Uint8Array(buffer_arg));
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
  listUoms: {
    path: '/v3.MetadataService/ListUoms',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: v3_metadata_pb.ListUomsResponse,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_v3_ListUomsResponse,
    responseDeserialize: deserialize_v3_ListUomsResponse,
  },
};

exports.MetadataServiceClient = grpc.makeGenericClientConstructor(MetadataServiceService);
