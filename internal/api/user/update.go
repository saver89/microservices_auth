package user

import (
	"context"
	"log/slog"

	desc "github.com/saver89/microservices_proto/pkg/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update updates a user
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	op := "userServer.Update"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	return &emptypb.Empty{}, nil
}
