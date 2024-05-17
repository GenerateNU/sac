package tags

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
)

func EventTags(eventParams types.RouteParams) {
	eventTagController := NewEventTagController(NewEventTagService(eventParams.ServiceParams))

	eventTags := eventParams.Router.Group("/:eventID/tags")

	eventTags.Get("/", eventTagController.GetEventTags)
	eventTags.Post(
		"/",
		middleware.AttachExtractor(eventParams.AuthMiddleware.EventAuthorizeById, middleware.ExtractFromParams("eventID")),
		eventTagController.CreateEventTags,
	)
	eventTags.Delete(
		"/:tagID",
		middleware.AttachExtractor(eventParams.AuthMiddleware.EventAuthorizeById, middleware.ExtractFromParams("eventID")),
		eventTagController.DeleteEventTag,
	)
}
