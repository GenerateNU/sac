package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func UserMember(usersRouter fiber.Router, userMembershipService services.UserMemberServiceInterface) {
	userMemberController := controllers.NewUserMemberController(userMembershipService)

	userMember := usersRouter.Group("/member")

	userMember.Get("/", userMemberController.GetMembership)

	clubID := userMember.Group("/:clubID")

	// api/v1/users/:userID/member/:clubID*
	clubID.Post("/", userMemberController.CreateMembership)
	clubID.Delete("/", userMemberController.DeleteMembership)
}
