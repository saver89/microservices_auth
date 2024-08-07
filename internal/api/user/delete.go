package user

import (
	"context"
	"log/slog"

	desc "github.com/saver89/microservices_proto/pkg/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	op := "userServer.Delete"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	return &emptypb.Empty{}, nil
}
