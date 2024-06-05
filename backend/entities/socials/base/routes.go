package base

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func Social(socialParams types.RouteParams) {
	socialController := NewSocialController(NewSocialService(socialParams.ServiceParams))

	// api/v1/socials/*
	socials := socialParams.Router.Group("/socials")

	socials.Get("/", socialParams.UtilityMiddleware.Paginator, socialController.GetSocials)
	socials.Get("/:socialID", socialController.GetSocial)
	socials.Delete("/:socialID", socialParams.AuthMiddleware.Authorize(auth.DeleteAll), socialController.DeleteSocial)
}
