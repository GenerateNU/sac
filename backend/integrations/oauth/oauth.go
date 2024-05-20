package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

type OAuthClient struct {
	ResourceClient 	OAuthResourceClientInterface
}

type OAuthClientInterface interface {
	Authorize(db *gorm.DB, jwt auth.JWTClientInterface, userId *uuid.UUID) string
	Token(requestBody models.OAuthTokenRequestBody, ctx *fasthttp.RequestCtx, userId *uuid.UUID, db *gorm.DB, jwtClient *auth.JWTClientInterface) error
}

// Return the authorization URL: 
func (client OAuthClient) Authorize(db *gorm.DB, jwt auth.JWTClientInterface, userId *uuid.UUID) (*string, error) {
	csrfToken, err := CreateCalendarToken(db, jwt, userId)
	if err != nil {
		return nil, err
	}
	url := client.ResourceClient.AuthorizeURL(*csrfToken)

	return &url, nil
}

func (client OAuthClient) Token(
	requestBody models.OAuthTokenRequestBody,
	ctx *fasthttp.RequestCtx,
	userId *uuid.UUID,
	db *gorm.DB,
	jwtClient *auth.JWTClientInterface,
) (*string, error) {
	// Validate the CSRFToken:
	if err := validateCSRFToken(db, userId, requestBody.State, *jwtClient); err != nil {
		return nil, err
	}

	// Throw error if the authorization code does not exist:
	if requestBody.Code == "" {
		return nil, errors.New("authorization code is required")
	}

	// Exchange the code for a token:
	token, err := client.TokenExchange(requestBody.Code)
	if err != nil {
		return nil, err
	}

	// Save the refresh token:
	if err = SaveRefreshToken(db, userId, token.RefreshToken); err != nil {
		return nil, err
	}

	return &token.AccessToken, nil
}

/**
 * Exchange an authorization code for an OAuth token.
**/
func (client OAuthClient) TokenExchange(code string) (*OAuthToken, error) {
	config := client.ResourceClient.GetConfig()
	tokenUrl := fmt.Sprintf("%s/token", config.BaseURL)

	formData := url.Values{
		"client_id": {config.ClientID.Expose()},
		"client_secret": {config.ClientSecret.Expose()},
		"code": {code},
		"redirect_uri": {config.RedirectURI},
		"grant_type": {"authorization_code"},
	}
	encodedData := formData.Encode();

	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(encodedData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Initiate the request:
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the data:
	var oauthToken OAuthToken
	if err = json.Unmarshal(body, &oauthToken); err != nil {
		return nil, err
	}

	return &oauthToken, nil
}

/**
 * Validate a CSRFToken.
**/
func validateCSRFToken(db *gorm.DB, userId *uuid.UUID, csrfToken string, jwtClient auth.JWTClientInterface) error {
	if csrfToken == "" {
		return errors.New("state is required")
	}
	
	// Get the CSRF token from the database:
	storedCSRFToken, err := GetCSRFToken(db, userId)
	if err != nil {
		return err
	}

	// Validate the token:
	_, err = jwtClient.IsTokenValid(csrfToken, "csrf")
	if err != nil {
		return err
	}

	// Compare the CSRF token with the stored CSRF token:
	if csrfToken != *storedCSRFToken {
		return errors.New("state does not match CSRF token")
	}

	return nil
}