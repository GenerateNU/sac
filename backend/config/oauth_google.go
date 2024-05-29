package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type intermediateGoogleOAuthSettings struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
	RedirectURI  string `env:"REDIRECT_URI"`
}

func (i *intermediateGoogleOAuthSettings) into() (*OAuthSettings, error) {
	secretClientID, err := m.NewSecret(i.ClientID)
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
		RedirectURI:          i.RedirectURI,
		IncludeGrantedScopes: "true",
		AccessType:           "offline",
		Prompt:               "consent",
	}, nil
}
