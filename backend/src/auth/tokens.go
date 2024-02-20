package auth

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"

	m "github.com/garrettladley/mattress"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CreateTokenPair(id string, role string, authSettings config.AuthSettings) (*string, *string, *errors.Error) {
	accessToken, catErr := CreateAccessToken(id, role, authSettings.AccessTokenExpiry, authSettings.AccessKey)
	if catErr != nil {
		return nil, nil, catErr
	}

	refreshToken, crtErr := CreateRefreshToken(id, authSettings.RefreshTokenExpiry, authSettings.RefreshKey)
	if crtErr != nil {
		return nil, nil, crtErr
	}

	return accessToken, refreshToken, nil
}

// CreateAccessToken creates a new access token for the user
func CreateAccessToken(id string, role string, accessExpiresAfter uint, accessToken *m.Secret[string]) (*string, *errors.Error) {
	if id == "" || role == "" {
		return nil, &errors.FailedToCreateAccessToken
	}

	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    id,
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(accessExpiresAfter)).Unix(),
		},
		Role: role,
	})

	returnedAccessToken, err := SignToken(accessTokenClaims, accessToken)
	if err != nil {
		return nil, err
	}

	return returnedAccessToken, nil
}

// CreateRefreshToken creates a new refresh token for the user
func CreateRefreshToken(id string, refreshExpiresAfter uint, refreshKey *m.Secret[string]) (*string, *errors.Error) {
	if id == "" {
		return nil, &errors.FailedToCreateRefreshToken
	}

	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		Issuer:    id,
		ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(refreshExpiresAfter)).Unix(),
	})

	returnedRefreshToken, err := SignToken(refreshTokenClaims, refreshKey)
	if err != nil {
		return nil, err
	}

	return returnedRefreshToken, nil
}

func SignToken(token *jwt.Token, key *m.Secret[string]) (*string, *errors.Error) {
	if token == nil || key.Expose() == "" {
		return nil, &errors.FailedToSignToken
	}

	tokenString, err := token.SignedString([]byte(key.Expose()))
	if err != nil {
		return nil, &errors.FailedToSignToken
	}
	return &tokenString, nil
}

// CreateCookie creates a new cookie
func CreateCookie(name string, value string, expires time.Time) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expires,
		HTTPOnly: true,
	}
}

// ExpireCookie expires a cookie
func ExpireCookie(name string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     name,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
}

// RefreshAccessToken refreshes the access token
func RefreshAccessToken(refreshCookie string, role string, refreshKey *m.Secret[string], accessExpiresAfter uint, accessKey *m.Secret[string]) (*string, *errors.Error) {
	// Parse the refresh token
	refreshToken, err := ParseRefreshToken(refreshCookie, refreshKey)
	if err != nil {
		return nil, &errors.FailedToParseRefreshToken
	}

	// Extract the claims from the refresh token
	claims, ok := refreshToken.Claims.(*jwt.StandardClaims)
	if !ok || !refreshToken.Valid {
		return nil, &errors.FailedToValidateRefreshToken
	}

	// Create a new access token
	accessToken, catErr := CreateAccessToken(claims.Issuer, role, accessExpiresAfter, accessKey)
	if catErr != nil {
		return nil, &errors.FailedToCreateAccessToken
	}

	return accessToken, nil
}

// ParseAccessToken parses the access token
func ParseAccessToken(cookie string, accessKey *m.Secret[string]) (*jwt.Token, error) {
	return jwt.ParseWithClaims(cookie, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessKey.Expose()), nil
	})
}

// ParseRefreshToken parses the refresh token
func ParseRefreshToken(cookie string, refreshKey *m.Secret[string]) (*jwt.Token, error) {
	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshKey.Expose()), nil
	})
}

// GetRoleFromToken gets the role from the custom claims
func GetRoleFromToken(tokenString string, accessKey *m.Secret[string]) (*string, error) {
	token, err := ParseAccessToken(tokenString, accessKey)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, &errors.FailedToValidateAccessToken
	}

	return &claims.Role, nil
}

// ExtractClaims extracts the claims from the token
func ExtractAccessClaims(tokenString string, accessKey *m.Secret[string]) (*CustomClaims, *errors.Error) {
	token, err := ParseAccessToken(tokenString, accessKey)
	if err != nil {
		return nil, &errors.FailedToParseAccessToken
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, &errors.FailedToValidateAccessToken
	}

	return claims, nil
}

// ExtractClaims extracts the claims from the token
func ExtractRefreshClaims(tokenString string, refreshKey *m.Secret[string]) (*jwt.StandardClaims, *errors.Error) {
	token, err := ParseRefreshToken(tokenString, refreshKey)
	if err != nil {
		return nil, &errors.FailedToParseRefreshToken
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, &errors.FailedToValidateRefreshToken
	}

	return claims, nil
}

func IsBlacklisted(token string) bool {
	return false
}
