package routes

import (
	p "github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ClubRoutes(router fiber.Router, db *gorm.DB, validate *validator.Validate, authMiddleware *middleware.AuthMiddlewareService) {
	clubIDRouter := Club(router, services.NewClubService(db, validate), authMiddleware)

	ClubTag(clubIDRouter, services.NewClubTagService(db, validate), authMiddleware)
	ClubFollower(clubIDRouter, services.NewClubFollowerService(db), authMiddleware)
	ClubMember(clubIDRouter, services.NewClubMemberService(db, validate), authMiddleware)
	ClubContact(clubIDRouter, services.NewClubContactService(db, validate), authMiddleware)
	ClubEvent(clubIDRouter, services.NewClubEventService(db))
}

func Club(router fiber.Router, clubService services.ClubServiceInterface, authMiddleware *middleware.AuthMiddlewareService) fiber.Router {
	clubController := controllers.NewClubController(clubService)

	// api/v1/clubs/*
	clubs := router.Group("/clubs")

	clubs.Get("/", clubController.GetClubs)
	clubs.Post("/", authMiddleware.Authorize(p.CreateAll), clubController.CreateClub)

	// api/v1/clubs/:clubID/*
	clubsID := clubs.Group("/:clubID")

	clubsID.Get("/", clubController.GetClub)
	clubsID.Patch("/", authMiddleware.ClubAuthorizeById, clubController.UpdateClub)
	clubsID.Delete("/", authMiddleware.Authorize(p.DeleteAll), clubController.DeleteClub)

	return clubsID
}
