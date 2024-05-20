package tags

import "github.com/google/uuid"

type CreateEventTagsRequestBody struct {
	Tags []uuid.UUID `json:"tags" validate:"required"`
}
