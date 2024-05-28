package recruitment

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubRecruitment(clubParams types.RouteParams) {
	clubRecruitmentController := NewClubRecruitmentController(NewClubRecruitmentService(clubParams.ServiceParams))

	// api/v1/clubs/:clubID/recruitment/*
	clubRecruitment := clubParams.Router.Group("/recruitment")

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

	// api/v1/clubs/:clubID/recruitment/:recruitmentID/applications/*
	clubRecruitmentIDApplications := clubRecruitment.Group("/:recruitmentID/applications")

	clubRecruitmentIDApplications.Get(
		"/",
		clubParams.UtilityMiddleware.Paginator,
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.GetClubRecruitmentApplications,
	)
	clubRecruitmentIDApplications.Post(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.CreateClubRecruitmentApplication,
	)

	// api/v1/clubs/:clubID/recruitment/:recruitmentID/applications/:applicationID/*
	clubRecruitmentIDApplicationsID := clubRecruitmentIDApplications.Group("/:applicationID")

	clubRecruitmentIDApplicationsID.Get(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.GetClubRecruitmentApplication,
	)
	clubRecruitmentIDApplications.Patch(
		"/:applicationID",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.UpdateClubRecruitmentApplication,
	)
	clubRecruitmentIDApplications.Delete(
		"/:applicationID",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubRecruitmentController.DeleteClubRecruitmentApplication,
	)
}
