// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: src/proto/hello.proto

package hello

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

type SayHelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speech string `protobuf:"bytes,1,opt,name=speech,proto3" json:"speech,omitempty"`
}

func (x *SayHelloReq) Reset() {
	*x = SayHelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloReq) ProtoMessage() {}

func (x *SayHelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloReq.ProtoReflect.Descriptor instead.
func (*SayHelloReq) Descriptor() ([]byte, []int) {
	return file_src_proto_hello_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloReq) GetSpeech() string {
	if x != nil {
		return x.Speech
	}
	return ""
}

type SayHelloRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speech string `protobuf:"bytes,1,opt,name=speech,proto3" json:"speech,omitempty"`
}

func (x *SayHelloRes) Reset() {
	*x = SayHelloRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRes) ProtoMessage() {}

func (x *SayHelloRes) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRes.ProtoReflect.Descriptor instead.
func (*SayHelloRes) Descriptor() ([]byte, []int) {
	return file_src_proto_hello_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloRes) GetSpeech() string {
	if x != nil {
		return x.Speech
	}
	return ""
}

var File_src_proto_hello_proto protoreflect.FileDescriptor

var file_src_proto_hello_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x25,
	0x0a, 0x0b, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x32, 0x42, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73,
	0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_hello_proto_rawDescOnce sync.Once
	file_src_proto_hello_proto_rawDescData = file_src_proto_hello_proto_rawDesc
)

func file_src_proto_hello_proto_rawDescGZIP() []byte {
	file_src_proto_hello_proto_rawDescOnce.Do(func() {
		file_src_proto_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_hello_proto_rawDescData)
	})
	return file_src_proto_hello_proto_rawDescData
}

var file_src_proto_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_proto_hello_proto_goTypes = []interface{}{
	(*SayHelloReq)(nil), // 0: hello.SayHelloReq
	(*SayHelloRes)(nil), // 1: hello.SayHelloRes
}
var file_src_proto_hello_proto_depIdxs = []int32{
	0, // 0: hello.HelloService.SayHello:input_type -> hello.SayHelloReq
	1, // 1: hello.HelloService.SayHello:output_type -> hello.SayHelloRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_proto_hello_proto_init() }
func file_src_proto_hello_proto_init() {
	if File_src_proto_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloReq); i {
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
		file_src_proto_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRes); i {
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
			RawDescriptor: file_src_proto_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_hello_proto_goTypes,
		DependencyIndexes: file_src_proto_hello_proto_depIdxs,
		MessageInfos:      file_src_proto_hello_proto_msgTypes,
	}.Build()
	File_src_proto_hello_proto = out.File
	file_src_proto_hello_proto_rawDesc = nil
	file_src_proto_hello_proto_goTypes = nil
	file_src_proto_hello_proto_depIdxs = nil
}
