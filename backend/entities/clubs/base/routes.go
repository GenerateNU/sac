package base

import (
	p "github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/clubs/contact"
	"github.com/GenerateNU/sac/backend/entities/clubs/event"
	"github.com/GenerateNU/sac/backend/entities/clubs/follower"
	"github.com/GenerateNU/sac/backend/entities/clubs/member"
	"github.com/GenerateNU/sac/backend/entities/clubs/poc"
	"github.com/GenerateNU/sac/backend/entities/clubs/tag"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/gofiber/fiber/v2"
)

func ClubRoutes(clubParams types.RouteParams) {
	clubParams.Router = ClubRouter(clubParams)

	tag.ClubTag(clubParams)
	follower.ClubFollower(clubParams)
	member.ClubMember(clubParams)
	contact.ClubContact(clubParams)
	event.ClubEvent(clubParams)
	poc.ClubPointOfContact(clubParams)
}

func ClubRouter(clubParams types.RouteParams) fiber.Router {
	clubController := NewClubController(NewClubService(clubParams.ServiceParams))

	// api/v1/clubs/*
	clubs := clubParams.Router.Group("/clubs")

	clubs.Get("/", utilities.Make(clubController.GetClubs))
	clubs.Post("/", clubParams.AuthMiddleware.Authorize(p.CreateAll), utilities.Make(clubController.CreateClub))

	// api/v1/clubs/:clubID/*
	clubID := clubs.Group("/:clubID")

	clubID.Get("/", utilities.Make(clubController.GetClub))
	clubID.Patch("/", clubParams.AuthMiddleware.ClubAuthorizeById, utilities.Make(clubController.UpdateClub))
	clubID.Delete("/", clubParams.AuthMiddleware.Authorize(p.DeleteAll), utilities.Make(clubController.DeleteClub))

	return clubID
}
