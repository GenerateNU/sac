package config

import (
	"errors"

	m "github.com/garrettladley/mattress"
)

type SuperUserSettings struct {
	Password *m.Secret[string]
}

type intermediateSuperUserSettings struct {
	Password string `yaml:"password"`
}

func (int *intermediateSuperUserSettings) into() (*SuperUserSettings, error) {
	password, err := m.NewSecret(int.Password)
	if err != nil {
		return nil, errors.New("failed to create secret from password")
	}

	return &SuperUserSettings{
		Password: password,
	}, nil
}
