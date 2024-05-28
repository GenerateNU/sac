package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type AuthSettings struct {
	AccessKey  *m.Secret[string]
	RefreshKey *m.Secret[string]
}

type intermediateAuthSettings struct {
	AccessKey  string `env:"ACCESS_KEY"`
	RefreshKey string `env:"REFRESH_KEY"`
}

func (i *intermediateAuthSettings) into() (*AuthSettings, error) {
	accessKey, err := m.NewSecret(i.AccessKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from access key: %s", err.Error())
	}

	refreshKey, err := m.NewSecret(i.RefreshKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from refresh key: %s", err.Error())
	}

	return &AuthSettings{
		AccessKey:  accessKey,
		RefreshKey: refreshKey,
	}, nil
}
