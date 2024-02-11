package controllers

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"

	"github.com/gofiber/fiber/v2"
)

type FileController struct {
	fileService services.FileServiceInterface
}

func NewFileController(fileService services.FileServiceInterface) *FileController {
	return &FileController{fileService: fileService}
}

// Create File
func (f *FileController) CreateFile(c *fiber.Ctx) error {
	var file models.File
	formFile, err := c.FormFile("img")
	if err != nil {
		return errors.FailedToProcessRequest.FiberError(c)
	}
	fileData, err := formFile.Open()
	if err != nil {
		return errors.FailedToOpenFile.FiberError(c)
	}
	buff := make([]byte, 512)
	if _, err = fileData.Read(buff); err != nil {
		return errors.InvalidImageFormat.FiberError(c)
	}
	if !((http.DetectContentType(buff) == "image/png") || (http.DetectContentType(buff) == "image/jpeg")) {
		return errors.FailedToValidatedData.FiberError(c)
	}
	defer fileData.Close()
	fileCreated, errFile := f.fileService.CreateFile(file, formFile, fileData)
	if errFile != nil {
		return errFile.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(fileCreated)
}

// // Delete File
// func (f *FileController) DeleteFile(c *fiber.Ctx) error {
// 	fileID := c.Params("fid")
// 	if err := f.fileService.DeleteFile(fileID, false); err != nil {
// 		return err
// 	}
// 	return c.SendStatus(fiber.StatusNoContent)
// }

// // Get File
// func (f *FileController) GetFile(c *fiber.Ctx) error {
// 	fileID := c.Params("fid")
// 	file, err := f.fileService.GetFile(fileID)
// 	if err != nil {
// 		return err.FiberError(c)
// 	}
// 	return c.Status(fiber.StatusOK).JSON(file)
// }
