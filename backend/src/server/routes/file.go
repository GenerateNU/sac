package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func File(router fiber.Router, fileService services.FileServiceInterface) fiber.Router {
	fileController := controllers.NewFileController(fileService)

	files := router.Group("/files")
	files.Post("/", fileController.CreateFile)
	files.Get("/:fileID", fileController.GetFile)
	files.Delete("/:fileID", fileController.DeleteFile)

	return files
}

