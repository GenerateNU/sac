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

// GetUser godoc
//
// @Summary		Gets specific user
// @Description	Returns specific user
// @ID			get-user
// @Tags      	user
// @Produce		json
// @Success		200	  {object}	  models.User
// @Failure     400   {string}    string "id must be a positive number"
// @Failure     404   {string}    string "record not found"
// @Failure     500   {string}    string
// @Router		/api/v1/users/  [get]
func (u *UserController) GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	user, err := u.userService.GetUser(userID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}
