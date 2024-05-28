package base

type OAuthTokenRequestBody struct {
	Code  string `json:"code" validate:"omitempty"`
	State string `json:"state" validate:"omitempty"`
}
