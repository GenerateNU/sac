package models

import (
	"time"

	"github.com/google/uuid"
)

type OAuthResource string

const (
	Google  OAuthResource = "google"
	Outlook OAuthResource = "outlook"
)

type UserOAuthTokens struct {
	UserID       uuid.UUID     `json:"user_id" validate:"required,uuid4"`
	RefreshToken string        `json:"refresh_token" validate:"max=255"`
	AccessToken  string        `json:"access_token" validate:"max=255"`
	CSRFToken    string        `json:"csrf_token" validate:"max=255"`
	ResourceType OAuthResource `json:"resource_type" validate:"required"`
	ExpiresAt    time.Time     `json:"expires_at" validate:"required"`
}

type OAuthToken struct {
	AccessToken  string `json:"access_token" validate:"required"`
	ExpiresIn    int    `json:"expires_in" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
	Scope        string `json:"scope" validate:"required"`
	TokenType    string `json:"token_type" validate:"required"`
}

func (UserOAuthTokens) TableName() string {
	return "user_oauth_tokens"
}
