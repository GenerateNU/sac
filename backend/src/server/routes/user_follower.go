package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func UserFollower(usersIDRouter fiber.Router, userFollowerService services.UserFollowerServiceInterface) {
	userFollowerController := controllers.NewUserFollowerController(userFollowerService)

	userFollower := usersIDRouter.Group("/follower")

	// api/users/:userID/follower/*
	userFollower.Put("/:clubID", userFollowerController.CreateFollowing)
	userFollower.Delete("/:clubID", userFollowerController.DeleteFollowing)
	userFollower.Get("/", userFollowerController.GetAllFollowing)
}
