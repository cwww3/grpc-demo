// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: gym.proto

package gym

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//谁在健身
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Actions []string `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gym_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_gym_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_gym_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

//结果
type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gym_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_gym_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_gym_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Reply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_gym_proto protoreflect.FileDescriptor

var file_gym_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x06, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x2d, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x32, 0x28, 0x0a, 0x03, 0x47, 0x79, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x42, 0x6f, 0x64,
	0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x07, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x79, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gym_proto_rawDescOnce sync.Once
	file_gym_proto_rawDescData = file_gym_proto_rawDesc
)

func file_gym_proto_rawDescGZIP() []byte {
	file_gym_proto_rawDescOnce.Do(func() {
		file_gym_proto_rawDescData = protoimpl.X.CompressGZIP(file_gym_proto_rawDescData)
	})
	return file_gym_proto_rawDescData
}

var file_gym_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gym_proto_goTypes = []interface{}{
	(*Person)(nil), // 0: Person
	(*Reply)(nil),  // 1: Reply
}
var file_gym_proto_depIdxs = []int32{
	0, // 0: Gym.BodyBuilding:input_type -> Person
	1, // 1: Gym.BodyBuilding:output_type -> Reply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gym_proto_init() }
func file_gym_proto_init() {
	if File_gym_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gym_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_gym_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_gym_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gym_proto_goTypes,
		DependencyIndexes: file_gym_proto_depIdxs,
		MessageInfos:      file_gym_proto_msgTypes,
	}.Build()
	File_gym_proto = out.File
	file_gym_proto_rawDesc = nil
	file_gym_proto_goTypes = nil
	file_gym_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GymClient is the client API for Gym service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GymClient interface {
	BodyBuilding(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Reply, error)
}

type gymClient struct {
	cc grpc.ClientConnInterface
}

func NewGymClient(cc grpc.ClientConnInterface) GymClient {
	return &gymClient{cc}
}

func (c *gymClient) BodyBuilding(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Gym/BodyBuilding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GymServer is the server API for Gym service.
type GymServer interface {
	BodyBuilding(context.Context, *Person) (*Reply, error)
}

// UnimplementedGymServer can be embedded to have forward compatible implementations.
type UnimplementedGymServer struct {
}

func (*UnimplementedGymServer) BodyBuilding(context.Context, *Person) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BodyBuilding not implemented")
}

func RegisterGymServer(s *grpc.Server, srv GymServer) {
	s.RegisterService(&_Gym_serviceDesc, srv)
}

func _Gym_BodyBuilding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GymServer).BodyBuilding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Gym/BodyBuilding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GymServer).BodyBuilding(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gym_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Gym",
	HandlerType: (*GymServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BodyBuilding",
			Handler:    _Gym_BodyBuilding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gym.proto",
}
