package base

import (
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/types"
)

func Auth(params types.RouteParams) {
	authController := NewAuthController(NewAuthService(params.ServiceParams))

	// api/v1/auth/*
	auth := params.Router.Group("/auth")

	auth.Post("/register", authController.Register)
	auth.Post("/logout", authController.Logout)
	auth.Post("/login", authController.Login)

	auth.Post("/send-code", params.UtilityMiddleware.Limiter(constants.RATE_LIMIT_MAX_REQUESTS, constants.RATE_LIMIT_DURATION), authController.SendCode)
	auth.Post("/verify-email", authController.VerifyEmail)

	auth.Post("/forgot-password", params.UtilityMiddleware.Limiter(constants.RATE_LIMIT_MAX_REQUESTS, constants.RATE_LIMIT_DURATION), authController.ForgotPassword)
	auth.Post("/verify-reset", authController.VerifyPasswordResetToken)
}
