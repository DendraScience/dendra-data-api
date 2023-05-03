# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from v3 import datapoints_pb2 as v3_dot_datapoints__pb2


class DatapointsServiceStub(object):
    """NOTES:
    ListAggregates (range, only count, sum, min and max)
    GetDatapoint (range | sort)
    StreamAggregates (range | sort | interval | only count, max, min, sum)
    StreamDatapoints (range | sort | limit)
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.StreamDatapoints = channel.unary_stream(
                '/v3.DatapointsService/StreamDatapoints',
                request_serializer=v3_dot_datapoints__pb2.StreamDatapointsRequest.SerializeToString,
                response_deserializer=v3_dot_datapoints__pb2.StreamDatapointsResponse.FromString,
                )


class DatapointsServiceServicer(object):
    """NOTES:
    ListAggregates (range, only count, sum, min and max)
    GetDatapoint (range | sort)
    StreamAggregates (range | sort | interval | only count, max, min, sum)
    StreamDatapoints (range | sort | limit)
    """

    def StreamDatapoints(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DatapointsServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'StreamDatapoints': grpc.unary_stream_rpc_method_handler(
                    servicer.StreamDatapoints,
                    request_deserializer=v3_dot_datapoints__pb2.StreamDatapointsRequest.FromString,
                    response_serializer=v3_dot_datapoints__pb2.StreamDatapointsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'v3.DatapointsService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DatapointsService(object):
    """NOTES:
    ListAggregates (range, only count, sum, min and max)
    GetDatapoint (range | sort)
    StreamAggregates (range | sort | interval | only count, max, min, sum)
    StreamDatapoints (range | sort | limit)
    """

    @staticmethod
    def StreamDatapoints(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/v3.DatapointsService/StreamDatapoints',
            v3_dot_datapoints__pb2.StreamDatapointsRequest.SerializeToString,
            v3_dot_datapoints__pb2.StreamDatapointsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
