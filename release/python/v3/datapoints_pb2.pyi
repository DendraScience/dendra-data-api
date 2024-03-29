from google.api import annotations_pb2 as _annotations_pb2
from protoc_gen_openapiv2.options import annotations_pb2 as _annotations_pb2_1
from v3 import types_pb2 as _types_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class StreamAggregatesRequest(_message.Message):
    __slots__ = ["bucket_width", "convert", "count_only", "datastream_id", "query"]
    BUCKET_WIDTH_FIELD_NUMBER: _ClassVar[int]
    CONVERT_FIELD_NUMBER: _ClassVar[int]
    COUNT_ONLY_FIELD_NUMBER: _ClassVar[int]
    DATASTREAM_ID_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    bucket_width: str
    convert: _types_pb2.ConvertArgs
    count_only: bool
    datastream_id: str
    query: _types_pb2.DatapointsQuery
    def __init__(self, datastream_id: _Optional[str] = ..., query: _Optional[_Union[_types_pb2.DatapointsQuery, _Mapping]] = ..., convert: _Optional[_Union[_types_pb2.ConvertArgs, _Mapping]] = ..., bucket_width: _Optional[str] = ..., count_only: bool = ...) -> None: ...

class StreamAggregatesResponse(_message.Message):
    __slots__ = ["aggregates"]
    AGGREGATES_FIELD_NUMBER: _ClassVar[int]
    aggregates: _containers.RepeatedCompositeFieldContainer[_types_pb2.Aggregate]
    def __init__(self, aggregates: _Optional[_Iterable[_Union[_types_pb2.Aggregate, _Mapping]]] = ...) -> None: ...

class StreamDatapointsRequest(_message.Message):
    __slots__ = ["convert", "datastream_id", "query"]
    CONVERT_FIELD_NUMBER: _ClassVar[int]
    DATASTREAM_ID_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    convert: _types_pb2.ConvertArgs
    datastream_id: str
    query: _types_pb2.DatapointsQuery
    def __init__(self, datastream_id: _Optional[str] = ..., query: _Optional[_Union[_types_pb2.DatapointsQuery, _Mapping]] = ..., convert: _Optional[_Union[_types_pb2.ConvertArgs, _Mapping]] = ...) -> None: ...

class StreamDatapointsResponse(_message.Message):
    __slots__ = ["datapoints"]
    DATAPOINTS_FIELD_NUMBER: _ClassVar[int]
    datapoints: _containers.RepeatedCompositeFieldContainer[_types_pb2.Datapoint]
    def __init__(self, datapoints: _Optional[_Iterable[_Union[_types_pb2.Datapoint, _Mapping]]] = ...) -> None: ...
