package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func UserFollower(usersRouter fiber.Router, userFollowerService services.UserFollowerServiceInterface) {
	userFollowerController := controllers.NewUserFollowerController(userFollowerService)

	userFollower := usersRouter.Group("/:userID/follower")

	userFollower.Get("/", userFollowerController.GetAllFollowing)

	clubID := userFollower.Group("/:clubID")

	// api/v1/users/:userID/follower/*
	clubID.Post("/", userFollowerController.CreateFollowing)
	clubID.Delete("/", userFollowerController.DeleteFollowing)
}
