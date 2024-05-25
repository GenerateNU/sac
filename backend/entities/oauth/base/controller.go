package base

import (
	"errors"
	"net/http"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type OAuthController struct {
	OAuthService OAuthServiceInterface
}

func NewOAuthController(oauthService OAuthServiceInterface) *OAuthController {
	return &OAuthController{OAuthService: oauthService}
}

func (oc *OAuthController) Authorize(c *fiber.Ctx) error {
	// Extract the resource type from the query params:
	resourceType := models.OAuthResource(c.Query("type"))
	if resourceType == "" {
		return errors.New("resource type is required")
	}

	// Extract the user making the call:
	userID, err := auth.UserIDFrom(c)
	if err != nil {
		return err
	}

	// Call the respective authorize method:
	authUrl, err := oc.OAuthService.Authorize(*userID, resourceType)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(*authUrl)
}

func (oc *OAuthController) Token(c *fiber.Ctx) error {
	// Extract the resource type from the query params:
	resourceType := models.OAuthResource(c.Query("type"))
	if resourceType == "" {
		return errors.New("resource type is required")
	}

	// Parse the body of the request:
	var tokenBody models.OAuthTokenRequestBody
	if err := c.BodyParser(&tokenBody); err != nil {
		return err
	}

	// Extract the user making the call:
	userID := c.Locals("userID").(string)
	idAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return err
	}

	// Call the respective token method:
	err = oc.OAuthService.Token(*idAsUUID, tokenBody, resourceType)
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
	userID := c.Locals("userID").(string)
	idAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return err
	}

	err = oc.OAuthService.Revoke(*idAsUUID, resourceType)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON("Successfully revoked token")
}
