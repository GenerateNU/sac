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

	OwnerID   uuid.UUID `json:"-"`
	OwnerType string    `json:"-"`

	FileName  string `json:"file_name"`
	FileType  string `json:"file_type"`
	FileSize  int    `json:"file_size"`
	FileURL   string `json:"file_url"`
	ObjectKey string `json:"object_key"`
}

type CreateFileRequestBody struct {
	OwnerID   uuid.UUID `json:"owner_id" validate:"required,uuid4"`
	OwnerType string    `json:"owner_type" validate:"required,max=255"`
}
