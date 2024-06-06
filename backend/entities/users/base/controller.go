package base

import (
	"net/http"

	authEntities "github.com/GenerateNU/sac/backend/entities/auth"
	"github.com/GenerateNU/sac/backend/locals"
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
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}		error
// @Failure     500   {object}      error
// @Router		/auth/me  [get]
func (u *UserController) GetMe(c *fiber.Ctx) error {
	userID, err := locals.UserIDFrom(c)
	if err != nil {
		return err
	}

	user, err := u.userService.GetMe(*userID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(user)
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

// UpdatePassword godoc
//
// @Summary		Update a user's password
// @Description	Updates a user's password
// @ID			update-password
// @Tags      	user
// @Accept		json
// @Produce		json
// @Param		userID		path	string	true	"User ID"
// @Param		passwordBody	body	authEntities.UpdatePasswordRequestBody	true	"Password Body"
// @Success		200	  {string}	    utilities.SuccessResponse
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/users/{userID}/password  [patch]
func (u *UserController) UpdatePassword(c *fiber.Ctx) error {
	var passwordBody authEntities.UpdatePasswordRequestBody

	if err := c.BodyParser(&passwordBody); err != nil {
		return utilities.InvalidJSON()
	}

	err := u.userService.UpdatePassword(c.Params("userID"), passwordBody)
	if err != nil {
		return err
	}

	return utilities.FiberMessage(c, http.StatusOK, "success")
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
