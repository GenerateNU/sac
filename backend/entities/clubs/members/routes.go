package members

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubMember(clubParams types.RouteParams) {
	clubMemberController := NewClubMemberController(NewClubMemberService(clubParams.ServiceParams))

	clubMembers := clubParams.Router.Group("/members")

	// api/v1/clubs/:clubID/members/*
	clubMembers.Get("/", clubParams.UtilityMiddleware.Paginator, clubMemberController.GetClubMembers)
	clubMembers.Post(
		"/:userID",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubMemberController.CreateClubMember,
	)
	clubMembers.Delete(
		"/:userID",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubMemberController.DeleteClubMember,
	)
}
