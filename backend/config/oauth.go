package config

import (
	"errors"
	"os"

	m "github.com/garrettladley/mattress"
)

type OAuthConfig struct {
	BaseURL 				string
	ClientID 				*m.Secret[string]
	ClientSecret 			*m.Secret[string]
	Scopes 					string
	RedirectURI 			string
	ResponseType			string
	ResponseMode    		string
	AccessType 				string
	IncludeGrantedScopes 	string
	Prompt 					string
}

type OAuthSettings struct {
	GoogleOAuthSettings 	*OAuthConfig
	OutlookOAuthSettings 	*OAuthConfig
}

func readGoogleOAuthSettings() (*OAuthConfig, error) {
	clientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	if clientID == "" {
		return nil, errors.New("GOOGLE_OAUTH_CLIENT_ID is not set")
	}

	secretClientID, err := m.NewSecret(clientID)
	if err != nil {
		return nil, errors.New("failed to create secret from client ID")
	}
	
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	if clientSecret == "" {
		return nil, errors.New("GOOGLE_OAUTH_CLIENT_SECRET is not set")
	}

	secretClientSecret, err := m.NewSecret(clientSecret)
	if err != nil {
		return nil, errors.New("failed to create secret from client secret")
	}

	return &OAuthConfig{
		BaseURL: "https://accounts.google.com/o/oauth2/v2",
		ClientID: secretClientID,
		ClientSecret: secretClientSecret,
		Scopes: "https://www.googleapis.com/auth/calendar.events https://www.googleapis.com/auth/calendar.readonly",
		ResponseType: "code",
		RedirectURI: "http://localhost:3000",
		IncludeGrantedScopes: "true",
		AccessType: "offline",
		Prompt: "consent",
	}, nil
}

func readOutlookOAuthSettings() (*OAuthConfig, error) {
	clientID := os.Getenv("OUTLOOK_OAUTH_CLIENT_ID")
	if clientID == "" {
		return nil, errors.New("OUTLOOK_OAUTH_CLIENT_ID is not set")
	}

	secretClientID, err := m.NewSecret(clientID)
	if err != nil {
		return nil, errors.New("failed to create secret from client ID")
	}

	clientSecret := os.Getenv("OUTLOOK_OAUTH_CLIENT_SECRET")
	if clientSecret == "" {
		return nil, errors.New("OUTLOOK_OAUTH_CLIENT_SECRET is not set")
	}

	secretClientSecret, err := m.NewSecret(clientSecret)
	if err != nil {
		return nil, errors.New("failed to create secret from client secret")
	}

	return &OAuthConfig{
		BaseURL: "https://login.microsoftonline.com/common/oauth2/v2.0",
		ClientID: secretClientID,
		ClientSecret: secretClientSecret,
		Scopes: "offline_access user.read mail.read",
		ResponseType: "code",
		RedirectURI: "http://localhost:3000",
		ResponseMode: "query",
		Prompt: "consent",
	}, nil
}
