package user_log

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/saver89/microservices_auth/internal/client/db"
	"github.com/saver89/microservices_auth/internal/model"
	"github.com/saver89/microservices_auth/internal/repository"
)

const (
	tableName = "user_logs"

	idColumn        = "id"
	userIDColumn    = "user_id"
	logColumn       = "log"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

// NewUserLogRepository creates a new user log repository
func NewUserLogRepository(db db.Client) repository.UserLogRepository {
	return &repo{
		db: db,
	}
}

// Create creates a new user log
func (r *repo) Create(ctx context.Context, log model.UserLogInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(userIDColumn, logColumn).
		Values(log.UserID, log.Log).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_log_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
