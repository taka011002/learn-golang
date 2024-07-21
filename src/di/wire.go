//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"learn-golang/src/db"
	"learn-golang/src/graph"
	"learn-golang/src/repository"
	"learn-golang/src/usecase"

	"github.com/google/wire"
)

func InitializeResolver(ctx context.Context) (*graph.Resolver, func(), error) {
	wire.Build(
		graph.NewResolver,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		repository.NewTimeGenerator,
		repository.NewUuidGenerator,
		db.NewConfig,
		db.NewQueries,
	)

	return &graph.Resolver{}, nil, nil
}
