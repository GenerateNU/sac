package base

import (
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/GenerateNU/sac/backend/types"
)

func OAuth(params types.RouteParams) {
	oauthController := NewOAuthController(
		oauth.NewHandler(
			params.ServiceParams.DB,
			params.ServiceParams.Stores,
			params.ServiceParams.Integrations.OAuth,
			params.ServiceParams.Validate,
		),
	)

	// api/v1/calendar/*
	calendar := params.Router.Group("/oauth")

	calendar.Get("/authorize", oauthController.Authorize)
	calendar.Post("/token", oauthController.AuthorizationToken)
	calendar.Delete("/revoke", oauthController.Revoke)
}
