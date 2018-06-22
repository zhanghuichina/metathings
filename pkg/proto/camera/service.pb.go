// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package camera

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

// Client API for CameraService service

type CameraServiceClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StopResponse, error)
	Show(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ShowResponse, error)
}

type cameraServiceClient struct {
	cc *grpc.ClientConn
}

func NewCameraServiceClient(cc *grpc.ClientConn) CameraServiceClient {
	return &cameraServiceClient{cc}
}

func (c *cameraServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camera.CameraService/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameraServiceClient) Stop(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camera.CameraService/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameraServiceClient) Show(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camera.CameraService/Show", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CameraService service

type CameraServiceServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *empty.Empty) (*StopResponse, error)
	Show(context.Context, *empty.Empty) (*ShowResponse, error)
}

func RegisterCameraServiceServer(s *grpc.Server, srv CameraServiceServer) {
	s.RegisterService(&_CameraService_serviceDesc, srv)
}

func _CameraService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameraServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camera.CameraService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameraServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameraService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameraServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camera.CameraService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameraServiceServer).Stop(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameraService_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameraServiceServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camera.CameraService/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameraServiceServer).Show(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CameraService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.camera.CameraService",
	HandlerType: (*CameraServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _CameraService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _CameraService_Stop_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _CameraService_Show_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_f8ac3da992eb591f) }

var fileDescriptor_service_f8ac3da992eb591f = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49, 0xcc, 0xd4, 0xcb, 0x4d,
	0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x83, 0x49, 0x26, 0x27, 0xe6, 0xa6, 0x16,
	0x25, 0x4a, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0xd5, 0x26, 0x95, 0xa6, 0xe9,
	0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0xb4, 0x4a, 0x71, 0x17, 0x97, 0x24, 0x16, 0x95, 0x40, 0x39,
	0x5c, 0xc5, 0x25, 0xf9, 0x05, 0x70, 0x76, 0x46, 0x7e, 0x39, 0x84, 0x6d, 0xd4, 0xcf, 0xc4, 0xc5,
	0xeb, 0x0c, 0x36, 0x2c, 0x18, 0x62, 0xb4, 0x50, 0x12, 0x17, 0x6b, 0x30, 0x48, 0xa3, 0x90, 0x96,
	0x1e, 0x3e, 0xbb, 0xf5, 0xc0, 0x8a, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0xb4, 0x89,
	0x52, 0x5b, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xaa, 0xc4, 0x20, 0xe4, 0xc3, 0xc5, 0x12, 0x5c, 0x92,
	0x5f, 0x20, 0x24, 0xa6, 0x07, 0xf1, 0x80, 0x1e, 0xcc, 0x03, 0x7a, 0xae, 0x20, 0x0f, 0x48, 0x11,
	0xb4, 0x3a, 0xbf, 0x00, 0xcd, 0xb4, 0x8c, 0xfc, 0x72, 0xb2, 0x4d, 0xcb, 0xc8, 0x2f, 0x47, 0x98,
	0xe6, 0xc4, 0x11, 0xc5, 0x06, 0x91, 0x48, 0x62, 0x03, 0x9b, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xd6, 0xda, 0x2b, 0x92, 0x93, 0x01, 0x00, 0x00,
}