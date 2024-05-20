package base

import "github.com/GenerateNU/sac/backend/types"

func OAuth(params types.RouteParams) {
	oauthController := NewOAuthController(params.ServiceParams)

	// api/v1/calendar/*
	calendar := params.Router.Group("/oauth")

	calendar.Get("/authorize", oauthController.Authorize)
	calendar.Post("/token", oauthController.Token)
}