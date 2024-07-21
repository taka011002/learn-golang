package graph

import "learn-golang/src/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userUseCase usecase.UserUseCase
}

func NewResolver(userUseCase usecase.UserUseCase) *Resolver {
	return &Resolver{
		userUseCase: userUseCase,
	}
}
