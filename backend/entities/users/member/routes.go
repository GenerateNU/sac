package member

import (
	"github.com/GenerateNU/sac/backend/types"
)

func UserMember(userParams types.RouteParams) {
	userMemberController := NewUserMemberController(NewUserMemberService(userParams.ServiceParams))

	// api/v1/users/:userID/member/*
	userMember := userParams.Router.Group("/member")

	userMember.Get("/", userMemberController.GetMembership)
	userMember.Post("/:clubID", userParams.AuthMiddleware.UserAuthorizeById, userMemberController.CreateMembership)
	userMember.Delete("/:clubID", userParams.AuthMiddleware.UserAuthorizeById, userMemberController.DeleteMembership)
}
