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
	user := router.Group("/users")

	user.Post("/", userController.CreateUser)
	user.Get("/", authMiddleware.Authorize(p.ReadAll), userController.GetUsers)

	// api/v1/users/:userID/*
	userID := user.Group("/:userID")
	userID.Use(authMiddleware.UserAuthorizeById)

	userID.Get("/", userController.GetUser)
	userID.Patch("/", userController.UpdateUser)
	userID.Delete("/", userController.DeleteUser)

<<<<<<< HEAD
	return userID
=======
	return usersID
>>>>>>> main
}
