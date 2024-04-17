package controllers

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService services.AuthServiceInterface
}

func NewAuthController(authService services.AuthServiceInterface) *AuthController {
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
// @Param		loginBody	body	models.LoginUserResponseBody	true	"Login Body"
// @Success		200	  {object}	    models.User
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}		errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/auth/login  [post]
func (a *AuthController) Login(c *fiber.Ctx) error {
	var userBody models.LoginUserResponseBody

	if err := c.BodyParser(&userBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	user, tokens, err := a.authService.Login(userBody)
	if err != nil {
		return err.FiberError(c)
	}

	err = auth.SetResponseTokens(c, tokens)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(user)
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
// @Failure     400   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/auth/refresh  [post]
func (a *AuthController) Refresh(c *fiber.Ctx) error {
	var refreshBody models.RefreshTokenRequestBody

	if err := c.BodyParser(&refreshBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	tokens, err := a.authService.Refresh(refreshBody.RefreshToken)
	if err != nil {
		return err.FiberError(c)
	}

	err = auth.SetResponseTokens(c, tokens)
	if err != nil {
		return err.FiberError(c)
	}

	return utilities.FiberMessage(c, fiber.StatusOK, "success")
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
	err := a.authService.Logout(c)
	if err != nil {
		return err.FiberError(c)
	}

	return utilities.FiberMessage(c, fiber.StatusOK, "success")
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
// @Failure     400   {object}       errors.Error
// @Failure     429   {object}       errors.Error
// @Failure     500   {object}       errors.Error
// @Router		/auth/forgot-password  [post]
func (a *AuthController) ForgotPassword(c *fiber.Ctx) error {
	var emailBody models.EmailRequestBody

	if err := c.BodyParser(&emailBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	if err := a.authService.ForgotPassword(emailBody.Email); err != nil {
		return err.FiberError(c)
	}

	return utilities.FiberMessage(c, fiber.StatusOK, "success")
}

// VerifyPasswordResetToken godoc
//
// @Summary		Verifies a password reset token
// @Description	Verifies a password reset token
// @ID			verify-password-reset-token
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Param		tokenBody	body	 models.VerifyPasswordResetTokenRequestBody	true	"Password Reset Token Body"
// @Security   Bearer
// @Success		200	  {object}	     utilities.SuccessResponse
// @Failure     400   {object}       errors.Error
// @Failure     429   {object}       errors.Error
// @Failure     500   {object}       errors.Error
// @Router		/auth/verify-reset  [post]
func (a *AuthController) VerifyPasswordResetToken(c *fiber.Ctx) error {
	var tokenBody models.VerifyPasswordResetTokenRequestBody

	if err := c.BodyParser(&tokenBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	if err := a.authService.VerifyPasswordResetToken(tokenBody); err != nil {
		return err.FiberError(c)
	}

	return utilities.FiberMessage(c, fiber.StatusOK, "success")
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
// @Failure     400   {object}       errors.Error
// @Failure     429   {object}       errors.Error
// @Failure     500   {object}       errors.Error
// @Router		/auth/send-code [post]
func (a *AuthController) SendCode(c *fiber.Ctx) error {
	var emailBody models.EmailRequestBody

	if err := c.BodyParser(&emailBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	if err := a.authService.SendCode(emailBody.Email); err != nil {
		return err.FiberError(c)
	}

	return utilities.FiberMessage(c, fiber.StatusOK, "success")
}

// VerifyEmail godoc
//
// @Summary		Verifies an email
// @Description	Verifies an email
// @ID			verify-email
// @Tags      	auth
// @Accept		json
// @Produce		json
// @Param		tokenBody	body	 models.VerifyEmailRequestBody	true	"Email Verification Token Body"
// @Success		200	  {object}	     utilities.SuccessResponse
// @Failure     400   {object}       errors.Error
// @Failure     429   {object}       errors.Error
// @Failure     500   {object}       errors.Error
// @Router		/auth/verify-email  [post]
func (a *AuthController) VerifyEmail(c *fiber.Ctx) error {
	var tokenBody models.VerifyEmailRequestBody

	if err := c.BodyParser(&tokenBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	if err := a.authService.VerifyEmail(tokenBody); err != nil {
		return err.FiberError(c)
	}

	return utilities.FiberMessage(c, fiber.StatusOK, "success")
}
