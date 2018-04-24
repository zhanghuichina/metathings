// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	core.proto
	create_core.proto
	delete_core.proto
	get_core.proto
	heartbeat.proto
	list_cores.proto
	list_cores_for_user.proto
	patch_core.proto
	send_unary_call.proto
	service.proto
	stream.proto
	unary_call.proto

It has these top-level messages:
	Core
	CreateCoreRequest
	CreateCoreResponse
	DeleteCoreRequest
	GetCoreRequest
	GetCoreResponse
	HeartbeatRequest
	ListCoresRequest
	ListCoresResponse
	ListCoresForUserRequest
	ListCoresForUserResponse
	PatchCoreRequest
	PatchCoreResponse
	SendUnaryCallRequest
	SendUnaryCallResponse
	StreamErrorResponsePayload
	StreamResponse
	StreamRequest
	UnaryCallRequestPayload
	UnaryCallResponsePayload
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Core) Validate() error {
	return nil
}
