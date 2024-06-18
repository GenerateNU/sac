package config

import m "github.com/garrettladley/mattress"

type SessionSettings struct {
	Redis      RedisSettings
	PassPhrase *m.Secret[string]
}

type intermediateSessionSettings struct {
	PassPhrase string `env:"PASS_PHRASE"`
}

func (i *intermediateSessionSettings) into(redis RedisSettings) (*SessionSettings, error) {
	passPhrase, err := m.NewSecret(i.PassPhrase)
	if err != nil {
		return nil, err
	}

	return &SessionSettings{
		Redis:      redis,
		PassPhrase: passPhrase,
	}, nil
}
