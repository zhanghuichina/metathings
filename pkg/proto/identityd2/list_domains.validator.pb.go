// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_domains.proto

package identityd2

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListDomainsRequest) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Alias != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return go_proto_validators.FieldError("Alias", err)
		}
	}
	if this.Parent != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Parent); err != nil {
			return go_proto_validators.FieldError("Parent", err)
		}
	}
	for _, item := range this.Children {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Children", err)
			}
		}
	}
	return nil
}
func (this *ListDomainsResponse) Validate() error {
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	return nil
}
