package model

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        string    `json; "id"`
	createdAt time.Time `json: "created_at"`
}

func NewBase() *Base {
	return &Base{
		ID:        uuid.New().String(),
		createdAt: time.Now(),
	}
}
