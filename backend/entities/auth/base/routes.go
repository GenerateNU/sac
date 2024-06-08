package base

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/oauth/goog"
	"github.com/GenerateNU/sac/backend/integrations/oauth/msft"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	m "github.com/garrettladley/mattress"
	"github.com/markbates/goth"
)

type Params struct {
	sessionSecret  *m.Secret[string]
	store          fiber.Storage
	msftSettings   config.MicrosoftOAuthSettings
	googSettings   config.GoogleOAuthSettings
	applicationURL string
	router         fiber.Router
	db             *gorm.DB
	emailer        email.Emailer
	validate       *validator.Validate
}

func NewParams(sessionSecret *m.Secret[string], store fiber.Storage, msftSettings config.MicrosoftOAuthSettings, googSettings config.GoogleOAuthSettings, applicationURL string, router fiber.Router, db *gorm.DB, emailer email.Emailer, validate *validator.Validate) Params {
	return Params{
		sessionSecret:  sessionSecret,
		store:          store,
		msftSettings:   msftSettings,
		googSettings:   googSettings,
		applicationURL: applicationURL,
		router:         router,
		db:             db,
		emailer:        emailer,
		validate:       validate,
	}
}

func Auth(params Params) {
	goth.UseProviders(
		msft.New(params.msftSettings.Key, params.msftSettings.Secret, fmt.Sprintf("%s/api/v1/auth/callback/microsoftonline", params.applicationURL), params.msftSettings.Tenant),
		goog.New(params.googSettings.Key, params.googSettings.Secret, fmt.Sprintf("%s/api/v1/auth/callback/google", params.applicationURL)),
	)

	handler := Handler{db: params.db, emailer: params.emailer, validate: params.validate}

	params.router.Route("/auth", func(r fiber.Router) {
		r.Get("/login/:provider", handler.Provider)
		r.Get("/callback/:provider", handler.ProviderCallback)
		r.Get("/logout", handler.ProviderLogout)
	})
}
