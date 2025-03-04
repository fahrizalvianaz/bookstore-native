package service

import (
	"context"
	"simple-project-native/cmd/myapp1/model"
	"simple-project-native/cmd/myapp1/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Register(ctx context.Context, user *model.User) error {
	return u.repo.Register(ctx, user)
}

func (u *UserService) Login(ctx context.Context, user *model.User) error {
	return u.repo.Login(ctx, user)
}
