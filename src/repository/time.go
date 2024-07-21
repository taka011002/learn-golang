package repository

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type TimeGenerator interface {
	Now() (*pgtype.Timestamp, error)
}
type timeGenerator struct{}

func NewTimeGenerator() TimeGenerator {
	return &timeGenerator{}
}

func (t *timeGenerator) Now() (*pgtype.Timestamp, error) {
	timestamp := pgtype.Timestamp{}
	err := timestamp.Scan(time.Now())
	if err != nil {
		return nil, err
	}

	return &timestamp, err
}

func timestampToTime(timestamp *pgtype.Timestamp) time.Time {
	return timestamp.Time
}
