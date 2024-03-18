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

	// update the router in params
	userParams.Router = usersRouter

	UserTag(userParams)
	UserFollower(userParams)
	UserMember(userParams)
}

func User(userParams types.RouteParams) fiber.Router {
	userController := controllers.NewUserController(services.NewUserService(userParams.ServiceParams))

	// api/v1/users/*
	users := userParams.Router.Group("/users")

	users.Post("/", userController.CreateUser)
	users.Get("/", userParams.AuthMiddleware.Authorize(p.ReadAll), userController.GetUsers)

	// api/v1/users/:userID/*
	usersID := users.Group("/:userID")
	usersID.Use(userParams.AuthMiddleware.UserAuthorizeById)

	usersID.Get("/", userController.GetUser)
	usersID.Patch("/", userController.UpdateUser)
	usersID.Delete("/", userController.DeleteUser)

	return usersID
}
