package routes

import (
	"github.com/GenerateNU/sac/backend/controllers"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/types"
)

func Auth(params types.RouteParams) {
	authController := controllers.NewAuthController(services.NewAuthService(params.ServiceParams))

	// api/v1/auth/*
	auth := params.Router.Group("/auth")

	auth.Post("/logout", authController.Logout)
	auth.Post("/login", authController.Login)
	auth.Post("/refresh", authController.Refresh)

	// TODO: rate limit
	auth.Post("/send-code", authController.SendCode)
	auth.Post("/verify-email", authController.VerifyEmail)

	// TODO: rate limit
	auth.Post("/forgot-password", authController.ForgotPassword)
	auth.Post("/verify-reset", authController.VerifyPasswordResetToken)
}
