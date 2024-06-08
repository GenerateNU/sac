package auth

import (
	"net/http"
	"slices"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/GenerateNU/sac/backend/integrations/oauth/sacic"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"

	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareHandler) IsSuper(c *fiber.Ctx) bool {
	user, err := locals.UserFrom(c)
	if err != nil {
		return false
	}
	if user == nil {
		return false
	}
	return user.Role == models.Super
}

func (m *AuthMiddlewareHandler) Authenticate(c *fiber.Ctx) error {
	// get the provider name from the request
	providerName, err := sacic.GetProviderName(c)
	if err != nil {
		return c.Redirect("/login", http.StatusTemporaryRedirect)
	}

	// get the provider
	provider, err := oauth.GetProvider(providerName)
	if err != nil {
		return c.Redirect("/login", http.StatusTemporaryRedirect)
	}

	// get the session value
	value, err := sacic.GetFromSession(providerName, c)
	if err != nil {
		return c.Redirect("/login", http.StatusTemporaryRedirect)
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return c.Redirect("/login", http.StatusTemporaryRedirect)
	}

	sacicUser, err := provider.FetchUser(sess)
	if err != nil {
		return c.Redirect("/login", http.StatusTemporaryRedirect)
	}

	user, err := users.GetUserByEmail(c.Context(), m.db, sacicUser.Email)
	if err != nil {
		return err
	}

	locals.SetUser(c, user)

	// user is authenticated, proceed to the next handler
	return c.Next()
}

func (m *AuthMiddlewareHandler) Authorize(requiredPermissions ...auth.Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authErr := m.Authenticate(c)
		if authErr != nil {
			return utilities.Unauthorized()
		}

		user, err := locals.UserFrom(c)
		if err != nil {
			return err
		}

		if user == nil {
			return utilities.Unauthorized()
		}

		if user.Role == models.Super {
			return c.Next()
		}

		userPermissions := auth.GetPermissions(user.Role)

		for _, requiredPermission := range requiredPermissions {
			if !slices.Contains(userPermissions, requiredPermission) {
				return utilities.Forbidden()
			}
		}

		return c.Next()
	}
}
