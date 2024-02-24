package routes

import (
	"fmt"
	"time"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Auth(router fiber.Router, authService services.AuthServiceInterface, settings config.AuthSettings, authMiddleware *middleware.AuthMiddlewareService) {
	authController := controllers.NewAuthController(authService, settings)

	// api/v1/auth/*
	auth := router.Group("/auth")

	auth.Post("/login", authController.Login)
	auth.Get("/logout", authController.Logout)
	auth.Get("/refresh", authController.Refresh)
	auth.Get("/me", authController.Me)
	auth.Post("/update-password/:userID", limiter.New(limiter.Config{
		Max:        2,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s-%s", c.IP(), c.Params("userId"))
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests",
			})
		},
	}), authMiddleware.UserAuthorizeById, authController.UpdatePassword)
	// auth.Post("/reset-password/:userID", middleware.Skip(authMiddleware.UserAuthorizeById), authController.ResetPassword)
}
