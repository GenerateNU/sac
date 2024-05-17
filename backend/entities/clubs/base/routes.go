package base

import (
	p "github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/clubs/contacts"
	"github.com/GenerateNU/sac/backend/entities/clubs/events"
	"github.com/GenerateNU/sac/backend/entities/clubs/followers"
	"github.com/GenerateNU/sac/backend/entities/clubs/members"
	"github.com/GenerateNU/sac/backend/entities/clubs/pocs"
	"github.com/GenerateNU/sac/backend/entities/clubs/tags"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/garrettladley/fiberpaginate"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/gofiber/fiber/v2"
)

func ClubRoutes(clubParams types.RouteParams) {
	clubParams.Router = ClubRouter(clubParams)

	tags.ClubTag(clubParams)
	followers.ClubFollower(clubParams)
	members.ClubMember(clubParams)
	contacts.ClubContact(clubParams)
	events.ClubEvent(clubParams)
	pocs.ClubPointOfContact(clubParams)
}

func ClubRouter(clubParams types.RouteParams) fiber.Router {
	clubController := NewClubController(NewClubService(clubParams.ServiceParams))

	// api/v1/clubs/*
	clubs := clubParams.Router.Group("/clubs")

	clubs.Get("/", fiberpaginate.New(), clubController.GetClubs)
	clubs.Post("/", clubParams.AuthMiddleware.Authorize(p.CreateAll), clubController.CreateClub)

	// api/v1/clubs/:clubID/*
	clubID := clubs.Group("/:clubID")

	clubID.Get("/", clubController.GetClub)
	clubID.Patch(
		"/",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		clubController.UpdateClub,
	)
	clubID.Delete("/", clubParams.AuthMiddleware.Authorize(p.DeleteAll), clubController.DeleteClub)

	return clubID
}
