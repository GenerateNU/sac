package oauth

import "github.com/GenerateNU/sac/backend/config"

type OutlookOAuthClient struct {
	OAuthConfig config.OAuthSettings
}

func (client *OutlookOAuthClient) AuthorizeURL(state string) string {
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

func (client *OutlookOAuthClient) GetConfig() config.OAuthSettings {
	return client.OAuthConfig
}

func (client *OutlookOAuthClient) Revoke(token string) error {
	/** Outlook does not expose a revokation endpoint :/ **/
	return nil
}
