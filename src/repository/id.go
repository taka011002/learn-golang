package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type IdGenerator interface {
	Generate() (*pgtype.UUID, error)
}

type uuidGenerator struct{}

func NewUuidGenerator() IdGenerator {
	return &uuidGenerator{}
}

func (u *uuidGenerator) Generate() (*pgtype.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &pgtype.UUID{Bytes: id, Valid: true}, nil
}

func uuidToString(uuid *pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid.Bytes[0:4], uuid.Bytes[4:6], uuid.Bytes[6:8], uuid.Bytes[8:10], uuid.Bytes[10:16])
}
