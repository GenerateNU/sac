package leadership

type CreateLeaderBody struct {
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Position string `json:"position" validate:"required,max=255"`
}

type UpdateLeaderBody struct {
	Name     string `json:"name" validate:"omitempty,max=255"`
	Email    string `json:"email" validate:"omitempty,email,max=255"`
	Position string `json:"position" validate:"omitempty,max=255"`
}
