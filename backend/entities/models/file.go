package models

import (
	"github.com/google/uuid"
)

type FileInfo struct {
	FileName  string
	FileType  string
	FileSize  int
	FileURL   string
	ObjectKey string
}

type File struct {
	Model

	OwnerID   uuid.UUID `json:"-" validate:"required,uuid4"`
	OwnerType string    `json:"-" validate:"required,max=255"`

	FileName  string `json:"file_name" validate:"required,max=255"`
	FileType  string `json:"file_type" validate:"required,max=255"`
	FileSize  int    `json:"file_size" validate:"required,min=1"`
	FileURL   string `json:"file_url" validate:"required,max=255"`
	ObjectKey string `json:"object_key" validate:"required,max=255"`
}

type CreateFileRequestBody struct {
	OwnerID   uuid.UUID `json:"owner_id" validate:"required,uuid4"`
	OwnerType string    `json:"owner_type" validate:"required,max=255"`
}
