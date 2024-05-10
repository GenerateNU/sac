package base

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func Auth(params types.RouteParams) {
	authController := NewAuthController(NewAuthService(params.ServiceParams))

	// api/v1/auth/*
	auth := params.Router.Group("/auth")

	auth.Post("/logout", utilities.Make(authController.Logout))
	auth.Post("/login", utilities.Make(authController.Login))
	auth.Post("/refresh", utilities.Make(authController.Refresh))

	// TODO: rate limit
	auth.Post("/send-code", utilities.Make(authController.SendCode))
	auth.Post("/verify-email", utilities.Make(authController.VerifyEmail))

	// TODO: rate limit
	auth.Post("/forgot-password", utilities.Make(authController.ForgotPassword))
	auth.Post("/verify-reset", utilities.Make(authController.VerifyPasswordResetToken))
}
