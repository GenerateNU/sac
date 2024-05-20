package base

import (
	"context"
	"net/http"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/constants"
	authEntities "github.com/GenerateNU/sac/backend/entities/auth"
	users "github.com/GenerateNU/sac/backend/entities/users/base"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService AuthServicer
}

func NewAuthController(authService AuthServicer) *AuthController {
	return &AuthController{authService: authService}
}

// Login godoc
//
// @Summary		Logs in a user
// @Description	Logs in a user
// @ID			login-user
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Param		loginBody	body	authEntities.LoginResponseBody	true	"Login Body"
// @Success		200	  {object}	    models.User
// @Failure     400   {object}      error
// @Failure     404   {object}		error
// @Failure     500   {object}      error
// @Router		/auth/login  [post]
func (a *AuthController) Login(c *fiber.Ctx) error {
	var userBody authEntities.LoginResponseBody

	if err := c.BodyParser(&userBody); err != nil {
		return utilities.InvalidJSON()
	}

	ctx, cancel := context.WithTimeout(c.Context(), constants.HANDLER_TIMEOUT)
	defer cancel()

	user, tokens, err := a.authService.Login(ctx, userBody)
	if err != nil {
		return err
	}

	err = auth.SetResponseTokens(c, tokens)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(user)
}

// Register godoc
//
// @Summary		Registers a user
// @Description	Registers a user
// @ID			register-user
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Param		userBody	body	CreateUserRequestBody	true	"User Body"
// @Success		201	  {object}	    models.User
// @Failure     400   {object}      error
// @Failure     404   {object}		error
// @Failure     500   {object}      error
// @Router		/auth/register  [post]
func (a *AuthController) Register(c *fiber.Ctx) error {
	var userBody users.CreateUserRequestBody

	if err := c.BodyParser(&userBody); err != nil {
		return utilities.InvalidJSON()
	}

	ctx, cancel := context.WithTimeout(c.Context(), constants.HANDLER_TIMEOUT)
	defer cancel()

	user, tokens, err := a.authService.Register(ctx, userBody)
	if err != nil {
		return err
	}

	err = auth.SetResponseTokens(c, tokens)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// Refresh godoc
//
// @Summary		Refreshes a user's access token
// @Description	Refreshes a user's access token
// @ID			refresh-user
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Success		200	  {object}	  utilities.SuccessResponse
// @Failure     400   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/auth/refresh  [post]
func (a *AuthController) Refresh(c *fiber.Ctx) error {
	var refreshBody authEntities.RefreshTokenRequestBody

	if err := c.BodyParser(&refreshBody); err != nil {
		return utilities.InvalidJSON()
	}

	ctx, cancel := context.WithTimeout(c.Context(), constants.HANDLER_TIMEOUT)
	defer cancel()

	tokens, err := a.authService.Refresh(ctx, refreshBody.RefreshToken)
	if err != nil {
		return err
	}

	err = auth.SetResponseTokens(c, tokens)
	if err != nil {
		return err
	}

	return utilities.FiberMessage(c, http.StatusOK, "success")
}

// Logout godoc
//
// @Summary		Logs out a user
// @Description	Logs out a user
// @ID			logout-user
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Success		200	  {object}	  utilities.SuccessResponse
// @Router		/auth/logout  [post]
func (a *AuthController) Logout(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), constants.HANDLER_TIMEOUT)
	defer cancel()

	err := a.authService.Logout(ctx, c)
	if err != nil {
		return err
	}

	return utilities.FiberMessage(c, http.StatusOK, "success")
}

// ForgotPassword godoc
//
// @Summary		Generates a password reset token
// @Description	Generates a password reset token
// @ID			forgot-password
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Param		email	body	string	true	"Email"
// @Success		200	  {object}	     utilities.SuccessResponse
// @Failure     400   {object}       error
// @Failure     429   {object}       error
// @Failure     500   {object}       error
// @Router		/auth/forgot-password  [post]
func (a *AuthController) ForgotPassword(c *fiber.Ctx) error {
	var emailBody EmailRequestBody

	if err := c.BodyParser(&emailBody); err != nil {
		return utilities.InvalidJSON()
	}

	ctx, cancel := context.WithTimeout(c.Context(), constants.HANDLER_TIMEOUT)
	defer cancel()

	if err := a.authService.ForgotPassword(ctx, emailBody.Email); err != nil {
		return err
	}

	return utilities.FiberMessage(c, http.StatusOK, "success")
}

// VerifyPasswordResetToken godoc
//
// @Summary		Verifies a password reset token
// @Description	Verifies a password reset token
// @ID			verify-password-reset-token
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Param		tokenBody	body	 VerifyPasswordResetTokenRequestBody	true	"Password Reset Token Body"
// @Security   Bearer
// @Success		200	  {object}	     utilities.SuccessResponse
// @Failure     400   {object}       error
// @Failure     429   {object}       error
// @Failure     500   {object}       error
// @Router		/auth/verify-reset  [post]
func (a *AuthController) VerifyPasswordResetToken(c *fiber.Ctx) error {
	var tokenBody VerifyPasswordResetTokenRequestBody

	if err := c.BodyParser(&tokenBody); err != nil {
		return utilities.InvalidJSON()
	}

	ctx, cancel := context.WithTimeout(c.Context(), constants.HANDLER_TIMEOUT)
	defer cancel()

	if err := a.authService.VerifyPasswordResetToken(ctx, tokenBody); err != nil {
		return err
	}

	return utilities.FiberMessage(c, http.StatusOK, "success")
}

// SendCode godoc
//
// @Summary		Sends a verification code
// @Description	Sends a verification code
// @ID			send-verification-code
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Param		email	body	string	true	"Email"
// @Success		200	  {object}	     utilities.SuccessResponse
// @Failure     400   {object}       error
// @Failure     429   {object}       error
// @Failure     500   {object}       error
// @Router		/auth/send-code [post]
func (a *AuthController) SendCode(c *fiber.Ctx) error {
	var emailBody EmailRequestBody

	if err := c.BodyParser(&emailBody); err != nil {
		return utilities.InvalidJSON()
	}

	ctx, cancel := context.WithTimeout(c.Context(), constants.HANDLER_TIMEOUT)
	defer cancel()

	if err := a.authService.SendCode(ctx, emailBody.Email); err != nil {
		return err
	}

	return utilities.FiberMessage(c, http.StatusOK, "success")
}

// VerifyEmail godoc
//
// @Summary		Verifies an email
// @Description	Verifies an email
// @ID			verify-email
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Param		tokenBody	body	 VerifyEmailRequestBody	true	"Email Verification Token Body"
// @Success		200	  {object}	     utilities.SuccessResponse
// @Failure     400   {object}       error
// @Failure     429   {object}       error
// @Failure     500   {object}       error
// @Router		/auth/verify-email  [post]
func (a *AuthController) VerifyEmail(c *fiber.Ctx) error {
	var tokenBody VerifyEmailRequestBody

	if err := c.BodyParser(&tokenBody); err != nil {
		return utilities.InvalidJSON()
	}

	ctx, cancel := context.WithTimeout(c.Context(), constants.HANDLER_TIMEOUT)
	defer cancel()

	if err := a.authService.VerifyEmail(ctx, tokenBody); err != nil {
		return err
	}

	return utilities.FiberMessage(c, http.StatusOK, "success")
}
