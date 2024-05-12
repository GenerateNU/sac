package series

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
)

func EventSeries(eventParams types.RouteParams) {
	eventSeriesController := NewEventSeriesController(NewEventSeriesService(eventParams.ServiceParams))

	// api/v1/events/series/*
	eventSeries := eventParams.Router.Group("/series")

	eventSeries.Post(
		"/",
		middleware.AttachExtractor(eventParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromBody("host")),
		eventSeriesController.CreateEventSeries,
	)

	eventSeriedID := eventSeries.Group("/:seriesID")

	eventSeriedID.Get("/", eventSeriesController.GetEventSeries)
	eventSeriedID.Delete(
		"/",
		middleware.AttachExtractor(eventParams.AuthMiddleware.ClubAuthorizeById, nil), // MARK: ExtractFromParamsIntoQuery
		eventSeriesController.DeleteEventSeries,
	)
}
