package cmd

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	service "github.com/bigdatagz/metathings/pkg/identity/service"
	pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

var (
	identityd_opts struct {
		bind      string
		ksBaseURL string
	}
)

var (
	identitydCmd = &cobra.Command{
		Use:    "identityd",
		Short:  "Identity Service Daemon",
		PreRun: globalPreRunHook,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runIdentityd(); err != nil {
				log.Fatalf("failed to run identityd: %v", err)
			}
		},
	}
)

func runIdentityd() error {
	lis, err := net.Listen("tcp", identityd_opts.bind)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	srv := service.NewIdentityService(
		service.SetKeystoneBaseURL(identityd_opts.ksBaseURL),
		service.SetLogLevel(V("log_level")),
	)

	pb.RegisterIdentityServiceServer(s, srv)
	log.Infof("metathings identity service listen on %v", identityd_opts.bind)
	return s.Serve(lis)
}

func init() {
	identitydCmd.Flags().StringVarP(&identityd_opts.bind, "bind", "b", "127.0.0.1:5000", "Metathings Identity Service binding address")
	identitydCmd.Flags().StringVar(&identityd_opts.ksBaseURL, "keystone-base-url", "http://localhost:35357", "Backend Keystone Base URL")

	RootCmd.AddCommand(identitydCmd)
}