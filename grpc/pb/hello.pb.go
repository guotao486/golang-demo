// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: pb/hello.proto

package pb

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

// 请求数据格式
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_pb_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 响应数据格式
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_pb_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// request
type StreamReqData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamReqData) Reset() {
	*x = StreamReqData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamReqData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamReqData) ProtoMessage() {}

func (x *StreamReqData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamReqData.ProtoReflect.Descriptor instead.
func (*StreamReqData) Descriptor() ([]byte, []int) {
	return file_pb_hello_proto_rawDescGZIP(), []int{2}
}

func (x *StreamReqData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// response
type StreamResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamResData) Reset() {
	*x = StreamResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResData) ProtoMessage() {}

func (x *StreamResData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResData.ProtoReflect.Descriptor instead.
func (*StreamResData) Descriptor() ([]byte, []int) {
	return file_pb_hello_proto_rawDescGZIP(), []int{3}
}

func (x *StreamResData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_pb_hello_proto protoreflect.FileDescriptor

var file_pb_hello_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x23, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xe9, 0x01, 0x0a, 0x07, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01,
	0x12, 0x38, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x0c, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_hello_proto_rawDescOnce sync.Once
	file_pb_hello_proto_rawDescData = file_pb_hello_proto_rawDesc
)

func file_pb_hello_proto_rawDescGZIP() []byte {
	file_pb_hello_proto_rawDescOnce.Do(func() {
		file_pb_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_hello_proto_rawDescData)
	})
	return file_pb_hello_proto_rawDescData
}

var file_pb_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_hello_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),  // 0: pb.HelloRequest
	(*HelloReply)(nil),    // 1: pb.HelloReply
	(*StreamReqData)(nil), // 2: pb.StreamReqData
	(*StreamResData)(nil), // 3: pb.StreamResData
}
var file_pb_hello_proto_depIdxs = []int32{
	0, // 0: pb.Greeter.SayHello:input_type -> pb.HelloRequest
	2, // 1: pb.Greeter.Streamserver:input_type -> pb.StreamReqData
	2, // 2: pb.Greeter.StreamClient:input_type -> pb.StreamReqData
	2, // 3: pb.Greeter.StreamAction:input_type -> pb.StreamReqData
	1, // 4: pb.Greeter.SayHello:output_type -> pb.HelloReply
	3, // 5: pb.Greeter.Streamserver:output_type -> pb.StreamResData
	3, // 6: pb.Greeter.StreamClient:output_type -> pb.StreamResData
	3, // 7: pb.Greeter.StreamAction:output_type -> pb.StreamResData
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_hello_proto_init() }
func file_pb_hello_proto_init() {
	if File_pb_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_pb_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_pb_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamReqData); i {
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
		file_pb_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResData); i {
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
			RawDescriptor: file_pb_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_hello_proto_goTypes,
		DependencyIndexes: file_pb_hello_proto_depIdxs,
		MessageInfos:      file_pb_hello_proto_msgTypes,
	}.Build()
	File_pb_hello_proto = out.File
	file_pb_hello_proto_rawDesc = nil
	file_pb_hello_proto_goTypes = nil
	file_pb_hello_proto_depIdxs = nil
}