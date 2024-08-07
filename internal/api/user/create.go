package user

import (
	"context"
	"log/slog"

	desc "github.com/saver89/microservices_proto/pkg/user/v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	op := "userServer.Create"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	return &desc.CreateResponse{}, nil
}
