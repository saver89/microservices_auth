package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/saver89/microservices_auth/internal/client/db"
	"github.com/saver89/microservices_auth/internal/model"
	"github.com/saver89/microservices_auth/internal/repository"
	"github.com/saver89/microservices_auth/internal/repository/user/converter"
	modelRepo "github.com/saver89/microservices_auth/internal/repository/user/model"
)

const (
	tableName = "users"

	idColumn        = "id"
	emailColumn     = "email"
	roleColumn      = "role"
	nameColumn      = "name"
	passwordColumn  = "password_hash"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// NewUserRepository creates a new user repository
func NewUserRepository(db db.Client) repository.UserRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, user model.CreateUserRequest) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(emailColumn, roleColumn, nameColumn, passwordColumn).
		Values(user.Info.Email, user.Info.Role, user.Info.Name, user.Password).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, emailColumn, roleColumn, nameColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Update(ctx context.Context, id int64, user model.UserInfo) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(updatedAtColumn, "NOW()").
		Where(sq.Eq{idColumn: id})

	if user.Email != "" {
		builder = builder.Set(emailColumn, user.Email)
	}

	if user.Role != "" {
		builder = builder.Set(roleColumn, user.Role)
	}

	if user.Name != "" {
		builder = builder.Set(nameColumn, user.Name)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
