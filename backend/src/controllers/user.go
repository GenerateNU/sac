package controllers

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

// GetAllUsers godoc
//
// @Summary		Gets all users
// @Description	Returns all users
// @ID			get-all-users
// @Tags      	user
// @Produce		json
// @Success		200	  {object}	  []models.User
// @Failure     500   {string}    string "failed to get all users"
// @Router		/api/v1/users/  [get]
func (u *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		return utilities.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// GetUser godoc
//
// @Summary		Gets a user
// @Description	Returns a user
// @ID			get-user-by-id
// @Tags      	user
// @Produce		json
// @Param		id	path	string	true	"User ID"
// @Success		200	  {object}	  models.User
// @Failure     404   {string}    string "user not found"
// @Failure		400   {string}    string "failed to validate id"
// @Failure     500   {string}    string "failed to get user"
// @Router		/api/v1/users/:id  [get]
func (u *UserController) GetUser(c *fiber.Ctx) error {
	user, err := u.userService.GetUser(c.Params("id"))
	if err != nil {
		return utilities.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// UpdateUser godoc
//
// @Summary		Updates a user
// @Description	Updates a user
// @ID			update-user-by-id
// @Tags      	user
// @Produce		json
// @Success		200	  {object}	  models.User
// @Failure     404   {string}    string "user not found"
// @Failure 	400   {string}    string "invalid request body"
// @Failure		400   {string}    string "failed to validate id"
// @Failure		500   {string}	  string "database error"
// @Failure		500   {string} 	  string "failed to hash password"
// @Router		/api/v1/users/:id  [patch]
func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	var user models.UserRequestBody

	if err := c.BodyParser(&user); err != nil {
		return utilities.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	updatedUser, err := u.userService.UpdateUser(c.Params("id"), user)
	if err != nil {
		return utilities.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	// Return the updated user details
	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

// DeleteUser godoc
//
// @Summary		Deletes the given userID
// @Description	Returns nil
// @ID			delete-user
// @Tags      	user
// @Produce		json
// @Success		200
// @Failure     500   {string}     string "failed to get all users"
// @Router		/api/v1/users/:id  [delete]
func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	err := u.userService.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}
