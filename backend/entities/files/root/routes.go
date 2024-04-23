package root

import (
	"github.com/GenerateNU/sac/backend/types"
)

func File(fileParams types.RouteParams) {
	fileController := NewFileController(NewFileService(fileParams.ServiceParams))

	file := fileParams.Router.Group("/files")
	file.Get("/", fileController.GetFiles)
	file.Get("/:fileID", fileController.GetFile)
	file.Post("/", fileController.CreateFile)
	file.Delete("/:fileID", fileController.DeleteFile)
}
