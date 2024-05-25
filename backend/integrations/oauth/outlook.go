package oauth

import (
	"fmt"
	"net/url"

	"github.com/GenerateNU/sac/backend/config"
)

type OutlookOAuthClient struct {
	OAuthConfig config.OAuthSettings
}

func (client *OutlookOAuthClient) AuthorizeURL(state string) string {
	values := url.Values{}
	values.Set("client_id", client.OAuthConfig.ClientID.Expose())
	values.Set("redirect_uri", client.OAuthConfig.RedirectURI)
	values.Set("response_type", client.OAuthConfig.ResponseType)
	values.Set("response_mode", client.OAuthConfig.ResponseMode)
	values.Set("scope", client.OAuthConfig.Scopes)
	values.Set("access_type", client.OAuthConfig.AccessType)
	values.Set("state", state)

	return fmt.Sprintf("%s/authorize?%s", client.OAuthConfig.BaseURL, values.Encode())
}

func (client *OutlookOAuthClient) GetConfig() config.OAuthSettings {
	return client.OAuthConfig
}

func (client *OutlookOAuthClient) Revoke(token string) error {
	// Outlook does not expose a revokation endpoint :/
	return nil
}
