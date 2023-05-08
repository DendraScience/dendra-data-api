# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from v3 import converter_pb2 as v3_dot_converter__pb2


class ConverterServiceStub(object):
    """NOTES:
    rpc ListUnits
    rpc ConvertOne
    rpc ConvertStream
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ConvertMany = channel.unary_unary(
                '/v3.ConverterService/ConvertMany',
                request_serializer=v3_dot_converter__pb2.ConvertManyRequest.SerializeToString,
                response_deserializer=v3_dot_converter__pb2.ConvertManyResponse.FromString,
                )


class ConverterServiceServicer(object):
    """NOTES:
    rpc ListUnits
    rpc ConvertOne
    rpc ConvertStream
    """

    def ConvertMany(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ConverterServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ConvertMany': grpc.unary_unary_rpc_method_handler(
                    servicer.ConvertMany,
                    request_deserializer=v3_dot_converter__pb2.ConvertManyRequest.FromString,
                    response_serializer=v3_dot_converter__pb2.ConvertManyResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'v3.ConverterService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ConverterService(object):
    """NOTES:
    rpc ListUnits
    rpc ConvertOne
    rpc ConvertStream
    """

    @staticmethod
    def ConvertMany(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/v3.ConverterService/ConvertMany',
            v3_dot_converter__pb2.ConvertManyRequest.SerializeToString,
            v3_dot_converter__pb2.ConvertManyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
