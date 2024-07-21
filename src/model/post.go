package model

import (
	"time"
)

type Post struct {
	ID        string
	UserId    string
	Title     string
	Content   string
	CreatedAt time.Time
}
