// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: volc.proto

package grpc

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

type WatchUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vid string `protobuf:"bytes,1,opt,name=vid,proto3" json:"vid,omitempty"`
}

func (x *WatchUploadRequest) Reset() {
	*x = WatchUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchUploadRequest) ProtoMessage() {}

func (x *WatchUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_volc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchUploadRequest.ProtoReflect.Descriptor instead.
func (*WatchUploadRequest) Descriptor() ([]byte, []int) {
	return file_volc_proto_rawDescGZIP(), []int{0}
}

func (x *WatchUploadRequest) GetVid() string {
	if x != nil {
		return x.Vid
	}
	return ""
}

type WatchUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WatchUploadResponse) Reset() {
	*x = WatchUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchUploadResponse) ProtoMessage() {}

func (x *WatchUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_volc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchUploadResponse.ProtoReflect.Descriptor instead.
func (*WatchUploadResponse) Descriptor() ([]byte, []int) {
	return file_volc_proto_rawDescGZIP(), []int{1}
}

func (x *WatchUploadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetVidInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vid string `protobuf:"bytes,1,opt,name=vid,proto3" json:"vid,omitempty"`
}

func (x *GetVidInfoRequest) Reset() {
	*x = GetVidInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVidInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVidInfoRequest) ProtoMessage() {}

func (x *GetVidInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_volc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVidInfoRequest.ProtoReflect.Descriptor instead.
func (*GetVidInfoRequest) Descriptor() ([]byte, []int) {
	return file_volc_proto_rawDescGZIP(), []int{2}
}

func (x *GetVidInfoRequest) GetVid() string {
	if x != nil {
		return x.Vid
	}
	return ""
}

type GetVidInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetVidInfoResponse) Reset() {
	*x = GetVidInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVidInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVidInfoResponse) ProtoMessage() {}

func (x *GetVidInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_volc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVidInfoResponse.ProtoReflect.Descriptor instead.
func (*GetVidInfoResponse) Descriptor() ([]byte, []int) {
	return file_volc_proto_rawDescGZIP(), []int{3}
}

func (x *GetVidInfoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_volc_proto protoreflect.FileDescriptor

var file_volc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x12,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x76, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x57, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x80, 0x01, 0x0a,
	0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x13, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x62,
	0x77, 0x30, 0x30, 0x34, 0x36, 0x2f, 0x76, 0x6f, 0x6c, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_volc_proto_rawDescOnce sync.Once
	file_volc_proto_rawDescData = file_volc_proto_rawDesc
)

func file_volc_proto_rawDescGZIP() []byte {
	file_volc_proto_rawDescOnce.Do(func() {
		file_volc_proto_rawDescData = protoimpl.X.CompressGZIP(file_volc_proto_rawDescData)
	})
	return file_volc_proto_rawDescData
}

var file_volc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_volc_proto_goTypes = []interface{}{
	(*WatchUploadRequest)(nil),  // 0: WatchUploadRequest
	(*WatchUploadResponse)(nil), // 1: WatchUploadResponse
	(*GetVidInfoRequest)(nil),   // 2: GetVidInfoRequest
	(*GetVidInfoResponse)(nil),  // 3: GetVidInfoResponse
}
var file_volc_proto_depIdxs = []int32{
	0, // 0: backend.WatchUpload:input_type -> WatchUploadRequest
	2, // 1: backend.GetVidInfo:input_type -> GetVidInfoRequest
	1, // 2: backend.WatchUpload:output_type -> WatchUploadResponse
	3, // 3: backend.GetVidInfo:output_type -> GetVidInfoResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_volc_proto_init() }
func file_volc_proto_init() {
	if File_volc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_volc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchUploadRequest); i {
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
		file_volc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchUploadResponse); i {
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
		file_volc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVidInfoRequest); i {
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
		file_volc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVidInfoResponse); i {
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
			RawDescriptor: file_volc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_volc_proto_goTypes,
		DependencyIndexes: file_volc_proto_depIdxs,
		MessageInfos:      file_volc_proto_msgTypes,
	}.Build()
	File_volc_proto = out.File
	file_volc_proto_rawDesc = nil
	file_volc_proto_goTypes = nil
	file_volc_proto_depIdxs = nil
}
