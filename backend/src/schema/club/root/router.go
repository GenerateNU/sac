package root

import (
	p "github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/gofiber/fiber/v2"
)

func Club(clubParams types.RouteParams) fiber.Router {
	clubController := NewClubController(NewClubService(clubParams.ServiceParams))

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
