// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.5
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

type ConvertManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUnit   string       `protobuf:"bytes,1,opt,name=from_unit,json=fromUnit,proto3" json:"from_unit,omitempty"`
	ToUnit     string       `protobuf:"bytes,2,opt,name=to_unit,json=toUnit,proto3" json:"to_unit,omitempty"`
	Datapoints []*Datapoint `protobuf:"bytes,3,rep,name=datapoints,proto3" json:"datapoints,omitempty"`
}

func (x *ConvertManyRequest) Reset() {
	*x = ConvertManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_converter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertManyRequest) ProtoMessage() {}

func (x *ConvertManyRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ConvertManyRequest.ProtoReflect.Descriptor instead.
func (*ConvertManyRequest) Descriptor() ([]byte, []int) {
	return file_v3_converter_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertManyRequest) GetFromUnit() string {
	if x != nil {
		return x.FromUnit
	}
	return ""
}

func (x *ConvertManyRequest) GetToUnit() string {
	if x != nil {
		return x.ToUnit
	}
	return ""
}

func (x *ConvertManyRequest) GetDatapoints() []*Datapoint {
	if x != nil {
		return x.Datapoints
	}
	return nil
}

type ConvertManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datapoints []*Datapoint `protobuf:"bytes,1,rep,name=datapoints,proto3" json:"datapoints,omitempty"`
}

func (x *ConvertManyResponse) Reset() {
	*x = ConvertManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_converter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertManyResponse) ProtoMessage() {}

func (x *ConvertManyResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ConvertManyResponse.ProtoReflect.Descriptor instead.
func (*ConvertManyResponse) Descriptor() ([]byte, []int) {
	return file_v3_converter_proto_rawDescGZIP(), []int{1}
}

func (x *ConvertManyResponse) GetDatapoints() []*Datapoint {
	if x != nil {
		return x.Datapoints
	}
	return nil
}

var File_v3_converter_proto protoreflect.FileDescriptor

var file_v3_converter_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x33, 0x1a, 0x0e, 0x76, 0x33, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x6f, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x22, 0x44, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x4d, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x32, 0x54, 0x0a, 0x10, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x16, 0x2e, 0x76,
	0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65,
	0x6e, 0x64, 0x72, 0x61, 0x73, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x6e, 0x64,
	0x72, 0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_v3_converter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v3_converter_proto_goTypes = []interface{}{
	(*ConvertManyRequest)(nil),  // 0: v3.ConvertManyRequest
	(*ConvertManyResponse)(nil), // 1: v3.ConvertManyResponse
	(*Datapoint)(nil),           // 2: v3.Datapoint
}
var file_v3_converter_proto_depIdxs = []int32{
	2, // 0: v3.ConvertManyRequest.datapoints:type_name -> v3.Datapoint
	2, // 1: v3.ConvertManyResponse.datapoints:type_name -> v3.Datapoint
	0, // 2: v3.ConverterService.ConvertMany:input_type -> v3.ConvertManyRequest
	1, // 3: v3.ConverterService.ConvertMany:output_type -> v3.ConvertManyResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v3_converter_proto_init() }
func file_v3_converter_proto_init() {
	if File_v3_converter_proto != nil {
		return
	}
	file_v3_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v3_converter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertManyRequest); i {
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
			switch v := v.(*ConvertManyResponse); i {
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
			NumMessages:   2,
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
