package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func ClubMember(clubIDRouter fiber.Router, clubMemberService services.ClubMemberServiceInterface, authMiddleware *middleware.AuthMiddlewareService) {
	clubMemberController := controllers.NewClubMemberController(clubMemberService)

	clubMember := clubIDRouter.Group("/members")

	// api/v1/clubs/:clubID/members/*
	clubMember.Get("/", authMiddleware.ClubAuthorizeById, clubMemberController.GetClubMembers)
}
