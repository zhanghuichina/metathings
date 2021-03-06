// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: patch_domain.proto

package identityd

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

func (this *PatchDomainRequest) Validate() error {
	if nil == this.DomainId {
		return go_proto_validators.FieldError("DomainId", fmt.Errorf("message must exist"))
	}
	if this.DomainId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.DomainId); err != nil {
			return go_proto_validators.FieldError("DomainId", err)
		}
	}
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Description != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return go_proto_validators.FieldError("Description", err)
		}
	}
	if this.Enabled != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Enabled); err != nil {
			return go_proto_validators.FieldError("Enabled", err)
		}
	}
	return nil
}
func (this *PatchDomainResponse) Validate() error {
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	return nil
}
