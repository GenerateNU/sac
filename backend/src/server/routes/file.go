package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func File(fileParams types.RouteParams) {
	fileController := controllers.NewFileController(services.NewFileService(fileParams.ServiceParams))

	file := fileParams.Router.Group("/files")
	file.Get("/", fileController.GetFiles)
	file.Get("/:fileID", fileController.GetFile)
	file.Post("/", fileController.CreateFile)
	file.Delete("/:fileID", fileController.DeleteFile)
}
