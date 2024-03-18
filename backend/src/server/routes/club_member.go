package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func ClubMember(clubParams types.RouteParams) {
	clubMemberController := controllers.NewClubMemberController(services.NewClubMemberService(clubParams.ServiceParams))

	clubMember := clubParams.Router.Group("/members")

	// api/v1/clubs/:clubID/members/*
	clubMember.Get("/", clubParams.AuthMiddleware.ClubAuthorizeById, clubMemberController.GetClubMembers)
}
