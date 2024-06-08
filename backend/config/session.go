package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type SessionSettings struct {
	Secret *m.Secret[string]
}

type intermediateSessionSettings struct {
	Secret string `env:"SECRET"`
}

func (i *intermediateSessionSettings) into() (*SessionSettings, error) {
	secret, err := m.NewSecret(i.Secret)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from secret: %s", err.Error())
	}

	return &SessionSettings{
		Secret: secret,
	}, nil
}
