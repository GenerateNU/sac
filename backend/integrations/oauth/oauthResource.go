package oauth

import "github.com/GenerateNU/sac/backend/config"

type OAuthResourceClientInterface interface {
	AuthorizeURL(state string) string
	GetConfig() config.OAuthConfig
}

type ResourceClient struct {
	Outlook 	OutlookOAuthClient
	Google 		GoogleOAuthClient
}

// Calendar Types: 
type OutlookOAuthClient struct {
	OAuthConfig config.OAuthConfig
}
type GoogleOAuthClient struct {
	OAuthConfig config.OAuthConfig
}

/**
* GOOGLE OAUTH CLIENT
**/
func (client GoogleOAuthClient) AuthorizeURL(state string) string {
	return (
		client.OAuthConfig.BaseURL + "/auth?" + 
		"client_id=" + client.OAuthConfig.ClientID.Expose() +
		"&redirect_uri=" + client.OAuthConfig.RedirectURI +
		"&response_type=" + client.OAuthConfig.ResponseType +
		"&scope=" + client.OAuthConfig.Scopes +
		"&access_type=" + client.OAuthConfig.AccessType +
		"&include_granted_scopes=" + client.OAuthConfig.IncludeGrantedScopes +
		"&prompt=" + client.OAuthConfig.Prompt +
		"&state=" + state)
}

func (client GoogleOAuthClient) GetConfig() config.OAuthConfig {
	return client.OAuthConfig
}

/**
* OUTLOOK OAUTH CLIENT
**/

func (client OutlookOAuthClient) AuthorizeURL(state string) string {
	return (
		client.OAuthConfig.BaseURL + "/authorize?" +
		"client_id=" + client.OAuthConfig.ClientID.Expose() +
		"&redirect_uri=" + client.OAuthConfig.RedirectURI +
		"&response_type=" + client.OAuthConfig.ResponseType +
		"&response_mode=" + client.OAuthConfig.ResponseMode +
		"&scope=" + client.OAuthConfig.Scopes +
		"&access_type=" + client.OAuthConfig.AccessType +
		"&state=" + state)
}

func (client OutlookOAuthClient) GetConfig() config.OAuthConfig {
	return client.OAuthConfig
}