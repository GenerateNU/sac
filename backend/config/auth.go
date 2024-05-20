package config

import (
	"errors"

	m "github.com/garrettladley/mattress"
)

type AuthSettings struct {
	AccessKey  *m.Secret[string]
	RefreshKey *m.Secret[string]
	CSRFKey    *m.Secret[string]
}

type intermediateAuthSettings struct {
	AccessKey  string `yaml:"accesskey"`
	RefreshKey string `yaml:"refreshkey"`
	CSRFKey	   string `yaml:"csrfkey"`
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

	csrfToken, err := m.NewSecret(int.CSRFKey)
	if err != nil {
		return nil, errors.New("failed to create secret from csrf key")
	}

	return &AuthSettings{
		AccessKey:  accessToken,
		RefreshKey: refreshToken,
		CSRFKey:    csrfToken,
	}, nil
}
