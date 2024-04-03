package routes

import (
	p "github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(userParams types.RouteParams) {
	usersRouter := User(userParams)

	userParams.Router = usersRouter

	UserTag(userParams)
	UserFollower(userParams)
	UserMember(userParams)
}

func User(userParams types.RouteParams) fiber.Router {
	userController := controllers.NewUserController(services.NewUserService(userParams.ServiceParams))

	// api/v1/users/*
	users := userParams.Router.Group("/users")

	users.Get("/", userParams.AuthMiddleware.Authorize(p.ReadAll), userController.GetUsers)
	users.Post("/", userController.CreateUser)
	users.Get("/me", userController.GetMe)

	// api/v1/users/:userID/*
	usersID := users.Group("/:userID")
	usersID.Use(userParams.AuthMiddleware.UserAuthorizeById)

	usersID.Get("/", userController.GetUser)
	usersID.Patch("/", userController.UpdateUser)
	usersID.Patch("/password", userController.UpdatePassword)
	usersID.Delete("/", userController.DeleteUser)

	return userID
}
