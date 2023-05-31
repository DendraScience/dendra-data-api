from v3 import types_pb2 as _types_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ConvertAggregatesRequest(_message.Message):
    __slots__ = ["aggregates", "convert"]
    AGGREGATES_FIELD_NUMBER: _ClassVar[int]
    CONVERT_FIELD_NUMBER: _ClassVar[int]
    aggregates: _containers.RepeatedCompositeFieldContainer[_types_pb2.Aggregate]
    convert: _types_pb2.ConvertArgs
    def __init__(self, convert: _Optional[_Union[_types_pb2.ConvertArgs, _Mapping]] = ..., aggregates: _Optional[_Iterable[_Union[_types_pb2.Aggregate, _Mapping]]] = ...) -> None: ...

class ConvertAggregatesResponse(_message.Message):
    __slots__ = ["aggregates"]
    AGGREGATES_FIELD_NUMBER: _ClassVar[int]
    aggregates: _containers.RepeatedCompositeFieldContainer[_types_pb2.Aggregate]
    def __init__(self, aggregates: _Optional[_Iterable[_Union[_types_pb2.Aggregate, _Mapping]]] = ...) -> None: ...

class ConvertDatapointsRequest(_message.Message):
    __slots__ = ["convert", "datapoints"]
    CONVERT_FIELD_NUMBER: _ClassVar[int]
    DATAPOINTS_FIELD_NUMBER: _ClassVar[int]
    convert: _types_pb2.ConvertArgs
    datapoints: _containers.RepeatedCompositeFieldContainer[_types_pb2.Datapoint]
    def __init__(self, convert: _Optional[_Union[_types_pb2.ConvertArgs, _Mapping]] = ..., datapoints: _Optional[_Iterable[_Union[_types_pb2.Datapoint, _Mapping]]] = ...) -> None: ...

class ConvertDatapointsResponse(_message.Message):
    __slots__ = ["datapoints"]
    DATAPOINTS_FIELD_NUMBER: _ClassVar[int]
    datapoints: _containers.RepeatedCompositeFieldContainer[_types_pb2.Datapoint]
    def __init__(self, datapoints: _Optional[_Iterable[_Union[_types_pb2.Datapoint, _Mapping]]] = ...) -> None: ...
