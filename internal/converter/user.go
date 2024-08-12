package converter

import (
	"github.com/saver89/microservices_auth/internal/model"
	desc "github.com/saver89/microservices_proto/pkg/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToCreateUserRequestFromDesc converts a CreateRequest to CreateUserRequest
func ToCreateUserRequestFromDesc(req *desc.CreateRequest) *model.CreateUserRequest {
	return &model.CreateUserRequest{
		Info:     ToUserInfoFromDesc(req.Info),
		Password: req.Password,
	}
}

// ToUserInfoFromDesc converts a UserInfo to UserInfo
func ToUserInfoFromDesc(info *desc.UserInfo) model.UserInfo {
	return model.UserInfo{
		Email: info.Email,
		Name:  info.Name,
		Role:  ToStringFromRoleDesc(info.Role),
	}
}

// ToStringFromRoleDesc converts a Role to string
func ToStringFromRoleDesc(role desc.Role) string {
	switch role {
	case desc.Role_ADMIN:
		return model.RoleAdmin
	case desc.Role_USER:
		return model.RoleUser
	case desc.Role_UNKNOWN:
		return ""
	default:
		return model.RoleUser
	}
}

// ToDescGetResponseFromService converts a User to GetResponse
func ToDescGetResponseFromService(user *model.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.GetResponse{
		User: &desc.User{
			Id:        user.ID,
			Info:      ToDescUserInfoFromService(user.Info),
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: updatedAt,
		},
	}
}

// ToDescUserInfoFromService converts a UserInfo to UserInfo
func ToDescUserInfoFromService(info model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Email: info.Email,
		Name:  info.Name,
		Role:  ToRoleDescFromString(info.Role),
	}
}

// ToRoleDescFromString converts a string to Role
func ToRoleDescFromString(role string) desc.Role {
	switch role {
	case model.RoleAdmin:
		return desc.Role_ADMIN
	case model.RoleUser:
		return desc.Role_USER
	default:
		return desc.Role_UNKNOWN
	}
}
