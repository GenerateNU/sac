package calendar

import (
	"github.com/GenerateNU/sac/backend/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleClient struct {
	OAuthClient *oauth2.Config
}

func NewGoogleOAuthClient(googleSettings *config.GoogleClientSettings) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     googleSettings.ClientID.Expose(),
		ClientSecret: googleSettings.ClientSecret.Expose(),
		Scopes: googleSettings.Scopes,
		Endpoint: google.Endpoint,
		RedirectURL: googleSettings.RedirectURI,
	}
}

func (client GoogleClient) Authorize() string { 
	return client.OAuthClient.AuthCodeURL("not-implemented", oauth2.AccessTypeOffline)
}

func (client GoogleClient) Token(grantType string, code *string) error { return nil }


// 	return &GoogleClient{
// 		OAuthClient: googleOAuthClient,
// 	}
// }

// func (c *GoogleClient) OAuthURL() string {
// 	return c.OAuthClient.AuthCodeURL("not-implemented", oauth2.AccessTypeOffline)
// }

// func (c *GoogleClient) OAuthCallback(ctx *fiber.Ctx) (*oauth2.Token, *errors.Error) {
// 	code := ctx.Query("code")
// 	token, err := c.OAuthClient.Exchange(ctx.Context(), code)
// 	if err != nil {
// 		return nil, &errors.FailedToExchangeAuthCode
// 	}

// 	return token, nil
// }



