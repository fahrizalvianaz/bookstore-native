package repository

import (
	"context"
	"database/sql"
	"simple-project-native/cmd/myapp1/model/dto"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Register(ctx context.Context, user *dto.RegisterRequest) error {
	_, err := u.db.ExecContext(ctx,
		"INSERT INTO customer (nama, email, username, password) VALUES ($1, $2, $3, $4)",
		user.Nama, user.Email, user.Username, user.Password,
	)
	return err
}

func (u *UserRepository) Login(ctx context.Context, user *dto.LoginRequest) (string, error) {
	var passwordUser *string
	err := u.db.QueryRowContext(ctx,
		"SELECT username, password FROM customer WHERE username = $1",
		user.Username,
	).Scan(&user.Username, &passwordUser)

	return *passwordUser, err
}
