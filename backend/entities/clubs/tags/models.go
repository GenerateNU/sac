package tags

import "github.com/google/uuid"

type CreateClubTagsRequestBody struct {
	Tags []uuid.UUID `json:"tags" validate:"required"`
}
