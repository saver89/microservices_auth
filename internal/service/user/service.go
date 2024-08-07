package user

import (
	"github.com/saver89/microservices_auth/internal/client/db"
	"github.com/saver89/microservices_auth/internal/repository"
	"github.com/saver89/microservices_auth/internal/service"
)

const (
	createLog = "create"
	updateLog = "update"
	deleteLog = "delete"
)

type serv struct {
	userRepository    repository.UserRepository
	userLogRepository repository.UserLogRepository
	txManager         db.TxManager
}

// NewUserService creates a new user service
func NewUserService(
	userRepository repository.UserRepository,
	userLogRepository repository.UserLogRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepository:    userRepository,
		userLogRepository: userLogRepository,
		txManager:         txManager,
	}
}
