// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var v3_provider_pb = require('../v3/provider_pb.js');
var v3_types_pb = require('../v3/types_pb.js');

function serialize_v3_ProviderStreamDatapointsRequest(arg) {
  if (!(arg instanceof v3_provider_pb.ProviderStreamDatapointsRequest)) {
    throw new Error('Expected argument of type v3.ProviderStreamDatapointsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_ProviderStreamDatapointsRequest(buffer_arg) {
  return v3_provider_pb.ProviderStreamDatapointsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_ProviderStreamDatapointsResponse(arg) {
  if (!(arg instanceof v3_provider_pb.ProviderStreamDatapointsResponse)) {
    throw new Error('Expected argument of type v3.ProviderStreamDatapointsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_ProviderStreamDatapointsResponse(buffer_arg) {
  return v3_provider_pb.ProviderStreamDatapointsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ProviderServiceService = exports.ProviderServiceService = {
  streamDatapoints: {
    path: '/v3.ProviderService/StreamDatapoints',
    requestStream: false,
    responseStream: true,
    requestType: v3_provider_pb.ProviderStreamDatapointsRequest,
    responseType: v3_provider_pb.ProviderStreamDatapointsResponse,
    requestSerialize: serialize_v3_ProviderStreamDatapointsRequest,
    requestDeserialize: deserialize_v3_ProviderStreamDatapointsRequest,
    responseSerialize: serialize_v3_ProviderStreamDatapointsResponse,
    responseDeserialize: deserialize_v3_ProviderStreamDatapointsResponse,
  },
};

exports.ProviderServiceClient = grpc.makeGenericClientConstructor(ProviderServiceService);
// NOTES:
// rpc GetAggregate (range | count, sum, min and max)
// rpc GetDatapoint (range | sort)
// rpc StreamAggregates (range | sort | interval | count, max, min, sum)
