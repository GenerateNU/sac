package base

import (
	p "github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/users/follower"
	"github.com/GenerateNU/sac/backend/entities/users/member"
	"github.com/GenerateNU/sac/backend/entities/users/tag"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(userParams types.RouteParams) {
	usersRouter := UsersRouter(userParams)

	userParams.Router = usersRouter

	tag.UserTag(userParams)
	follower.UserFollower(userParams)
	member.UserMember(userParams)
}

func UsersRouter(userParams types.RouteParams) fiber.Router {
	userController := NewUserController(NewUserService(userParams.ServiceParams))

	// api/v1/users/*
	users := userParams.Router.Group("/users")

	users.Get("/", userParams.AuthMiddleware.Authorize(p.ReadAll), utilities.Make(userController.GetUsers))
	users.Post("/", utilities.Make(userController.CreateUser))
	users.Get("/me", userParams.AuthMiddleware.Authorize(p.UserRead), utilities.Make(userController.GetMe))

	// api/v1/users/:userID/*
	usersID := users.Group("/:userID")

	usersID.Get("/", utilities.Make(userController.GetUser))
	usersID.Patch("/", userParams.AuthMiddleware.UserAuthorizeById, utilities.Make(userController.UpdateUser))
	usersID.Patch("/password", userParams.AuthMiddleware.UserAuthorizeById, utilities.Make(userController.UpdatePassword))
	usersID.Delete("/", userParams.AuthMiddleware.UserAuthorizeById, utilities.Make(userController.DeleteUser))

	return usersID
}
