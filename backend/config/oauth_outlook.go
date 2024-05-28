package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type intermdeiateOutlookOAuthSettings struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
	RedirectURI  string `env:"REDIRECT_URI"`
}

func (i *intermdeiateOutlookOAuthSettings) into() (*OAuthSettings, error) {
	secretClientID, err := m.NewSecret(i.ClientID)
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
		RedirectURI:  i.RedirectURI,
		ResponseMode: "query",
		Prompt:       "consent",
	}, nil
}
