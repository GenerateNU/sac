package base

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type FileController struct {
	fileService FileServiceInterface
}

func NewFileController(fileService FileServiceInterface) *FileController {
	return &FileController{fileService: fileService}
}

// GetFiles godoc
//
// @Summary		Retrieve all files
// @Description	Retrieves all files
// @ID			get-files
// @Tags      	file
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	  []models.File
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/files/  [get]
func (f *FileController) GetFiles(c *fiber.Ctx) error {
	pagination, err := middleware.GetPagination(c)
	if err != nil {
		return err
	}

	files, err := f.fileService.GetFiles(*pagination)
	if err != nil {
		return err
	}

	return c.JSON(files)
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
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/files/:fileID  [get]
func (f *FileController) GetFile(c *fiber.Ctx) error {
	fileID := c.Params("fileID")

	file, err := f.fileService.GetFile(fileID)
	if err != nil {
		return err
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
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/files/  [post]
func (f *FileController) CreateFile(c *fiber.Ctx) error {
	var fileBody models.CreateFileRequestBody

	if err := c.BodyParser(&fileBody); err != nil {
		return utilities.InvalidJSON()
	}

	formFile, err := c.FormFile("file")
	if err != nil {
		return err
	}

	file, err := f.fileService.CreateFile(&fileBody, formFile)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(file)
}

// DeleteFile godoc
//
// @Summary		Delete a file
// @Description	Deletes a file
// @ID			delete-file
// @Tags      	file
// @Produce		json
// @Param		fileID		path	string	true	"File ID"
// @Success		201	  {object}	  models.File
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/files/:fileID  [delete]
func (f *FileController) DeleteFile(c *fiber.Ctx) error {
	fileID := c.Params("fileID")

	err := f.fileService.DeleteFile(fileID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
