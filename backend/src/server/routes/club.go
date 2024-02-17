package routes

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Club(router fiber.Router, clubService services.ClubServiceInterface, middlewareService middleware.MiddlewareInterface) fiber.Router {
	clubController := controllers.NewClubController(clubService)

	clubs := router.Group("/clubs")

	clubs.Get("/", middlewareService.Authorize(auth.ClubReadAll), clubController.GetAllClubs)
	clubs.Post("/", clubController.CreateClub)

	// api/v1/clubs/:clubID/*
	clubsID := clubs.Group("/:clubID")
	clubsID.Use(middleware.SuperSkipper(middlewareService.UserAuthorizeById))

	clubsID.Get("/", clubController.GetClub)
	clubsID.Patch("/", middlewareService.Authorize(auth.ClubWrite), clubController.UpdateClub)
	clubsID.Delete("/", middlewareService.Authorize(auth.ClubDelete), clubController.DeleteClub)

	// api/v1/clubs/:clubID/poc/*
	pointOfContact := router.Group("/clubs/:clubID/poc")
	pointOfContact.Get("/", clubController.GetAllPointOfContact)
	pointOfContact.Get("/:pocID", clubController.GetPointOfContact)
	pointOfContact.Put("/", clubController.UpsertPointOfContact)
	pointOfContact.Delete("/:pocID", clubController.DeletePointOfContact)

	return clubsID
}
