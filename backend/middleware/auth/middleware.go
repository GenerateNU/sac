package auth

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/GenerateNU/sac/backend/permission"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"gorm.io/gorm"
)

type AuthMiddlewareService interface {
	ClubAuthorizeById(c *fiber.Ctx, extractor ExtractID) error
	UserAuthorizeById(c *fiber.Ctx) error
	EventAuthorizeById(c *fiber.Ctx, extractor ExtractID) error
	Authorize(requiredPermissions ...permission.Permission) fiber.Handler
}

type AuthMiddlewareHandler struct {
	db       *gorm.DB
	validate *validator.Validate
	provider goth.Provider
}

func New(db *gorm.DB, validate *validator.Validate, provider goth.Provider) AuthMiddlewareService {
	return &AuthMiddlewareHandler{
		db:       db,
		validate: validate,
		provider: provider,
	}
}

func (m *AuthMiddlewareHandler) Authenticate(c *fiber.Ctx) error {
	r, err := adaptor.ConvertRequest(c, false)
	if err != nil {
		return err
	}

	session, err := gothic.Store.Get(r, gothic.SessionName)
	if err != nil {
		return err
	}

	u := session.Values["user"]
	if u == nil {
		return utilities.Unauthorized()
	}

	user, ok := u.(models.User)
	if !ok {
		return utilities.Unauthorized()
	}

	locals.SetUser(c, &user)

	return nil
}
