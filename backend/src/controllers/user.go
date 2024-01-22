package controllers

import (
	"github.com/GenerateNU/sac/backend/src/services"

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
		return err
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// DeleteUser godoc
//
// @Summary		Deletes the given userID
// @Description	Returns nil
// @ID			delete-user
// @Tags      	user
// @Produce		json
// @Success		204   {string}     string "No Content"
// @Failure     500   {string}     string "failed to get all users"
// @Router		/api/v1/users/:id  [delete]
func (u * UserController) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	err := u.userService.DeleteUser(userID)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
