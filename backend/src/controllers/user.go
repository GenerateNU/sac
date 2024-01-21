package controllers

import (
	"strconv"
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
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

// GetUser godoc
//
// @Summary		Gets a user
// @Description	Returns a user
// @ID			get-user
// @Tags      	user
// @Produce		json
// @Param		id	path	int	true	"User ID"
// @Success		200	  {object}	  models.User
// @Failure     400   {string}    string "Invalid user id"
// @Failure     500   {string}    string "Failed to fetch user"
// @Router		/api/v1/users/{id}  [get]
func (u *UserController) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	userID := uint(id)

	user, err := u.userService.GetUser(userID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (u *UserController) Register(c *fiber.Ctx) error {
	var userBody models.CreateUserResponseBody

	if err := c.BodyParser(&userBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse body",
		})
	}

	user, err := u.userService.Register(userBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to register user",
		})
	}

	// TODO: Should we login the user after registering?
	return c.Status(fiber.StatusOK).JSON(user)
}

func (u *UserController) Login(c *fiber.Ctx) error {
	var userBody models.LoginUserResponseBody

	if err := c.BodyParser(&userBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse body",
		})
	}

	user, err := u.userService.Login(userBody)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	accessToken, err := auth.CreateAccessToken(strconv.Itoa(int(user.ID)), string(user.Role))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	refreshToken, err := auth.CreateRefreshToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Set the tokens in the response
	c.Cookie(auth.CreateCookie("access_token", *accessToken, time.Now().Add(time.Minute*15)))
	c.Cookie(auth.CreateCookie("refresh_token", *refreshToken, time.Now().Add(time.Hour*24*30)))

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func (u *UserController) Refresh(c *fiber.Ctx) error {
	// Extract token values from cookies
	accessTokenValue := c.Cookies("access_token")
	refreshTokenValue := c.Cookies("refresh_token")

	accessToken, err := auth.RefreshAccessToken(accessTokenValue, refreshTokenValue)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Set the access token in the response (e.g., in a cookie or JSON response)
	c.Cookie(auth.CreateCookie("access_token", *accessToken, time.Now().Add(time.Minute*15)))

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func (u *UserController) Logout(c *fiber.Ctx) error {
	// var blacklist []string
	// Extract token values from cookies
	// accessTokenValue := c.Cookies("access_token")
	// refreshTokenValue := c.Cookies("refresh_token")

	// TODO: Implement blacklist, ideally with Redis
	// blacklist = append(blacklist, accessTokenValue)
	// blacklist = append(blacklist, refreshTokenValue)

	// Expire and clear the cookies
	auth.ExpireCookie("access_token")
	auth.ExpireCookie("refresh_token")

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
