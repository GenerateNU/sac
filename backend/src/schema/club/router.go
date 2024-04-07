package club

import (
	"github.com/GenerateNU/sac/backend/src/schema/club/root"
	"github.com/GenerateNU/sac/backend/src/types"
)

func ClubRoutes(clubParams types.RouteParams) {
	clubIDRouter := root.Club(clubParams)

	// update the router in params
	clubParams.Router = clubIDRouter

	ClubTag(clubParams)
	follower.ClubFollower(clubParams)
	ClubMember(clubParams)
	ClubContact(clubParams)
	ClubEvent(clubParams)
	ClubPointOfContact(clubParams)
}
