package client_helper

import (
	"fmt"
	"strings"

	"google.golang.org/grpc"

	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
	camera_pb "github.com/nayotta/metathings/pkg/proto/camera"
	camerad_pb "github.com/nayotta/metathings/pkg/proto/camerad"
	agent_pb "github.com/nayotta/metathings/pkg/proto/core_agent"
	cored_pb "github.com/nayotta/metathings/pkg/proto/cored"
	echo_pb "github.com/nayotta/metathings/pkg/proto/echo"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd"
	motor_pb "github.com/nayotta/metathings/pkg/proto/motor"
	sensor_pb "github.com/nayotta/metathings/pkg/proto/sensor"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
	servo_pb "github.com/nayotta/metathings/pkg/proto/servo"
	streamd_pb "github.com/nayotta/metathings/pkg/proto/streamd"
	switcher_pb "github.com/nayotta/metathings/pkg/proto/switcher"
)

const (
	DEFAULT_CONFIG = iota
	IDENTITYD_CONFIG
	CORED_CONFIG
	CAMERAD_CONFIG
	SENSORD_CONFIG
	AGENT_CONFIG
	ECHO_CONFIG
	SWITCHER_CONFIG
	MOTOR_CONFIG
	CAMERA_CONFIG
	SERVO_CONFIG
	SENSOR_CONFIG
	STREAMD_CONFIG
)

func parseAddress(addr string) string {
	if !strings.Contains(addr, ":") {
		addr = fmt.Sprintf("%v:%v", addr, constant_helper.CONSTANT_METATHINGSD_DEFAULT_PORT)
	}
	return addr
}

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

	conn, err := grpc.Dial(parseAddress(cfg.Address), opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (f *ClientFactory) NewCoredServiceClient(opts ...grpc.DialOption) (cored_pb.CoredServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(CORED_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return cored_pb.NewCoredServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewIdentitydServiceClient(opts ...grpc.DialOption) (identityd_pb.IdentitydServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(IDENTITYD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return identityd_pb.NewIdentitydServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewCoreAgentServiceClient(opts ...grpc.DialOption) (agent_pb.CoreAgentServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(AGENT_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return agent_pb.NewCoreAgentServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewEchoServiceClient(opts ...grpc.DialOption) (echo_pb.EchoServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(ECHO_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return echo_pb.NewEchoServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewSwitcherServiceClient(opts ...grpc.DialOption) (switcher_pb.SwitcherServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(SWITCHER_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return switcher_pb.NewSwitcherServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewMotorServiceClient(opts ...grpc.DialOption) (motor_pb.MotorServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(MOTOR_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return motor_pb.NewMotorServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewCameraServiceClient(opts ...grpc.DialOption) (camera_pb.CameraServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(CAMERA_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return camera_pb.NewCameraServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewCameradServiceClient(opts ...grpc.DialOption) (camerad_pb.CameradServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(CAMERAD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return camerad_pb.NewCameradServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewServoServiceClient(opts ...grpc.DialOption) (servo_pb.ServoServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(SERVO_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return servo_pb.NewServoServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewSensorServiceClient(opts ...grpc.DialOption) (sensor_pb.SensorServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(SENSOR_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return sensor_pb.NewSensorServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewSensordServiceClient(opts ...grpc.DialOption) (sensord_pb.SensordServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(SENSORD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return sensord_pb.NewSensordServiceClient(conn), func() { conn.Close() }, nil
}

func (f *ClientFactory) NewStreamdServiceClient(opts ...grpc.DialOption) (streamd_pb.StreamdServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(STREAMD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return streamd_pb.NewStreamdServiceClient(conn), func() { conn.Close() }, nil
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

func NewDefaultServiceConfigs(addr string) ServiceConfigs {
	return ServiceConfigs{
		DEFAULT_CONFIG: ServiceConfig{addr},
	}
}

func WithInsecureOptionFunc() DialOptionFn {
	return func() []grpc.DialOption {
		return []grpc.DialOption{grpc.WithInsecure()}
	}
}
