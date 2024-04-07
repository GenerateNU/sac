package root

import (
	p "github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/gofiber/fiber/v2"
)

func User(userParams types.RouteParams) fiber.Router {
	userController := NewUserController(NewUserService(userParams.ServiceParams))

	// api/v1/users/*
	users := userParams.Router.Group("/users")

	users.Get("/", userParams.AuthMiddleware.Authorize(p.ReadAll), userController.GetUsers)
	users.Post("/", userController.CreateUser)
	users.Get("/me", userController.GetMe)
	users.Get("/:userID", userController.GetUser)

	// api/v1/users/:userID/*
	usersID := users.Group("/:userID")
	usersID.Use(userParams.AuthMiddleware.UserAuthorizeById)

	usersID.Patch("/", userController.UpdateUser)
	usersID.Patch("/password", userController.UpdatePassword)
	usersID.Delete("/", userController.DeleteUser)

	return usersID
}
