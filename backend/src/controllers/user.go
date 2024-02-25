package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"

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
// @Failure     500   {object}    errors.Error
// @Router		/users/  [post]
func (u *UserController) CreateUser(c *fiber.Ctx) error {
	var userBody models.CreateUserRequestBody

	if err := c.BodyParser(&userBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	user, err := u.userService.CreateUser(userBody)
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
	defaultLimit := 10
	defaultPage := 1

	categories, err := u.userService.GetUsers(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&categories)
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

	// Return the updated user details
	return c.Status(fiber.StatusOK).JSON(updatedUser)
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
