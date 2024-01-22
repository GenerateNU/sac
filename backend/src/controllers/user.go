package controllers

import (
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
	var user models.UpdateUserRequestBody

	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	userID := c.Params("id")

	updatedUser, err := u.userService.UpdateUser(userID, user)

	if err != nil {
		return err
	}

	// Return the updated user details
	return c.Status(fiber.StatusOK).JSON(updatedUser)
}
