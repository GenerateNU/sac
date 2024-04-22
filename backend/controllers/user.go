package controllers

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/models"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

// CreateUser godoc
//
// @Summary		Create a user
// @Description	Creates a user
// @ID			create-user
// @Tags      	user
// @Accept		json
// @Produce		json
// @Param		userBody	body	models.CreateUserRequestBody	true	"User Body"
// @Success		201	  {object}	  models.User
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     409   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/users/  [post]
func (u *UserController) CreateUser(c *fiber.Ctx) error {
	var userBody models.CreateUserRequestBody

	if err := c.BodyParser(&userBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	user, tokens, err := u.userService.CreateUser(userBody)
	if err != nil {
		return err.FiberError(c)
	}

	err = auth.SetResponseTokens(c, tokens)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUsers godoc
//
// @Summary		Retrieve all users
// @Description	Retrieves all users
// @ID			get-users
// @Tags      	user
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	    []models.User
// @Failure     401   {object}      errors.Error
// @Failure     400   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/  [get]
func (u *UserController) GetUsers(c *fiber.Ctx) error {
	categories, err := u.userService.GetUsers(c.Query("limit", constants.DEFAULT_LIMIT_STRING), c.Query("page", constants.DEFAULT_PAGE_STRING))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&categories)
}

// Me godoc
//
// @Summary		Retrieves the currently authenticated user
// @Description	Retrieves the currently authenticated user
// @ID			get-me
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Security   Bearer
// @Success		200	  {object}	    models.User
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}		errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/auth/me  [get]
func (u *UserController) GetMe(c *fiber.Ctx) error {
	claims, err := auth.From(c)
	if err != nil {
		return err.FiberError(c)
	}

	user, err := u.userService.GetMe(claims.Issuer)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// GetUser godoc
//
// @Summary		Retrieve a user
// @Description	Retrieves a user
// @ID			get-user
// @Tags      	user
// @Produce		json
// @Param		userID	path	string	true	"User ID"
// @Success		200	  {object}	    models.User
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/{userID}/  [get]
func (u *UserController) GetUser(c *fiber.Ctx) error {
	user, err := u.userService.GetUser(c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// UpdateUser godoc
//
// @Summary		Update a user
// @Description	Updates a user
// @ID			update-user
// @Tags      	user
// @Accept		json
// @Produce		json
// @Param		userID		path	string	true	"User ID"
// @Param		userBody	body	models.UpdateUserRequestBody	true	"User Body"
// @Success		200	  {object}	  models.User
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/users/{userID}/  [patch]
func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	var user models.UpdateUserRequestBody

	if err := c.BodyParser(&user); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedUser, err := u.userService.UpdateUser(c.Params("userID"), user)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

// UpdatePassword godoc
//
// @Summary		Update a user's password
// @Description	Updates a user's password
// @ID			update-password
// @Tags      	user
// @Accept		json
// @Produce		json
// @Param		userID		path	string	true	"User ID"
// @Param		passwordBody	body	models.UpdatePasswordRequestBody	true	"Password Body"
// @Success		200	  {string}	    utilities.SuccessResponse
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/{userID}/password  [patch]
func (u *UserController) UpdatePassword(c *fiber.Ctx) error {
	var passwordBody models.UpdatePasswordRequestBody

	if err := c.BodyParser(&passwordBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	err := u.userService.UpdatePassword(c.Params("userID"), passwordBody)
	if err != nil {
		return err.FiberError(c)
	}

	return utilities.FiberMessage(c, fiber.StatusOK, "success")
}

// DeleteUser godoc
//
// @Summary		Delete a user
// @Description	Deletes a user
// @ID			delete-user
// @Tags      	user
// @Produce		json
// @Param		userID	path	string	true	"User ID"
// @Success		204	  {string}	    utilities.SuccessResponse
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/{userID}/  [delete]
func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	err := u.userService.DeleteUser(c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
