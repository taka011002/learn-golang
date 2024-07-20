package usecase

import (
	"context"
	"learn-golang/src/db"
	"learn-golang/src/db/sqlc"
	"learn-golang/src/graph/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (u *UseCase) GetUser(ctx context.Context, name string) (*model.User, error) {
	user, err := u.queries.GetUser(ctx, name)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:   db.UUIDtoString(&user.ID),
		Name: user.Name,
	}, nil
}

func (u *UseCase) CreateUser(ctx context.Context, name string) (*model.User, error) {
	uuidValue, err := uuid.NewV7()
	if err != nil {
		return nil, err

	}
	id := pgtype.UUID{Bytes: uuidValue, Valid: true}

	user, err := u.queries.CreateUser(ctx, sqlc.CreateUserParams{ID: id, Name: name})
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:   db.UUIDtoString(&user.ID),
		Name: user.Name,
	}, nil
}
