package graph

import "learn-golang/src/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	useCase usecase.UseCase
}

func NewResolver(useCase usecase.UseCase) *Resolver {
	return &Resolver{
		useCase: useCase,
	}
}
