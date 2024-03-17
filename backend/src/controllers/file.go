package controllers

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/gofiber/fiber/v2"
)

type FileController struct {
	fileService services.FileServiceInterface
}

func NewFileController(fileService services.FileServiceInterface) *FileController {
	return &FileController{fileService: fileService}
}

// GetFile godoc
//
// @Summary		Retrieve a file
// @Description	Retrieves a file
// @ID			get-file
// @Tags      	file
// @Produce		json
// @Param		fileID		path	string	true	"File ID"
// @Success		200	  {object}	  models.File
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/files/:fileID  [get]
func (f *FileController) GetFile(c *fiber.Ctx) error {
	fileID := c.Params("fileID")
	
	file, err := f.fileService.GetFile(fileID)
	if err != nil {
		return err.FiberError(c)
	}

	return c.JSON(file)
}

// CreateFile godoc
//
// @Summary		Create a file
// @Description	Creates a file
// @ID			create-file
// @Tags      	file
// @Accept		multipart/form-data
// @Produce		json
// @Param		file	body	models.CreateFileRequestBody	true	"File"
// @Success		201	  {object}	  models.File
// @Failure     400   {object}    errors.Erro
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/files/  [post]
func (f *FileController) CreateFile(c *fiber.Ctx) error {
	var fileBody models.CreateFileRequestBody

	if parseErr := c.BodyParser(&fileBody); parseErr != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	formFile, parseErr := c.FormFile("file")
	if parseErr != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	file, err := f.fileService.CreateFile(&fileBody, formFile)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(file)
}
