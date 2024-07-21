package usecase

import "learn-golang/src/db/sqlc"

type UseCase struct {
	queries *sqlc.Queries
}

func NewUseCase(queries *sqlc.Queries) *UseCase {
	return &UseCase{
		queries: queries,
	}
}
