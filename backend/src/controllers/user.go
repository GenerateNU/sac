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
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	user, err := u.userService.CreateUser(userBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
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
		return err.FiberError(c)
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
	var user models.UpdateUserRequestBody

	if err := c.BodyParser(&user); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedUser, err := u.userService.UpdateUser(c.Params("id"), user)
	if err != nil {
		return err.FiberError(c)
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
// @Success		204   {string}     string "no content"
// @Failure     500   {string}     string "failed to get all users"
// @Router		/api/v1/users/:id  [delete]
func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	err := u.userService.DeleteUser(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (u *UserController) GetUserTags(c *fiber.Ctx) error {
	tags, err := u.userService.GetUserTags(c.Params("uid"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(&tags)
}

func (u *UserController) CreateUserTags(c *fiber.Ctx) error {
	var requestBody models.CreateUserTagsBody
	if err := c.BodyParser(&requestBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	tags, err := u.userService.CreateUserTags(c.Params("uid"), requestBody)
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusCreated).JSON(&tags)
}

func (u *UserController) CreateFollowing(c *fiber.Ctx) error {
	err := u.userService.CreateFollowing(c.Params("user_id"), c.Params("club_id"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (u *UserController) DeleteFollowing(c *fiber.Ctx) error {

	err := u.userService.DeleteFollowing(c.Params("user_id"), c.Params("club_id"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (u *UserController) GetAllFollowing(c *fiber.Ctx) error {
	clubs, err := u.userService.GetFollowing(c.Params("user_id"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(clubs)
}
