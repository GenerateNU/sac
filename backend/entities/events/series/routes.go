package series

import (
	"github.com/GenerateNU/sac/backend/extractors"
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func EventSeries(eventParams types.RouteParams) {
	eventSeriesController := NewEventSeriesController(NewEventSeriesService(eventParams.ServiceParams))

	// api/v1/events/series/*
	eventSeries := eventParams.Router.Group("/series")

	eventSeries.Post(
		"/",
		authMiddleware.AttachExtractor(eventParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromBody("host")),
		eventSeriesController.CreateEventSeries,
	)

	eventSeriedID := eventSeries.Group("/:seriesID")

	eventSeriedID.Get("/", eventSeriesController.GetEventSeries)
	eventSeriedID.Delete(
		"/",
		authMiddleware.AttachExtractor(
			eventParams.AuthMiddleware.ClubAuthorizeById,
			authMiddleware.FromExtractFromParamIntoQuery("seriesID", eventParams.ServiceParams.DB, extractors.GetSeriesHost),
		),
		eventSeriesController.DeleteEventSeries,
	)
}
