// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.5
// source: v3/transformer.proto

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

type TransformManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transform  *TransformArgs `protobuf:"bytes,1,opt,name=transform,proto3" json:"transform,omitempty"`
	Datapoints []*Datapoint   `protobuf:"bytes,2,rep,name=datapoints,proto3" json:"datapoints,omitempty"`
}

func (x *TransformManyRequest) Reset() {
	*x = TransformManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_transformer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformManyRequest) ProtoMessage() {}

func (x *TransformManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v3_transformer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformManyRequest.ProtoReflect.Descriptor instead.
func (*TransformManyRequest) Descriptor() ([]byte, []int) {
	return file_v3_transformer_proto_rawDescGZIP(), []int{0}
}

func (x *TransformManyRequest) GetTransform() *TransformArgs {
	if x != nil {
		return x.Transform
	}
	return nil
}

func (x *TransformManyRequest) GetDatapoints() []*Datapoint {
	if x != nil {
		return x.Datapoints
	}
	return nil
}

type TransformManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datapoints []*Datapoint `protobuf:"bytes,1,rep,name=datapoints,proto3" json:"datapoints,omitempty"`
}

func (x *TransformManyResponse) Reset() {
	*x = TransformManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_transformer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformManyResponse) ProtoMessage() {}

func (x *TransformManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v3_transformer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformManyResponse.ProtoReflect.Descriptor instead.
func (*TransformManyResponse) Descriptor() ([]byte, []int) {
	return file_v3_transformer_proto_rawDescGZIP(), []int{1}
}

func (x *TransformManyResponse) GetDatapoints() []*Datapoint {
	if x != nil {
		return x.Datapoints
	}
	return nil
}

var File_v3_transformer_proto protoreflect.FileDescriptor

var file_v3_transformer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x33, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x33, 0x1a, 0x0e, 0x76, 0x33, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x14, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x41, 0x72, 0x67, 0x73, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x2d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x22, 0x46, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x4d,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x32, 0x5c, 0x0a, 0x12, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x61, 0x6e,
	0x79, 0x12, 0x18, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x33,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x63, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x6e, 0x64, 0x72, 0x61, 0x2d, 0x64, 0x61, 0x74, 0x61,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f,
	0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v3_transformer_proto_rawDescOnce sync.Once
	file_v3_transformer_proto_rawDescData = file_v3_transformer_proto_rawDesc
)

func file_v3_transformer_proto_rawDescGZIP() []byte {
	file_v3_transformer_proto_rawDescOnce.Do(func() {
		file_v3_transformer_proto_rawDescData = protoimpl.X.CompressGZIP(file_v3_transformer_proto_rawDescData)
	})
	return file_v3_transformer_proto_rawDescData
}

var file_v3_transformer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v3_transformer_proto_goTypes = []interface{}{
	(*TransformManyRequest)(nil),  // 0: v3.TransformManyRequest
	(*TransformManyResponse)(nil), // 1: v3.TransformManyResponse
	(*TransformArgs)(nil),         // 2: v3.TransformArgs
	(*Datapoint)(nil),             // 3: v3.Datapoint
}
var file_v3_transformer_proto_depIdxs = []int32{
	2, // 0: v3.TransformManyRequest.transform:type_name -> v3.TransformArgs
	3, // 1: v3.TransformManyRequest.datapoints:type_name -> v3.Datapoint
	3, // 2: v3.TransformManyResponse.datapoints:type_name -> v3.Datapoint
	0, // 3: v3.TransformerService.TransformMany:input_type -> v3.TransformManyRequest
	1, // 4: v3.TransformerService.TransformMany:output_type -> v3.TransformManyResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v3_transformer_proto_init() }
func file_v3_transformer_proto_init() {
	if File_v3_transformer_proto != nil {
		return
	}
	file_v3_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v3_transformer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformManyRequest); i {
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
		file_v3_transformer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformManyResponse); i {
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
			RawDescriptor: file_v3_transformer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v3_transformer_proto_goTypes,
		DependencyIndexes: file_v3_transformer_proto_depIdxs,
		MessageInfos:      file_v3_transformer_proto_msgTypes,
	}.Build()
	File_v3_transformer_proto = out.File
	file_v3_transformer_proto_rawDesc = nil
	file_v3_transformer_proto_goTypes = nil
	file_v3_transformer_proto_depIdxs = nil
}