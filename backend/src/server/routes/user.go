package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func User(router fiber.Router, userService services.UserServiceInterface, middlewareService middleware.MiddlewareInterface) fiber.Router {
	userController := controllers.NewUserController(userService)

	// api/v1/users/*
	users := router.Group("/users")
	users.Post("/", userController.CreateUser)
	users.Get("/", userController.GetUsers)

	// api/v1/users/:userID/*
	usersID := users.Group("/:userID")
	usersID.Use(middlewareService.UserAuthorizeById)

	usersID.Get("/", userController.GetUser)
	usersID.Patch("/", userController.UpdateUser)
	usersID.Delete("/", userController.DeleteUser)

	return users
}
