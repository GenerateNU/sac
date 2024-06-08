package base

import (
	p "github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/clubs/events"
	"github.com/GenerateNU/sac/backend/entities/clubs/followers"
	"github.com/GenerateNU/sac/backend/entities/clubs/leadership"
	"github.com/GenerateNU/sac/backend/entities/clubs/members"
	"github.com/GenerateNU/sac/backend/entities/clubs/recruitment"
	"github.com/GenerateNU/sac/backend/entities/clubs/socials"
	"github.com/GenerateNU/sac/backend/entities/clubs/tags"
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/gofiber/fiber/v2"
)

func ClubRoutes(clubParams types.RouteParams) {
	clubParams.Router = ClubRouter(clubParams)

	tags.ClubTag(clubParams)
	followers.ClubFollower(clubParams)
	members.ClubMember(clubParams)
	leadership.ClubLeader(clubParams)
	events.ClubEvent(clubParams)
	socials.ClubSocial(clubParams)
	recruitment.ClubRecruitment(clubParams)
}

func ClubRouter(clubParams types.RouteParams) fiber.Router {
	clubController := NewClubController(NewClubService(clubParams.ServiceParams))

	// api/v1/clubs/*
	clubs := clubParams.Router.Group("/clubs")

	clubs.Get("/", clubParams.UtilityMiddleware.Paginator, clubController.GetClubs)
	clubs.Post("/", clubParams.AuthMiddleware.Authorize(p.CreateAll), clubController.CreateClub)

	// api/v1/clubs/:clubID/*
	clubID := clubs.Group("/:clubID")

	clubID.Get("/", clubController.GetClub)
	clubID.Patch(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubController.UpdateClub,
	)
	clubID.Delete("/", clubParams.AuthMiddleware.Authorize(p.DeleteAll), clubController.DeleteClub)

	return clubID
}
