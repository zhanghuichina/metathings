// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package switcher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SwitcherService service

type SwitcherServiceClient interface {
	Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetResponse, error)
	Turn(ctx context.Context, in *TurnRequest, opts ...grpc.CallOption) (*TurnResponse, error)
}

type switcherServiceClient struct {
	cc *grpc.ClientConn
}

func NewSwitcherServiceClient(cc *grpc.ClientConn) SwitcherServiceClient {
	return &switcherServiceClient{cc}
}

func (c *switcherServiceClient) Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.switcher.SwitcherService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *switcherServiceClient) Turn(ctx context.Context, in *TurnRequest, opts ...grpc.CallOption) (*TurnResponse, error) {
	out := new(TurnResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.switcher.SwitcherService/Turn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SwitcherService service

type SwitcherServiceServer interface {
	Get(context.Context, *empty.Empty) (*GetResponse, error)
	Turn(context.Context, *TurnRequest) (*TurnResponse, error)
}

func RegisterSwitcherServiceServer(s *grpc.Server, srv SwitcherServiceServer) {
	s.RegisterService(&_SwitcherService_serviceDesc, srv)
}

func _SwitcherService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwitcherServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.switcher.SwitcherService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwitcherServiceServer).Get(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwitcherService_Turn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwitcherServiceServer).Turn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.switcher.SwitcherService/Turn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwitcherServiceServer).Turn(ctx, req.(*TurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SwitcherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.switcher.SwitcherService",
	HandlerType: (*SwitcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SwitcherService_Get_Handler,
		},
		{
			MethodName: "Turn",
			Handler:    _SwitcherService_Turn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_0646b177a101f2bd) }

var fileDescriptor_service_0646b177a101f2bd = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xcc, 0xd4, 0xcb, 0x4d,
	0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x83, 0x49, 0x16, 0x97, 0x67, 0x96, 0x24,
	0x67, 0xa4, 0x16, 0x49, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x55, 0x27, 0x95,
	0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0x34, 0x4b, 0x71, 0xa6, 0xa7, 0x96, 0x40, 0x99,
	0x5c, 0x25, 0xa5, 0x45, 0x79, 0x10, 0xb6, 0xd1, 0x11, 0x46, 0x2e, 0xfe, 0x60, 0xa8, 0x01, 0xc1,
	0x10, 0x03, 0x85, 0x7c, 0xb8, 0x98, 0xdd, 0x53, 0x4b, 0x84, 0xc4, 0xf4, 0x20, 0xe6, 0xe9, 0xc1,
	0xcc, 0xd3, 0x73, 0x05, 0x99, 0x27, 0xa5, 0xad, 0x87, 0xdf, 0x1d, 0x7a, 0xee, 0xa9, 0x25, 0x41,
	0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0x0c, 0x42, 0xc9, 0x5c, 0x2c, 0x21, 0xa5, 0x45,
	0x79, 0x42, 0x04, 0xb5, 0x81, 0x54, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0xe9, 0x10,
	0xa7, 0x18, 0x66, 0x89, 0x13, 0x57, 0x14, 0x07, 0x4c, 0x2a, 0x89, 0x0d, 0xec, 0x5e, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xdf, 0x03, 0x3c, 0x3e, 0x01, 0x00, 0x00,
}
