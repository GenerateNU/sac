package base

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/oauth/sacic"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Service interface {
	Provider(c *fiber.Ctx) error
	ProviderCallback(c *fiber.Ctx) error
	ProviderLogout(c *fiber.Ctx) error
	SendCode(c *fiber.Ctx) error
	VerifyEmail(c *fiber.Ctx) error
}

type Handler struct {
	db       *gorm.DB
	emailer  email.Emailer
	validate *validator.Validate
}

func (h *Handler) Provider(c *fiber.Ctx) error {
	// try to get the user without re-authenticating
	if sacUser, err := sacic.CompleteUserAuth(c); err == nil {
		slog.Info("user auth callback", "user", sacUser)
		sacUser, err := FindOrCreateUser(c.Context(), h.db, sacUser)
		if err != nil {
			return err
		}
		slog.Info("user auth callback", "user", sacUser)
		locals.SetUser(c, sacUser)
	} else {
		sacic.BeginAuthHandler(c)
	}
	return nil
}

func (h *Handler) ProviderCallback(c *fiber.Ctx) error {
	user, err := sacic.CompleteUserAuth(c)
	if err != nil {
		return err
	}

	slog.Info("user auth callback", "user", user)

	return nil
}

func (h *Handler) ProviderLogout(c *fiber.Ctx) error {
	if err := sacic.Logout(c); err != nil {
		return err
	}

	return c.Redirect("/", http.StatusTemporaryRedirect)
}

func (h *Handler) SendCode(c *fiber.Ctx) error {
	var body emailRequestBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	user, err := users.GetUserByEmail(h.db, body.Email)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return errors.New("user is already verified")
	}

	activeOTP, err := GetActiveTokenByUserID(h.db, user.ID)
	if err != nil {
		if !utilities.IsNotFound(err) {
			return err
		}
	}

	if activeOTP != nil {
		if err := h.emailer.SendVerification(context.TODO(), user.Email, activeOTP.Token); err != nil {
			return err
		}

		return nil
	}

	otp, err := auth.GenerateOTP(constants.OTP_LENGTH)
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	if err := SaveToken(h.db, user.ID, *otp, time.Now().Add(constants.OTP_EXPIRY).UTC()); err != nil {
		return err
	}

	return nil
}

func (h *Handler) VerifyEmail(c *fiber.Ctx) error {
	var body verifyEmailRequestBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	if err := utilities.Validate(h.validate, body); err != nil {
		return err
	}

	user, err := users.GetUserByEmail(h.db, body.Email)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return errors.New("user is already verified")
	}

	otp, err := GetToken(h.db, body.Token)
	if err != nil {
		return err
	}

	if otp.Token != body.Token {
		slog.Error("invalid otp", "otp", otp.Token, "token", body.Token)
		return utilities.BadRequest(errors.New("invalid otp"))
	}

	if otp.ExpiresAt.Before(time.Now().UTC()) {
		return utilities.BadRequest(errors.New("otp expired"))
	}

	return users.VerifyEmail(c.Context(), h.db, otp.Token, user.ID)
}
