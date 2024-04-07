package user

import (
	"github.com/GenerateNU/sac/backend/src/schema/user/root"
	"github.com/GenerateNU/sac/backend/src/types"
)

func UserRoutes(userParams types.RouteParams) {
	usersRouter := root.User(userParams)

	userParams.Router = usersRouter

	UserTag(userParams)
	UserFollower(userParams)
	UserMember(userParams)
}
