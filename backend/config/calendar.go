package config

import (
	"errors"
	"os"

	m "github.com/garrettladley/mattress"
)

type GoogleClientSettings struct {
	ClientID				*m.Secret[string]
	ClientSecret 			*m.Secret[string]
	ResponseType 			string
	Scopes 					[]string
	RedirectURI 			string
	IncludeGrantedScopes 	string
	AccessType 				string
	Prompt 					string
}

type OutlookClientSettings struct {
	ClientID 				string
	ResponseType 			string
	Scopes 					string
	RedirectURI 			string
	Tenant 					string
	ResponseMode 			string
	Prompt 					string
}

func readGoogleOAuthSettings() (*GoogleClientSettings, error) {
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

	return &GoogleClientSettings{
		ClientID: secretClientID,
		ClientSecret: secretClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/calendar.events",
			"https://www.googleapis.com/auth/calendar.readonly",
		},
		RedirectURI: "http://localhost:3000",
		IncludeGrantedScopes: "true",
		AccessType: "offline",
	}, nil
}

func readOutlookOAuthSettings() (*OutlookClientSettings, error) {
	clientID := os.Getenv("OUTLOOK_OAUTH_CLIENT_ID")
	if clientID == "" {
		return nil, errors.New("OUTLOOK_OAUTH_CLIENT_ID is not set")
	}

	clientSecret := os.Getenv("OUTLOOK_OAUTH_CLIENT_SECRET")
	if clientSecret == "" {
		return nil, errors.New("OUTLOOK_OAUTH_CLIENT_SECRET is not set")
	}

	return &OutlookClientSettings{
		ClientID: clientID,
		ResponseType: "code",
		RedirectURI: "http://localhost:3000",
		Scopes: "offline_access user.read mail.read",
		Tenant: "common",
		ResponseMode: "query",
		Prompt: "consent",
	}, nil
}
