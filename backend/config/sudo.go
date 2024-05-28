package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type SuperUserSettings struct {
	Password *m.Secret[string]
}
type intermediateSuperUserSettings struct {
	password string `env:"PASSWORD"`
}

func (i *intermediateSuperUserSettings) into() (*SuperUserSettings, error) {
	password, err := m.NewSecret(i.password)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from password: %s", err.Error())
	}

	return &SuperUserSettings{
		Password: password,
	}, nil
}
