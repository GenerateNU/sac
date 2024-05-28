package config

import m "github.com/garrettladley/mattress"

type IntermediateSuperUserSettings struct {
	Password string `env:"PASSWORD"`
}

type SuperUserSettings struct {
	Password *m.Secret[string] `env:"PASSWORD"`
}

func (i *IntermediateSuperUserSettings) Into() (*SuperUserSettings, error) {
	password, err := m.NewSecret(i.Password)
	if err != nil {
		return nil, nil
	}

	return &SuperUserSettings{
		Password: password,
	}, nil
}
