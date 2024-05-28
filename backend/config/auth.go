package config

import m "github.com/garrettladley/mattress"

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
		return nil, err
	}

	refreshKey, err := m.NewSecret(i.RefreshKey)
	if err != nil {
		return nil, err
	}

	return &AuthSettings{
		AccessKey:  accessKey,
		RefreshKey: refreshKey,
	}, nil
}
