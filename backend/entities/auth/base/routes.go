package base

import (
	"time"

	"github.com/GenerateNU/sac/backend/types"
)

func Auth(params types.RouteParams) {
	authController := NewAuthController(NewAuthService(params.ServiceParams))

	// api/v1/auth/*
	auth := params.Router.Group("/auth")

	auth.Post("/register", authController.Register)
	auth.Post("/logout", authController.Logout)
	auth.Post("/login", authController.Login)
	auth.Post("/refresh", authController.Refresh)

	auth.Post("/send-code", params.UtilityMiddleware.Limiter(5, 5*time.Minute), authController.SendCode)
	auth.Post("/verify-email", authController.VerifyEmail)

	auth.Post("/forgot-password", params.UtilityMiddleware.Limiter(5, 5*time.Minute), authController.ForgotPassword)
	auth.Post("/verify-reset", authController.VerifyPasswordResetToken)
}
