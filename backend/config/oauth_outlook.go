package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type OutlookOAuthSettings struct {
	ClientId     string `env:"OUTLOOK_OAUTH_CLIENT_ID"`
	ClientSecret string `env:"OUTLOOK_OAUTH_CLIENT_SECRET"`
}

func (i *OutlookOAuthSettings) into() (*OAuthSettings, error) {
	secretClientID, err := m.NewSecret(i.ClientId)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from client ID: %s", err.Error())
	}

	secretClientSecret, err := m.NewSecret(i.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from client secret: %s", err.Error())
	}

	return &OAuthSettings{
		BaseURL:      "https://login.microsoftonline.com/common/oauth2/v2.0",
		TokenURL:     "https://login.microsoftonline.com/common/oauth2/v2.0",
		ClientID:     secretClientID,
		ClientSecret: secretClientSecret,
		Scopes:       "offline_access user.read calendars.readwrite",
		ResponseType: "code",
		RedirectURI:  "http://127.0.0.1:3000",
		ResponseMode: "query",
		Prompt:       "consent",
	}, nil
}
