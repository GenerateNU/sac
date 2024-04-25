package mocks

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/golang-jwt/jwt"
)

type JWTMockClient struct{}

func NewJWTMockClient() auth.JWTClientInterface {
	return &JWTMockClient{}
}

func (c *JWTMockClient) GenerateTokenPair(accessClaims, refreshClaims auth.Claims) (*auth.Token, *errors.Error) {
	return &auth.Token{}, nil
}

func (c *JWTMockClient) GenerateToken(claims auth.Claims, tokenType auth.JWTType) ([]byte, *errors.Error) {
	return []byte{}, nil
}

func (c *JWTMockClient) RefreshToken(token, refreshToken string, tokenType auth.JWTType, newClaims jwt.MapClaims) ([]byte, *errors.Error) {
	return []byte{}, nil
}

func (c *JWTMockClient) ExtractClaims(tokenString string, tokenType auth.JWTType) (jwt.MapClaims, *errors.Error) {
	return jwt.MapClaims{}, nil
}

func (c *JWTMockClient) ParseToken(tokenString string, tokenType auth.JWTType) (*jwt.Token, *errors.Error) {
	return &jwt.Token{}, nil
}

func (c *JWTMockClient) IsTokenValid(tokenString string, tokenType auth.JWTType) (bool, *errors.Error) {
	return true, nil
}
