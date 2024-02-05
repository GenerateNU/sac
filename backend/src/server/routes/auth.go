package routes

import (
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Auth(router fiber.Router, authService services.AuthServiceInterface, authSettings config.AuthSettings) {
	authController := controllers.NewAuthController(authService, authSettings)

	// api/v1/auth/*
	auth := router.Group("/auth")

	auth.Post("/login", authController.Login)
	auth.Get("/logout", authController.Logout)
	auth.Get("/refresh", authController.Refresh)
	auth.Get("/me", authController.Me)
}
