package config

import m "github.com/garrettladley/mattress"

const (
	tenantID string = "a8eec281-aaa3-4dae-ac9b-9a398b9215e7"
)

type MicrosoftOAuthSettings struct {
	Key    *m.Secret[string]
	Secret *m.Secret[string]
	Tenant string
}

type intermediateMicrosoftOAuthSetting struct {
	Key string `env:"KEY"`
}

func (i *intermediateMicrosoftOAuthSetting) into() (*MicrosoftOAuthSettings, error) {
	secretKey, err := m.NewSecret(i.Key)
	if err != nil {
		return nil, err
	}

	return &MicrosoftOAuthSettings{
		Key:    secretKey,
		Tenant: tenantID,
	}, nil
}
