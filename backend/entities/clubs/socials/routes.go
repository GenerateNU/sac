package socials

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubSocial(clubParams types.RouteParams) {
	clubSocialController := NewClubSocialController(NewClubSocialService(clubParams.ServiceParams))

	clubSocials := clubParams.Router.Group("/socials")

	// api/v1/clubs/:clubID/socials/*
	clubSocials.Get("/", clubSocialController.GetClubSocials)
	clubSocials.Put(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubSocialController.PutSocial,
	)
}
