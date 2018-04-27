// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heartbeat.proto

package core_agent

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

func (this *HeartbeatRequest) Validate() error {
	if nil == this.EntityId {
		return go_proto_validators.FieldError("EntityId", fmt.Errorf("message must exist"))
	}
	if this.EntityId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.EntityId); err != nil {
			return go_proto_validators.FieldError("EntityId", err)
		}
	}
	return nil
}
