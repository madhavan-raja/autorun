package server

import (
	"context"
	"log/slog"

	"github.com/madhavan-raja/autorun/internal"
	"github.com/madhavan-raja/autorun/internal/ardaemon"
	"github.com/madhavan-raja/autorun/pb"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger.WithGroup("server")
}

type ArDaemonServer struct {
	pb.UnimplementedArDaemonServer
	ArDaemon *ardaemon.ArDaemon
}

func (a *ArDaemonServer) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	logger.Info("Received List Request")
	processes := []*pb.Process{}

	for _, p := range a.ArDaemon.Jobs {
		processes = append(processes, &pb.Process{
			Id: p.Id,
			Name: p.Name,
			Description: p.Description,
			CronSchedule: "",
			Command: p.Cmd,
		})
	}

	return &pb.ListResponse{Processes: processes}, nil
}

func (a *ArDaemonServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	logger.Info("Received Add Request", "req", req)

	id, err := a.ArDaemon.Add(req.Name, req.Description, req.Command, req.CronSchedule)
	if err != nil {
		return nil, err
	}

	return &pb.AddResponse{Id: id}, nil
}

func (a *ArDaemonServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	logger.Info("Received Update Request", "req", req)
	return &pb.UpdateResponse{}, nil
}

func (a *ArDaemonServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	logger.Info("Received Delete Request", "req", req)
	return &pb.DeleteResponse{}, nil
}

func (a *ArDaemonServer) Trigger(ctx context.Context, req *pb.TriggerRequest) (*pb.TriggerResponse, error) {
	logger.Info("Received Trigger Request", "req", req)
	return &pb.TriggerResponse{}, nil
}
