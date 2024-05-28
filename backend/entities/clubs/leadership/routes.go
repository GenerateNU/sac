package leadership

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubLeader(leaderParams types.RouteParams) {
	clubLeaderController := NewClubLeaderController(NewClubLeaderService(leaderParams.ServiceParams))

	clubLeaders := leaderParams.Router.Group("/leadership")

	// api/v1/clubs/:clubID/leadership/*
	clubLeaders.Get("/", clubLeaderController.GetClubLeadership)
	clubLeaders.Get("/:leaderID", clubLeaderController.GetClubLeader)
	clubLeaders.Post(
		"/",
		authMiddleware.AttachExtractor(leaderParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubLeaderController.CreateClubLeader,
	)
	clubLeaders.Patch(
		"/:leaderID",
		authMiddleware.AttachExtractor(leaderParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubLeaderController.UpdateClubLeader,
	)
	clubLeaders.Patch(
		"/:leaderID/photo",
		authMiddleware.AttachExtractor(leaderParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubLeaderController.UpdateClubLeaderPhoto,
	)
	clubLeaders.Delete(
		"/:leaderID",
		authMiddleware.AttachExtractor(leaderParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubLeaderController.DeleteClubLeader,
	)
}
