package models

import (
	"github.com/google/uuid"
)

type File struct {
	Model

	OwnerID   uuid.UUID `gorm:"uniqueIndex:compositeindex;index;not null;foreignKey:OwnerID" json:"-" validate:"uuid4"`
	OwnerType string    `gorm:"uniqueIndex:compositeindex;index;not null;type:varchar(255)" json:"-" validate:"required,max=255"`

	FileName string `gorm:"type:varchar(255)" json:"file_name" validate:"required,max=255"`
	FileType string `gorm:"type:varchar(255)" json:"file_type" validate:"required,max=255"`
	FileSize int    `gorm:"type:int" json:"file_size" validate:"required,min=1"`
	FileURL  string `gorm:"type:varchar(255)" json:"file_url" validate:"required,max=255"`
	ObjectKey string `gorm:"type:varchar(255)" json:"object_key" validate:"required,max=255"`
}

type CreateFileRequestBody struct {
	OwnerID   uuid.UUID `json:"owner_id" validate:"required,uuid4"`
	OwnerType string    `json:"owner_type" validate:"required,max=255"`
}
