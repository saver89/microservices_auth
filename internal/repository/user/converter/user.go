package converter

import (
	"github.com/saver89/microservices_auth/internal/model"
	modelRepo "github.com/saver89/microservices_auth/internal/repository/user/model"
)

// ToUserFromRepo converts a user from repository to model
func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ToUserInfoFromRepo converts a user info from repository to model
func ToUserInfoFromRepo(user modelRepo.UserInfo) model.UserInfo {
	return model.UserInfo{
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
	}
}
