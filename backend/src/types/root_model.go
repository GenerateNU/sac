package types

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primarykey" json:"id" example:"1"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at" example:"2023-09-20T16:34:50Z"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at" example:"2023-09-20T16:34:50Z"`
}
