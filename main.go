package main

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/madhavan-raja/autorun/internal"
	"github.com/madhavan-raja/autorun/internal/ardaemon"
	"github.com/madhavan-raja/autorun/internal/server"
	"github.com/madhavan-raja/autorun/pb"
	"google.golang.org/grpc"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger.WithGroup("main")
}

func main() {
	port := uint32(5678)

	a := ardaemon.NewArDaemon()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error("Cannot create listener: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterAutorunServer(s, &server.ArDaemonServer{Autorun: a})
	if err = s.Serve(lis); err != nil {
		logger.Error("Cannot serve: %s", err)
	}
}
