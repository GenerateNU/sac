package member

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/garrettladley/fiberpaginate"
)

func ClubMember(clubParams types.RouteParams) {
	clubMemberController := NewClubMemberController(NewClubMemberService(clubParams.ServiceParams))

	clubMembers := clubParams.Router.Group("/members")

	// api/v1/clubs/:clubID/members/*
	clubMembers.Get("/", fiberpaginate.New(), clubMemberController.GetClubMembers)
	clubMembers.Post("/:userID", clubParams.AuthMiddleware.ClubAuthorizeById, clubMemberController.CreateClubMember)
	clubMembers.Delete("/:userID", clubParams.AuthMiddleware.ClubAuthorizeById, clubMemberController.DeleteClubMember)
}
