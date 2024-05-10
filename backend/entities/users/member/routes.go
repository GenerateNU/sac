package member

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func UserMember(userParams types.RouteParams) {
	userMemberController := NewUserMemberController(NewUserMemberService(userParams.ServiceParams))

	// api/v1/users/:userID/member/*
	userMember := userParams.Router.Group("/member")

	userMember.Get("/", utilities.Make(userMemberController.GetMembership))
}
