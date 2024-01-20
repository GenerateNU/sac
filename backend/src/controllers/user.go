package controllers

import (
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/models"
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

// Create User godoc
//
// @Summary		Creates a User
// @Description	Creates a user
// @ID			create-user
// @Tags      	user
// @Accept      json
// @Produce		json
// @Success		201	  {object}	  models.User
// @Failure     400   {string}    string "failed to create user"
// @Failure     500   {string}    string "internal server error"
// @Router		/api/v1/users/  [post]
func (u *UserController) CreateUser(c *fiber.Ctx) error {
	var userBody models.CreateUserRequestBody

	if err := c.BodyParser(&userBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to process the request")
	}

	user, err := u.userService.CreateUser(userBody)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
