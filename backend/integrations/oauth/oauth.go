package oauth

import "github.com/GenerateNU/sac/backend/config"

type OAuthResourceClientInterface interface {
	AuthorizeURL(state string) string
	GetConfig() config.OAuthSettings
	Revoke(token string) error
}

type OAuthClient struct {
	ResourceClient 	OAuthResourceClientInterface
}

type OAuthResources struct {
	Outlook 	OutlookOAuthClient
	Google 		GoogleOAuthClient
}
