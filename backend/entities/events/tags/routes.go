package tags

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func EventTags(eventParams types.RouteParams) {
	eventTagController := NewEventTagController(NewEventTagService(eventParams.ServiceParams))

	eventTags := eventParams.Router.Group("/:eventID/tags")

	eventTags.Get("/", eventTagController.GetEventTags)
	eventTags.Post(
		"/",
		authMiddleware.AttachExtractor(eventParams.AuthMiddleware.EventAuthorizeById, authMiddleware.ExtractFromParams("eventID")),
		eventTagController.CreateEventTags,
	)
	eventTags.Delete(
		"/:tagID",
		authMiddleware.AttachExtractor(eventParams.AuthMiddleware.EventAuthorizeById, authMiddleware.ExtractFromParams("eventID")),
		eventTagController.DeleteEventTag,
	)
}
