package converter

import (
	"github.com/saver89/microservices_auth/internal/model"
	modelRepo "github.com/saver89/microservices_auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		Id:        user.Id,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserInfoFromRepo(user modelRepo.UserInfo) model.UserInfo {
	return model.UserInfo{
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
	}
}
