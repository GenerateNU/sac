package base

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func File(fileParams types.RouteParams) {
	fileController := NewFileController(NewFileService(fileParams.ServiceParams))

	file := fileParams.Router.Group("/files")

	file.Get("/", utilities.Make(fileController.GetFiles))
	file.Get("/:fileID", utilities.Make(fileController.GetFile))
	file.Post("/", utilities.Make(fileController.CreateFile))
	file.Delete("/:fileID", utilities.Make(fileController.DeleteFile))
}
