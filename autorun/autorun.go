package autorun

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/madhavan-raja/autorun/internal"
	"github.com/madhavan-raja/autorun/internal/types"
	"google.golang.org/grpc"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger
}

type autorunServer struct {
	UnimplementedAutorunServer
}

func (a *autorunServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return &CreateResponse{}, nil
}

type Autorun struct {
	Target    string
	Port      uint32
	processes []types.Process
}

func New(target string) Autorun {
	return Autorun{target, 5678, []types.Process{}}
}

func (a *Autorun) LoadProcesses() {
	a.processes = []types.Process{
		{
			Name:        "Test",
			Description: "Test Process",
			Cmd:         "echo 'Hello World'",
			RunOnStart:  false,
			Repeat:      true,
			Interval:    5,
		},
	}
}

func (a *Autorun) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	if err != nil {
		logger.Error("Cannot create listener: %v", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &autorunServer{}

	RegisterAutorunServer(serverRegistrar, service)
	if err = serverRegistrar.Serve(lis); err != nil {
		logger.Error("Cannot serve: %s", err)
	}
}

func (a *Autorun) Run() {
	wg := sync.WaitGroup{}

	for _, p := range a.processes {
		wg.Add(1)
		go func() {
			defer wg.Done()

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

	wg.Wait()
}
