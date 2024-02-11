package config

import (
	"errors"

	m "github.com/garrettladley/mattress"
)

type AuthSettings struct {
	AccessKey          *m.Secret[string]
	RefreshKey         *m.Secret[string]
	AccessTokenExpiry  uint
	RefreshTokenExpiry uint
}

type intermediateAuthSettings struct {
	AccessKey          string `yaml:"accesskey"`
	RefreshKey         string `yaml:"refreshkey"`
	AccessTokenExpiry  uint   `yaml:"accesstokenexpiry"`
	RefreshTokenExpiry uint   `yaml:"refreshtokenexpiry"`
}

func (int *intermediateAuthSettings) into() (*AuthSettings, error) {
	accessToken, err := m.NewSecret(int.AccessKey)
	if err != nil {
		return nil, errors.New("failed to create secret from access key")
	}

	refreshToken, err := m.NewSecret(int.RefreshKey)
	if err != nil {
		return nil, errors.New("failed to create secret from refresh key")
	}

	return &AuthSettings{
		AccessKey:          accessToken,
		RefreshKey:         refreshToken,
		AccessTokenExpiry:  int.AccessTokenExpiry,
		RefreshTokenExpiry: int.RefreshTokenExpiry,
	}, nil
}
