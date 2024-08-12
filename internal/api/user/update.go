package user

import (
	"context"
	"log/slog"

	"github.com/saver89/microservices_auth/internal/converter"
	"github.com/saver89/microservices_auth/internal/log"
	desc "github.com/saver89/microservices_proto/pkg/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update updates a user
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	op := "userServer.Update"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	err := i.userService.Update(ctx, req.Id, converter.ToUserInfoFromDesc(req.Info))
	if err != nil {
		i.log.ErrorContext(ctx, op, log.Err(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
