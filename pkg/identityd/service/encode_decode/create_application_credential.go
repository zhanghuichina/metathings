package encode_decode

import (
	"context"
	"time"

	"github.com/parnurzeal/gorequest"

	pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

type _createApplicationCredentialRequestBody_role struct {
	Id   *string `json:"id,ommitempty"`
	Name *string `json:"name,ommitempty"`
}

type createApplicationCredentialRequestBody struct {
	ApplicationCredential struct {
		Name         string                                         `json:"name"`
		Secret       *string                                        `json:"secret,omitempty"`
		Description  *string                                        `json:"description,omitempty"`
		ExpiresAt    *string                                        `json:"expires_at,omitempty"`
		Roles        []_createApplicationCredentialRequestBody_role `json:"roles,omitempty"`
		Unrestricted *bool                                          `json:"unrestricted,omitempty"`
	} `json:"application_credential"`
}

func EncodeCreateApplicationCredential(ctx context.Context, req *pb.CreateApplicationCredentialRequest) (res interface{}, err error) {
	body := &createApplicationCredentialRequestBody{}
	body.ApplicationCredential.Name = req.GetName().GetValue()

	secret := req.GetSecret()
	if secret != nil {
		body.ApplicationCredential.Secret = &secret.Value
	}

	description := req.GetDescription()
	if description != nil {
		body.ApplicationCredential.Description = &description.Value
	}

	expires_at := req.GetExpiresAt()
	if expires_at != nil {
		time_str := time.Unix(expires_at.GetSeconds(), 0).Format(time.RFC3339)
		body.ApplicationCredential.ExpiresAt = &time_str
	}

	roles := req.GetRoles()
	if len(roles) > 0 {
		rs := []_createApplicationCredentialRequestBody_role{}
		for _, r := range roles {
			id := r.GetId()
			name := r.GetName()
			if id != nil {
				rs = append(rs, _createApplicationCredentialRequestBody_role{Id: &id.Value})
			} else if name != nil {
				rs = append(rs, _createApplicationCredentialRequestBody_role{Name: &name.Value})
			}
		}
		body.ApplicationCredential.Roles = rs
	}

	unrestricted := req.GetUnrestricted()
	if unrestricted != nil {
		body.ApplicationCredential.Unrestricted = &unrestricted.Value
	}

	return body, nil
}

func DecodeCreateApplicationCredential(res gorequest.Response, body string) (*pb.CreateApplicationCredentialResponse, error) {
	app_cred, err := decodeApplicationCredential(res, body)
	if err != nil {
		return nil, err
	}

	return &pb.CreateApplicationCredentialResponse{ApplicationCredential: app_cred}, nil
}
