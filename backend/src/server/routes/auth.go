package routes

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func Auth(params types.RouteParams) {
	authController := controllers.NewAuthController(services.NewAuthService(params.ServiceParams), params.Settings)

	// api/v1/auth/*
	auth := params.Router.Group("/auth")

	auth.Post("/login", authController.Login)
	auth.Get("/logout", authController.Logout)
	auth.Get("/refresh", authController.Refresh)
	auth.Get("/me", authController.Me)
	auth.Post("/update-password/:userID", params.AuthMiddleware.Limiter(2, 1*time.Minute), params.AuthMiddleware.UserAuthorizeById, authController.UpdatePassword)
	auth.Post("/send-code/:userID", authController.SendCode)
	auth.Post("/verify-email", authController.VerifyEmail)
	auth.Post("/forgot-password", authController.ForgotPassword)
	auth.Post("/verify-reset", authController.VerifyPasswordResetToken)
}
