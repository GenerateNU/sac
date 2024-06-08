package config

import (
	m "github.com/garrettladley/mattress"
)

type OAuthSettings struct {
	Key    *m.Secret[string]
	Secret *m.Secret[string]
}

type intermediateOAuthSettings struct {
	Key    string `env:"KEY"`
	Secret string `env:"SECRET"`
}

func (i *intermediateOAuthSettings) into() (*OAuthSettings, error) {
	secretKey, err := m.NewSecret(i.Key)
	if err != nil {
		return nil, err
	}

	secretSecret, err := m.NewSecret(i.Secret)
	if err != nil {
		return nil, err
	}

	return &OAuthSettings{
		Key:    secretKey,
		Secret: secretSecret,
	}, nil
}

type OAuthProviderSettings struct {
	Google          *GoogleOAuthSettings
	OutlookSettings *MicrosoftOAuthSettings
}
