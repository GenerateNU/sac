package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

const (
	tenantID = "a8eec281-aaa3-4dae-ac9b-9a398b9215e7"
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
		BaseURL:      fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0", tenantID),
		TokenURL:     fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0", tenantID),
		ClientID:     secretClientID,
		ClientSecret: secretClientSecret,
		Scopes:       "offline_access user.read calendars.readwrite openid email profile",
		ResponseType: "code",
		RedirectURI:  i.RedirectURI,
		ResponseMode: "query",
		Prompt:       "consent",
	}, nil
}
