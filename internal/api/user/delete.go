package user

import (
	"context"
	"log/slog"

	"github.com/saver89/microservices_auth/internal/log"
	desc "github.com/saver89/microservices_proto/pkg/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes a user
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	op := "userServer.Delete"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		i.log.ErrorContext(ctx, op, log.Err(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
