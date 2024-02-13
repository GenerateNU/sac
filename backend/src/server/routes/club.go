package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/gofiber/fiber/v2"
)

func Club(router fiber.Router, clubService services.ClubServiceInterface, middlewareService middleware.MiddlewareInterface) fiber.Router {
	clubController := controllers.NewClubController(clubService)

	clubs := router.Group("/clubs")

	clubs.Get("/", middlewareService.Authorize(types.ClubReadAll), clubController.GetAllClubs)
	clubs.Post("/", clubController.CreateClub)

	// api/v1/clubs/:clubID/*
	clubsID := clubs.Group("/:clubID")
	clubsID.Use(middleware.SuperSkipper(middlewareService.UserAuthorizeById))

	clubsID.Get("/", clubController.GetClub)
	clubsID.Patch("/", middlewareService.Authorize(types.ClubWrite), clubController.UpdateClub)
	clubsID.Delete("/", middlewareService.Authorize(types.ClubDelete), clubController.DeleteClub)

	// TODO: refactor into club_event vertical (controller <-> service <-> model)
	// api/v1/clubs/:clubID/events/*
	events := clubsID.Group("/events")
	events.Get("/", clubController.GetClubEvents)

	return clubsID
}
