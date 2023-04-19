// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var v3_provider_pb = require('../v3/provider_pb.js');
var v3_types_pb = require('../v3/types_pb.js');

function serialize_v3_StreamDatapointsRequest(arg) {
  if (!(arg instanceof v3_provider_pb.StreamDatapointsRequest)) {
    throw new Error('Expected argument of type v3.StreamDatapointsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_StreamDatapointsRequest(buffer_arg) {
  return v3_provider_pb.StreamDatapointsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_StreamDatapointsResponse(arg) {
  if (!(arg instanceof v3_provider_pb.StreamDatapointsResponse)) {
    throw new Error('Expected argument of type v3.StreamDatapointsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_StreamDatapointsResponse(buffer_arg) {
  return v3_provider_pb.StreamDatapointsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ProviderServiceService = exports.ProviderServiceService = {
  streamDatapoints: {
    path: '/v3.ProviderService/StreamDatapoints',
    requestStream: false,
    responseStream: true,
    requestType: v3_provider_pb.StreamDatapointsRequest,
    responseType: v3_provider_pb.StreamDatapointsResponse,
    requestSerialize: serialize_v3_StreamDatapointsRequest,
    requestDeserialize: deserialize_v3_StreamDatapointsRequest,
    responseSerialize: serialize_v3_StreamDatapointsResponse,
    responseDeserialize: deserialize_v3_StreamDatapointsResponse,
  },
};

exports.ProviderServiceClient = grpc.makeGenericClientConstructor(ProviderServiceService);
// NOTES:
// rpc GetAggregate (range | count, sum, min and max)
// rpc GetDatapoint (range | sort)
// rpc StreamAggregates (range | sort | interval | count, max, min, sum)
