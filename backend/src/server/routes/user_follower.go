package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func UserFollower(userRouter fiber.Router, userFollowerService services.UserFollowerServiceInterface) {
	userFollowerController := controllers.NewUserFollowerController(userFollowerService)

	// api/v1/users/:userID/follower/*
	userFollower := userRouter.Group("/follower")

	userFollower.Get("/", userFollowerController.GetAllFollowing)
	userFollower.Post("/:clubID", userFollowerController.CreateFollowing)
	userFollower.Delete("/:clubID", userFollowerController.DeleteFollowing)
}
