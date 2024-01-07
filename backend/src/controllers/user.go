package controllers

import (
	"backend/src/services"

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
// @Failure     404   {string}    string "Failed to fetch users"
// @Router		/api/users/  [get]
func (u *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Failed to fetch users")
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
