package client_helper

import (
	"google.golang.org/grpc"

	cored_pb "github.com/bigdatagz/metathings/pkg/proto/core"
	identityd_pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

const (
	DEFAULT_CONFIG = iota
	IDENTITYD_CONFIG
	CORED_CONFIG
)

type CloseFn func()
type ServiceConfigs map[int]ServiceConfig
type DialOptionFn func() []grpc.DialOption

type ServiceConfig struct {
	Address string
}

type ClientFactory struct {
	defaultDialOptionFn DialOptionFn
	configs             ServiceConfigs
}

func (f *ClientFactory) NewConnection(cfg_val int, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if f.defaultDialOptionFn != nil {
		opts = append(opts, f.defaultDialOptionFn()...)
	}

	cfg, ok := f.configs[cfg_val]
	if !ok {
		cfg = f.configs[DEFAULT_CONFIG]
	}

	conn, err := grpc.Dial(cfg.Address, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (f *ClientFactory) NewCoreServiceClient(opts ...grpc.DialOption) (cored_pb.CoreServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(CORED_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		conn.Close()
	}

	return cored_pb.NewCoreServiceClient(conn), closeFn, nil
}

func (f *ClientFactory) NewIdentityServiceClient(opts ...grpc.DialOption) (identityd_pb.IdentityServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(IDENTITYD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		conn.Close()
	}

	return identityd_pb.NewIdentityServiceClient(conn), closeFn, nil
}

func NewClientFactory(configs ServiceConfigs, optFn DialOptionFn) (*ClientFactory, error) {
	if _, ok := configs[DEFAULT_CONFIG]; !ok {
		return nil, ErrMissingDefaultConfig
	}

	return &ClientFactory{
		configs:             configs,
		defaultDialOptionFn: optFn,
	}, nil
}
