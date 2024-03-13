package routes

import (
	p "github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(router fiber.Router, db *gorm.DB, validate *validator.Validate, authMiddleware *middleware.AuthMiddlewareService) {
	userIDRouter := User(router, services.NewUserService(db, validate), authMiddleware)

	UserTag(userIDRouter, services.NewUserTagService(db, validate))
	UserFollower(userIDRouter, services.NewUserFollowerService(db, validate))
	UserMember(userIDRouter, services.NewUserMemberService(db))
}

func User(router fiber.Router, userService services.UserServiceInterface, authMiddleware *middleware.AuthMiddlewareService) fiber.Router {
	userController := controllers.NewUserController(userService)

	// api/v1/users/*
	users := router.Group("/users")

	users.Post("/", userController.CreateUser)
	users.Get("/", authMiddleware.Authorize(p.ReadAll), userController.GetUsers)

	// api/v1/users/:userID/*
	usersID := users.Group("/:userID")
	usersID.Use(authMiddleware.UserAuthorizeById)

	usersID.Get("/", userController.GetUser)
	usersID.Patch("/", userController.UpdateUser)
	usersID.Delete("/", userController.DeleteUser)

	return usersID
}
