package tests

import (
	"testing"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/golang-jwt/jwt"
)

func TestCreateTokenPairSuccess(t *testing.T) {
	id := "user123"
	role := "admin"

	accessToken, refreshToken, err := auth.CreateTokenPair(id, role, 60, 30)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if accessToken == nil || refreshToken == nil {
		t.Errorf("Expected both tokens to be non-nil, got: %v, %v", accessToken, refreshToken)
	}
}

func TestCreateTokenPairFailure(t *testing.T) {
	id := "user123"
	role := ""

	accessToken, refreshToken, err := auth.CreateTokenPair(id, role, 60, 30)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if accessToken != nil || refreshToken != nil {
		t.Errorf("Expected both tokens to be nil, got: %v, %v", accessToken, refreshToken)
	}
}

func TestCreateAccessTokenSuccess(t *testing.T) {
	id := "user123"
	role := "admin"

	accessToken, err := auth.CreateAccessToken(id, role, 60)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if accessToken == nil {
		t.Errorf("Expected token to be non-nil, got: %v", accessToken)
	}
}

func TestCreateAccessTokenFailure(t *testing.T) {
	id := "user123"
	role := ""

	accessToken, err := auth.CreateAccessToken(id, role, 60)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if accessToken != nil {
		t.Errorf("Expected token to be nil, got: %v", accessToken)
	}
}

func TestCreateRefreshTokenSuccess(t *testing.T) {
	id := "user123"

	refreshToken, err := auth.CreateRefreshToken(id, 30)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if refreshToken == nil {
		t.Errorf("Expected token to be non-nil, got: %v", refreshToken)
	}
}

func TestCreateRefreshTokenFailure(t *testing.T) {
	id := ""

	refreshToken, err := auth.CreateRefreshToken(id, 30)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if refreshToken != nil {
		t.Errorf("Expected token to be nil, got: %v", refreshToken)
	}
}

func TestSignTokenSuccess(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	if token == nil {
		t.Fatal("Failed to create JWT token")
	}

	token.Claims = jwt.MapClaims{
		"sub": "user123",
		"exp": 1234567890,
		"iat": 1234567890,
		"iss": "sac",
	}

	signedToken, err := auth.SignToken(token, "secret")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if signedToken == nil {
		t.Errorf("Expected token to be non-nil, got: %v", signedToken)
	}
}

func TestSignTokenFailure(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	if token == nil {
		t.Fatal("Failed to create JWT token")
	}

	token.Claims = jwt.MapClaims{
		"sub": "user123",
		"exp": 1234567890,
		"iat": 1234567890,
		"iss": "sac",
	}

	signedToken, err := auth.SignToken(token, "")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if signedToken != nil {
		t.Errorf("Expected token to be nil, got: %v", signedToken)
	}
}
