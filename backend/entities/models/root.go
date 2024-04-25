package models

import (
	"time"

	"github.com/google/uuid"
)

type Tabler interface {
	TableName() string
}

type Model struct {
	ID        uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	CreatedAt time.Time `json:"created_at" example:"2023-09-20T16:34:50Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-09-20T16:34:50Z"`
}
