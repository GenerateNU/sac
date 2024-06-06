package oauth

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/go-playground/validator/v10"
	go_json "github.com/goccy/go-json"
	"gorm.io/gorm"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
)

type Service interface {
	Authorize(ctx context.Context, userID uuid.UUID, resource models.OAuthResource) (*string, error)
	AuthorizationToken(ctx context.Context, userID uuid.UUID, code string, state string, resource models.OAuthResource) error
	RefreshToken(ctx context.Context, userID uuid.UUID, resource models.OAuthResource) error
	Revoke(ctx context.Context, userID uuid.UUID, resource models.OAuthResource) error
}

type Handler struct {
	db        *gorm.DB
	stores    *store.Stores
	providers OAuthProviders
	validate  *validator.Validate
}

func NewHandler(db *gorm.DB, stores *store.Stores, providers OAuthProviders, validate *validator.Validate) Service {
	return &Handler{
		db: db,
	}
}

func (oa *Handler) Authorize(ctx context.Context, userID uuid.UUID, resource models.OAuthResource) (*string, error) {
	client, err := oa.getResourceClient(resource)
	if err != nil {
		return nil, err
	}

	csrfToken, err := auth.GenerateURLSafeToken(constants.CSRF_TOKEN_LENGTH)
	if err != nil {
		return nil, err
	}

	if err := CreateOAuthToken(oa.db,
		models.UserOAuthTokens{
			UserID:       userID,
			CSRFToken:    *csrfToken,
			ResourceType: resource,
		}); err != nil {
		return nil, err
	}

	authorizeURL := client.AuthorizeURL(*csrfToken)

	return &authorizeURL, nil
}

func (oa *Handler) AuthorizationToken(ctx context.Context, userID uuid.UUID, code string, state string, resource models.OAuthResource) error {
	client, err := oa.getResourceClient(resource)
	if err != nil {
		return err
	}

	oauthToken, err := GetOAuthToken(oa.db, userID, resource)
	if err != nil {
		return err
	}

	if oauthToken.CSRFToken != state || oauthToken.CSRFToken == "" {
		return errors.New("invalid CSRF token")
	}

	formData := url.Values{
		"client_id":     {client.GetConfig().ClientID.Expose()},
		"client_secret": {client.GetConfig().ClientSecret.Expose()},
		"code":          {code},
		"redirect_uri":  {client.GetConfig().RedirectURI},
		"grant_type":    {"authorization_code"},
	}

	token, err := oa.performTokenExchange(client, formData)
	if err != nil {
		return err
	}

	return oa.saveOAuthTokens(userID, *token, resource)
}

func (oa *Handler) RefreshToken(ctx context.Context, userID uuid.UUID, resource models.OAuthResource) error {
	client, err := oa.getResourceClient(resource)
	if err != nil {
		return err
	}

	oauthToken, err := GetOAuthToken(oa.db, userID, resource)
	if err != nil {
		return err
	}

	if oauthToken.RefreshToken == "" {
		return errors.New("refresh token is empty")
	}

	formData := url.Values{
		"client_id":     {client.GetConfig().ClientID.Expose()},
		"client_secret": {client.GetConfig().ClientSecret.Expose()},
		"refresh_token": {oauthToken.RefreshToken},
		"grant_type":    {"refresh_token"},
	}

	newToken, err := oa.performTokenExchange(client, formData)
	if err != nil {
		return err
	}

	if err := oa.saveOAuthTokens(userID, *newToken, resource); err != nil {
		return err
	}

	return oa.stores.Refresher.Set(ctx, userID.String(), newToken.RefreshToken, constants.REFRESH_TOKEN_EXPIRY)
}

func (oa *Handler) Revoke(ctx context.Context, userID uuid.UUID, resource models.OAuthResource) error {
	client, err := oa.getResourceClient(resource)
	if err != nil {
		return err
	}

	oauthToken, err := GetOAuthToken(oa.db, userID, resource)
	if err != nil {
		return err
	}

	if oauthToken.AccessToken == "" {
		return errors.New("user does not have an access token")
	}

	if err := client.Revoke(oauthToken.AccessToken); err != nil {
		return err
	}

	if err := DeleteOAuthToken(oa.db, userID, resource); err != nil {
		return err
	}

	return nil
}

func (oa *Handler) getResourceClient(resourceType models.OAuthResource) (*OAuthClient, error) {
	var resourceClient OAuthProvider
	switch resourceType {
	case "google":
		resourceClient = &oa.providers.Google
	case "outlook":
		resourceClient = &oa.providers.Outlook
	default:
		return nil, errors.New("invalid resource type")
	}

	return &OAuthClient{
		resourceClient,
	}, nil
}

func (oa *Handler) performTokenExchange(client *OAuthClient, formData url.Values) (*models.OAuthToken, error) {
	config := client.GetConfig()
	tokenUrl := fmt.Sprintf("%s/token", config.TokenURL)

	encodedData := formData.Encode()

	resp, err := utilities.Request(http.MethodPost, tokenUrl, []byte(encodedData), utilities.FormURLEncoded())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !utilities.IsOk(resp.StatusCode) {
		return nil, errors.New("failed to exchange token")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var oauthToken models.OAuthToken
	if err = go_json.Unmarshal(body, &oauthToken); err != nil {
		return nil, err
	}

	if err := utilities.Validate(oa.validate, oauthToken); err != nil {
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

func (oa *Handler) saveOAuthTokens(userID uuid.UUID, token models.OAuthToken, resource models.OAuthResource) error {
	return SaveOAuthTokens(oa.db, userID, token, resource)
}
