package service

import (
	"context"

	"github.com/saver89/microservices_auth/internal/model"
)

// UserService is the interface for user service
type UserService interface {
	Create(ctx context.Context, req model.CreateUserRequest) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, id int64, user model.UserInfo) error
	Delete(ctx context.Context, id int64) error
}

// UserLogService is the interface for user log service
type UserLogService interface {
	Create(ctx context.Context, req model.UserLogInfo) (int64, error)
}
