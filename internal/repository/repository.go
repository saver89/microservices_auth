package repository

import (
	"context"

	"github.com/saver89/microservices_auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, req model.CreateUserRequest) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, id int64, req model.UserInfo) error
	Delete(ctx context.Context, id int64) error
}

type UserLogRepository interface {
	Create(ctx context.Context, req model.UserLogInfo) (int64, error)
}
