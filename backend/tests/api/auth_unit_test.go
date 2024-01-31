package tests

import (
	"testing"
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/golang-jwt/jwt"
)

func TestCreateTokenPairSuccess(t *testing.T) {
    id := "user123"
    role := "admin"

    accessToken, refreshToken, err := auth.CreateTokenPair(id, role)
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

	accessToken, refreshToken, err := auth.CreateTokenPair(id, role)
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

	accessToken, err := auth.CreateAccessToken(id, role)
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

	accessToken, err := auth.CreateAccessToken(id, role)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if accessToken != nil {
		t.Errorf("Expected token to be nil, got: %v", accessToken)
	}
}

func TestCreateRefreshTokenSuccess(t *testing.T) {
	id := "user123"

	refreshToken, err := auth.CreateRefreshToken(id)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if refreshToken == nil {
		t.Errorf("Expected token to be non-nil, got: %v", refreshToken)
	}
}

func TestCreateRefreshTokenFailure(t *testing.T) {
	id := ""

	refreshToken, err := auth.CreateRefreshToken(id)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if refreshToken != nil {
		t.Errorf("Expected token to be nil, got: %v", refreshToken)
	}
}

func TestSignTokenSuccess(t *testing.T) {	
	tokenString := &jwt.Token{
		Header: map[string]interface{}{
			"alg": "HS256",
			"typ": "JWT",
		},
		Claims: jwt.MapClaims{
			"sub": "user123",
			"exp": 1234567890,
			"iat": 1234567890,
			"iss": "sac",
		},
	}

	signedToken, err := auth.SignToken(tokenString, "secret")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if signedToken == nil {
		t.Errorf("Expected token to be non-nil, got: %v", signedToken)
	}
}

func TestSignTokenFailure(t *testing.T) {
	tokenString := &jwt.Token{
		Header: map[string]interface{}{
			"alg": "HS256",
			"typ": "JWT",
		},
		Claims: jwt.MapClaims{
			"sub": "user123",
			"exp": 1234567890,
			"iat": 1234567890,
			"iss": "sac",
		},
	}

	signedToken, err := auth.SignToken(tokenString, "")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if signedToken != nil {
		t.Errorf("Expected token to be nil, got: %v", signedToken)
	}
}

func TestCreateAndExpireCookieSuccess(t *testing.T) {
	cookie := auth.CreateCookie("name", "value", time.Now().Add(time.Hour))
	if cookie == nil {
		t.Errorf("Expected cookie to be non-nil, got: %v", cookie)
	}

	if cookie.Name != "name" {
		t.Errorf("Expected cookie name to be 'name', got: %v", cookie.Name)
	}

	if cookie.Value != "value" {
		t.Errorf("Expected cookie value to be 'value', got: %v", cookie.Value)
	}

	if cookie.Expires.IsZero() {
		t.Errorf("Expected cookie expiration to be non-zero, got: %v", cookie.Expires)
	}

	if !cookie.HTTPOnly {
		t.Errorf("Expected cookie HTTPOnly to be true, got: %v", cookie.HTTPOnly)
	}

	// clear cookie
	cookie.Expires = time.Now().Add(-time.Hour)
}
