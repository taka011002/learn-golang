package repository

import (
	"context"
	"learn-golang/src/db/sqlc"
	"learn-golang/src/model"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, name string) (*model.User, error)
}

type userRepository struct {
	queries       *sqlc.Queries
	idGenerator   IdGenerator
	timeGenerator TimeGenerator
}

func NewUserRepository(queries *sqlc.Queries, idGenerator IdGenerator, timeGenerator TimeGenerator) UserRepository {
	return &userRepository{
		queries:       queries,
		idGenerator:   idGenerator,
		timeGenerator: timeGenerator,
	}
}

func (u *userRepository) GetUser(ctx context.Context, id string) (*model.User, error) {
	dbId := pgtype.UUID{}
	err := dbId.Scan(id)
	if err != nil {
		return nil, err
	}
	user, err := u.queries.GetUser(ctx, dbId)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Id:        uuidToString(&user.ID),
		Name:      user.Name,
		CreatedAt: timestampToTime(&user.CreatedAt),
	}, nil
}

func (u *userRepository) CreateUser(ctx context.Context, name string) (*model.User, error) {
	id, err := u.idGenerator.Generate()
	if err != nil {
		return nil, err
	}
	createdAt, err := u.timeGenerator.Now()
	if err != nil {
		return nil, err
	}

	user, err := u.queries.CreateUser(ctx, sqlc.CreateUserParams{ID: *id, Name: name, CreatedAt: *createdAt})
	if err != nil {
		return nil, err
	}

	return &model.User{
		Id:        uuidToString(&user.ID),
		Name:      user.Name,
		CreatedAt: timestampToTime(&user.CreatedAt),
	}, nil
}
