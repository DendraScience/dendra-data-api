// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var v3_converter_pb = require('../v3/converter_pb.js');
var v3_types_pb = require('../v3/types_pb.js');

function serialize_v3_ConvertManyRequest(arg) {
  if (!(arg instanceof v3_converter_pb.ConvertManyRequest)) {
    throw new Error('Expected argument of type v3.ConvertManyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_ConvertManyRequest(buffer_arg) {
  return v3_converter_pb.ConvertManyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v3_ConvertManyResponse(arg) {
  if (!(arg instanceof v3_converter_pb.ConvertManyResponse)) {
    throw new Error('Expected argument of type v3.ConvertManyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v3_ConvertManyResponse(buffer_arg) {
  return v3_converter_pb.ConvertManyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ConverterServiceService = exports.ConverterServiceService = {
  convertMany: {
    path: '/v3.ConverterService/ConvertMany',
    requestStream: false,
    responseStream: false,
    requestType: v3_converter_pb.ConvertManyRequest,
    responseType: v3_converter_pb.ConvertManyResponse,
    requestSerialize: serialize_v3_ConvertManyRequest,
    requestDeserialize: deserialize_v3_ConvertManyRequest,
    responseSerialize: serialize_v3_ConvertManyResponse,
    responseDeserialize: deserialize_v3_ConvertManyResponse,
  },
};

exports.ConverterServiceClient = grpc.makeGenericClientConstructor(ConverterServiceService);
// NOTES:
// rpc ListUnits
// rpc ConvertOne
// rpc ConvertStream
