package base

import "github.com/google/uuid"

type CreateFileRequestBody struct {
	OwnerID   uuid.UUID `json:"owner_id" validate:"required,uuid4"`
	OwnerType string    `json:"owner_type" validate:"required,max=255"`
}
