# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v3/types.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0ev3/types.proto\x12\x02v3\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\";\n\x07\x45ndTime\x12%\n\x01t\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\t\n\x01o\x18\x02 \x01(\x11\"\xfa\x02\n\tDatapoint\x12%\n\x01t\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\t\n\x01o\x18\x02 \x01(\x11\x12&\n\x02lt\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\"\n\x01\x64\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12#\n\x02\x64\x61\x18\x05 \x03(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x01v\x18\x06 \x01(\x01H\x00\x88\x01\x01\x12\n\n\x02va\x18\x07 \x03(\x01\x12\x0f\n\x02uv\x18\x08 \x01(\x01H\x01\x88\x01\x01\x12\"\n\x01g\x18\n \x01(\x0b\x32\x17.google.protobuf.Struct\x12\n\n\x02gc\x18\x0b \x03(\x01\x12\"\n\x01q\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12#\n\x02ti\x18\x10 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x17\n\x02\x65t\x18\x11 \x01(\x0b\x32\x0b.v3.EndTimeB\x04\n\x02_vB\x05\n\x03_uv\"m\n\x11\x41nnotationActions\x12\'\n\x06\x61ttrib\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x10\n\x08\x65valuate\x18\x02 \x01(\t\x12\x0f\n\x07\x65xclude\x18\x03 \x01(\x08\x12\x0c\n\x04\x66lag\x18\x04 \x03(\t\"\xc7\x02\n\x18\x44\x61tapointsConfigInstance\x12\x32\n\tbegins_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x00\x88\x01\x01\x12\x34\n\x0b\x65nds_before\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x01\x88\x01\x01\x12\'\n\x06params\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nconnection\x18\x04 \x01(\t\x12\x0c\n\x04path\x18\x05 \x01(\t\x12\x10\n\x03ref\x18\x06 \x01(\rH\x02\x88\x01\x01\x12&\n\x07\x61\x63tions\x18\n \x01(\x0b\x32\x15.v3.AnnotationActions\x12\x16\n\x0e\x61nnotation_ids\x18\x0b \x03(\tB\x0c\n\n_begins_atB\x0e\n\x0c_ends_beforeB\x06\n\x04_ref\"\xa7\x02\n\x0f\x44\x61tapointsQuery\x12\x10\n\x08sort_asc\x18\x01 \x01(\x08\x12\x12\n\x05limit\x18\x02 \x01(\rH\x00\x88\x01\x01\x12\x30\n\x07lt_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x01\x88\x01\x01\x12\x30\n\x07gt_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x02\x88\x01\x01\x12\x10\n\x08lt_equal\x18\x05 \x01(\x08\x12\x10\n\x08gt_equal\x18\x06 \x01(\x08\x12\x10\n\x08is_local\x18\x07 \x01(\x08\x12\x10\n\x03lat\x18\x08 \x01(\x01H\x03\x88\x01\x01\x12\x10\n\x03lon\x18\t \x01(\x01H\x04\x88\x01\x01\x42\x08\n\x06_limitB\n\n\x08_lt_timeB\n\n\x08_gt_timeB\x06\n\x04_latB\x06\n\x04_lon\"#\n\rStationLookup\x12\x12\n\nutc_offset\x18\x01 \x01(\x11\"\xdc\x02\n\nDatastream\x12\x0b\n\x03_id\x18\x01 \x01(\t\x12\x12\n\nversion_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x19\n\x11\x64\x65rivation_method\x18\x04 \x01(\t\x12#\n\x1b\x64\x65rived_from_datastream_ids\x18\x05 \x03(\t\x12\x37\n\x11\x64\x61tapoints_config\x18\n \x03(\x0b\x32\x1c.v3.DatapointsConfigInstance\x12=\n\x17\x64\x61tapoints_config_built\x18\x0b \x03(\x0b\x32\x1c.v3.DatapointsConfigInstance\x12<\n\x16\x64\x61tapoints_config_refd\x18\x0c \x03(\x0b\x32\x1c.v3.DatapointsConfigInstance\x12)\n\x0estation_lookup\x18\r \x01(\x0b\x32\x11.v3.StationLookup\"\x9a\x01\n\x03Uom\x12\x0b\n\x03_id\x18\x01 \x01(\t\x12\x12\n\nversion_id\x18\x02 \x01(\t\x12\x0e\n\x06som_id\x18\x03 \x01(\t\x12\x1e\n\x16\x63onvertible_to_uom_ids\x18\x04 \x03(\t\x12\x11\n\tunit_tags\x18\x05 \x03(\t\x12/\n\x0elibrary_config\x18\x06 \x01(\x0b\x32\x17.google.protobuf.StructB8Z6github.com/dendrascience/dendra-data-api/release/go/v3b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'v3.types_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z6github.com/dendrascience/dendra-data-api/release/go/v3'
  _ENDTIME._serialized_start=85
  _ENDTIME._serialized_end=144
  _DATAPOINT._serialized_start=147
  _DATAPOINT._serialized_end=525
  _ANNOTATIONACTIONS._serialized_start=527
  _ANNOTATIONACTIONS._serialized_end=636
  _DATAPOINTSCONFIGINSTANCE._serialized_start=639
  _DATAPOINTSCONFIGINSTANCE._serialized_end=966
  _DATAPOINTSQUERY._serialized_start=969
  _DATAPOINTSQUERY._serialized_end=1264
  _STATIONLOOKUP._serialized_start=1266
  _STATIONLOOKUP._serialized_end=1301
  _DATASTREAM._serialized_start=1304
  _DATASTREAM._serialized_end=1652
  _UOM._serialized_start=1655
  _UOM._serialized_end=1809
# @@protoc_insertion_point(module_scope)
