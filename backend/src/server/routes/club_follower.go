package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func ClubFollower(clubsIDRouter fiber.Router, clubFollowerService services.ClubFollowerServiceInterface) {
	clubFollowerController := controllers.NewClubFollowerController(clubFollowerService)

	clubFollower := clubsIDRouter.Group("/followers")

	// api/clubs/:clubID/followers/*
	clubFollower.Get("/", clubFollowerController.GetClubFollowers)
}
