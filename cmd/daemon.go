package cmd

import (
	"context"
	"fmt"
	"net"

	"github.com/madhavan-raja/autorun/internal/ardaemon"
	"github.com/madhavan-raja/autorun/internal/server"
	"github.com/madhavan-raja/autorun/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start the Autorun Daemon",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		a := ardaemon.NewArDaemon(ctx)

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			logger.Error("Cannot create listener: %v", err)
		}

		s := grpc.NewServer()

		pb.RegisterArDaemonServer(s, &server.ArDaemonServer{ArDaemon: a})
		if err = s.Serve(lis); err != nil {
			logger.Error("Cannot serve: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
