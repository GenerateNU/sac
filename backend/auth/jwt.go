package auth

import (
	"fmt"
	"time"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/constants"

	"github.com/GenerateNU/sac/backend/utilities"
	m "github.com/garrettladley/mattress"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
}

type JWTType string

const (
	AccessToken  JWTType = "access"
	RefreshToken JWTType = "refresh"
)

type Token struct {
	AccessToken  []byte
	RefreshToken []byte
}

type Claims struct {
	StandardClaims *jwt.StandardClaims
	CustomClaims   *jwt.MapClaims
}

type JWTClientInterface interface {
	GenerateTokenPair(accessClaims, refreshClaims Claims) (*Token, error)
	GenerateToken(claims Claims, tokenType JWTType) ([]byte, error)
	RefreshToken(token, refreshToken string, tokenType JWTType, newClaims jwt.MapClaims) ([]byte, error)
	ExtractClaims(tokenString string, tokenType JWTType) (jwt.MapClaims, error)
	ParseToken(tokenString string, tokenType JWTType) (*jwt.Token, error)
	IsTokenValid(tokenString string, tokenType JWTType) (bool, error)
}

type JWTClient struct {
	RefreshExp    time.Duration
	AccessExp     time.Duration
	RefreshKey    *m.Secret[string]
	AccessKey     *m.Secret[string]
	SigningMethod jwt.SigningMethod
}

func NewJWTClient(authSettings config.AuthSettings, signingMethod jwt.SigningMethod) JWTClientInterface {
	return &JWTClient{
		RefreshExp:    constants.REFRESH_TOKEN_EXPIRY,
		AccessExp:     constants.ACCESS_TOKEN_EXPIRY,
		RefreshKey:    authSettings.RefreshKey,
		AccessKey:     authSettings.AccessKey,
		SigningMethod: signingMethod,
	}
}

func (j *JWTClient) GenerateTokenPair(accessClaims, refreshClaims Claims) (*Token, error) {
	accessToken, err := j.GenerateToken(accessClaims, AccessToken)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.GenerateToken(refreshClaims, RefreshToken)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// GenerateToken generates a token with the claims passed in.
// It returns the token if successful, otherwise it returns an error.
func (j *JWTClient) GenerateToken(claims Claims, tokenType JWTType) ([]byte, error) {
	// create a new map to store the combined claims
	combinedClaims := make(jwt.MapClaims)

	// copy the standard claims to the combined claims if they are present
	if claims.CustomClaims != nil {
		copyCustomClaims(&combinedClaims, *claims.CustomClaims)
	}

	// copy the custom claims to the combined claims if they are present
	if claims.StandardClaims != nil {
		copyStandardClaims(&combinedClaims, *claims.StandardClaims)
	}

	// get the secret key for the token type
	secretKey, err := j.getSecretKey(tokenType)
	if err != nil {
		return nil, err
	}

	// get the expiry for the token type
	exp := j.getExpiry(tokenType)

	// set the expiry of the token if its a value in the client, otherwise use whatever was passed in
	if exp != 0 {
		combinedClaims["exp"] = time.Now().Add(exp).Unix()
	}

	// create a new token with the combined claims
	token := jwt.NewWithClaims(j.SigningMethod, combinedClaims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %w", err)
	}

	return []byte(signedToken), nil
}

// ParseToken parses the token string and returns the token if successful, otherwise it returns an error.
// It uses the secret key for the token type to parse the token.
func (j *JWTClient) ParseToken(tokenString string, tokenType JWTType) (*jwt.Token, error) {
	secretKey, err := j.getSecretKey(tokenType)
	if err != nil {
		return nil, err
	}

	token, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if parseErr != nil {
		return nil, fmt.Errorf("failed to parse token: %w", parseErr)
	}

	return token, nil
}

// ExtractClaims extracts the claims from the token.
// It returns the claims if successful, otherwise it returns an error.
func (j *JWTClient) ExtractClaims(tokenString string, tokenType JWTType) (jwt.MapClaims, error) {
	token, err := j.ParseToken(tokenString, tokenType)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract claims from token. got: %T", token.Claims)
	}

	return claims, nil
}

