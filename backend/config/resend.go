package config

import m "github.com/garrettladley/mattress"

type ResendSettings struct {
	APIKey *m.Secret[string] `env:"API_KEY"`
}

type IntermediateResendSettings struct {
	APIKey string `env:"API_KEY"`
}

func (i *IntermediateResendSettings) Into() (*ResendSettings, error) {
	apiKey, err := m.NewSecret(i.APIKey)
	if err != nil {
		return nil, nil
	}

	return &ResendSettings{
		APIKey: apiKey,
	}, nil
}
