package models

import "github.com/google/uuid"

type GrantType string

const (
	AuthorizationCode   ContactType = "authorization_code"
	RefreshToken  ContactType = "refresh_token"
)

type OAuthTokenRequestBody struct {
	Code	  	string    `json:"code" validate:"omitempty"`
	State 		string    `json:"state" validate:"omitempty"`
}

type OAuthToken struct {
	UserID 			uuid.UUID 	`json:"user_id" validate:"required,uuid4"`
	RefreshToken    string		`json:"refresh_token" validate:"max=255"`
	CSRFToken     	string		`json:"csrf_token" validate:"max=255"`
}
