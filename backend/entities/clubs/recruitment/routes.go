package recruitment

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubRecruitment(clubParams types.RouteParams) {
	clubRecruitmentController := NewClubRecruitmentController(NewClubRecruitmentService(clubParams.ServiceParams))

	clubRecruitment := clubParams.Router.Group("/recruitment")

	// api/v1/clubs/:clubID/recruitment/*
	clubRecruitment.Post(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.CreateClubRecruitment,
	)
	clubRecruitment.Patch(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.UpdateClubRecruitment,
	)

	clubRecruitment.Delete(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.DeleteClubRecruitment,
	)
}
