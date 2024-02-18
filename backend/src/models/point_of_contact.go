package models

import "github.com/google/uuid"

type PointOfContact struct {
	Model

	Name        string `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`
	Email       string `gorm:"uniqueIndex:compositeindex;index;not null;type:varchar(255)" json:"email" validate:"required,email,max=255"`
	Photo       string `gorm:"type:varchar(255);default:NULL" json:"photo" validate:"url,max=255"` // S3 URL, fallback to default logo if null
	Position    string `gorm:"type:varchar(255);" json:"position" validate:"required,max=255"`

	ClubID      uuid.UUID `gorm:"uniqueIndex:compositeindex;index;not null;foreignKey:ClubID" json:"-" validate:"min=1"`
	PhotoFileID uuid.UUID `gorm:"uniqueIndex:compositeindex;index;not null;foreignKey:FileID" json:"photo_file_id" validate:"min=1"`
}

type CreatePointOfContactBody struct {
	Name           string `json:"name" validate:"required,max=255"`
	Email          string `json:"email" validate:"required,email,max=255"`
	PhotoFileID    string `json:"photo_file_id" validate:"uuid4, max=255"` // S3 URL, fallback to default logo if null
	Position       string `json:"position" validate:"required,max=255"`
}
