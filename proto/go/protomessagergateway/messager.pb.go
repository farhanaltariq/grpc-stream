// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.28.2
// source: source/messagergateway/messager.proto

package protomessagergateway

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

type MapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alt string `protobuf:"bytes,1,opt,name=alt,proto3" json:"alt,omitempty"`
	Lat string `protobuf:"bytes,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon string `protobuf:"bytes,3,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *MapRequest) Reset() {
	*x = MapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_messagergateway_messager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRequest) ProtoMessage() {}

func (x *MapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_source_messagergateway_messager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapRequest.ProtoReflect.Descriptor instead.
func (*MapRequest) Descriptor() ([]byte, []int) {
	return file_source_messagergateway_messager_proto_rawDescGZIP(), []int{0}
}

func (x *MapRequest) GetAlt() string {
	if x != nil {
		return x.Alt
	}
	return ""
}

func (x *MapRequest) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *MapRequest) GetLon() string {
	if x != nil {
		return x.Lon
	}
	return ""
}

type MapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Alt    string `protobuf:"bytes,2,opt,name=alt,proto3" json:"alt,omitempty"`
	Lat    string `protobuf:"bytes,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon    string `protobuf:"bytes,4,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *MapResponse) Reset() {
	*x = MapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_messagergateway_messager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapResponse) ProtoMessage() {}

func (x *MapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_source_messagergateway_messager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapResponse.ProtoReflect.Descriptor instead.
func (*MapResponse) Descriptor() ([]byte, []int) {
	return file_source_messagergateway_messager_proto_rawDescGZIP(), []int{1}
}

func (x *MapResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MapResponse) GetAlt() string {
	if x != nil {
		return x.Alt
	}
	return ""
}

func (x *MapResponse) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *MapResponse) GetLon() string {
	if x != nil {
		return x.Lon
	}
	return ""
}

var File_source_messagergateway_messager_proto protoreflect.FileDescriptor

var file_source_messagergateway_messager_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x42, 0x0a,
	0x0a, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6c, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f,
	0x6e, 0x22, 0x5b, 0x0a, 0x0b, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x32, 0x58,
	0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x03,
	0x4d, 0x61, 0x70, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x4d, 0x61, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x72, 0x68, 0x61, 0x6e, 0x61, 0x6c, 0x74,
	0x61, 0x72, 0x69, 0x71, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x72, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_source_messagergateway_messager_proto_rawDescOnce sync.Once
	file_source_messagergateway_messager_proto_rawDescData = file_source_messagergateway_messager_proto_rawDesc
)

func file_source_messagergateway_messager_proto_rawDescGZIP() []byte {
	file_source_messagergateway_messager_proto_rawDescOnce.Do(func() {
		file_source_messagergateway_messager_proto_rawDescData = protoimpl.X.CompressGZIP(file_source_messagergateway_messager_proto_rawDescData)
	})
	return file_source_messagergateway_messager_proto_rawDescData
}

var file_source_messagergateway_messager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_source_messagergateway_messager_proto_goTypes = []interface{}{
	(*MapRequest)(nil),  // 0: protomessagergateway.MapRequest
	(*MapResponse)(nil), // 1: protomessagergateway.MapResponse
}
var file_source_messagergateway_messager_proto_depIdxs = []int32{
	0, // 0: protomessagergateway.MapService.Map:input_type -> protomessagergateway.MapRequest
	1, // 1: protomessagergateway.MapService.Map:output_type -> protomessagergateway.MapResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_source_messagergateway_messager_proto_init() }
func file_source_messagergateway_messager_proto_init() {
	if File_source_messagergateway_messager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_source_messagergateway_messager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapRequest); i {
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
		file_source_messagergateway_messager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapResponse); i {
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
			RawDescriptor: file_source_messagergateway_messager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_source_messagergateway_messager_proto_goTypes,
		DependencyIndexes: file_source_messagergateway_messager_proto_depIdxs,
		MessageInfos:      file_source_messagergateway_messager_proto_msgTypes,
	}.Build()
	File_source_messagergateway_messager_proto = out.File
	file_source_messagergateway_messager_proto_rawDesc = nil
	file_source_messagergateway_messager_proto_goTypes = nil
	file_source_messagergateway_messager_proto_depIdxs = nil
}