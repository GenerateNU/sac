package calendar

import (
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/gofiber/fiber/v2"
)

type GoogleClient struct {
	Settings config.GoogleClientSettings
}

func (client GoogleClient) Authorize() *string { return nil }

func (client GoogleClient) Token(c *fiber.Ctx) *errors.Error { return nil }

// func NewGoogleClient(googleSettings config.GoogleOAuthSettings) *GoogleClient {
// 	googleOAuthClient := &oauth2.Config{
// 		ClientID:     googleSettings.ClientID.Expose(),
// 		ClientSecret: googleSettings.ClientSecret.Expose(),
// 		Scopes: googleSettings.Scopes,
// 		Endpoint: google.Endpoint,
// 		RedirectURL: googleSettings.RedirectURI,
// 	}

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



