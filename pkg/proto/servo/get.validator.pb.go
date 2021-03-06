// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

/*
Package servo is a generated protocol buffer package.

It is generated from these files:
	get.proto
	list.proto
	service.proto
	servo.proto
	stream.proto

It has these top-level messages:
	GetRequest
	GetResponse
	ListRequest
	ListResponse
	Servo
	StreamPingRequest
	StreamPingResponse
	StreamSetStateRequest
	StreamSetAngleRequest
	StreamRequest
	StreamRequests
	StreamResponse
*/
package servo

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetRequest) Validate() error {
	if nil == this.Name {
		return go_proto_validators.FieldError("Name", fmt.Errorf("message must exist"))
	}
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	return nil
}
func (this *GetResponse) Validate() error {
	if this.Servo != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Servo); err != nil {
			return go_proto_validators.FieldError("Servo", err)
		}
	}
	return nil
}
