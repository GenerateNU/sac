package auth

type LoginResponseBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=255"` // MARK: must be validated manually
}

type UpdatePasswordRequestBody struct {
	OldPassword string `json:"old_password" validate:"required,max=255"`                                    // MARK: must be validated manually
	NewPassword string `json:"new_password" validate:"required,not_equal_if_not_empty=OldPassword,max=255"` // MARK: must be validated manually
}

type RefreshTokenRequestBody struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
