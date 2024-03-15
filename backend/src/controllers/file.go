package controllers

// type FileController struct {
// 	fileService services.FileServiceInterface
// }

// func NewFileController(fileService services.FileServiceInterface) *FileController {
// 	return &FileController{fileService: fileService}
// }

// // CreateFile godoc
// //
// // @Summary		Create a file
// // @Description	Creates a file
// // @ID			create-file
// // @Tags      	file
// // @Accept		multipart/form-data
// // @Accept		json
// // @Produce		json
// // @Param		file	body	models.CreateFileRequestBody	true	"File"
// // @Success		201	  {object}	  models.File
// // @Failure     400   {object}    errors.Error
// // @Failure     500   {object}    errors.Error
// // @Router		/files/  [post]
// func (f *FileController) CreateFile(c *fiber.Ctx) error {
// 	var fileBody models.CreateFileRequestBody

// 	if parseErr := c.BodyParser(&fileBody); parseErr != nil {
// 		return errors.FailedToParseRequestBody.FiberError(c)
// 	}

// 	formFile, parseErr := c.FormFile("file")
// 	if parseErr != nil {
// 		return errors.FailedToParseRequestBody.FiberError(c)
// 	}

// 	file, err := f.fileService.CreateFile(&fileBody, formFile)
// 	if err != nil {
// 		return err.FiberError(c)
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(file)
// }
