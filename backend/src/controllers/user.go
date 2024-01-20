package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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
// @Success		200	  {object}	  []models.User
// @Failure     400   {string}    string "invalid user id"
// @Failure     500   {string}    string "failed to fetch user"
// @Router		/api/v1/users/{id}  [get]
func (u *UserController) GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	if err := claims.Valid(); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	userID := uint(id)

	user, err := u.userService.GetUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (u *UserController) CurrentUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	if err := claims.Valid(); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fmt.Println(claims.Issuer)

	userID, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	user, err := u.userService.GetUser(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (u *UserController) Register(c *fiber.Ctx) error {
	var userBody models.CreateUserResponseBody

	if err := c.BodyParser(&userBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse body")
	}

	// TODO: fiber.NewError vs fiber.Map
	user, err := u.userService.Register(userBody)
	if err != nil {
		return err
	}

	// TODO: Should we login the user after registering?

	return c.Status(fiber.StatusOK).JSON(user)
}

type CustomClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
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

	// Create Access Token with Custom Claims
	// accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	// 	Issuer:    strconv.Itoa(int(user.ID)),
	// 	ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // Short-lived access token
	// })

	// &CustomClaims{
	// 	StandardClaims: jwt.StandardClaims{
	// 		Issuer:    strconv.Itoa(int(user.ID)),
	// 		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // Short-lived access token
	// 	},
	// 	Role: string(user.Role),
	// }
	// accessToken, err := accessTokenClaims.SignedString([]byte("access_secret"))
	// if err != nil {
	// 	return err
	// }

	// Create Refresh Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(), // Long-lived refresh token
	})
	refreshToken, err := token.SignedString([]byte("token"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// // Set Access Token Cookie
	// accessTokenCookie := fiber.Cookie{
	// 	Name:     "access_token",
	// 	Value:    accessToken,
	// 	Expires:  time.Now().Add(time.Minute * 15),
	// 	HTTPOnly: true,
	// }
	// c.Cookie(&accessTokenCookie)

	// Set Refresh Token Cookie
	tokenCookie := fiber.Cookie{
		Name:     "token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
	}
	c.Cookie(&tokenCookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func (u *UserController) Refresh(c *fiber.Ctx) error {
	// Extract the refresh token from the request (e.g., from a cookie or request body)
	refreshTokenValue := c.Cookies("refresh_token")

	// Validate the refresh token
	claims := &CustomClaims{}
	refreshToken, err := jwt.ParseWithClaims(refreshTokenValue, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("refresh_secret"), nil
	})

	if err != nil || !refreshToken.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired refresh token")
	}

	// At this point, the refresh token is valid, and you can generate a new access token
	newAccessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    claims.Issuer,
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // Short-lived access token
	})

	newAccessToken, err := newAccessTokenClaims.SignedString([]byte("access_secret"))
	if err != nil {
		return err
	}

	// Set the new access token in the response (e.g., in a cookie or JSON response)
	accessTokenCookie := fiber.Cookie{
		Name:     "access_token",
		Value:    newAccessToken,
		Expires:  time.Now().Add(time.Minute * 15),
		HTTPOnly: true,
	}
	c.Cookie(&accessTokenCookie)

	return c.JSON(fiber.Map{
		"message": "Token refreshed successfully",
	})
}

func (u *UserController) Logout(c *fiber.Ctx) error {
	// var blacklist []string
	// Extract token values from cookies
	// accessTokenValue := c.Cookies("access_token")
	// refreshTokenValue := c.Cookies("refresh_token")

	// // Add tokens to the blacklist or invalidate references (implementation depends on your approach)
	// // For simplicity, let's assume a global blacklist (not recommended for production)
	// // You may use a distributed cache, database, or another mechanism in a real-world scenario.
	// blacklist = append(blacklist, accessTokenValue)
	// blacklist = append(blacklist, refreshTokenValue)

	// Expire and clear the cookies
	// c.Cookie(&fiber.Cookie{
	// 	Name:     "access_token",
	// 	Value:    "",
	// 	Expires:  time.Now().Add(-time.Hour),
	// 	HTTPOnly: true,
	// })
	c.Cookie(&fiber.Cookie{
		Name:     "token_cookie",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
