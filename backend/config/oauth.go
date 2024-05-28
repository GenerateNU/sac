package config

import (
	"errors"

	m "github.com/garrettladley/mattress"
)

type OAuthSettings struct {
	BaseURL              string
	TokenURL             string
	ClientID             *m.Secret[string]
	ClientSecret         *m.Secret[string]
	Scopes               string
	RedirectURI          string
	ResponseType         string
	ResponseMode         string
	AccessType           string
	IncludeGrantedScopes string
	Prompt               string
}

type OAuthResources struct {
	GoogleOAuthSettings  *OAuthSettings
	OutlookOAuthSettings *OAuthSettings
}

type GoogleOAuthSettings struct {
	ClientId     string `env:"GOOGLE_OAUTH_CLIENT_ID"`
	ClientSecret string `env:"GOOGLE_OAUTH_CLIENT_SECRET"`
}

func (i *GoogleOAuthSettings) Into() (*OAuthSettings, error) {
	secretClientID, err := m.NewSecret(i.ClientId)
	if err != nil {
		return nil, errors.New("failed to create secret from client ID")
	}

	secretClientSecret, err := m.NewSecret(i.ClientSecret)
	if err != nil {
		return nil, errors.New("failed to create secret from client secret")
	}

	return &OAuthSettings{
		BaseURL:              "https://accounts.google.com/o/oauth2/v2",
		TokenURL:             "https://oauth2.googleapis.com",
		ClientID:             secretClientID,
		ClientSecret:         secretClientSecret,
		Scopes:               "https://www.googleapis.com/auth/calendar.events https://www.googleapis.com/auth/calendar.readonly",
		ResponseType:         "code",
		RedirectURI:          "http://localhost:3000",
		IncludeGrantedScopes: "true",
		AccessType:           "offline",
		Prompt:               "consent",
	}, nil
}

type OutlookOAuthSettings struct {
	ClientId     string `env:"OUTLOOK_OAUTH_CLIENT_ID"`
	ClientSecret string `env:"OUTLOOK_OAUTH_CLIENT_SECRET"`
}

func (i *OutlookOAuthSettings) Into() (*OAuthSettings, error) {
	secretClientID, err := m.NewSecret(i.ClientId)
	if err != nil {
		return nil, errors.New("failed to create secret from client ID")
	}

	secretClientSecret, err := m.NewSecret(i.ClientSecret)
	if err != nil {
		return nil, errors.New("failed to create secret from client secret")
	}

	return &OAuthSettings{
		BaseURL:      "https://login.microsoftonline.com/common/oauth2/v2.0",
		TokenURL:     "https://login.microsoftonline.com/common/oauth2/v2.0",
		ClientID:     secretClientID,
		ClientSecret: secretClientSecret,
		Scopes:       "offline_access user.read calendars.readwrite",
		ResponseType: "code",
		RedirectURI:  "http://localhost:3000",
		ResponseMode: "query",
		Prompt:       "consent",
	}, nil
}
