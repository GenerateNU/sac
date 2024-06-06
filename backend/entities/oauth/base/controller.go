package base

import (
	"errors"
	"net/http"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/GenerateNU/sac/backend/locals"

	"github.com/gofiber/fiber/v2"
)

type OAuthController struct {
	oauth.Service
}

func NewOAuthController(Service oauth.Service) *OAuthController {
	return &OAuthController{Service}
}

func (oc *OAuthController) Authorize(c *fiber.Ctx) error {
	// Extract the resource type from the query params:
	resourceType := models.OAuthResource(c.Query("type"))
	if resourceType == "" {
		return errors.New("resource type is required")
	}

	// Extract the user making the call:
	userID, err := locals.UserIDFrom(c)
	if err != nil {
		return err
	}

	// Call the respective authorize method:
	authUrl, err := oc.Service.Authorize(c.Context(), *userID, resourceType)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(*authUrl)
}

func (oc *OAuthController) AuthorizationToken(c *fiber.Ctx) error {
	// Extract the resource type from the query params:
	resourceType := models.OAuthResource(c.Query("type"))
	if resourceType == "" {
		return errors.New("resource type is required")
	}

	// Parse the body of the request:
	var tokenBody OAuthTokenRequestBody
	if err := c.BodyParser(&tokenBody); err != nil {
		return err
	}

	// Extract the user making the call:
	userID, err := locals.UserIDFrom(c)
	if err != nil {
		return err
	}

	// Call the respective token method:
	err = oc.Service.AuthorizationToken(c.Context(), *userID, tokenBody.Code, tokenBody.State, resourceType)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

func (oc *OAuthController) Revoke(c *fiber.Ctx) error {
	// Extract the resource type from the query params:
	resourceType := models.OAuthResource(c.Query("type"))
	if resourceType == "" {
		return errors.New("resource type is required")
	}

	// Extract the user making the call:
	userID, err := locals.UserIDFrom(c)
	if err != nil {
		return err
	}

	err = oc.Service.Revoke(c.Context(), *userID, resourceType)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON("Successfully revoked token")
}
