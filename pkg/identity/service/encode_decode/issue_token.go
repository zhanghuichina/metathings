package encode_decode

import (
	"context"
	"encoding/json"

	"github.com/parnurzeal/gorequest"

	pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

type _domain struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type _user struct {
	Id       string   `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Password string   `json:"password,omitempty"`
	Domain   *_domain `json:"domain,omitempty"`
}

type pwdIdentity struct {
	Methods  []string `json:"methods"`
	Password struct {
		User _user `json:"user"`
	} `json:"password"`
}

type tokenIdentity struct {
	Methods []string `json:"methods"`
	Token   struct {
		Id string `json:"id"`
	} `json:"token"`
}

type _appCredUser struct {
	Id     string   `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Domain *_domain `json:"domain,omitempty"`
}

type appCredIdentity struct {
	Methods               []string `json:"methods"`
	ApplicationCredential struct {
		Id     string        `json:"id,omitempty"`
		Name   string        `json:"name,omitempty"`
		Secret string        `json:"secret"`
		User   *_appCredUser `json:"user,omitempty"`
	} `json:"application_credential"`
}

type _scope struct {
	Project struct {
		Id string `json:"id"`
	} `json:"project"`
}

type issueTokenViaPasswordRequestBody struct {
	Auth struct {
		Identity pwdIdentity `json:"identity"`
		Scope    *_scope     `json:"scope,omitempty"`
	} `json:"auth"`
}

type issueTokenViaTokenRequestBody struct {
	Auth struct {
		Identity tokenIdentity `json:"identity"`
		Scope    *_scope       `json:"scope,omitempty"`
	} `json:"auth"`
}

type issueTokenViaApplicationCredentialRequestBody struct {
	Auth struct {
		Identity appCredIdentity `json:"identity"`
	} `json:"auth"`
}

type issueTokenResponseBody struct {
	Token struct {
		Project struct {
			Domain _domain `json:"domain"`
			Id     string  `json:"id"`
			Name   string  `json:"name"`
		} `json:"project"`
		User struct {
			Domain _domain `json:"domain"`
			Id     string  `json:"id"`
			Name   string  `json:"name"`
		} `json:"user"`
		Methods []string `json:"methods"`
		Roles   []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"roles"`
		IsDomain  bool   `json:"is_domain"`
		ExipresAt string `json:"exipres_at"`
		IssuedAt  string `json:"issued_at"`

		ApplicationCredentialRestricted bool `json:"application_credential_restricted"`
	} `json:"token"`
}

func EncodeIssueTokenRequest(ctx context.Context, req *pb.IssueTokenRequest) (interface{}, error) {
	switch req.GetMethod() {
	case pb.AUTH_METHOD_PASSWORD:
		return encodeIssueTokenViaPasswordRequest(ctx, req)
	case pb.AUTH_METHOD_TOKEN:
		return encodeIssueTokenViaTokenRequest(ctx, req)
	case pb.AUTH_METHOD_APPLICATION_CREDENTIAL:
		return encodeIssueTokenViaApplicationCredentialRequest(ctx, req)
	default:
		return nil, Unimplemented
	}
}

func encodeIssueTokenViaPasswordRequest(ctx context.Context, req *pb.IssueTokenRequest) (interface{}, error) {
	payload := req.GetPayload().(*pb.IssueTokenRequest_Password).Password
	user_id := payload.GetId()
	username := payload.GetUsername()
	password := payload.GetPassword()
	domain_id := payload.GetDomainId()
	domain_name := payload.GetDomainName()
	scope := payload.GetScope()
	body := issueTokenViaPasswordRequestBody{}
	idt := pwdIdentity{}

	idt.Methods = []string{"password"}

	if user_id != nil {
		idt.Password.User.Id = user_id.Value
	} else if username != nil {
		idt.Password.User.Name = username.Value
		if domain_id != nil || domain_name != nil {
			idt.Password.User.Domain = &_domain{}
		}
		if domain_id != nil {
			idt.Password.User.Domain.Id = domain_id.Value
		} else if domain_name != nil {
			idt.Password.User.Domain.Name = domain_name.Value
		}
	}

	if password != nil {
		idt.Password.User.Password = password.Value
	}

	if scope != nil {
		project_id := scope.GetProjectId()

		if project_id != nil {
			body.Auth.Scope = &_scope{}
			body.Auth.Scope.Project.Id = project_id.Value
		}
	}

	body.Auth.Identity = idt

	return &body, nil

}

func encodeIssueTokenViaTokenRequest(ctx context.Context, req *pb.IssueTokenRequest) (interface{}, error) {
	payload := req.GetPayload().(*pb.IssueTokenRequest_Token).Token
	token_id := payload.GetTokenId()
	scope := payload.GetScope()
	body := issueTokenViaTokenRequestBody{}
	idt := tokenIdentity{}

	idt.Methods = []string{"token"}

	if token_id != nil {
		idt.Token.Id = token_id.Value
	}

	if scope != nil {
		project_id := scope.GetProjectId()

		if project_id != nil {
			body.Auth.Scope = &_scope{}
			body.Auth.Scope.Project.Id = project_id.Value
		}
	}

	body.Auth.Identity = idt

	return &body, nil
}

func encodeIssueTokenViaApplicationCredentialRequest(ctx context.Context, req *pb.IssueTokenRequest) (interface{}, error) {
	payload := req.GetPayload().(*pb.IssueTokenRequest_ApplicationCredential).ApplicationCredential
	id := payload.GetId()
	name := payload.GetName()
	secret := payload.GetSecret()
	username := payload.GetUsername()
	user_id := payload.GetUserId()
	domain_id := payload.GetDomainId()
	body := issueTokenViaApplicationCredentialRequestBody{}
	idt := appCredIdentity{}

	idt.Methods = []string{"application_credential"}

	if id != nil {
		idt.ApplicationCredential.Id = id.Value
	} else if name != nil {
		idt.ApplicationCredential.Name = name.Value
		if user_id != nil || domain_id != nil {
			idt.ApplicationCredential.User = &_appCredUser{}
		}
		if user_id != nil {
			idt.ApplicationCredential.User.Id = user_id.Value
		} else if domain_id != nil {
			idt.ApplicationCredential.User.Domain = &_domain{}
			idt.ApplicationCredential.User.Domain.Id = domain_id.Value
			if username != nil {
				idt.ApplicationCredential.User.Name = username.Value
			}
		}
	}

	if secret != nil {
		idt.ApplicationCredential.Secret = secret.Value
	}

	body.Auth.Identity = idt

	return &body, nil
}

func DecodeIssueTokenResponse(_ gorequest.Response, body string) (*pb.IssueTokenResponse, error) {
	b := issueTokenResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}
	t := b.Token
	res := &pb.IssueTokenResponse{}
	token := &pb.Token{
		Methods:  t.Methods,
		IsDomain: t.IsDomain,
		User: &pb.Token__User{
			Id:   t.User.Id,
			Name: t.User.Name,
			Domain: &pb.Token__Domain{
				Id:   t.User.Domain.Id,
				Name: t.User.Domain.Name,
			},
		},
		Roles: []*pb.Token__Role{},
		Project: &pb.Token__Project{
			Id:   t.Project.Id,
			Name: t.Project.Name,
			Domain: &pb.Token__Domain{
				Id:   t.Project.Domain.Id,
				Name: t.Project.Domain.Name,
			},
		},
	}
	for _, r := range t.Roles {
		token.Roles = append(token.Roles, &pb.Token__Role{
			Id:   r.Id,
			Name: r.Name,
		})
	}
	res.Token = token

	return res, nil
}