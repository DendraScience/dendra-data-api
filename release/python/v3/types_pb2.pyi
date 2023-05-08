from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AnnotationActions(_message.Message):
    __slots__ = ["attrib", "evaluate", "exclude", "flag"]
    ATTRIB_FIELD_NUMBER: _ClassVar[int]
    EVALUATE_FIELD_NUMBER: _ClassVar[int]
    EXCLUDE_FIELD_NUMBER: _ClassVar[int]
    FLAG_FIELD_NUMBER: _ClassVar[int]
    attrib: _struct_pb2.Struct
    evaluate: str
    exclude: bool
    flag: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, attrib: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., evaluate: _Optional[str] = ..., exclude: bool = ..., flag: _Optional[_Iterable[str]] = ...) -> None: ...

class ConvertArgs(_message.Message):
    __slots__ = ["from_unit", "library", "to_unit"]
    FROM_UNIT_FIELD_NUMBER: _ClassVar[int]
    LIBRARY_FIELD_NUMBER: _ClassVar[int]
    TO_UNIT_FIELD_NUMBER: _ClassVar[int]
    from_unit: str
    library: str
    to_unit: str
    def __init__(self, from_unit: _Optional[str] = ..., to_unit: _Optional[str] = ..., library: _Optional[str] = ...) -> None: ...

class Datapoint(_message.Message):
    __slots__ = ["d", "da", "et", "g", "gc", "lt", "o", "q", "t", "ti", "uv", "v", "va"]
    DA_FIELD_NUMBER: _ClassVar[int]
    D_FIELD_NUMBER: _ClassVar[int]
    ET_FIELD_NUMBER: _ClassVar[int]
    GC_FIELD_NUMBER: _ClassVar[int]
    G_FIELD_NUMBER: _ClassVar[int]
    LT_FIELD_NUMBER: _ClassVar[int]
    O_FIELD_NUMBER: _ClassVar[int]
    Q_FIELD_NUMBER: _ClassVar[int]
    TI_FIELD_NUMBER: _ClassVar[int]
    T_FIELD_NUMBER: _ClassVar[int]
    UV_FIELD_NUMBER: _ClassVar[int]
    VA_FIELD_NUMBER: _ClassVar[int]
    V_FIELD_NUMBER: _ClassVar[int]
    d: _struct_pb2.Struct
    da: _containers.RepeatedCompositeFieldContainer[_struct_pb2.Struct]
    et: EndTime
    g: _struct_pb2.Struct
    gc: _containers.RepeatedScalarFieldContainer[float]
    lt: _timestamp_pb2.Timestamp
    o: int
    q: _struct_pb2.Struct
    t: _timestamp_pb2.Timestamp
    ti: _struct_pb2.Struct
    uv: float
    v: float
    va: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, t: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., o: _Optional[int] = ..., lt: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., d: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., da: _Optional[_Iterable[_Union[_struct_pb2.Struct, _Mapping]]] = ..., v: _Optional[float] = ..., va: _Optional[_Iterable[float]] = ..., uv: _Optional[float] = ..., g: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., gc: _Optional[_Iterable[float]] = ..., q: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., ti: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., et: _Optional[_Union[EndTime, _Mapping]] = ...) -> None: ...

