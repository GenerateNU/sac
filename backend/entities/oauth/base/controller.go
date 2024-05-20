package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OAuthController struct {
	serviceSettings types.ServiceParams
}

func NewOAuthController(serviceSettings types.ServiceParams) *OAuthController {
	return &OAuthController{serviceSettings: serviceSettings}
}

func (oauth *OAuthController) Authorize(c *fiber.Ctx) error {
	// Extract the resource type from the query params:
	resourceType := c.Query("type")
	resource, err := oauth.parseCalendarType(resourceType)
	if err != nil {
		return err
	}

	userID, err := oauth.parseUserIDFromRequest(c)
	if err != nil {
		return err
	}

	// Call the respective authorize method:
	authUrl, err := resource.Authorize(oauth.serviceSettings.DB, oauth.serviceSettings.JWT, userID)
	if err != nil {
		return err
	}

	return c.Redirect(*authUrl, fiber.StatusFound)
}

func (cal *OAuthController) Token(c *fiber.Ctx) error {
	// Extract the resource type from the query params:
	resourceType := c.Query("type")
	calendar, err := cal.parseCalendarType(resourceType)
	if err != nil {
		return err
	}

	// Parse the body of the request:
	var tokenBody models.OAuthTokenRequestBody
	if err = c.BodyParser(&tokenBody); err != nil {
		return err
	}

	userID, err := cal.parseUserIDFromRequest(c)
	if err != nil {
		return err
	}

	// Call the respective token method:
	access_token, err := calendar.Token(tokenBody, c.Context(), userID, cal.serviceSettings.DB, &cal.serviceSettings.JWT)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(access_token)
}

func (cal *OAuthController) parseCalendarType(resourceType string) (*oauth.OAuthClient, error) {
	var resourceClient oauth.OAuthResourceClientInterface
	switch resourceType {
		case "google":
			resourceClient = cal.serviceSettings.Integrations.OAuth2.Google
		case "outlook":
			resourceClient = cal.serviceSettings.Integrations.OAuth2.Outlook
		default:
			return nil, errors.New("invalid resource type")
	}

	return &oauth.OAuthClient{
		ResourceClient: resourceClient,
	}, nil
}

func (cal *OAuthController) parseUserIDFromRequest(c *fiber.Ctx) (*uuid.UUID, error) {
	// Extract the user making the call:
	claims, err := auth.From(c)
	if err != nil {
		return nil, err
	}
	issuerIDAsUUID, err := utilities.ValidateID(claims.Issuer)
	if err != nil {
		return nil, err
	}

	return issuerIDAsUUID, nil
}




