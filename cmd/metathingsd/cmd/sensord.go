package cmd

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pb "github.com/nayotta/metathings/pkg/proto/sensord"
	service "github.com/nayotta/metathings/pkg/sensord/service"
)

type _hubOptions struct {
	Name string
}

type _sensordOptions struct {
	_rootOptions  `mapstructure:",squash"`
	Listen        string
	Storage       _storageOptions
	ServiceConfig _serviceConfigOptions `mapstructure:"service_config"`
	Hub           _hubOptions
}

var (
	sensord_opts *_sensordOptions
)

var (
	sensordCmd = &cobra.Command{
		Use:   "sensord",
		Short: "Sensor Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			cmd_helper.UnmarshalConfig(sensord_opts)
			root_opts = &sensord_opts._rootOptions
			sensord_opts.Service = "sensord"
			sensord_opts.Stage = cmd_helper.GetStageFromEnv()
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runSensord(); err != nil {
				log.WithError(err).Fatalf("failed to run sensord")
			}
		},
	}
)

func runSensord() error {
	lis, err := net.Listen("tcp", sensord_opts.Listen)
	if err != nil {
		return err
	}

	hub_v := cmd_helper.GetFromStage().Sub("hub")
	hub_opts := opt_helper.NewOption(
		"name", hub_v.GetString("name"),
		"options", hub_v,
	)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(nil)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(nil)),
	)
	srv, err := service.NewSensordService(
		service.SetLogLevel(sensord_opts.Log.Level),
		service.SetIdentitydAddr(sensord_opts.ServiceConfig.Identityd.Address),
		service.SetCoredAddr(sensord_opts.ServiceConfig.Cored.Address),
		service.SetStorage(sensord_opts.Storage.Driver, sensord_opts.Storage.Uri),
		service.SetApplicationCredential(
			sensord_opts.ApplicationCredential.Id,
			sensord_opts.ApplicationCredential.Secret,
		),
		service.SetHub(hub_opts),
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new sensor service")
		return err
	}

	pb.RegisterSensordServiceServer(s, srv)

	log.WithField("listen", sensord_opts.Listen).Infof("metathings sensor service listening")
	return s.Serve(lis)
}

func init() {
	sensord_opts = &_sensordOptions{}

	sensordCmd.Flags().StringVarP(&sensord_opts.Listen, "listen", "l", "127.0.0.1:5003", "Metathings Sensor Service listening address")
	sensordCmd.Flags().StringVar(&sensord_opts.ServiceConfig.Identityd.Address, "identityd-addr", "mt-api.nayotta.com", "MetaThings Identity Service Address")
	sensordCmd.Flags().StringVar(&sensord_opts.ServiceConfig.Cored.Address, "cored-addr", "mt-api.nayotta.com", "MetaThings Core Service Address")
	sensordCmd.Flags().StringVar(&sensord_opts.Storage.Driver, "storage-driver", "sqlite3", "Storage Driver [sqlite3]")
	sensordCmd.Flags().StringVar(&sensord_opts.Storage.Uri, "storage-uri", "", "Storage URI")

	RootCmd.AddCommand(sensordCmd)
}
