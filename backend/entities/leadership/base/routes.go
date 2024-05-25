package base

import (
	"github.com/GenerateNU/sac/backend/types"
)

func Leader(leaderParams types.RouteParams) {
	LeaderController := NewLeaderController(NewLeaderService(leaderParams.ServiceParams))

	// api/v1/leader/*
	Leader := leaderParams.Router.Group("/leadership")

	Leader.Get("/", leaderParams.UtilityMiddleware.Paginator, LeaderController.GetLeadership)
	Leader.Get("/:leaderID", LeaderController.GetLeader)
	// Leader.Get("/:leaderID/file", LeaderController.GetPointOfContacFileInfo))
}
