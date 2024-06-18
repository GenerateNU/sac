package config

import m "github.com/garrettladley/mattress"

const (
	tenantID string = "a8eec281-aaa3-4dae-ac9b-9a398b9215e7"
)

type MicrosoftWebOAuthSettings struct {
	Key    *m.Secret[string]
	Secret *m.Secret[string]
	Tenant string
}

type MicrosoftMobileOAuthSettings struct {
	Key    *m.Secret[string]
	Tenant string
}

type intermediateMicrosoftWebOAuthSettings struct {
	Key    string `env:"KEY"`
	Secret string `env:"SECRET"`
}

func (i *intermediateMicrosoftWebOAuthSettings) into() (*MicrosoftWebOAuthSettings, error) {
	secretKey, err := m.NewSecret(i.Key)
	if err != nil {
		return nil, err
	}

	secretSecret, err := m.NewSecret(i.Secret)
	if err != nil {
		return nil, err
	}

	return &MicrosoftWebOAuthSettings{
		Key:    secretKey,
		Secret: secretSecret,
		Tenant: tenantID,
	}, nil
}

type intermediateMicrosoftMobileOAuthSettings struct {
	Key string `env:"KEY"`
}

func (i *intermediateMicrosoftMobileOAuthSettings) into() (*MicrosoftMobileOAuthSettings, error) {
	secretKey, err := m.NewSecret(i.Key)
	if err != nil {
		return nil, err
	}

	return &MicrosoftMobileOAuthSettings{
		Key:    secretKey,
		Tenant: tenantID,
	}, nil
}
