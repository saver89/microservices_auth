package user

import (
	"log/slog"

	"github.com/saver89/microservices_auth/internal/service"
	desc "github.com/saver89/microservices_proto/pkg/user/v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server

	log         *slog.Logger
	userService service.UserService
}

func NewImplementation(log *slog.Logger, userService service.UserService) *Implementation {
	return &Implementation{
		log:         log,
		userService: userService,
	}
}
