package controllers

import (
	"strconv"
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
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

// Register godoc
//
// @Summary		Registers a user
// @Description	Registers a user
// @ID			register-user
// @Tags      	user
// @Accept		json
// @Produce		json
// @Param		userBody	body	[]string	true	"User Body"
// @Success		201	  {object}	  models.User
// @Failure     400   {string}    string "failed to parse body"
// @Failure     400   {string}    string "failed to register user"
// @Router		/api/v1/users/auth/register  [post]
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
			"message": err.Error(),
		})
	}

	// TODO: Should we login the user after registering?
	return c.Status(fiber.StatusOK).JSON(user)
}

// Login godoc
//
// @Summary		Logs in a user
// @Description	Logs in a user
// @ID			login-user
// @Tags      	user
// @Accept		json
// @Produce		json
// @Param		userBody	body	[]string	true	"User Body"
// @Success		200	  {object}	  string
// @Failure     400   {string}    string "failed to parse body"
// @Failure     401   {string}    string "failed to login user"
// @Router		/api/v1/users/auth/login  [post]
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

// Refresh godoc
//
// @Summary		Refreshes a user's access token
// @Description	Refreshes a user's access token
// @ID			refresh-user
// @Tags      	user
// @Accept		json
// @Produce		json
// @Success		200	  {object}	  string
// @Failure     401   {string}    string "failed to refresh access token"
// @Router		/api/v1/users/auth/refresh  [get]
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

// Logout godoc
//
// @Summary		Logs out a user
// @Description	Logs out a user
// @ID			logout-user
// @Tags      	user
// @Accept		json
// @Produce		json
// @Success		200	  {object}	  string
// @Failure     401   {string}    string "failed to logout user"
// @Router		/api/v1/users/auth/logout  [get]
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
