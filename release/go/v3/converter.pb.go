// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.1
// source: v3/converter.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConvertAggregatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Convert    *ConvertArgs `protobuf:"bytes,1,opt,name=convert,proto3" json:"convert,omitempty"`
	Aggregates []*Aggregate `protobuf:"bytes,2,rep,name=aggregates,proto3" json:"aggregates,omitempty"`
}

func (x *ConvertAggregatesRequest) Reset() {
	*x = ConvertAggregatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_converter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertAggregatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertAggregatesRequest) ProtoMessage() {}

func (x *ConvertAggregatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v3_converter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertAggregatesRequest.ProtoReflect.Descriptor instead.
func (*ConvertAggregatesRequest) Descriptor() ([]byte, []int) {
	return file_v3_converter_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertAggregatesRequest) GetConvert() *ConvertArgs {
	if x != nil {
		return x.Convert
	}
	return nil
}

func (x *ConvertAggregatesRequest) GetAggregates() []*Aggregate {
	if x != nil {
		return x.Aggregates
	}
	return nil
}

type ConvertAggregatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aggregates []*Aggregate `protobuf:"bytes,1,rep,name=aggregates,proto3" json:"aggregates,omitempty"`
}

func (x *ConvertAggregatesResponse) Reset() {
	*x = ConvertAggregatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_converter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertAggregatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertAggregatesResponse) ProtoMessage() {}

func (x *ConvertAggregatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v3_converter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertAggregatesResponse.ProtoReflect.Descriptor instead.
func (*ConvertAggregatesResponse) Descriptor() ([]byte, []int) {
	return file_v3_converter_proto_rawDescGZIP(), []int{1}
}

func (x *ConvertAggregatesResponse) GetAggregates() []*Aggregate {
	if x != nil {
		return x.Aggregates
	}
	return nil
}

type ConvertDatapointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Convert    *ConvertArgs `protobuf:"bytes,1,opt,name=convert,proto3" json:"convert,omitempty"`
	Datapoints []*Datapoint `protobuf:"bytes,2,rep,name=datapoints,proto3" json:"datapoints,omitempty"`
}

func (x *ConvertDatapointsRequest) Reset() {
	*x = ConvertDatapointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_converter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertDatapointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertDatapointsRequest) ProtoMessage() {}

func (x *ConvertDatapointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v3_converter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertDatapointsRequest.ProtoReflect.Descriptor instead.
func (*ConvertDatapointsRequest) Descriptor() ([]byte, []int) {
	return file_v3_converter_proto_rawDescGZIP(), []int{2}
}

func (x *ConvertDatapointsRequest) GetConvert() *ConvertArgs {
	if x != nil {
		return x.Convert
	}
	return nil
}

func (x *ConvertDatapointsRequest) GetDatapoints() []*Datapoint {
	if x != nil {
		return x.Datapoints
	}
	return nil
}

type ConvertDatapointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datapoints []*Datapoint `protobuf:"bytes,1,rep,name=datapoints,proto3" json:"datapoints,omitempty"`
}

func (x *ConvertDatapointsResponse) Reset() {
	*x = ConvertDatapointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_converter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertDatapointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertDatapointsResponse) ProtoMessage() {}

func (x *ConvertDatapointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v3_converter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertDatapointsResponse.ProtoReflect.Descriptor instead.
func (*ConvertDatapointsResponse) Descriptor() ([]byte, []int) {
	return file_v3_converter_proto_rawDescGZIP(), []int{3}
}

func (x *ConvertDatapointsResponse) GetDatapoints() []*Datapoint {
	if x != nil {
		return x.Datapoints
	}
	return nil
}

var File_v3_converter_proto protoreflect.FileDescriptor

var file_v3_converter_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x33, 0x1a, 0x0e, 0x76, 0x33, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x41, 0x72, 0x67, 0x73, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x12,
	0x2d, 0x0a, 0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x52, 0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x22, 0x4a,
	0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x0a,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x22, 0x74, 0x0a, 0x18, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x41, 0x72, 0x67, 0x73, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x22, 0x4a, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x32, 0xba, 0x01, 0x0a,
	0x10, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x52, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x63,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x6e, 0x64, 0x72, 0x61, 0x2d, 0x64, 0x61, 0x74,
	0x61, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f,
	0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v3_converter_proto_rawDescOnce sync.Once
	file_v3_converter_proto_rawDescData = file_v3_converter_proto_rawDesc
)

func file_v3_converter_proto_rawDescGZIP() []byte {
	file_v3_converter_proto_rawDescOnce.Do(func() {
		file_v3_converter_proto_rawDescData = protoimpl.X.CompressGZIP(file_v3_converter_proto_rawDescData)
	})
	return file_v3_converter_proto_rawDescData
}

var file_v3_converter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v3_converter_proto_goTypes = []interface{}{
	(*ConvertAggregatesRequest)(nil),  // 0: v3.ConvertAggregatesRequest
	(*ConvertAggregatesResponse)(nil), // 1: v3.ConvertAggregatesResponse
	(*ConvertDatapointsRequest)(nil),  // 2: v3.ConvertDatapointsRequest
	(*ConvertDatapointsResponse)(nil), // 3: v3.ConvertDatapointsResponse
	(*ConvertArgs)(nil),               // 4: v3.ConvertArgs
	(*Aggregate)(nil),                 // 5: v3.Aggregate
	(*Datapoint)(nil),                 // 6: v3.Datapoint
}
var file_v3_converter_proto_depIdxs = []int32{
	4, // 0: v3.ConvertAggregatesRequest.convert:type_name -> v3.ConvertArgs
	5, // 1: v3.ConvertAggregatesRequest.aggregates:type_name -> v3.Aggregate
	5, // 2: v3.ConvertAggregatesResponse.aggregates:type_name -> v3.Aggregate
	4, // 3: v3.ConvertDatapointsRequest.convert:type_name -> v3.ConvertArgs
	6, // 4: v3.ConvertDatapointsRequest.datapoints:type_name -> v3.Datapoint
	6, // 5: v3.ConvertDatapointsResponse.datapoints:type_name -> v3.Datapoint
	0, // 6: v3.ConverterService.ConvertAggregates:input_type -> v3.ConvertAggregatesRequest
	2, // 7: v3.ConverterService.ConvertDatapoints:input_type -> v3.ConvertDatapointsRequest
	1, // 8: v3.ConverterService.ConvertAggregates:output_type -> v3.ConvertAggregatesResponse
	3, // 9: v3.ConverterService.ConvertDatapoints:output_type -> v3.ConvertDatapointsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v3_converter_proto_init() }
func file_v3_converter_proto_init() {
	if File_v3_converter_proto != nil {
		return
	}
	file_v3_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v3_converter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertAggregatesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v3_converter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertAggregatesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v3_converter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertDatapointsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v3_converter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertDatapointsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v3_converter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v3_converter_proto_goTypes,
		DependencyIndexes: file_v3_converter_proto_depIdxs,
		MessageInfos:      file_v3_converter_proto_msgTypes,
	}.Build()
	File_v3_converter_proto = out.File
	file_v3_converter_proto_rawDesc = nil
	file_v3_converter_proto_goTypes = nil
	file_v3_converter_proto_depIdxs = nil
}
