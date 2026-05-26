package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/madhavan-raja/autorun/autorun"
	"github.com/madhavan-raja/autorun/internal"
	"github.com/madhavan-raja/autorun/pb"
	"google.golang.org/grpc"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger.WithGroup("main")
}

type autorunServer struct {
	pb.UnimplementedAutorunServer
	autorun *autorun.Autorun
}

func (a *autorunServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	logger.Info("Received Add Request", "req", req)

	id, err := a.autorun.Add(req.Name, req.Description, req.Command, req.CronSchedule)
	if err != nil {
		return nil, err
	}

	return &pb.AddResponse{Id: id}, nil
}

func (a *autorunServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	logger.Info("Received Update Request", "req", req)
	return &pb.UpdateResponse{}, nil
}

func (a *autorunServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	logger.Info("Received Delete Request", "req", req)
	return &pb.DeleteResponse{}, nil
}

func (a *autorunServer) Trigger(ctx context.Context, req *pb.TriggerRequest) (*pb.TriggerResponse, error) {
	logger.Info("Received Trigger Request", "req", req)
	return &pb.TriggerResponse{}, nil
}

func main() {
	port := uint32(5678)

	a := autorun.NewAutorun()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error("Cannot create listener: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterAutorunServer(s, &autorunServer{autorun: a})
	if err = s.Serve(lis); err != nil {
		logger.Error("Cannot serve: %s", err)
	}
}
