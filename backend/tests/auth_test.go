package tests

import (
	"testing"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/config"

	m "github.com/garrettladley/mattress"
	"github.com/golang-jwt/jwt"
	"github.com/huandu/go-assert"
)

func AuthSettings() (*config.AuthSettings, error) {
	accessKey, err := m.NewSecret("g(r|##*?>\\Qp}h37e+,T2")
	if err != nil {
		return nil, err
	}

	refreshKey, err := m.NewSecret("amk*2!gG}1i\"8D9RwJS$p")
	if err != nil {
		return nil, err
	}

	return &config.AuthSettings{
		AccessKey:          accessKey,
		AccessTokenExpiry:  60,
		RefreshKey:         refreshKey,
		RefreshTokenExpiry: 30,
	}, nil
}

func TestCreateTokenPairSuccess(t *testing.T) {
	assert := assert.New(t)

	id := "user123"
	role := "admin"

	authSettings, err := AuthSettings()
	assert.NilError(err)

	accessToken, refreshToken, authErr := auth.CreateTokenPair(id, role, *authSettings)

	assert.Assert(authErr == nil)

	assert.Assert(accessToken != nil)
	assert.Assert(refreshToken != nil)
}

func TestCreateTokenPairFailure(t *testing.T) {
	assert := assert.New(t)

	id := "user123"
	role := ""

	authSettings, err := AuthSettings()

	assert.NilError(err)

	accessToken, refreshToken, authErr := auth.CreateTokenPair(id, role, *authSettings)

	assert.Assert(authErr != nil)

	assert.Assert(accessToken == nil)
	assert.Assert(refreshToken == nil)
}

func TestCreateAccessTokenSuccess(t *testing.T) {
	assert := assert.New(t)

	id := "user123"
	role := "admin"

	authSettings, err := AuthSettings()

	assert.NilError(err)

	accessToken, authErr := auth.CreateAccessToken(id, role, authSettings.AccessTokenExpiry, authSettings.AccessKey)

	assert.Assert(authErr == nil)

	assert.Assert(accessToken != nil)
}

func TestCreateAccessTokenFailure(t *testing.T) {
	assert := assert.New(t)

	id := "user123"
	role := ""

	authSettings, err := AuthSettings()

	assert.NilError(err)

	accessToken, authErr := auth.CreateAccessToken(id, role, authSettings.AccessTokenExpiry, authSettings.AccessKey)

	assert.Assert(authErr != nil)

	assert.Assert(accessToken == nil)
}

func TestCreateRefreshTokenSuccess(t *testing.T) {
	assert := assert.New(t)

	id := "user123"

	authSettings, err := AuthSettings()

	assert.NilError(err)

	refreshToken, authErr := auth.CreateRefreshToken(id, authSettings.RefreshTokenExpiry, authSettings.RefreshKey)

	assert.Assert(authErr == nil)

	assert.Assert(refreshToken != nil)
}

func TestCreateRefreshTokenFailure(t *testing.T) {
	assert := assert.New(t)

	id := ""

	authSettings, err := AuthSettings()

	assert.NilError(err)

	refreshToken, authErr := auth.CreateRefreshToken(id, authSettings.RefreshTokenExpiry, authSettings.RefreshKey)

	assert.Assert(authErr != nil)

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

	key, err := m.NewSecret("secret")

	assert.NilError(err)

	signedToken, authErr := auth.SignToken(token, key)

	assert.NilError(authErr == nil)

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

	key, err := m.NewSecret("")

	assert.NilError(err)

	signedToken, authErr := auth.SignToken(token, key)

	assert.Assert(authErr != nil)

	assert.Assert(signedToken == nil)
}
