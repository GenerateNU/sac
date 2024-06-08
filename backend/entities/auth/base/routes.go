package base

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/GenerateNU/sac/backend/integrations/oauth/goog"
	"github.com/GenerateNU/sac/backend/integrations/oauth/msft"
	"github.com/GenerateNU/sac/backend/integrations/oauth/sacic"
	"github.com/GenerateNU/sac/backend/middleware/utility"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	m "github.com/garrettladley/mattress"
)

type AuthParams struct {
	sessionSecret     *m.Secret[string]
	store             fiber.Storage
	msftSettings      config.MicrosoftOAuthSettings
	googSettings      config.GoogleOAuthSettings
	applicationURL    string
	router            fiber.Router
	db                *gorm.DB
	utilityMiddleware utility.Service
	emailer           email.Emailer
	validate          *validator.Validate
}

func Auth(params AuthParams) {
	sacic.Init(params.sessionSecret, params.store)

	oauth.UseProviders(
		msft.New(params.msftSettings.Key, params.msftSettings.Secret, fmt.Sprintf("%s/auth/microsoftonline/callback", params.applicationURL), params.msftSettings.Tenant),
		goog.New(params.googSettings.Key, params.googSettings.Secret, fmt.Sprintf("%s/auth/google/callback", params.applicationURL)),
	)

	handler := Handler{db: params.db, emailer: params.emailer, validate: params.validate}

	params.router.Route("/auth", func(r fiber.Router) {
		r.Get("/:provider", handler.Provider)
		r.Get("/:provider/callback", handler.ProviderCallback)
		r.Get("/:provider/logout", handler.ProviderLogout)
		r.Get("/send-code", params.utilityMiddleware.Limiter(constants.RATE_LIMIT_MAX_REQUESTS, constants.RATE_LIMIT_DURATION), handler.SendCode)
		r.Get("/verify-email", handler.VerifyEmail)
	})
}
