package member

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/garrettladley/fiberpaginate"
)

func ClubMember(clubParams types.RouteParams) {
	clubMemberController := NewClubMemberController(NewClubMemberService(clubParams.ServiceParams))

	clubMembers := clubParams.Router.Group("/members")

	// api/v1/clubs/:clubID/members/*
	clubMembers.Get("/", fiberpaginate.New(), clubMemberController.GetClubMembers)
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
