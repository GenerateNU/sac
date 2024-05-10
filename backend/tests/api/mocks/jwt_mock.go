package mocks

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/golang-jwt/jwt"
)

type JWTMockClient struct{}

func NewJWTMockClient() auth.JWTClientInterface {
	return &JWTMockClient{}
}

func (c *JWTMockClient) GenerateTokenPair(accessClaims, refreshClaims auth.Claims) (*auth.Token, error) {
	return &auth.Token{}, nil
}

func (c *JWTMockClient) GenerateToken(claims auth.Claims, tokenType auth.JWTType) ([]byte, error) {
	return []byte{}, nil
}

func (c *JWTMockClient) RefreshToken(token, refreshToken string, tokenType auth.JWTType, newClaims jwt.MapClaims) ([]byte, error) {
	return []byte{}, nil
}

func (c *JWTMockClient) ExtractClaims(tokenString string, tokenType auth.JWTType) (jwt.MapClaims, error) {
	return jwt.MapClaims{}, nil
}

func (c *JWTMockClient) ParseToken(tokenString string, tokenType auth.JWTType) (*jwt.Token, error) {
	return &jwt.Token{}, nil
}

func (c *JWTMockClient) IsTokenValid(tokenString string, tokenType auth.JWTType) (bool, error) {
	return true, nil
}
