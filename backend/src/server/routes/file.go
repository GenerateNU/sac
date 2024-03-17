package routes

import "github.com/gofiber/fiber/v2"

func File(router fiber.Router, fileService services.FileServiceInterface) {
	fileController := controllers.NewFileController(fileService)

	file := router.Group("/files")
	file.Post("/", fileController.CreateFile)
	files.Get("/:fileID", fileController.GetFile)
	files.Delete("/:fileID", fileController.DeleteFile)
}

