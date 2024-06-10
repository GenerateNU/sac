package base

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService UserServiceInterface
}

func NewUserController(userService UserServiceInterface) *UserController {
	return &UserController{userService: userService}
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
// @Failure     401   {object}      error
// @Failure     400   {object}      error
// @Failure     500   {object}      error
// @Router		/users/  [get]
func (u *UserController) GetUsers(c *fiber.Ctx) error {
	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPageInfo
	}

	users, err := u.userService.GetUsers(*pageInfo)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(&users)
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
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/users/{userID}/  [get]
func (u *UserController) GetUser(c *fiber.Ctx) error {
	user, err := u.userService.GetUser(c.Params("userID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(user)
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
// @Param		userBody	body	UpdateUserRequestBody	true	"User Body"
// @Success		200	  {object}	  models.User
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/users/{userID}/  [patch]
func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	var user UpdateUserRequestBody

	if err := c.BodyParser(&user); err != nil {
		return utilities.InvalidJSON()
	}

	updatedUser, err := u.userService.UpdateUser(c.Params("userID"), user)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(updatedUser)
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
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/users/{userID}/  [delete]
func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	err := u.userService.DeleteUser(c.Params("userID"))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
