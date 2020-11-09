// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: rng.proto

package dtrngpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// RequestRngs - request some rngs
type RequestRngs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nums     int32  `protobuf:"varint,1,opt,name=nums,proto3" json:"nums,omitempty"`
	Gamecode string `protobuf:"bytes,2,opt,name=gamecode,proto3" json:"gamecode,omitempty"`
}

func (x *RequestRngs) Reset() {
	*x = RequestRngs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rng_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRngs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRngs) ProtoMessage() {}

func (x *RequestRngs) ProtoReflect() protoreflect.Message {
	mi := &file_rng_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRngs.ProtoReflect.Descriptor instead.
func (*RequestRngs) Descriptor() ([]byte, []int) {
	return file_rng_proto_rawDescGZIP(), []int{0}
}

func (x *RequestRngs) GetNums() int32 {
	if x != nil {
		return x.Nums
	}
	return 0
}

func (x *RequestRngs) GetGamecode() string {
	if x != nil {
		return x.Gamecode
	}
	return ""
}

// ReplyRngs - reply rngs
type ReplyRngs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rngs []int64 `protobuf:"varint,1,rep,packed,name=rngs,proto3" json:"rngs,omitempty"`
}

func (x *ReplyRngs) Reset() {
	*x = ReplyRngs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rng_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyRngs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyRngs) ProtoMessage() {}

func (x *ReplyRngs) ProtoReflect() protoreflect.Message {
	mi := &file_rng_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyRngs.ProtoReflect.Descriptor instead.
func (*ReplyRngs) Descriptor() ([]byte, []int) {
	return file_rng_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyRngs) GetRngs() []int64 {
	if x != nil {
		return x.Rngs
	}
	return nil
}

var File_rng_proto protoreflect.FileDescriptor

var file_rng_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x74, 0x72,
	0x6e, 0x67, 0x70, 0x62, 0x22, 0x3d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x1f, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x6e, 0x67, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04,
	0x72, 0x6e, 0x67, 0x73, 0x32, 0x3e, 0x0a, 0x05, 0x44, 0x54, 0x52, 0x6e, 0x67, 0x12, 0x35, 0x0a,
	0x07, 0x67, 0x65, 0x74, 0x52, 0x6e, 0x67, 0x73, 0x12, 0x14, 0x2e, 0x64, 0x74, 0x72, 0x6e, 0x67,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x6e, 0x67, 0x73, 0x1a, 0x12,
	0x2e, 0x64, 0x74, 0x72, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x6e,
	0x67, 0x73, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x73, 0x30, 0x30, 0x37, 0x2f, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x67,
	0x61, 0x6d, 0x65, 0x63, 0x6f, 0x72, 0x65, 0x37, 0x2f, 0x64, 0x74, 0x72, 0x6e, 0x67, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rng_proto_rawDescOnce sync.Once
	file_rng_proto_rawDescData = file_rng_proto_rawDesc
)

func file_rng_proto_rawDescGZIP() []byte {
	file_rng_proto_rawDescOnce.Do(func() {
		file_rng_proto_rawDescData = protoimpl.X.CompressGZIP(file_rng_proto_rawDescData)
	})
	return file_rng_proto_rawDescData
}

var file_rng_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rng_proto_goTypes = []interface{}{
	(*RequestRngs)(nil), // 0: dtrngpb.RequestRngs
	(*ReplyRngs)(nil),   // 1: dtrngpb.ReplyRngs
}
var file_rng_proto_depIdxs = []int32{
	0, // 0: dtrngpb.DTRng.getRngs:input_type -> dtrngpb.RequestRngs
	1, // 1: dtrngpb.DTRng.getRngs:output_type -> dtrngpb.ReplyRngs
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rng_proto_init() }
func file_rng_proto_init() {
	if File_rng_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rng_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRngs); i {
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
		file_rng_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyRngs); i {
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
			RawDescriptor: file_rng_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rng_proto_goTypes,
		DependencyIndexes: file_rng_proto_depIdxs,
		MessageInfos:      file_rng_proto_msgTypes,
	}.Build()
	File_rng_proto = out.File
	file_rng_proto_rawDesc = nil
	file_rng_proto_goTypes = nil
	file_rng_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DTRngClient is the client API for DTRng service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DTRngClient interface {
	// getRngs - get rngs
	GetRngs(ctx context.Context, in *RequestRngs, opts ...grpc.CallOption) (*ReplyRngs, error)
}

type dTRngClient struct {
	cc grpc.ClientConnInterface
}

func NewDTRngClient(cc grpc.ClientConnInterface) DTRngClient {
	return &dTRngClient{cc}
}

func (c *dTRngClient) GetRngs(ctx context.Context, in *RequestRngs, opts ...grpc.CallOption) (*ReplyRngs, error) {
	out := new(ReplyRngs)
	err := c.cc.Invoke(ctx, "/dtrngpb.DTRng/getRngs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DTRngServer is the server API for DTRng service.
type DTRngServer interface {
	// getRngs - get rngs
	GetRngs(context.Context, *RequestRngs) (*ReplyRngs, error)
}

// UnimplementedDTRngServer can be embedded to have forward compatible implementations.
type UnimplementedDTRngServer struct {
}

func (*UnimplementedDTRngServer) GetRngs(context.Context, *RequestRngs) (*ReplyRngs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRngs not implemented")
}

func RegisterDTRngServer(s *grpc.Server, srv DTRngServer) {
	s.RegisterService(&_DTRng_serviceDesc, srv)
}

func _DTRng_GetRngs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRngs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DTRngServer).GetRngs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtrngpb.DTRng/GetRngs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DTRngServer).GetRngs(ctx, req.(*RequestRngs))
	}
	return interceptor(ctx, in, info, handler)
}

var _DTRng_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dtrngpb.DTRng",
	HandlerType: (*DTRngServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getRngs",
			Handler:    _DTRng_GetRngs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rng.proto",
}
