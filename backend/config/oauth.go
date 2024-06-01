package config

import (
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

type OauthProviderSettings struct {
	GoogleOAuthSettings  *OAuthSettings
	OutlookOAuthSettings *OAuthSettings
}
