package tags

import "github.com/google/uuid"

type CreateUserTagsBody struct {
	Tags []uuid.UUID `json:"tags" validate:"required"`
}
