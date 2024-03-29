from v3 import types_pb2 as _types_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ProviderStreamDatapointsRequest(_message.Message):
    __slots__ = ["config_instance", "query"]
    CONFIG_INSTANCE_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    config_instance: _types_pb2.DatapointsConfigInstance
    query: _types_pb2.DatapointsQuery
    def __init__(self, query: _Optional[_Union[_types_pb2.DatapointsQuery, _Mapping]] = ..., config_instance: _Optional[_Union[_types_pb2.DatapointsConfigInstance, _Mapping]] = ...) -> None: ...

class ProviderStreamDatapointsResponse(_message.Message):
    __slots__ = ["datapoints"]
    DATAPOINTS_FIELD_NUMBER: _ClassVar[int]
    datapoints: _containers.RepeatedCompositeFieldContainer[_types_pb2.Datapoint]
    def __init__(self, datapoints: _Optional[_Iterable[_Union[_types_pb2.Datapoint, _Mapping]]] = ...) -> None: ...
