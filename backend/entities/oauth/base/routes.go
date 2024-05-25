package base

import "github.com/GenerateNU/sac/backend/types"

func OAuth(params types.RouteParams) {
	oauthController := NewOAuthController(&OAuthService{ServiceParams: params.ServiceParams})

	// api/v1/calendar/*
	calendar := params.Router.Group("/oauth")

	calendar.Get("/authorize", oauthController.Authorize)
	calendar.Post("/token", oauthController.Token)
	calendar.Delete("/revoke", oauthController.Revoke)
}
