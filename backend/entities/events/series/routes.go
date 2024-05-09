package series

import "github.com/GenerateNU/sac/backend/types"

func EventSeries(eventParams types.RouteParams) {
	eventSeriesController := NewEventSeriesController(NewEventSeriesService(eventParams.ServiceParams))

	// api/v1/events/series/*
	eventSeries := eventParams.Router.Group("/series")

	eventSeries.Post("/", eventSeriesController.CreateEventSeries)

	eventSeriedID := eventSeries.Group("/:seriesID")

	eventSeriedID.Get("/", eventSeriesController.GetEventSeries)
	eventSeriedID.Patch("/", eventSeriesController.UpdateEventSeries)
	eventSeriedID.Delete("/", eventSeriesController.DeleteEventSeries)
}
