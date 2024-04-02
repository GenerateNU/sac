package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func UserMember(userParams types.RouteParams) {
	userMemberController := controllers.NewUserMemberController(services.NewUserMemberService(userParams.ServiceParams))

	// api/v1/users/:userID/member/*
	userMember := userParams.Router.Group("/member")

	userMember.Get("/", userMemberController.GetMembership)
	userMember.Post("/:clubID", userMemberController.CreateMembership)
	userMember.Delete("/:clubID", userMemberController.DeleteMembership)
}
