// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var v3_datapoints_pb = require('../v3/datapoints_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var protoc$gen$openapiv2_options_annotations_pb = require('../protoc-gen-openapiv2/options/annotations_pb.js');
var v3_types_pb = require('../v3/types_pb.js');

function serialize_v3_StreamAggregatesRequest(arg) {
  if (!(arg instanceof v3_datapoints_pb.StreamAggregatesRequest)) {
    throw new Error('Expected argument of type v3.StreamAggregatesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_StreamAggregatesRequest(buffer_arg) {
  return v3_datapoints_pb.StreamAggregatesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_StreamAggregatesResponse(arg) {
  if (!(arg instanceof v3_datapoints_pb.StreamAggregatesResponse)) {
    throw new Error('Expected argument of type v3.StreamAggregatesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_StreamAggregatesResponse(buffer_arg) {
  return v3_datapoints_pb.StreamAggregatesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_StreamDatapointsRequest(arg) {
  if (!(arg instanceof v3_datapoints_pb.StreamDatapointsRequest)) {
    throw new Error('Expected argument of type v3.StreamDatapointsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_StreamDatapointsRequest(buffer_arg) {
  return v3_datapoints_pb.StreamDatapointsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_StreamDatapointsResponse(arg) {
  if (!(arg instanceof v3_datapoints_pb.StreamDatapointsResponse)) {
    throw new Error('Expected argument of type v3.StreamDatapointsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_StreamDatapointsResponse(buffer_arg) {
  return v3_datapoints_pb.StreamDatapointsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var DatapointsServiceService = exports.DatapointsServiceService = {
  streamAggregates: {
    path: '/v3.DatapointsService/StreamAggregates',
    requestStream: false,
    responseStream: true,
    requestType: v3_datapoints_pb.StreamAggregatesRequest,
    responseType: v3_datapoints_pb.StreamAggregatesResponse,
    requestSerialize: serialize_v3_StreamAggregatesRequest,
    requestDeserialize: deserialize_v3_StreamAggregatesRequest,
    responseSerialize: serialize_v3_StreamAggregatesResponse,
    responseDeserialize: deserialize_v3_StreamAggregatesResponse,
  },
  streamDatapoints: {
    path: '/v3.DatapointsService/StreamDatapoints',
    requestStream: false,
    responseStream: true,
    requestType: v3_datapoints_pb.StreamDatapointsRequest,
    responseType: v3_datapoints_pb.StreamDatapointsResponse,
    requestSerialize: serialize_v3_StreamDatapointsRequest,
    requestDeserialize: deserialize_v3_StreamDatapointsRequest,
    responseSerialize: serialize_v3_StreamDatapointsResponse,
    responseDeserialize: deserialize_v3_StreamDatapointsResponse,
  },
};

exports.DatapointsServiceClient = grpc.makeGenericClientConstructor(DatapointsServiceService);
