package auth

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/golang-jwt/jwt"
)

type JWTMockClient struct{}

func NewJWTMockClient() JWTClientInterface {
	return &JWTMockClient{}
}

func (c *JWTMockClient) GenerateTokenPair(accessClaims, refreshClaims Claims) (*Token, *errors.Error) {
	return &Token{}, nil
}

func (c *JWTMockClient) GenerateToken(claims Claims, tokenType JWTType) ([]byte, *errors.Error) {
	return []byte{}, nil
}

func (c *JWTMockClient) RefreshToken(token, refreshToken string, tokenType JWTType, newClaims jwt.MapClaims) ([]byte, *errors.Error) {
	return []byte{}, nil
}

func (c *JWTMockClient) ExtractClaims(tokenString string, tokenType JWTType) (jwt.MapClaims, *errors.Error) {
	return jwt.MapClaims{}, nil
}

func (c *JWTMockClient) ParseToken(tokenString string, tokenType JWTType) (*jwt.Token, *errors.Error) {
	return &jwt.Token{}, nil
}

func (c *JWTMockClient) IsTokenValid(tokenString string, tokenType JWTType) (bool, *errors.Error) {
	return true, nil
}
