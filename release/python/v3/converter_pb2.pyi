from v3 import types_pb2 as _types_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ConvertManyRequest(_message.Message):
    __slots__ = ["datapoints", "from_unit", "to_unit"]
    DATAPOINTS_FIELD_NUMBER: _ClassVar[int]
    FROM_UNIT_FIELD_NUMBER: _ClassVar[int]
    TO_UNIT_FIELD_NUMBER: _ClassVar[int]
    datapoints: _containers.RepeatedCompositeFieldContainer[_types_pb2.Datapoint]
    from_unit: str
    to_unit: str
    def __init__(self, from_unit: _Optional[str] = ..., to_unit: _Optional[str] = ..., datapoints: _Optional[_Iterable[_Union[_types_pb2.Datapoint, _Mapping]]] = ...) -> None: ...

class ConvertManyResponse(_message.Message):
    __slots__ = ["datapoints"]
    DATAPOINTS_FIELD_NUMBER: _ClassVar[int]
    datapoints: _containers.RepeatedCompositeFieldContainer[_types_pb2.Datapoint]
    def __init__(self, datapoints: _Optional[_Iterable[_Union[_types_pb2.Datapoint, _Mapping]]] = ...) -> None: ...
