package usecase

import (
	"context"
	"learn-golang/src/graph/model"
)

func (u *UseCase) GetUser(ctx context.Context, name string) (*model.User, error) {
	return &model.User{
		ID:   "1",
		Name: name,
	}, nil
}
