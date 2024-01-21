package auth

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/types"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// CreateAccessToken creates a new access token for the user
func CreateAccessToken(id string, role string) (*string, error) {
	var settings config.Settings

	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    id,
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
		Role: role,
	})

	accessToken, err := SignToken(accessTokenClaims, settings.AuthSecret.AccessToken)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

// CreateRefreshToken creates a new refresh token for the user
func CreateRefreshToken() (*string, error) {
	var settings config.Settings

	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	})

	refreshToken, err := SignToken(refreshTokenClaims, settings.AuthSecret.RefreshToken)
	if err != nil {
		return nil, err
	}

	return refreshToken, nil
}

func SignToken(token *jwt.Token, secret string) (*string, error) {
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
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
func RefreshAccessToken(accessCookie, refreshCookie string) (*string, error) {
	var settings config.Settings
	// Parse the access token
	accessToken, err := ParseAccessToken(accessCookie)
	if err != nil {
		return nil, err
	}

	// Parse the refresh token
	refreshToken, err := ParseRefreshToken(refreshCookie)
	if err != nil {
		return nil, err
	}

	// Check if the access token is valid
	if _, ok := accessToken.Claims.(*types.CustomClaims); ok && accessToken.Valid {
		// Access token is already valid, no need for refresh return current access token
		tokenString, err := SignToken(accessToken, settings.AuthSecret.AccessToken)
		if err != nil {
			return nil, err
		}
		return tokenString, nil
	}

	// Check if the refresh token is valid
	if _, ok := refreshToken.Claims.(*jwt.StandardClaims); !ok || !refreshToken.Valid {
		// Refresh token is invalid, return unauthorized
		return nil, jwt.ErrInvalidKey
	}

	// Refresh the access token
	claims := refreshToken.Claims.(*types.CustomClaims)

	// Create Access Token with Custom Claims
	newAccessToken, err := CreateAccessToken(claims.Issuer, claims.Role)
	if err != nil {
		return nil, err
	}

	return newAccessToken, nil
}


// ParseAccessToken parses the access token
func ParseAccessToken(cookie string) (*jwt.Token, error) {
	var settings config.Settings

	return jwt.ParseWithClaims(cookie, &types.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.AuthSecret.AccessToken), nil
	})
}

// ParseRefreshToken parses the refresh token
func ParseRefreshToken(cookie string) (*jwt.Token, error) {
	var settings config.Settings

	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.AuthSecret.RefreshToken), nil
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
		return nil, jwt.ErrInvalidKey
	}

	return &claims.Role, nil
}
