package base

import (
	p "github.com/GenerateNU/sac/backend/auth"
	authBase "github.com/GenerateNU/sac/backend/entities/auth/base"
	"github.com/GenerateNU/sac/backend/entities/users/follower"
	"github.com/GenerateNU/sac/backend/entities/users/member"
	"github.com/GenerateNU/sac/backend/entities/users/tag"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(userParams types.RouteParams, authService authBase.AuthServiceInterface) {
	usersRouter := UsersRouter(userParams, authService)

	userParams.Router = usersRouter

	tag.UserTag(userParams)
	follower.UserFollower(userParams)
	member.UserMember(userParams)
}

func UsersRouter(userParams types.RouteParams, authService authBase.AuthServiceInterface) fiber.Router {
	userController := NewUserController(NewUserService(userParams.ServiceParams, authService))

	// api/v1/users/*
	users := userParams.Router.Group("/users")

	users.Get("/", userParams.AuthMiddleware.Authorize(p.ReadAll), middleware.Paginate, userController.GetUsers)
	users.Post("/", userController.CreateUser)
	users.Get("/me", userParams.AuthMiddleware.Authorize(p.UserRead), userController.GetMe)

	// api/v1/users/:userID/*
	usersID := users.Group("/:userID")

	usersID.Get("/", userController.GetUser)
	usersID.Patch("/", userParams.AuthMiddleware.UserAuthorizeById, userController.UpdateUser)
	usersID.Patch("/password", userParams.AuthMiddleware.UserAuthorizeById, userController.UpdatePassword)
	usersID.Delete("/", userParams.AuthMiddleware.UserAuthorizeById, userController.DeleteUser)

	return usersID
}
