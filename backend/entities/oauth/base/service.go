package base

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	go_json "github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
)

type OAuthServiceInterface interface {
	Authorize(userID uuid.UUID, resource models.OAuthResource) (*string, error)
	Token(userID uuid.UUID, tokenBody OAuthTokenRequestBody, resource models.OAuthResource) error
	Revoke(userID uuid.UUID, resource models.OAuthResource) error
}

type OAuthService struct {
	ServiceParams types.ServiceParams
}

func New(serviceParams types.ServiceParams) OAuthServiceInterface {
	return &OAuthService{
		ServiceParams: serviceParams,
	}
}

func (e *OAuthService) Authorize(userID uuid.UUID, resource models.OAuthResource) (*string, error) {
	// Parse the resource client:
	client, err := e.parseResourceClient(resource)
	if err != nil {
		return nil, err
	}

	// Create a CSRF Token:
	csrfToken, err := auth.GenerateURLSafeToken(constants.CSRF_TOKEN_LENGTH)
	if err != nil {
		return nil, err
	}

	// Save the CSRF Token to the database:
	if err := CreateOAuthToken(e.ServiceParams.DB,
		models.UserOAuthTokens{
			UserID:       userID,
			CSRFToken:    *csrfToken,
			ResourceType: resource,
		}); err != nil {
		return nil, err
	}

	// Retrieve the authorize URL:
	authorizeURL := client.ResourceClient.AuthorizeURL(*csrfToken)

	return &authorizeURL, nil
}

func (e *OAuthService) Token(userID uuid.UUID, tokenBody OAuthTokenRequestBody, resource models.OAuthResource) error {
	// Parse the resource client:
	client, err := e.parseResourceClient(resource)
	if err != nil {
		return err
	}

	// Retrieve the CSRF Token from the database:
	oauthToken, err := GetOAuthToken(e.ServiceParams.DB, userID, resource)
	if err != nil {
		return err
	}

	// Validate the CSRF Token:
	if oauthToken.CSRFToken != tokenBody.State || oauthToken.CSRFToken == "" {
		return errors.New("invalid CSRF token")
	}

	// Exchange the code for an access token:
	token, err := e.tokenExchange(client, tokenBody.Code)
	if err != nil {
		return err
	}

	// Save the refresh token to the database:
	if err := SaveOAuthTokens(e.ServiceParams.DB, userID, *token, resource); err != nil {
		return err
	}

	return nil
}

func (e *OAuthService) Revoke(userID uuid.UUID, resource models.OAuthResource) error {
	// Parse the resource client:
	client, err := e.parseResourceClient(resource)
	if err != nil {
		return err
	}

	// Retrieve access token:
	oauthToken, err := GetOAuthToken(e.ServiceParams.DB, userID, resource)
	if err != nil {
		return err
	}

	if oauthToken.AccessToken == "" {
		return errors.New("user does not have an access token")
	}

	// Revoke the token on resource server:
	if err := client.ResourceClient.Revoke(oauthToken.AccessToken); err != nil {
		return err
	}

	// Remove refresh token from the database:
	if err := DeleteOAuthToken(e.ServiceParams.DB, userID, resource); err != nil {
		return err
	}

	return nil
}

func (e *OAuthService) parseResourceClient(resourceType models.OAuthResource) (*oauth.OAuthClient, error) {
	var resourceClient oauth.OAuthResourceClientInterface
	switch resourceType {
	case "google":
		resourceClient = &e.ServiceParams.Integrations.OAuth.Google
	case "outlook":
		resourceClient = &e.ServiceParams.Integrations.OAuth.Outlook
	default:
		return nil, errors.New("invalid resource type")
	}

	return &oauth.OAuthClient{
		ResourceClient: resourceClient,
	}, nil
}

func (e *OAuthService) tokenExchange(client *oauth.OAuthClient, code string) (*models.OAuthToken, error) {
	// Retrieve the client config
	config := client.ResourceClient.GetConfig()
	tokenUrl := fmt.Sprintf("%s/token", config.TokenURL)

	// Create the form data:
	formData := url.Values{
		"client_id":     {config.ClientID.Expose()},
		"client_secret": {config.ClientSecret.Expose()},
		"code":          {code},
		"redirect_uri":  {config.RedirectURI},
		"grant_type":    {"authorization_code"},
	}

	encodedData := formData.Encode()

	resp, err := utilities.Request(http.MethodPost, tokenUrl, []byte(encodedData), utilities.FormURLEncoded())
	if err != nil {
		return nil, err
	}

	if !utilities.IsOk(resp.StatusCode) {
		return nil, errors.New("failed to exchange code for token")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the data:
	var oauthToken models.OAuthToken
	if err = go_json.Unmarshal(body, &oauthToken); err != nil {
		return nil, err
	}

	if err := utilities.Validate(e.ServiceParams.Validate, oauthToken); err != nil {
		return nil, err
	}

	if oauthToken.AccessToken == "" {
		return nil, errors.New("access token is empty")
	}

	if oauthToken.RefreshToken == "" {
		return nil, errors.New("refresh token is empty")
	}

	return &oauthToken, nil
}
