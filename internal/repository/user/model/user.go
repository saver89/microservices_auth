package model

import (
	"database/sql"
	"time"
)

// UserInfo is the user info model
type UserInfo struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Role  string `db:"role"`
}

// User is the user model
type User struct {
	ID        int64        `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
