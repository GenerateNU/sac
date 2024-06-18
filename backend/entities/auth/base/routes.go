package base

import (
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Params struct {
	webAuthProvider    soth.Provider
	mobileAuthProvider soth.Provider
	providers          []soth.Provider
	applicationURL     string
	router             fiber.Router
	db                 *gorm.DB
}

func NewParams(webAuthProvider soth.Provider, mobileAuthProvider soth.Provider, applicationURL string, router fiber.Router, db *gorm.DB, emailer email.Emailer, validate *validator.Validate, providers ...soth.Provider,
) Params {
	return Params{
		webAuthProvider:    webAuthProvider,
		mobileAuthProvider: mobileAuthProvider,
		providers:          providers,
		applicationURL:     applicationURL,
		router:             router,
		db:                 db,
	}
}

func Auth(params Params) {
	soth.UseProviders(
		append(params.providers, params.webAuthProvider, params.mobileAuthProvider)...,
	)

	handler := Handler{db: params.db, webAuthProvider: params.webAuthProvider, mobileAuthProvider: params.mobileAuthProvider}

	params.router.Route("/auth", func(r fiber.Router) {
		r.Get("/login", handler.Login)
		r.Get("/:provider", handler.Provider)
		r.Get(":provider/callback", handler.ProviderCallback)
		r.Get("/logout", handler.ProviderLogout)
	})
}
