package repository

import (
	"context"
	"database/sql"
	"simple-project-native/cmd/myapp1/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Register(ctx context.Context, user *model.User) error {
	_, err := u.db.ExecContext(ctx,
		"INSERT INTO user (username, password) VALUES ($1, $2)",
		user.Username, user.Password,
	)
	return err
}

func (u *UserRepository) Login(ctx context.Context, user *model.User) error {

	err := u.db.QueryRowContext(ctx,
		"SELECT id_user, username,password FROM user WHERE username = $1",
		user.Username,
	).Scan(&user.ID, &user.Username, user.Password)

	return err
}