// IsTokenValid checks if the token is valid.
// It returns true if the token is valid, otherwise it returns false.
func (j *JWTClient) IsTokenValid(tokenString string, tokenType JWTType) (bool, error) {
	token, err := j.ParseToken(tokenString, tokenType)
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

// RefreshToken generates a new access token using the refresh token.
// It checks if the refresh token is valid and extracts the claims from the access token.
// It then updates the issued at and expires at claims and gives the new claims priority over the old claims.
// It returns the new access token if successful, otherwise it returns an error.
func (j *JWTClient) RefreshToken(token, refreshToken string, tokenType JWTType, newClaims jwt.MapClaims) ([]byte, error) {
	ok, err := j.IsTokenValid(refreshToken, RefreshToken)
	if err != nil || !ok {
		return nil, err
	}

	claims, err := j.ExtractClaims(token, tokenType)
	if err != nil {
		return nil, err
	}

	claims = updateIssuedAt(claims)
	claims = updateExpiresAt(claims, j.getExpiry(tokenType))
	claims = emendClaims(claims, newClaims)

	newToken, err := j.GenerateToken(Claims{CustomClaims: &claims}, tokenType)
	if err != nil {
		return nil, err
	}

	return newToken, nil
}

// getSecretKey returns the secret key for the token type.
// If the token type is not present in the client, it returns an error.
func (j *JWTClient) getSecretKey(tokenType JWTType) (string, error) {
	switch tokenType {
	case AccessToken:
		return j.AccessKey.Expose(), nil
	case RefreshToken:
		return j.RefreshKey.Expose(), nil
	}

	return "", utilities.BadRequest(fmt.Errorf("invalid token type: %s", tokenType))
}

// getExpiry returns the expiry time for the token type.
// If the token type is not present in the client, it returns 0.
func (j *JWTClient) getExpiry(tokenType JWTType) time.Duration {
	switch tokenType {
	case AccessToken:
		return j.AccessExp
	case RefreshToken:
		return j.RefreshExp
	}

	return 0
}

// updateIssuedAt updates the issued at claim of the token.
// If the issued at claim is not present, it won't be added.
func updateIssuedAt(claims jwt.MapClaims) jwt.MapClaims {
	if _, ok := claims["iat"]; ok {
		claims["iat"] = time.Now().Unix()
	}
	return claims
}

// updateExpiresAt updates the expires at claim of the token.
// If the expires at claim is not present, it won't be added.
func updateExpiresAt(claims jwt.MapClaims, exp time.Duration) jwt.MapClaims {
	if _, ok := claims["exp"]; ok {
		claims["exp"] = time.Now().Add(exp).Unix()
	}
	return claims
}

// emendClaims updates the claims of the token with the new claims.
// If the claim is not present in the original claims, it won't be added.
func emendClaims(originalClaims, newClaims jwt.MapClaims) jwt.MapClaims {
	for key, value := range newClaims {
		if _, ok := originalClaims[key]; ok {
			originalClaims[key] = value
		}
	}
	return originalClaims
}

// copyStandardClaims copies the standard claims from a jwt.StandardClaims instance to a jwt.MapClaims instance.
// It is a utility function used to copy standard claims to the token claims.
func copyStandardClaims(claims *jwt.MapClaims, standardClaims jwt.StandardClaims) {
	claimMapping := map[string]interface{}{
		"exp": standardClaims.ExpiresAt,
		"iss": standardClaims.Issuer,
		"aud": standardClaims.Audience,
		"iat": standardClaims.IssuedAt,
		"nbf": standardClaims.NotBefore,
		"sub": standardClaims.Subject,
		"jti": standardClaims.Id,
	}

	for key, value := range claimMapping {
		if intValue, ok := value.(int64); ok && intValue != 0 {
			(*claims)[key] = value
		} else if strValue, ok := value.(string); ok && strValue != "" {
			(*claims)[key] = value
		}
	}
}

// copyCustomClaims copies the custom claims from a map[string]interface{} instance to a jwt.MapClaims instance.
// It is a utility function used to copy custom claims to the token claims.
func copyCustomClaims(claims *jwt.MapClaims, customClaims map[string]interface{}) {
	for key, value := range customClaims {
		(*claims)[key] = value
	}
}

func SetResponseTokens(c *fiber.Ctx, tokens *Token) {
	c.Set("Authorization", fmt.Sprintf("Bearer %s", tokens.AccessToken))
}

func ExpireResponseTokens(c *fiber.Ctx) {
	c.Set("Authorization", "")
}
