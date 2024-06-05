package recruitment

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubRecruitment(clubParams types.RouteParams) {
	clubRecruitmentController := NewClubRecruitmentController(NewClubRecruitmentService(clubParams.ServiceParams))

	// api/v1/clubs/:clubID/recruitment/*
	clubRecruitment := clubParams.Router.Group("/recruitment")

	clubRecruitment.Get(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.GetClubRecruitment,
	)
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

	// api/v1/clubs/:clubID/recruitment/applications/*
	clubRecruitmentApplications := clubRecruitment.Group("/applications")

	clubRecruitmentApplications.Get(
		"/",
		clubParams.UtilityMiddleware.Paginator,
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.GetClubRecruitmentApplications,
	)
	clubRecruitmentApplications.Post(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.CreateClubRecruitmentApplication,
	)

	// api/v1/clubs/:clubID/recruitment/applications/:applicationID/*
	clubRecruitmentApplicationsID := clubRecruitmentApplications.Group("/:applicationID")

	clubRecruitmentApplicationsID.Get(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.GetClubRecruitmentApplication,
	)
	clubRecruitmentApplicationsID.Patch(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.UpdateClubRecruitmentApplication,
	)
	clubRecruitmentApplicationsID.Delete(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.DeleteClubRecruitmentApplication,
	)
}
