package routes

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Auth(router fiber.Router, authService services.AuthServiceInterface, settings config.AuthSettings, authMiddleware *middleware.AuthMiddlewareService) {
	authController := controllers.NewAuthController(authService, settings)

	// api/v1/auth/*
	auth := router.Group("/auth")

	auth.Post("/login", authController.Login)
	auth.Get("/logout", authController.Logout)
	auth.Get("/refresh", authController.Refresh)
	auth.Get("/me", authController.Me)
	auth.Post("/update-password/:userID", authMiddleware.Limiter(2, 1*time.Minute), authMiddleware.UserAuthorizeById, authController.UpdatePassword)
	auth.Post("/send-code/:userID", authController.SendCode)
	auth.Post("/verify-email", authController.VerifyEmail)
	auth.Post("/forgot-password", authController.ForgotPassword)
	auth.Post("/verify-reset", authController.VerifyPasswordResetToken)
}
