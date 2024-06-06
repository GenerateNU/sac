package oauth

import "github.com/GenerateNU/sac/backend/config"

type OAuthProvider interface {
	AuthorizeURL(state string) string
	GetConfig() config.OAuthSettings
	Revoke(token string) error
}

type OAuthClient struct {
	OAuthProvider
}

type OAuthProviders struct {
	Outlook OutlookOAuthClient
	Google  GoogleOAuthClient
}
