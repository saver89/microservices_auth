package repository

import (
	"context"

	"github.com/saver89/microservices_auth/internal/model"
)

// UserRepository is the interface for user repository
type UserRepository interface {
	Create(ctx context.Context, req model.CreateUserRequest) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, id int64, req model.UserInfo) error
	Delete(ctx context.Context, id int64) error
}

// UserLogRepository is the interface for user log repository
type UserLogRepository interface {
	Create(ctx context.Context, req model.UserLogInfo) (int64, error)
}
