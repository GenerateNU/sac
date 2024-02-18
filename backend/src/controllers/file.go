package controllers

import (
	"net/http"

	"strings"

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
	var fileRequestBody models.FileBody

	if err := c.BodyParser(&fileRequestBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

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
	print(file.AssociationID.String())
	fileCreated, errFile := f.fileService.CreateFile(fileRequestBody, file, formFile, fileData)
	if errFile != nil {
		return errFile.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(fileCreated)
}

// Get File
func (f *FileController) GetFile(c *fiber.Ctx) error {
	fileID := c.Params("fileID")
	file, err := f.fileService.GetFile(fileID)
	if err != nil {
		return err.FiberError(c)
	}
	arr := strings.SplitAfter(file.FileName, ".")
	lenArr := len(arr)
	print(arr[lenArr-1])
	c.Set("Content-Type", "image/jpeg")
	return c.Send(file.FileData)
}

// Get File Info
func (f *FileController) GetFileInfo(c *fiber.Ctx) error {
	days := c.Params("days")
	if days == "" {
		days = "7"
	}
	fileID := c.Params("fileID")
	fileInfo, err := f.fileService.GetFileInfo(fileID, days)
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(fileInfo)
}

// Delete File
func (f *FileController) DeleteFile(c *fiber.Ctx) error {
	fileID := c.Params("fileID")
	if err := f.fileService.DeleteFile(fileID, false); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
