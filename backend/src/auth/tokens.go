package auth

import (
	"fmt"
	"time"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/types"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CreateTokenPair(id string, role string, authSettings config.AuthSettings) (*string, *string, *errors.Error) {
	accessToken, catErr := CreateAccessToken(id, role, authSettings.AccessTokenExpiry, authSettings.AccessToken)
	if catErr != nil {
		return nil, nil, catErr
	}

	refreshToken, crtErr := CreateRefreshToken(id, authSettings.RefreshTokenExpiry, authSettings.RefreshToken)
	if crtErr != nil {
		return nil, nil, crtErr
	}

	return accessToken, refreshToken, nil
}

// CreateAccessToken creates a new access token for the user
func CreateAccessToken(id string, role string, accessExpiresAfter uint, accessTokenSecret string) (*string, *errors.Error) {
	if id == "" || role == "" {
		return nil, &errors.FailedToCreateAccessToken
	}

	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    id,
			ExpiresAt: time.Now().Add(time.Duration(accessExpiresAfter)).Unix(),
		},
		Role: role,
	})

	accessToken, err := SignToken(accessTokenClaims, accessTokenSecret)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

// CreateRefreshToken creates a new refresh token for the user
func CreateRefreshToken(id string, refreshExpiresAfter uint, refreshTokenSecret string) (*string, *errors.Error) {
	if id == "" {
		return nil, &errors.FailedToCreateRefreshToken
	}

	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		Issuer:    id,
		ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(refreshExpiresAfter)).Unix(),
	})

	refreshToken, err := SignToken(refreshTokenClaims, refreshTokenSecret)
	if err != nil {
		return nil, err
	}

	return refreshToken, nil
}

func SignToken(token *jwt.Token, secret string) (*string, *errors.Error) {
	if token == nil || secret == "" {
		fmt.Println(token)
		return nil, &errors.FailedToSignToken
	}

	tokenString, err := token.SignedString([]byte(secret))
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
func RefreshAccessToken(refreshCookie string, role string, accessExpiresAfter uint, accessTokenSecret string) (*string, *errors.Error) {
	// Parse the refresh token
	refreshToken, err := ParseRefreshToken(refreshCookie)
	if err != nil {
		return nil, &errors.FailedToParseRefreshToken
	}

	// Extract the claims from the refresh token
	claims, ok := refreshToken.Claims.(*jwt.StandardClaims)
	if !ok || !refreshToken.Valid {
		return nil, &errors.FailedToValidateRefreshToken
	}

	// Create a new access token
	accessToken, catErr := CreateAccessToken(claims.Issuer, role, accessExpiresAfter, accessTokenSecret)
	if catErr != nil {
		return nil, &errors.FailedToCreateAccessToken
	}

	return accessToken, nil
}

// ParseAccessToken parses the access token
func ParseAccessToken(cookie string) (*jwt.Token, error) {
	var settings config.Settings

	return jwt.ParseWithClaims(cookie, &types.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.Auth.AccessToken), nil
	})
}

// ParseRefreshToken parses the refresh token
func ParseRefreshToken(cookie string) (*jwt.Token, error) {
	var settings config.Settings

	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.Auth.RefreshToken), nil
	})
}

// GetRoleFromToken gets the role from the custom claims
func GetRoleFromToken(tokenString string) (*string, error) {
	token, err := ParseAccessToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*types.CustomClaims)
	if !ok || !token.Valid {
		return nil, &errors.FailedToValidateAccessToken
	}

	return &claims.Role, nil
}

// ExtractClaims extracts the claims from the token
func ExtractAccessClaims(tokenString string) (*types.CustomClaims, *errors.Error) {
	token, err := ParseAccessToken(tokenString)
	if err != nil {
		return nil, &errors.FailedToParseAccessToken
	}

	claims, ok := token.Claims.(*types.CustomClaims)
	if !ok || !token.Valid {
		return nil, &errors.FailedToValidateAccessToken
	}

	return claims, nil
}

// ExtractClaims extracts the claims from the token
func ExtractRefreshClaims(tokenString string) (*jwt.StandardClaims, *errors.Error) {
	token, err := ParseRefreshToken(tokenString)
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
