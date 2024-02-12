package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func ClubMember(clubsIDRouter fiber.Router, clubMemberService services.ClubMemberServiceInterface) {
	clubMemberController := controllers.NewClubMemberController(clubMemberService)

	clubMember := clubsIDRouter.Group("/member")

	// api/v1/clubs/:clubID/member/*
	clubMember.Get("/", clubMemberController.GetClubMembers)
}
