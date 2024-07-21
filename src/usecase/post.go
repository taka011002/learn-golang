package usecase

import (
	"context"
	"learn-golang/src/model"
	"learn-golang/src/repository"
)

type PostUseCase interface {
	GetPost(ctx context.Context, id string) (*model.Post, error)
	CreatePost(ctx context.Context, title string, content *string) (*model.Post, error)
}

type postUseCase struct {
	repo repository.PostRepository
}

func NewPostUseCase(repo repository.PostRepository) PostUseCase {
	return &postUseCase{repo: repo}
}

func (u *postUseCase) GetPost(ctx context.Context, id string) (*model.Post, error) {
	return u.repo.GetPost(ctx, id)
}

func (u *postUseCase) CreatePost(ctx context.Context, title string, content *string) (*model.Post, error) {
	return u.repo.CreatePost(ctx, title, content)
}
