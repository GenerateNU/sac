package auth

type LoginResponseBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"` // MARK: must be validated manually
}

type UpdatePasswordRequestBody struct {
	OldPassword string `json:"old_password" validate:"required"`                                    // MARK: must be validated manually
	NewPassword string `json:"new_password" validate:"required,not_equal_if_not_empty=OldPassword"` // MARK: must be validated manually
}
