package metathings_identityd2_service

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	log "github.com/sirupsen/logrus"
)

func (self *MetathingsIdentitydService) IssueTokenByPassword(ctx context.Context, req *pb.IssueTokenByPasswordRequest) (*pb.IssueTokenByPasswordResponse, error) {
	var ent_s *storage.Entity
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	ent := req.GetEntity()

	doms := ent.GetDomains()
	if len(doms) != 1 {
		err = errors.New("entity.domains must be set 1 domain with id")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	dom := doms[0]
	dom_id := dom.GetId()
	if dom_id == nil {
		err = errors.New("entity.domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	dom_id_str := dom_id.GetValue()

	ent_passwd := ent.GetPassword()
	if ent_passwd == nil {
		err = errors.New("entity.password is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	ent_passwd_str := ent_passwd.GetValue()

	ent_id := ent.GetId()
	ent_name := ent.GetName()
	if ent_name == nil && ent_id == nil {
		err = errors.New("entity.id and entity.name is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if ent_id != nil {
		if ent_s, err = self.storage.GetEntity(ent_id.GetValue()); err != nil {
			self.logger.WithError(err).Errorf("failed to find entity by id in storage")
			return nil, status.Errorf(codes.Unauthenticated, ErrUnauthenticated.Error())
		}
	} else {
		if ent_s, err = self.storage.GetEntityByName(ent_name.GetValue()); err != nil {
			self.logger.WithError(err).Errorf("failed to find entity by name in storage")
			return nil, status.Errorf(codes.Unauthenticated, ErrUnauthenticated.Error())
		}
	}

	if !domain_in_entity(ent_s, dom_id_str) {
		err = ErrUnauthenticated
		self.logger.WithError(err).Warningf("failed to find domain in entity")
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if !validate_password(*ent_s.Password, ent_passwd_str) {
		err = ErrUnauthenticated
		self.logger.WithError(err).Warningf("failed to validate password")
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	tkn_id_str := id_helper.NewId()
	now := time.Now()
	expires_at := now.Add(self.opt.TokenExpireInterval)

	tkn := &storage.Token{
		Id:        &tkn_id_str,
		DomainId:  &dom_id_str,
		EntityId:  ent_s.Id,
		IssuedAt:  &now,
		ExpiresAt: &expires_at,
		Text:      &tkn_id_str, // token text is token id now.
	}

	if tkn, err = self.storage.CreateToken(tkn); err != nil {
		self.logger.WithError(err).Errorf("failed to create token in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.IssueTokenByPasswordResponse{
		Token: copy_token(tkn),
	}

	self.logger.WithFields(log.Fields{
		"entity_id": *ent_s.Id,
		"domain_id": dom_id_str,
	}).Infof("issue token by password")

	return res, nil
}
