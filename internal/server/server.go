package server

import (
	"context"
	"log/slog"

	"github.com/madhavan-raja/autorun/internal"
	"github.com/madhavan-raja/autorun/internal/autorun"
	"github.com/madhavan-raja/autorun/pb"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger.WithGroup("server")
}

type AutorunServer struct {
	pb.UnimplementedAutorunServer
	Autorun *autorun.Autorun
}

func (a *AutorunServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	logger.Info("Received Add Request", "req", req)

	id, err := a.Autorun.Add(req.Name, req.Description, req.Command, req.CronSchedule)
	if err != nil {
		return nil, err
	}

	return &pb.AddResponse{Id: id}, nil
}

func (a *AutorunServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	logger.Info("Received Update Request", "req", req)
	return &pb.UpdateResponse{}, nil
}

func (a *AutorunServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	logger.Info("Received Delete Request", "req", req)
	return &pb.DeleteResponse{}, nil
}

func (a *AutorunServer) Trigger(ctx context.Context, req *pb.TriggerRequest) (*pb.TriggerResponse, error) {
	logger.Info("Received Trigger Request", "req", req)
	return &pb.TriggerResponse{}, nil
}
