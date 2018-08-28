// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

/*
Package streamd is a generated protocol buffer package.

It is generated from these files:
	create.proto
	delete.proto
	get.proto
	list.proto
	list_for_user.proto
	service.proto
	start.proto
	stop.proto
	stream.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	DeleteRequest
	GetRequest
	GetResponse
	ListRequest
	ListResponse
	ListForUserRequest
	ListForUserResponse
	StartRequest
	StartResponse
	StopRequest
	StopResponse
	Stream
	Pod
	Source
	Filter
	Sink
*/
package streamd

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateRequest) Validate() error {
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	for _, item := range this.Pods {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Pods", err)
			}
		}
	}
	return nil
}
func (this *CreateResponse) Validate() error {
	if this.Stream != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Stream); err != nil {
			return go_proto_validators.FieldError("Stream", err)
		}
	}
	return nil
}