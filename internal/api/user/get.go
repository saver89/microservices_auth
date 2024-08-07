package user

import (
	"context"
	"log/slog"

	desc "github.com/saver89/microservices_proto/pkg/user/v1"
)

// Get gets a user
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	op := "userServer.Get"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	return &desc.GetResponse{}, nil
}
