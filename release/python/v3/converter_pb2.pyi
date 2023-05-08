from v3 import types_pb2 as _types_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ConvertManyRequest(_message.Message):
    __slots__ = ["convert", "datapoints"]
    CONVERT_FIELD_NUMBER: _ClassVar[int]
    DATAPOINTS_FIELD_NUMBER: _ClassVar[int]
    convert: _types_pb2.ConvertArgs
    datapoints: _containers.RepeatedCompositeFieldContainer[_types_pb2.Datapoint]
    def __init__(self, convert: _Optional[_Union[_types_pb2.ConvertArgs, _Mapping]] = ..., datapoints: _Optional[_Iterable[_Union[_types_pb2.Datapoint, _Mapping]]] = ...) -> None: ...

class ConvertManyResponse(_message.Message):
    __slots__ = ["datapoints"]
    DATAPOINTS_FIELD_NUMBER: _ClassVar[int]
    datapoints: _containers.RepeatedCompositeFieldContainer[_types_pb2.Datapoint]
    def __init__(self, datapoints: _Optional[_Iterable[_Union[_types_pb2.Datapoint, _Mapping]]] = ...) -> None: ...
