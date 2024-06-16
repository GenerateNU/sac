package base

import (
	"context"
	"log/slog"
	"net/http"

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

		strUser, err := user.Marshal()
		if err != nil {
			return err
		}

		if err := sothic.StoreInSession("user", *strUser, c); err != nil {
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

		strUser, err := user.Marshal()
		if err != nil {
			return err
		}

		if err := sothic.StoreInSession("user", *strUser, c); err != nil {
			return err
		}

		return nil
	} else {
		return sothic.BeginAuthHandler(c)
	}
}

func (h *Handler) ProviderCallback(c *fiber.Ctx) error {
	gfUser, err := sothic.CompleteUserAuth(c)
	if err != nil {
		return err
	}

	user, err := FindOrCreateUser(context.TODO(), h.db, gfUser)
	if err != nil {
		return err
	}

	strUser, err := user.Marshal()
	if err != nil {
		return err
	}

	if err := sothic.StoreInSession("user", *strUser, c); err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

func (h *Handler) ProviderLogout(c *fiber.Ctx) error {
	if err := sothic.Logout(c); err != nil {
		slog.Error("auth", "logout", err)
	}

	return c.SendStatus(http.StatusOK)
}
