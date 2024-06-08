package base

type emailRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}

type verifyEmailRequestBody struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required,len=6"`
}
