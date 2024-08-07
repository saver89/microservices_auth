package model

import (
	"database/sql"
	"time"
)

// UserInfo is the user info model
type UserInfo struct {
	Name  string
	Email string
	Role  string
}

// User is the user model
type User struct {
	ID        int64
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// CreateUserRequest is the request model for creating a user
type CreateUserRequest struct {
	Info     UserInfo
	Password string
}
