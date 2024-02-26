package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func UserMember(usersRouter fiber.Router, userMembershipService services.UserMemberServiceInterface) {
	userMemberController := controllers.NewUserMemberController(userMembershipService)

	// api/v1/users/:userID/member/*
	userMember := usersRouter.Group("/member")

	userMember.Get("/", userMemberController.GetMembership)
	userMember.Post("/:clubID", userMemberController.CreateMembership)
	userMember.Delete("/:clubID", userMemberController.DeleteMembership)
}
