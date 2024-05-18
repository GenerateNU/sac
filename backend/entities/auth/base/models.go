package base

type VerifyEmailRequestBody struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required,len=6"`
}

type VerifyPasswordResetTokenRequestBody struct {
	Token             string `json:"token" validate:"required"`
	NewPassword       string `json:"new_password" validate:"required,min=8,password"`
	VerifyNewPassword string `json:"verify_new_password" validate:"required,min=8,password,eqfield=NewPassword"`
}

type EmailRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}
