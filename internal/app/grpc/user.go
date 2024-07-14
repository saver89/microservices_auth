package grpc

import (
	"context"
	"log/slog"

	desc "github.com/saver89/microservices_proto/pkg/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userServer struct {
	log *slog.Logger
	desc.UnimplementedUserV1Server
}

// NewUserServer creates a new user server
func NewUserServer(log *slog.Logger) *userServer {
	return &userServer{
		log: log,
	}
}

func (cs *userServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	op := "userServer.Create"
	cs.log.InfoContext(ctx, op, slog.Any("req", req))
	return &desc.CreateResponse{}, nil
}

func (cs *userServer) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	op := "userServer.Get"
	cs.log.InfoContext(ctx, op, slog.Any("req", req))
	return &desc.GetResponse{}, nil
}

func (cs *userServer) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	op := "userServer.Update"
	cs.log.InfoContext(ctx, op, slog.Any("req", req))
	return &emptypb.Empty{}, nil
}

func (cs *userServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	op := "userServer.Delete"
	cs.log.InfoContext(ctx, op, slog.Any("req", req))
	return &emptypb.Empty{}, nil
}
