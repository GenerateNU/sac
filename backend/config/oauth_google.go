package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type GoogleOAuthSettings struct {
	ClientId     string `env:"GOOGLE_OAUTH_CLIENT_ID"`
	ClientSecret string `env:"GOOGLE_OAUTH_CLIENT_SECRET"`
}

func (i *GoogleOAuthSettings) into() (*OAuthSettings, error) {
	secretClientID, err := m.NewSecret(i.ClientId)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from client ID: %s", err.Error())
	}

	secretClientSecret, err := m.NewSecret(i.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from client secret: %s", err.Error())
	}

	return &OAuthSettings{
		BaseURL:              "https://accounts.google.com/o/oauth2/v2",
		TokenURL:             "https://oauth2.googleapis.com",
		ClientID:             secretClientID,
		ClientSecret:         secretClientSecret,
		Scopes:               "https://www.googleapis.com/auth/calendar.events https://www.googleapis.com/auth/calendar.readonly",
		ResponseType:         "code",
		RedirectURI:          "http://127.0.0.1:3000",
		IncludeGrantedScopes: "true",
		AccessType:           "offline",
		Prompt:               "consent",
	}, nil
}
