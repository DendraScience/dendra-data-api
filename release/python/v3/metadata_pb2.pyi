from v3 import types_pb2 as _types_pb2
from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetDatastreamRequest(_message.Message):
    __slots__ = ["datastream_id"]
    DATASTREAM_ID_FIELD_NUMBER: _ClassVar[int]
    datastream_id: str
    def __init__(self, datastream_id: _Optional[str] = ...) -> None: ...

class GetDatastreamResponse(_message.Message):
    __slots__ = ["datastream"]
    DATASTREAM_FIELD_NUMBER: _ClassVar[int]
    datastream: _types_pb2.Datastream
    def __init__(self, datastream: _Optional[_Union[_types_pb2.Datastream, _Mapping]] = ...) -> None: ...

class ListUomsResponse(_message.Message):
    __slots__ = ["uoms"]
    UOMS_FIELD_NUMBER: _ClassVar[int]
    uoms: _containers.RepeatedCompositeFieldContainer[_types_pb2.Uom]
    def __init__(self, uoms: _Optional[_Iterable[_Union[_types_pb2.Uom, _Mapping]]] = ...) -> None: ...
