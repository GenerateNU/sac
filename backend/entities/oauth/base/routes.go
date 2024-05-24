package base

import "github.com/GenerateNU/sac/backend/types"

func OAuth(params types.RouteParams) {
	oauthController := NewOAuthController(&OAuthService{ ServiceParams: params.ServiceParams })

	// api/v1/calendar/*
	calendar := params.Router.Group("/oauth")

	calendar.Get("/authorize", params.AuthMiddleware.UserAuthorizeAndExtractID, oauthController.Authorize)
	calendar.Post("/token", params.AuthMiddleware.UserAuthorizeAndExtractID, oauthController.Token)
	calendar.Delete("/revoke", params.AuthMiddleware.UserAuthorizeAndExtractID, oauthController.Revoke)
}
