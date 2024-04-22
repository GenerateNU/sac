package routes

import (
	p "github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/controllers"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/gofiber/fiber/v2"
)

func ClubRoutes(clubParams types.RouteParams) {
	clubIDRouter := Club(clubParams)

	// update the router in params
	clubParams.Router = clubIDRouter

	ClubTag(clubParams)
	ClubFollower(clubParams)
	ClubMember(clubParams)
	ClubContact(clubParams)
	ClubEvent(clubParams)
	ClubPointOfContact(clubParams)
}

func Club(clubParams types.RouteParams) fiber.Router {
	clubController := controllers.NewClubController(services.NewClubService(clubParams.ServiceParams))

	// api/v1/clubs/*
	clubs := clubParams.Router.Group("/clubs")

	clubs.Get("/", clubController.GetClubs)
	clubs.Post("/", clubParams.AuthMiddleware.Authorize(p.CreateAll), clubController.CreateClub)

	// api/v1/clubs/:clubID/*
	clubID := clubs.Group("/:clubID")

	clubID.Get("/", clubController.GetClub)
	clubID.Patch("/", clubParams.AuthMiddleware.ClubAuthorizeById, clubController.UpdateClub)
	clubID.Delete("/", clubParams.AuthMiddleware.Authorize(p.DeleteAll), clubController.DeleteClub)

	return clubID
}
