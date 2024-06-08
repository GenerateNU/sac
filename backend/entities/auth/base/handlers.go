package base

import (
	"log/slog"
	"net/http"

	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/oauth/sacic"

	"github.com/GenerateNU/sac/backend/locals"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Service interface {
	Provider(c *fiber.Ctx) error
	ProviderCallback(c *fiber.Ctx) error
	ProviderLogout(c *fiber.Ctx) error
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
