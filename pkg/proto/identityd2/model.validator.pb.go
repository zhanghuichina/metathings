// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package identityd2

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Domain) Validate() error {
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
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *OpDomain) Validate() error {
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
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Policy) Validate() error {
	if this.Role != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Role); err != nil {
			return go_proto_validators.FieldError("Role", err)
		}
	}
	return nil
}
func (this *OpPolicy) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Role != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Role); err != nil {
			return go_proto_validators.FieldError("Role", err)
		}
	}
	if this.Rule != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Rule); err != nil {
			return go_proto_validators.FieldError("Rule", err)
		}
	}
	if this.Description != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return go_proto_validators.FieldError("Description", err)
		}
	}
	return nil
}
func (this *Role) Validate() error {
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	for _, item := range this.Policies {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Policies", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *OpRole) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
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
	if this.Description != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return go_proto_validators.FieldError("Description", err)
		}
	}
	for _, item := range this.Policies {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Policies", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Entity) Validate() error {
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	for _, item := range this.Groups {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Groups", err)
			}
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *OpEntity) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	for _, item := range this.Groups {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Groups", err)
			}
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
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
	if this.Password != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Password); err != nil {
			return go_proto_validators.FieldError("Password", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Group) Validate() error {
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	for _, item := range this.Entities {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Entities", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *OpGroup) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	for _, item := range this.Entities {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Entities", err)
			}
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
	if this.Description != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return go_proto_validators.FieldError("Description", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Credential) Validate() error {
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	if this.Entity != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Entity); err != nil {
			return go_proto_validators.FieldError("Entity", err)
		}
	}
	if this.ExpiresAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ExpiresAt); err != nil {
			return go_proto_validators.FieldError("ExpiresAt", err)
		}
	}
	return nil
}
func (this *OpCredential) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	if this.Entity != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Entity); err != nil {
			return go_proto_validators.FieldError("Entity", err)
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
	if this.Secret != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return go_proto_validators.FieldError("Secret", err)
		}
	}
	if this.Description != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return go_proto_validators.FieldError("Description", err)
		}
	}
	if this.ExpiresAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ExpiresAt); err != nil {
			return go_proto_validators.FieldError("ExpiresAt", err)
		}
	}
	return nil
}
func (this *Token) Validate() error {
	if this.IssuedAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.IssuedAt); err != nil {
			return go_proto_validators.FieldError("IssuedAt", err)
		}
	}
	if this.ExpiresAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ExpiresAt); err != nil {
			return go_proto_validators.FieldError("ExpiresAt", err)
		}
	}
	if this.Entity != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Entity); err != nil {
			return go_proto_validators.FieldError("Entity", err)
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	if this.Credential != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Credential); err != nil {
			return go_proto_validators.FieldError("Credential", err)
		}
	}
	return nil
}
func (this *OpToken) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.IssuedAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.IssuedAt); err != nil {
			return go_proto_validators.FieldError("IssuedAt", err)
		}
	}
	if this.ExpiresAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ExpiresAt); err != nil {
			return go_proto_validators.FieldError("ExpiresAt", err)
		}
	}
	if this.Entity != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Entity); err != nil {
			return go_proto_validators.FieldError("Entity", err)
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	if this.Credential != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Credential); err != nil {
			return go_proto_validators.FieldError("Credential", err)
		}
	}
	if this.Text != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Text); err != nil {
			return go_proto_validators.FieldError("Text", err)
		}
	}
	return nil
}