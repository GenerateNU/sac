package oauth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/GenerateNU/sac/backend/config"
)

type GoogleOAuthClient struct {
	OAuthConfig config.OAuthSettings
}

func (client *GoogleOAuthClient) AuthorizeURL(state string) string {
	return (client.OAuthConfig.BaseURL + "/auth?" +
		"client_id=" + client.OAuthConfig.ClientID.Expose() +
		"&redirect_uri=" + client.OAuthConfig.RedirectURI +
		"&response_type=" + client.OAuthConfig.ResponseType +
		"&scope=" + client.OAuthConfig.Scopes +
		"&access_type=" + client.OAuthConfig.AccessType +
		"&include_granted_scopes=" + client.OAuthConfig.IncludeGrantedScopes +
		"&prompt=" + client.OAuthConfig.Prompt +
		"&state=" + state)
}

func (client *GoogleOAuthClient) GetConfig() config.OAuthSettings {
	return client.OAuthConfig
}

func (client *GoogleOAuthClient) Revoke(token string) error {
	revokeUrl := fmt.Sprintf("%s/revoke?token=%s", client.OAuthConfig.TokenURL, token)

	req, err := http.NewRequest("POST", revokeUrl, strings.NewReader(""))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
        return err
    }
	if resp.StatusCode > 299 {
		return errors.New("failed to revoke token")
	}

	return nil
}
