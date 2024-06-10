package base

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/GenerateNU/sac/backend/integrations/oauth/soth"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Service interface {
	Login(c *fiber.Ctx) error
	Provider(c *fiber.Ctx) error
	ProviderCallback(c *fiber.Ctx) error
	ProviderLogout(c *fiber.Ctx) error
}

type Handler struct {
	db           *gorm.DB
	authProvider soth.Provider
}

func (h *Handler) Login(c *fiber.Ctx) error {
	sothic.SetProvider(c, h.authProvider.Name())
	if gfUser, err := sothic.CompleteUserAuth(c); err == nil {
		user, err := FindOrCreateUser(context.TODO(), h.db, gfUser)
		if err != nil {
			return err
		}

		if err := sothic.StoreInSession("user", user.Marshal(), c); err != nil {
			return err
		}

		return nil
	} else {
		return sothic.BeginAuthHandler(c)
	}
}

func (h *Handler) Provider(c *fiber.Ctx) error {
	if gfUser, err := sothic.CompleteUserAuth(c); err == nil {
		user, err := FindOrCreateUser(context.TODO(), h.db, gfUser)
		if err != nil {
			return err
		}

		if err := sothic.StoreInSession("user", user.Marshal(), c); err != nil {
			return err
		}

		return nil
	} else {
		return sothic.BeginAuthHandler(c)
	}
}

func (h *Handler) ProviderCallback(c *fiber.Ctx) error {
	defer func() {
		c.Cookie(&fiber.Cookie{
			Name:    "redirect",
			Value:   "",
			Expires: time.Now().Add(-1 * time.Hour), // expire the cookie immediately
			// MARK: secure should be true in prod
			// use go build tags to do this
			HTTPOnly: true,
		})
	}()

	gfUser, err := sothic.CompleteUserAuth(c)
	if err != nil {
		return err
	}

	user, err := FindOrCreateUser(context.TODO(), h.db, gfUser)
	if err != nil {
		return err
	}

	if err := sothic.StoreInSession("user", user.Marshal(), c); err != nil {
		return err
	}

	return c.Redirect(c.Cookies("redirect", "/"))
}

func (h *Handler) ProviderLogout(c *fiber.Ctx) error {
	if err := sothic.Logout(c); err != nil {
		slog.Error("auth", "logout", err)
	}

	return c.SendStatus(http.StatusOK)
}
