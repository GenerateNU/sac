package base

import (
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Params struct {
	providers      []soth.Provider
	applicationURL string
	router         fiber.Router
	db             *gorm.DB
	emailer        email.Emailer
	validate       *validator.Validate
}

func NewParams(applicationURL string, router fiber.Router, db *gorm.DB, emailer email.Emailer, validate *validator.Validate, providerFns ...func() soth.Provider) Params {
	providers := make([]soth.Provider, len(providerFns))
	for i, providerFn := range providerFns {
		providers[i] = providerFn()
	}

	return Params{
		providers:      providers,
		applicationURL: applicationURL,
		router:         router,
		db:             db,
		emailer:        emailer,
		validate:       validate,
	}
}

func Auth(params Params) {
	soth.UseProviders(
		params.providers...,
	)

	handler := Handler{db: params.db, emailer: params.emailer, validate: params.validate}

	params.router.Route("/auth", func(r fiber.Router) {
		r.Get("/:provider", handler.Provider)
		r.Get(":provider/callback/", handler.ProviderCallback)
		r.Get("/logout", handler.ProviderLogout)
	})
}
