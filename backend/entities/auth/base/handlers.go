package base

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/GenerateNU/sac/backend/integrations/email"

	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"

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
	if gfUser, err := sothic.CompleteUserAuth(c, sothic.CompleteUserAuthOptions{ShouldLogout: false}); err == nil {
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
	gfUser, err := sothic.CompleteUserAuth(c, sothic.CompleteUserAuthOptions{ShouldLogout: false})
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

	redirectURL, err := sothic.GetFromSession("redirect", c)
	if err != nil {
		redirectURL = "/"
	}

	if err := sothic.SessionStore.Delete("redirect"); err != nil {
		slog.Error("auth", "provider_callback", err)
	}

	return c.Redirect(redirectURL)
}

func (h *Handler) ProviderLogout(c *fiber.Ctx) error {
	if err := sothic.Logout(c); err != nil {
		slog.Error("auth", "logout", err)
	}

	return c.SendString("logout")
}

// implements http.ResponseWriter
type FiberResponseWriter struct {
	c *fiber.Ctx
}

func (w *FiberResponseWriter) WriteHeader(statusCode int) {
	w.c.Status(statusCode)
}

func (w *FiberResponseWriter) Write(data []byte) (int, error) {
	return w.c.Write(data)
}

func (w *FiberResponseWriter) Header() http.Header {
	header := http.Header{}
	w.c.Response().Header.VisitAll(func(key, value []byte) {
		header.Add(string(key), string(value))
	})
	return header
}
