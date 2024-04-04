package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func File(fileParams types.RouteParams) {
	fileController := controllers.NewFileController(services.NewFileService(fileParams.ServiceParams))

	file := fileParams.Router.Group("/files")
	file.Post("/", fileController.CreateFile)
	file.Get("/:fileID", fileController.GetFile)
	// file.Delete("/:fileID", fileController.DeleteFile)
}
