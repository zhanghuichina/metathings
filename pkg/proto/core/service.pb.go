// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package core

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

// Client API for CoreService service

type CoreServiceClient interface {
	CreateCore(ctx context.Context, in *CreateCoreRequest, opts ...grpc.CallOption) (*CreateCoreResponse, error)
	DeleteCore(ctx context.Context, in *DeleteCoreRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PatchCore(ctx context.Context, in *PatchCoreRequest, opts ...grpc.CallOption) (*PatchCoreResponse, error)
	GetCore(ctx context.Context, in *GetCoreRequest, opts ...grpc.CallOption) (*GetCoreResponse, error)
	ListCores(ctx context.Context, in *ListCoresRequest, opts ...grpc.CallOption) (*ListCoresResponse, error)
	CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*CreateEntityResponse, error)
	DeleteEntity(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PatchEntity(ctx context.Context, in *PatchEntityRequest, opts ...grpc.CallOption) (*PatchEntityResponse, error)
	GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*GetEntityResponse, error)
	ListEntities(ctx context.Context, in *ListEntitiesRequest, opts ...grpc.CallOption) (*ListEntitiesResponse, error)
	ListEntitiesForCore(ctx context.Context, in *ListEntitiesForCoreRequest, opts ...grpc.CallOption) (*ListEntitiesForCoreResponse, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// NOTE: input is response body form core agentd,
	//       output is request body to core agentd.
	Stream(ctx context.Context, opts ...grpc.CallOption) (CoreService_StreamClient, error)
	ListCoresForUser(ctx context.Context, in *ListCoresForUserRequest, opts ...grpc.CallOption) (*ListCoresForUserResponse, error)
	UnaryCall(ctx context.Context, in *UnaryCallRequest, opts ...grpc.CallOption) (*UnaryCallResponse, error)
}

type coreServiceClient struct {
	cc *grpc.ClientConn
}

func NewCoreServiceClient(cc *grpc.ClientConn) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) CreateCore(ctx context.Context, in *CreateCoreRequest, opts ...grpc.CallOption) (*CreateCoreResponse, error) {
	out := new(CreateCoreResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/CreateCore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteCore(ctx context.Context, in *DeleteCoreRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/DeleteCore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) PatchCore(ctx context.Context, in *PatchCoreRequest, opts ...grpc.CallOption) (*PatchCoreResponse, error) {
	out := new(PatchCoreResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/PatchCore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetCore(ctx context.Context, in *GetCoreRequest, opts ...grpc.CallOption) (*GetCoreResponse, error) {
	out := new(GetCoreResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/GetCore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) ListCores(ctx context.Context, in *ListCoresRequest, opts ...grpc.CallOption) (*ListCoresResponse, error) {
	out := new(ListCoresResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/ListCores", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*CreateEntityResponse, error) {
	out := new(CreateEntityResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/CreateEntity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteEntity(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/DeleteEntity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) PatchEntity(ctx context.Context, in *PatchEntityRequest, opts ...grpc.CallOption) (*PatchEntityResponse, error) {
	out := new(PatchEntityResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/PatchEntity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*GetEntityResponse, error) {
	out := new(GetEntityResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/GetEntity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) ListEntities(ctx context.Context, in *ListEntitiesRequest, opts ...grpc.CallOption) (*ListEntitiesResponse, error) {
	out := new(ListEntitiesResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/ListEntities", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) ListEntitiesForCore(ctx context.Context, in *ListEntitiesForCoreRequest, opts ...grpc.CallOption) (*ListEntitiesForCoreResponse, error) {
	out := new(ListEntitiesForCoreResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/ListEntitiesForCore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/Heartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (CoreService_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CoreService_serviceDesc.Streams[0], c.cc, "/ai.metathings.service.core.CoreService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreServiceStreamClient{stream}
	return x, nil
}

type CoreService_StreamClient interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ClientStream
}

type coreServiceStreamClient struct {
	grpc.ClientStream
}

func (x *coreServiceStreamClient) Send(m *StreamResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *coreServiceStreamClient) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreServiceClient) ListCoresForUser(ctx context.Context, in *ListCoresForUserRequest, opts ...grpc.CallOption) (*ListCoresForUserResponse, error) {
	out := new(ListCoresForUserResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/ListCoresForUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) UnaryCall(ctx context.Context, in *UnaryCallRequest, opts ...grpc.CallOption) (*UnaryCallResponse, error) {
	out := new(UnaryCallResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CoreService service

type CoreServiceServer interface {
	CreateCore(context.Context, *CreateCoreRequest) (*CreateCoreResponse, error)
	DeleteCore(context.Context, *DeleteCoreRequest) (*empty.Empty, error)
	PatchCore(context.Context, *PatchCoreRequest) (*PatchCoreResponse, error)
	GetCore(context.Context, *GetCoreRequest) (*GetCoreResponse, error)
	ListCores(context.Context, *ListCoresRequest) (*ListCoresResponse, error)
	CreateEntity(context.Context, *CreateEntityRequest) (*CreateEntityResponse, error)
	DeleteEntity(context.Context, *DeleteEntityRequest) (*empty.Empty, error)
	PatchEntity(context.Context, *PatchEntityRequest) (*PatchEntityResponse, error)
	GetEntity(context.Context, *GetEntityRequest) (*GetEntityResponse, error)
	ListEntities(context.Context, *ListEntitiesRequest) (*ListEntitiesResponse, error)
	ListEntitiesForCore(context.Context, *ListEntitiesForCoreRequest) (*ListEntitiesForCoreResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*empty.Empty, error)
	// NOTE: input is response body form core agentd,
	//       output is request body to core agentd.
	Stream(CoreService_StreamServer) error
	ListCoresForUser(context.Context, *ListCoresForUserRequest) (*ListCoresForUserResponse, error)
	UnaryCall(context.Context, *UnaryCallRequest) (*UnaryCallResponse, error)
}

func RegisterCoreServiceServer(s *grpc.Server, srv CoreServiceServer) {
	s.RegisterService(&_CoreService_serviceDesc, srv)
}

func _CoreService_CreateCore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CreateCore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/CreateCore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CreateCore(ctx, req.(*CreateCoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteCore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteCore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/DeleteCore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteCore(ctx, req.(*DeleteCoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_PatchCore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchCoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).PatchCore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/PatchCore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).PatchCore(ctx, req.(*PatchCoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetCore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetCore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/GetCore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetCore(ctx, req.(*GetCoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_ListCores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ListCores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/ListCores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ListCores(ctx, req.(*ListCoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_CreateEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CreateEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/CreateEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CreateEntity(ctx, req.(*CreateEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/DeleteEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteEntity(ctx, req.(*DeleteEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_PatchEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).PatchEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/PatchEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).PatchEntity(ctx, req.(*PatchEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/GetEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetEntity(ctx, req.(*GetEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_ListEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ListEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/ListEntities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ListEntities(ctx, req.(*ListEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_ListEntitiesForCore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntitiesForCoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ListEntitiesForCore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/ListEntitiesForCore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ListEntitiesForCore(ctx, req.(*ListEntitiesForCoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreServiceServer).Stream(&coreServiceStreamServer{stream})
}

type CoreService_StreamServer interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ServerStream
}

type coreServiceStreamServer struct {
	grpc.ServerStream
}

func (x *coreServiceStreamServer) Send(m *StreamRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *coreServiceStreamServer) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CoreService_ListCoresForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCoresForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ListCoresForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/ListCoresForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ListCoresForUser(ctx, req.(*ListCoresForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).UnaryCall(ctx, req.(*UnaryCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.core.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCore",
			Handler:    _CoreService_CreateCore_Handler,
		},
		{
			MethodName: "DeleteCore",
			Handler:    _CoreService_DeleteCore_Handler,
		},
		{
			MethodName: "PatchCore",
			Handler:    _CoreService_PatchCore_Handler,
		},
		{
			MethodName: "GetCore",
			Handler:    _CoreService_GetCore_Handler,
		},
		{
			MethodName: "ListCores",
			Handler:    _CoreService_ListCores_Handler,
		},
		{
			MethodName: "CreateEntity",
			Handler:    _CoreService_CreateEntity_Handler,
		},
		{
			MethodName: "DeleteEntity",
			Handler:    _CoreService_DeleteEntity_Handler,
		},
		{
			MethodName: "PatchEntity",
			Handler:    _CoreService_PatchEntity_Handler,
		},
		{
			MethodName: "GetEntity",
			Handler:    _CoreService_GetEntity_Handler,
		},
		{
			MethodName: "ListEntities",
			Handler:    _CoreService_ListEntities_Handler,
		},
		{
			MethodName: "ListEntitiesForCore",
			Handler:    _CoreService_ListEntitiesForCore_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _CoreService_Heartbeat_Handler,
		},
		{
			MethodName: "ListCoresForUser",
			Handler:    _CoreService_ListCoresForUser_Handler,
		},
		{
			MethodName: "UnaryCall",
			Handler:    _CoreService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _CoreService_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_d7f29a35718d0785) }

var fileDescriptor_service_d7f29a35718d0785 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0x4b, 0x6f, 0x13, 0x31,
	0x10, 0xc7, 0x13, 0x09, 0x05, 0x65, 0x1a, 0x20, 0x38, 0x12, 0x12, 0x0b, 0xa7, 0x9c, 0x78, 0x34,
	0x4e, 0x45, 0x11, 0xdc, 0x09, 0x69, 0x39, 0x70, 0x40, 0x44, 0x39, 0xc0, 0x25, 0x72, 0xb6, 0xd3,
	0x64, 0xab, 0x4d, 0xbc, 0xd8, 0x0e, 0x52, 0x4f, 0x5c, 0xf9, 0xa8, 0x7c, 0x0c, 0xb4, 0x7e, 0xc5,
	0x3e, 0x74, 0xd7, 0x3d, 0xe6, 0xbf, 0xbf, 0x99, 0xbf, 0xe7, 0x61, 0x07, 0x1e, 0x49, 0x14, 0xbf,
	0x8b, 0x1c, 0x69, 0x25, 0xb8, 0xe2, 0x24, 0x63, 0x05, 0xdd, 0xa1, 0x62, 0x6a, 0x5b, 0xec, 0x37,
	0x92, 0xba, 0x8f, 0x39, 0x17, 0x98, 0xbd, 0xd8, 0x70, 0xbe, 0x29, 0x71, 0xaa, 0xc9, 0xf5, 0xe1,
	0x7a, 0x8a, 0xbb, 0x4a, 0xdd, 0x9a, 0xc0, 0xec, 0x69, 0x2e, 0x90, 0x29, 0x5c, 0xd5, 0xa4, 0x93,
	0xae, 0xb0, 0xc4, 0x58, 0x1a, 0x56, 0x4c, 0xe5, 0xdb, 0x50, 0x79, 0xbc, 0x41, 0x15, 0x11, 0x65,
	0x21, 0x8d, 0x20, 0xad, 0x32, 0xb2, 0x99, 0x71, 0xaf, 0x0a, 0x6f, 0x37, 0xb2, 0xb9, 0x23, 0x91,
	0x98, 0xec, 0x91, 0x36, 0xac, 0xf3, 0xc7, 0xa1, 0xda, 0x41, 0x4b, 0x85, 0x37, 0x79, 0x19, 0x89,
	0xab, 0x6b, 0x2e, 0xc2, 0x43, 0x3d, 0xd9, 0x22, 0x13, 0x6a, 0x8d, 0x4c, 0x59, 0x61, 0x20, 0x95,
	0x40, 0xb6, 0xb3, 0xbf, 0x9e, 0x1f, 0xcf, 0xac, 0x23, 0x0f, 0x12, 0x85, 0xb3, 0x3f, 0xec, 0x99,
	0xb8, 0x5d, 0xe5, 0xac, 0x2c, 0x8d, 0xf2, 0xee, 0xdf, 0x00, 0x4e, 0x66, 0x5c, 0xe0, 0xc2, 0xb4,
	0x96, 0xec, 0x00, 0x66, 0xba, 0xc0, 0x5a, 0x24, 0x13, 0x7a, 0xf7, 0x00, 0xe8, 0x91, 0xfb, 0x8e,
	0xbf, 0x0e, 0x28, 0x55, 0x46, 0x53, 0x71, 0x59, 0xf1, 0xbd, 0xc4, 0x71, 0x87, 0x2c, 0x01, 0x3e,
	0xeb, 0xd6, 0xb5, 0xdb, 0x1d, 0x39, 0x67, 0xf7, 0x8c, 0x9a, 0x15, 0xa0, 0x6e, 0x05, 0xe8, 0xbc,
	0x5e, 0x81, 0x71, 0x87, 0xdc, 0x40, 0xff, 0x5b, 0xdd, 0x7c, 0x9d, 0xf5, 0xb4, 0x29, 0xab, 0xc7,
	0x5c, 0xd2, 0x49, 0x22, 0xed, 0x4b, 0xb8, 0x82, 0x87, 0x97, 0xa8, 0xb4, 0xd3, 0x9b, 0xa6, 0x58,
	0x0b, 0x39, 0x9f, 0xb7, 0x49, 0xac, 0x77, 0xb9, 0x81, 0xfe, 0xd7, 0x42, 0x6a, 0x55, 0x36, 0x57,
	0xe4, 0xb1, 0xa4, 0x8a, 0x02, 0xda, 0x7b, 0x49, 0x18, 0x98, 0x61, 0xcd, 0xf5, 0xa2, 0x92, 0x69,
	0xfb, 0x58, 0x0d, 0xe9, 0x1c, 0xcf, 0xd2, 0x03, 0xbc, 0xe9, 0x0f, 0x18, 0x98, 0x09, 0xa7, 0x98,
	0x86, 0x64, 0xfb, 0x36, 0x54, 0x70, 0xa2, 0x07, 0x67, 0x33, 0xd3, 0xd6, 0x09, 0xc7, 0x89, 0xa7,
	0xc9, 0x7c, 0x38, 0xad, 0x4b, 0x54, 0xd6, 0xef, 0xb4, 0x65, 0xd2, 0xb1, 0xdb, 0x24, 0x91, 0x0e,
	0xa7, 0x55, 0x0f, 0x71, 0x6e, 0x1f, 0x8b, 0xe6, 0xc6, 0x85, 0x64, 0xd2, 0xb4, 0xe2, 0x00, 0x6f,
	0xfa, 0xb7, 0x0b, 0xa3, 0xf0, 0xd3, 0x05, 0x17, 0xfa, 0x06, 0x7c, 0x48, 0xcd, 0x65, 0x03, 0xdc,
	0x19, 0x3e, 0xde, 0x3b, 0xce, 0x1f, 0x65, 0x01, 0xfd, 0x2f, 0xee, 0x3d, 0x6c, 0xee, 0xb5, 0xc7,
	0xda, 0x57, 0x06, 0xa1, 0xb7, 0xd0, 0x6f, 0x6a, 0xf3, 0x9d, 0x36, 0x8c, 0x3b, 0x4c, 0xf6, 0x3a,
	0x85, 0xd5, 0xd6, 0xe3, 0xce, 0xab, 0xee, 0x59, 0x97, 0xfc, 0x81, 0xa1, 0xbf, 0x80, 0x17, 0x5c,
	0x2c, 0x25, 0x0a, 0x72, 0x9e, 0x74, 0x5d, 0x2d, 0xed, 0x2a, 0x79, 0x7f, 0xbf, 0xa0, 0x70, 0x51,
	0x97, 0xf5, 0x5f, 0xc2, 0x8c, 0x95, 0x65, 0x73, 0xf3, 0x3c, 0x96, 0xb4, 0xa8, 0x01, 0xed, 0xbc,
	0x3e, 0xf5, 0x7e, 0x3e, 0xa8, 0xbf, 0xad, 0x7b, 0xba, 0xdb, 0xe7, 0xff, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x9c, 0xce, 0xb0, 0x4b, 0xec, 0x07, 0x00, 0x00,
}
