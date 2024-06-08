package config

import m "github.com/garrettladley/mattress"

type GoogleOAuthSettings struct {
	Key    *m.Secret[string]
	Secret *m.Secret[string]
}

type intermediateGoogleOAuthSettings struct {
	Key    string `env:"KEY"`
	Secret string `env:"SECRET"`
}

func (i *intermediateGoogleOAuthSettings) into() (*GoogleOAuthSettings, error) {
	secretKey, err := m.NewSecret(i.Key)
	if err != nil {
		return nil, err
	}

	secretSecret, err := m.NewSecret(i.Secret)
	if err != nil {
		return nil, err
	}

	return &GoogleOAuthSettings{
		Key:    secretKey,
		Secret: secretSecret,
	}, nil
}