class DatapointsConfigInstance(_message.Message):
    __slots__ = ["actions", "annotation_ids", "begins_at", "connection", "ends_before", "params", "path", "ref"]
    ACTIONS_FIELD_NUMBER: _ClassVar[int]
    ANNOTATION_IDS_FIELD_NUMBER: _ClassVar[int]
    BEGINS_AT_FIELD_NUMBER: _ClassVar[int]
    CONNECTION_FIELD_NUMBER: _ClassVar[int]
    ENDS_BEFORE_FIELD_NUMBER: _ClassVar[int]
    PARAMS_FIELD_NUMBER: _ClassVar[int]
    PATH_FIELD_NUMBER: _ClassVar[int]
    REF_FIELD_NUMBER: _ClassVar[int]
    actions: AnnotationActions
    annotation_ids: _containers.RepeatedScalarFieldContainer[str]
    begins_at: _timestamp_pb2.Timestamp
    connection: str
    ends_before: _timestamp_pb2.Timestamp
    params: _struct_pb2.Struct
    path: str
    ref: int
    def __init__(self, begins_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., ends_before: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., params: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., connection: _Optional[str] = ..., path: _Optional[str] = ..., ref: _Optional[int] = ..., actions: _Optional[_Union[AnnotationActions, _Mapping]] = ..., annotation_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class DatapointsQuery(_message.Message):
    __slots__ = ["gt_equal", "gt_time", "is_local", "lat", "limit", "lon", "lt_equal", "lt_time", "sort_asc"]
    GT_EQUAL_FIELD_NUMBER: _ClassVar[int]
    GT_TIME_FIELD_NUMBER: _ClassVar[int]
    IS_LOCAL_FIELD_NUMBER: _ClassVar[int]
    LAT_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    LON_FIELD_NUMBER: _ClassVar[int]
    LT_EQUAL_FIELD_NUMBER: _ClassVar[int]
    LT_TIME_FIELD_NUMBER: _ClassVar[int]
    SORT_ASC_FIELD_NUMBER: _ClassVar[int]
    gt_equal: bool
    gt_time: _timestamp_pb2.Timestamp
    is_local: bool
    lat: float
    limit: int
    lon: float
    lt_equal: bool
    lt_time: _timestamp_pb2.Timestamp
    sort_asc: bool
    def __init__(self, sort_asc: bool = ..., limit: _Optional[int] = ..., lt_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., gt_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., lt_equal: bool = ..., gt_equal: bool = ..., is_local: bool = ..., lat: _Optional[float] = ..., lon: _Optional[float] = ...) -> None: ...

class Datastream(_message.Message):
    __slots__ = ["_id", "datapoints_config", "datapoints_config_built", "datapoints_config_refd", "derivation_method", "derived_from_datastream_ids", "name", "station_lookup", "terms", "terms_info", "version_id"]
    DATAPOINTS_CONFIG_BUILT_FIELD_NUMBER: _ClassVar[int]
    DATAPOINTS_CONFIG_FIELD_NUMBER: _ClassVar[int]
    DATAPOINTS_CONFIG_REFD_FIELD_NUMBER: _ClassVar[int]
    DERIVATION_METHOD_FIELD_NUMBER: _ClassVar[int]
    DERIVED_FROM_DATASTREAM_IDS_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    STATION_LOOKUP_FIELD_NUMBER: _ClassVar[int]
    TERMS_FIELD_NUMBER: _ClassVar[int]
    TERMS_INFO_FIELD_NUMBER: _ClassVar[int]
    VERSION_ID_FIELD_NUMBER: _ClassVar[int]
    _ID_FIELD_NUMBER: _ClassVar[int]
    _id: str
    datapoints_config: _containers.RepeatedCompositeFieldContainer[DatapointsConfigInstance]
    datapoints_config_built: _containers.RepeatedCompositeFieldContainer[DatapointsConfigInstance]
    datapoints_config_refd: _containers.RepeatedCompositeFieldContainer[DatapointsConfigInstance]
    derivation_method: str
    derived_from_datastream_ids: _containers.RepeatedScalarFieldContainer[str]
    name: str
    station_lookup: StationLookup
    terms: _struct_pb2.Struct
    terms_info: TermsInfo
    version_id: str
    def __init__(self, _id: _Optional[str] = ..., version_id: _Optional[str] = ..., name: _Optional[str] = ..., derivation_method: _Optional[str] = ..., derived_from_datastream_ids: _Optional[_Iterable[str]] = ..., datapoints_config: _Optional[_Iterable[_Union[DatapointsConfigInstance, _Mapping]]] = ..., datapoints_config_built: _Optional[_Iterable[_Union[DatapointsConfigInstance, _Mapping]]] = ..., datapoints_config_refd: _Optional[_Iterable[_Union[DatapointsConfigInstance, _Mapping]]] = ..., station_lookup: _Optional[_Union[StationLookup, _Mapping]] = ..., terms: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., terms_info: _Optional[_Union[TermsInfo, _Mapping]] = ...) -> None: ...

class EndTime(_message.Message):
    __slots__ = ["o", "t"]
    O_FIELD_NUMBER: _ClassVar[int]
    T_FIELD_NUMBER: _ClassVar[int]
    o: int
    t: _timestamp_pb2.Timestamp
    def __init__(self, t: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., o: _Optional[int] = ...) -> None: ...

class StationLookup(_message.Message):
    __slots__ = ["utc_offset"]
    UTC_OFFSET_FIELD_NUMBER: _ClassVar[int]
    utc_offset: int
    def __init__(self, utc_offset: _Optional[int] = ...) -> None: ...

class TermsInfo(_message.Message):
    __slots__ = ["class_keys", "class_tags", "unit_tag"]
    CLASS_KEYS_FIELD_NUMBER: _ClassVar[int]
    CLASS_TAGS_FIELD_NUMBER: _ClassVar[int]
    UNIT_TAG_FIELD_NUMBER: _ClassVar[int]
    class_keys: _containers.RepeatedScalarFieldContainer[str]
    class_tags: _containers.RepeatedScalarFieldContainer[str]
    unit_tag: str
    def __init__(self, class_keys: _Optional[_Iterable[str]] = ..., class_tags: _Optional[_Iterable[str]] = ..., unit_tag: _Optional[str] = ...) -> None: ...

class Uom(_message.Message):
    __slots__ = ["_id", "convertible_to_uom_ids", "library_config", "som_id", "unit_tags", "version_id"]
    CONVERTIBLE_TO_UOM_IDS_FIELD_NUMBER: _ClassVar[int]
    LIBRARY_CONFIG_FIELD_NUMBER: _ClassVar[int]
    SOM_ID_FIELD_NUMBER: _ClassVar[int]
    UNIT_TAGS_FIELD_NUMBER: _ClassVar[int]
    VERSION_ID_FIELD_NUMBER: _ClassVar[int]
    _ID_FIELD_NUMBER: _ClassVar[int]
    _id: str
    convertible_to_uom_ids: _containers.RepeatedScalarFieldContainer[str]
    library_config: _struct_pb2.Struct
    som_id: str
    unit_tags: _containers.RepeatedScalarFieldContainer[str]
    version_id: str
    def __init__(self, _id: _Optional[str] = ..., version_id: _Optional[str] = ..., som_id: _Optional[str] = ..., convertible_to_uom_ids: _Optional[_Iterable[str]] = ..., unit_tags: _Optional[_Iterable[str]] = ..., library_config: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...
