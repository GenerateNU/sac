package base

import "github.com/GenerateNU/sac/backend/types"

func Calendar(params types.RouteParams) {
	calendarController := NewCalendarController(params.ServiceParams)

	// api/v1/calendar/*
	calendar := params.Router.Group("/calendar")

	calendar.Post("/authorize", calendarController.Authorize)
	calendar.Post("/token", calendarController.Token)
}