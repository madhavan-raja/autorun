package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/madhavan-raja/autorun/autorun"
	"github.com/madhavan-raja/autorun/internal"
	"github.com/madhavan-raja/autorun/internal/types"
	"google.golang.org/grpc"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger
}

type autorunServer struct {
	autorun.UnimplementedAutorunServer
	processes []types.Process
}

func (a *autorunServer) RunProcesses() {
	for _, p := range a.processes {
		go func() {
			if p.RunOnStart {
				logger.Info("Running Process On Start", "process", p.Name)
			}

			if p.Repeat {
				ticker := time.NewTicker(time.Duration(p.Interval) * time.Second)
				defer ticker.Stop()

				for range ticker.C {
					logger.Info("Running Process", "process", p.Name)
				}
			}
		}()
	}
}

func (a *autorunServer) Create(ctx context.Context, req *autorun.CreateRequest) (*autorun.CreateResponse, error) {
	logger.Info("Received", "create_request", req)
	return &autorun.CreateResponse{}, nil
}


func main() {
	port := uint32(5678)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error("Cannot create listener: %v", err)
	}

	serverRegistrar := grpc.NewServer()
	server := &autorunServer{
		processes: []types.Process{
			{
				Name:        "Test",
				Description: "Test Process",
				Cmd:         "echo 'Hello World'",
				RunOnStart:  false,
				Repeat:      true,
				Interval:    5,
			},
		},
	}

	server.RunProcesses()

	autorun.RegisterAutorunServer(serverRegistrar, server)
	if err = serverRegistrar.Serve(lis); err != nil {
		logger.Error("Cannot serve: %s", err)
	}
}
