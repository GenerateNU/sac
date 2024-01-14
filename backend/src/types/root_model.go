package types

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID `gorm:"primarykey;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
