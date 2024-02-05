package controllers

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService  services.AuthServiceInterface
	blacklist    []string
	AuthSettings config.AuthSettings
}

func NewAuthController(authService services.AuthServiceInterface, authSettings config.AuthSettings) *AuthController {
	return &AuthController{authService: authService, blacklist: []string{}, AuthSettings: authSettings}
}

// Me godoc
//
// @Summary		Gets the current user
// @Description	Returns the current user
// @ID			get-current-user
// @Tags      	user
// @Produce		json
// @Success		200	  {object}	  models.User
// @Failure     401   {string}    string "failed to get current user"
// @Router		/api/v1/auth/me  [get]
func (a *AuthController) Me(c *fiber.Ctx) error {
	claims, err := types.From(c)
	if err != nil {
		return err.FiberError(c)
	}
	user, err := a.authService.Me(claims.Issuer)
	if err != nil {
		return err.FiberError(c)
	}

	return c.JSON(user)
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
// @Success		200	  {object}	  string "success"
// @Failure     400   {string}    string "failed to parse body"
// @Failure     401   {string}    string "failed to login user"
// @Router		/api/v1/auth/login  [post]
func (a *AuthController) Login(c *fiber.Ctx) error {
	var userBody models.LoginUserResponseBody

	if err := c.BodyParser(&userBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	user, err := a.authService.Login(userBody)
	if err != nil {
		return err.FiberError(c)
	}

	accessToken, refreshToken, err := auth.CreateTokenPair(user.ID.String(), string(user.Role), a.AuthSettings)
	if err != nil {
		return err.FiberError(c)
	}

	// Set the tokens in the response
	c.Cookie(auth.CreateCookie("access_token", *accessToken, time.Now().Add(time.Minute*time.Duration(a.AuthSettings.AccessTokenExpiry))))
	c.Cookie(auth.CreateCookie("refresh_token", *refreshToken, time.Now().Add(time.Hour*time.Duration(a.AuthSettings.RefreshTokenExpiry))))

	return utilities.FiberMessage(c, fiber.StatusOK, "success")
}

// Refresh godoc
//
// @Summary		Refreshes a user's access token
// @Description	Refreshes a user's access token
// @ID			refresh-user
// @Tags      	user
// @Accept		json
// @Produce		json
// @Success		200	  {object}	  string "success"
// @Failure     401   {string}    string "failed to refresh access token"
// @Router		/api/v1/auth/refresh  [get]
func (a *AuthController) Refresh(c *fiber.Ctx) error {
	// Extract token values from cookies
	refreshTokenValue := c.Cookies("refresh_token")

	// Extract id from refresh token
	claims, err := auth.ExtractRefreshClaims(refreshTokenValue, a.AuthSettings.RefreshToken)
	if err != nil {
		return err.FiberError(c)
	}

	role, err := a.authService.GetRole(claims.Issuer)
	if err != nil {
		return err.FiberError(c)
	}

	accessToken, err := auth.RefreshAccessToken(refreshTokenValue, string(*role), a.AuthSettings.RefreshToken, a.AuthSettings.AccessTokenExpiry, a.AuthSettings.AccessToken)
	if err != nil {
		return err.FiberError(c)
	}

	// Set the access token in the response (e.g., in a cookie or JSON response)
	c.Cookie(auth.CreateCookie("access_token", *accessToken, time.Now().Add(time.Minute*60)))

	return utilities.FiberMessage(c, fiber.StatusOK, "success")
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
// @Router		/api/v1/auth/logout  [get]
func (a *AuthController) Logout(c *fiber.Ctx) error {
	// Extract token values from cookies
	accessTokenValue := c.Cookies("access_token")
	refreshTokenValue := c.Cookies("refresh_token")

	// TODO: Redis
	a.blacklist = append(a.blacklist, accessTokenValue)
	a.blacklist = append(a.blacklist, refreshTokenValue)

	// Expire and clear the cookies
	c.Cookie(auth.ExpireCookie("access_token"))
	c.Cookie(auth.ExpireCookie("refresh_token"))

	return utilities.FiberMessage(c, fiber.StatusOK, "success")
}