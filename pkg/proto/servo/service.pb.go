// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package servo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Client API for ServoService service

type ServoServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (ServoService_StreamClient, error)
}

type servoServiceClient struct {
	cc *grpc.ClientConn
}

func NewServoServiceClient(cc *grpc.ClientConn) ServoServiceClient {
	return &servoServiceClient{cc}
}

func (c *servoServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.servo.ServoService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servoServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.servo.ServoService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servoServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (ServoService_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ServoService_serviceDesc.Streams[0], c.cc, "/ai.metathings.service.servo.ServoService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &servoServiceStreamClient{stream}
	return x, nil
}

type ServoService_StreamClient interface {
	Send(*StreamRequests) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type servoServiceStreamClient struct {
	grpc.ClientStream
}

func (x *servoServiceStreamClient) Send(m *StreamRequests) error {
	return x.ClientStream.SendMsg(m)
}

func (x *servoServiceStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ServoService service

type ServoServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Stream(ServoService_StreamServer) error
}

func RegisterServoServiceServer(s *grpc.Server, srv ServoServiceServer) {
	s.RegisterService(&_ServoService_serviceDesc, srv)
}

func _ServoService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServoServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.servo.ServoService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServoServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.servo.ServoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServoServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServoService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServoServiceServer).Stream(&servoServiceStreamServer{stream})
}

type ServoService_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequests, error)
	grpc.ServerStream
}

type servoServiceStreamServer struct {
	grpc.ServerStream
}

func (x *servoServiceStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *servoServiceStreamServer) Recv() (*StreamRequests, error) {
	m := new(StreamRequests)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ServoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.servo.ServoService",
	HandlerType: (*ServoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ServoService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ServoService_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _ServoService_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_660b8c7f3beac646) }

var fileDescriptor_service_660b8c7f3beac646 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4e, 0xcc, 0xd4, 0xcb, 0x4d,
	0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x83, 0x49, 0x82, 0xe8, 0x7c, 0x29, 0xae,
	0x9c, 0xcc, 0xe2, 0x12, 0x88, 0x42, 0x29, 0xce, 0xf4, 0x54, 0x18, 0x93, 0xa7, 0xb8, 0xa4, 0x28,
	0x35, 0x31, 0x17, 0xc2, 0x33, 0xda, 0xcc, 0xc4, 0xc5, 0x13, 0x0c, 0x52, 0x1e, 0x0c, 0xd1, 0x2b,
	0x14, 0xcb, 0xc5, 0xe2, 0x93, 0x59, 0x5c, 0x22, 0xa4, 0xa1, 0x87, 0xc7, 0x6c, 0x3d, 0x90, 0x92,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29, 0x4d, 0x22, 0x54, 0x16, 0x17, 0xe4, 0xe7, 0x15,
	0xa7, 0x2a, 0x31, 0x08, 0x45, 0x71, 0x31, 0xbb, 0xa7, 0x96, 0x08, 0xa9, 0xe3, 0xd5, 0xe3, 0x9e,
	0x0a, 0x37, 0x5c, 0x83, 0xb0, 0x42, 0xb8, 0xd9, 0x19, 0x5c, 0x6c, 0xc1, 0x60, 0xbf, 0x09, 0x69,
	0xe3, 0xd5, 0x05, 0x51, 0x04, 0xb5, 0xa1, 0x58, 0x8a, 0x38, 0xc5, 0x30, 0x5b, 0x34, 0x18, 0x0d,
	0x18, 0x9d, 0xd8, 0xa3, 0x58, 0xc1, 0xb2, 0x49, 0x6c, 0xe0, 0x50, 0x34, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x45, 0x40, 0x9b, 0xd2, 0x98, 0x01, 0x00, 0x00,
}
