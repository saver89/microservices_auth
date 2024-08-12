package user

import (
	"context"
	"log/slog"

	"github.com/saver89/microservices_auth/internal/converter"
	"github.com/saver89/microservices_auth/internal/log"
	desc "github.com/saver89/microservices_proto/pkg/user/v1"
)

// Create creates a user
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	op := "userServer.Create"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	id, err := i.userService.Create(ctx, *converter.ToCreateUserRequestFromDesc(req))
	if err != nil {
		i.log.ErrorContext(ctx, op, log.Err(err))
		return nil, err
	}
	res := &desc.CreateResponse{
		Id: id,
	}

	return res, nil
}
