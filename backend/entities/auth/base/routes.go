package base

import (
	"github.com/GenerateNU/sac/backend/types"
)

func Auth(params types.RouteParams) AuthServiceInterface {
	authService := NewAuthService(params.ServiceParams)
	authController := NewAuthController(authService)

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

	return authService
}
