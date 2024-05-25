package oauth

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/utilities"
)

type GoogleOAuthClient struct {
	OAuthConfig config.OAuthSettings
}

func (client *GoogleOAuthClient) AuthorizeURL(state string) string {
	values := url.Values{}
	values.Set("client_id", client.OAuthConfig.ClientID.Expose())
	values.Set("redirect_uri", client.OAuthConfig.RedirectURI)
	values.Set("response_type", client.OAuthConfig.ResponseType)
	values.Set("scope", client.OAuthConfig.Scopes)
	values.Set("access_type", client.OAuthConfig.AccessType)
	values.Set("include_granted_scopes", client.OAuthConfig.IncludeGrantedScopes)
	values.Set("prompt", client.OAuthConfig.Prompt)
	values.Set("state", state)

	return fmt.Sprintf("%s/auth?%s", client.OAuthConfig.BaseURL, values.Encode())
}

func (client *GoogleOAuthClient) GetConfig() config.OAuthSettings {
	return client.OAuthConfig
}

func (client *GoogleOAuthClient) Revoke(token string) error {
	resp, err := utilities.Request(http.MethodPost, fmt.Sprintf("%s/revoke?token=%s", client.OAuthConfig.TokenURL, token), nil, utilities.FormURLEncoded())
	if err != nil {
		return err
	}

	if !utilities.IsOk(resp.StatusCode) {
		return errors.New("failed to revoke token")
	}

	return nil
}
