package tags

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func EventTags(eventParams types.RouteParams) {
	eventTagController := NewEventTagController(NewEventTagService(eventParams.ServiceParams))

	eventTags := eventParams.Router.Group("/:eventID/tags")

	eventTags.Get("/", utilities.Make(eventTagController.GetEventTags))
	eventTags.Post(
		"/",
		middleware.AttachExtractor(eventParams.AuthMiddleware.EventAuthorizeById, middleware.ExtractFromParams("eventID")),
		utilities.Make(eventTagController.CreateEventTags),
	)
	eventTags.Delete(
		"/:tagID",
		middleware.AttachExtractor(eventParams.AuthMiddleware.EventAuthorizeById, middleware.ExtractFromParams("eventID")),
		utilities.Make(eventTagController.DeleteEventTag),
	)
}
