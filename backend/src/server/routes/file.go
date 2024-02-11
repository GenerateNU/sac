package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func fileRoutes(router fiber.Router, fileService services.FileServiceInterface) {
	fileController := controllers.NewFileController(fileService)

	file := router.Group("/file")
	file.Post("/", fileController.CreateFile)
}

