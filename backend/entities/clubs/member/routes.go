package member

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubMember(clubParams types.RouteParams) {
	clubMemberController := NewClubMemberController(NewClubMemberService(clubParams.ServiceParams))

	clubMembers := clubParams.Router.Group("/members")

	// api/v1/clubs/:clubID/members/*
	clubMembers.Get("/", clubMemberController.GetClubMembers)
	clubMembers.Post(
		"/:userID",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		clubMemberController.CreateClubMember,
	)
	clubMembers.Delete(
		"/:userID",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		clubMemberController.DeleteClubMember,
	)
}
