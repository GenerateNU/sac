package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func UserTag(usersRouter fiber.Router, userTagService services.UserTagServiceInterface) {
	userTagController := controllers.NewUserTagController(userTagService)

	// api/v1/user/:userID/tags/*
	userTags := usersRouter.Group("/tags")

	userTags.Post("/", userTagController.CreateUserTags)
	userTags.Get("/", userTagController.GetUserTags)
	
	tagID := userTags.Group("/:tagID")
	tagID.Delete("/", userTagController.DeleteUserTag)
}
