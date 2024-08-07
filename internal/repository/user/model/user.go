package model

import (
	"database/sql"
	"time"
)

type UserInfo struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Role  string `db:"role"`
}

type User struct {
	Id        int64        `db:"id"`
	Info      UserInfo     `db:"info"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
