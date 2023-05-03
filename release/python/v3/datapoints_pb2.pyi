from v3 import types_pb2 as _types_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class StreamDatapointsRequest(_message.Message):
    __slots__ = ["datastream_id", "query"]
    DATASTREAM_ID_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    datastream_id: str
    query: _types_pb2.DatapointsQuery
    def __init__(self, datastream_id: _Optional[str] = ..., query: _Optional[_Union[_types_pb2.DatapointsQuery, _Mapping]] = ...) -> None: ...

class StreamDatapointsResponse(_message.Message):
    __slots__ = ["datapoints"]
    DATAPOINTS_FIELD_NUMBER: _ClassVar[int]
    datapoints: _containers.RepeatedCompositeFieldContainer[_types_pb2.Datapoint]
    def __init__(self, datapoints: _Optional[_Iterable[_Union[_types_pb2.Datapoint, _Mapping]]] = ...) -> None: ...
