package usecase

import (
	"context"
	"learn-golang/src/model"
	"learn-golang/src/repository"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, name string) (*model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) GetUser(ctx context.Context, id string) (*model.User, error) {
	return u.repo.GetUser(ctx, id)
}

func (u *userUseCase) CreateUser(ctx context.Context, name string) (*model.User, error) {
	return u.repo.CreateUser(ctx, name)
}
