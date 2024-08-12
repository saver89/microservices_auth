package user

import (
	"context"
	"log/slog"

	"github.com/saver89/microservices_auth/internal/converter"
	"github.com/saver89/microservices_auth/internal/log"
	desc "github.com/saver89/microservices_proto/pkg/user/v1"
)

// Get gets a user
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	op := "userServer.Get"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	user, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		i.log.ErrorContext(ctx, op, log.Err(err))
		return nil, err
	}

	return converter.ToDescGetResponseFromService(user), nil
}
