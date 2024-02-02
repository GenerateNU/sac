package tests

import (
	"testing"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/config"

	"github.com/golang-jwt/jwt"
	"github.com/huandu/go-assert"
)

func AuthSettings() config.AuthSettings {
	return config.AuthSettings{
		AccessToken:        "g(r|##*?>\\Qp}h37e+,T2",
		AccessTokenExpiry:  60,
		RefreshToken:       "amk*2!gG}1i\"8D9RwJS$p",
		RefreshTokenExpiry: 30,
	}
}

func TestCreateTokenPairSuccess(t *testing.T) {
	assert := assert.New(t)

	id := "user123"
	role := "admin"

	accessToken, refreshToken, err := auth.CreateTokenPair(id, role, AuthSettings())

	assert.Assert(err == nil)

	assert.Assert(accessToken != nil)
	assert.Assert(refreshToken != nil)
}

func TestCreateTokenPairFailure(t *testing.T) {
	assert := assert.New(t)

	id := "user123"
	role := ""

	accessToken, refreshToken, err := auth.CreateTokenPair(id, role, AuthSettings())

	assert.Assert(err != nil)

	assert.Assert(accessToken == nil)
	assert.Assert(refreshToken == nil)
}

func TestCreateAccessTokenSuccess(t *testing.T) {
	assert := assert.New(t)

	id := "user123"
	role := "admin"

	authSettings := AuthSettings()

	accessToken, err := auth.CreateAccessToken(id, role, authSettings.AccessTokenExpiry, authSettings.AccessToken)

	assert.Assert(err == nil)

	assert.Assert(accessToken != nil)
}

func TestCreateAccessTokenFailure(t *testing.T) {
	assert := assert.New(t)

	id := "user123"
	role := ""

	authSettings := AuthSettings()

	accessToken, err := auth.CreateAccessToken(id, role, authSettings.AccessTokenExpiry, authSettings.AccessToken)

	assert.Assert(err != nil)

	assert.Assert(accessToken == nil)
}

func TestCreateRefreshTokenSuccess(t *testing.T) {
	assert := assert.New(t)

	id := "user123"

	authSettings := AuthSettings()

	refreshToken, err := auth.CreateRefreshToken(id, authSettings.RefreshTokenExpiry, authSettings.RefreshToken)

	assert.Assert(err == nil)

	assert.Assert(refreshToken != nil)
}

func TestCreateRefreshTokenFailure(t *testing.T) {
	assert := assert.New(t)

	id := ""

	authSettings := AuthSettings()

	refreshToken, err := auth.CreateRefreshToken(id, authSettings.RefreshTokenExpiry, authSettings.RefreshToken)

	assert.Assert(err != nil)

	assert.Assert(refreshToken == nil)
}

func TestSignTokenSuccess(t *testing.T) {
	assert := assert.New(t)

	token := jwt.New(jwt.SigningMethodHS256)

	assert.Assert(token != nil)

	token.Claims = jwt.MapClaims{
		"sub": "user123",
		"exp": 1234567890,
		"iat": 1234567890,
		"iss": "sac",
	}

	signedToken, err := auth.SignToken(token, "secret")

	assert.Assert(err == nil)

	assert.Assert(signedToken != nil)
}

func TestSignTokenFailure(t *testing.T) {
	assert := assert.New(t)

	token := jwt.New(jwt.SigningMethodHS256)

	assert.Assert(token != nil)

	token.Claims = jwt.MapClaims{
		"sub": "user123",
		"exp": 1234567890,
		"iat": 1234567890,
		"iss": "sac",
	}

	signedToken, err := auth.SignToken(token, "")

	assert.Assert(err != nil)

	assert.Assert(signedToken == nil)
}
