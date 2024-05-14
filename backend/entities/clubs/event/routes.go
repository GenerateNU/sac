package event

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubEvent(clubParams types.RouteParams) {
	clubEventController := NewClubEventController(NewClubEventService(clubParams.ServiceParams))

	// api/v1/clubs/:clubID/events/*
	events := clubParams.Router.Group("/events")

	events.Get("/", middleware.Paginate, clubEventController.GetClubEvents)
}
