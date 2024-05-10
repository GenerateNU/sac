package tags

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func EventTags(eventParams types.RouteParams) {
	eventTagController := NewEventTagController(NewEventTagService(eventParams.ServiceParams))

	eventTags := eventParams.Router.Group("/:eventID/tags")

	eventTags.Get("/", utilities.Make(eventTagController.GetEventTags))
	eventTags.Post("/", eventParams.AuthMiddleware.EventAuthorizeById, utilities.Make(eventTagController.CreateEventTags))
	eventTags.Delete("/:tagID", eventParams.AuthMiddleware.EventAuthorizeById, utilities.Make(eventTagController.DeleteEventTag))
}
