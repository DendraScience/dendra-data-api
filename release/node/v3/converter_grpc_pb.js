// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var v3_converter_pb = require('../v3/converter_pb.js');
var v3_types_pb = require('../v3/types_pb.js');

function serialize_v3_ConvertAggregatesRequest(arg) {
  if (!(arg instanceof v3_converter_pb.ConvertAggregatesRequest)) {
    throw new Error('Expected argument of type v3.ConvertAggregatesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_ConvertAggregatesRequest(buffer_arg) {
  return v3_converter_pb.ConvertAggregatesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_ConvertAggregatesResponse(arg) {
  if (!(arg instanceof v3_converter_pb.ConvertAggregatesResponse)) {
    throw new Error('Expected argument of type v3.ConvertAggregatesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_ConvertAggregatesResponse(buffer_arg) {
  return v3_converter_pb.ConvertAggregatesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_ConvertDatapointsRequest(arg) {
  if (!(arg instanceof v3_converter_pb.ConvertDatapointsRequest)) {
    throw new Error('Expected argument of type v3.ConvertDatapointsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_ConvertDatapointsRequest(buffer_arg) {
  return v3_converter_pb.ConvertDatapointsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_ConvertDatapointsResponse(arg) {
  if (!(arg instanceof v3_converter_pb.ConvertDatapointsResponse)) {
    throw new Error('Expected argument of type v3.ConvertDatapointsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_ConvertDatapointsResponse(buffer_arg) {
  return v3_converter_pb.ConvertDatapointsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ConverterServiceService = exports.ConverterServiceService = {
  convertAggregates: {
    path: '/v3.ConverterService/ConvertAggregates',
    requestStream: false,
    responseStream: false,
    requestType: v3_converter_pb.ConvertAggregatesRequest,
    responseType: v3_converter_pb.ConvertAggregatesResponse,
    requestSerialize: serialize_v3_ConvertAggregatesRequest,
    requestDeserialize: deserialize_v3_ConvertAggregatesRequest,
    responseSerialize: serialize_v3_ConvertAggregatesResponse,
    responseDeserialize: deserialize_v3_ConvertAggregatesResponse,
  },
  convertDatapoints: {
    path: '/v3.ConverterService/ConvertDatapoints',
    requestStream: false,
    responseStream: false,
    requestType: v3_converter_pb.ConvertDatapointsRequest,
    responseType: v3_converter_pb.ConvertDatapointsResponse,
    requestSerialize: serialize_v3_ConvertDatapointsRequest,
    requestDeserialize: deserialize_v3_ConvertDatapointsRequest,
    responseSerialize: serialize_v3_ConvertDatapointsResponse,
    responseDeserialize: deserialize_v3_ConvertDatapointsResponse,
  },
};

exports.ConverterServiceClient = grpc.makeGenericClientConstructor(ConverterServiceService);
// NOTES:
// rpc ListUnits
