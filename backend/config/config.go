package config

import (
	"reflect"

	"github.com/caarlos0/env/v11"
	m "github.com/garrettladley/mattress"
	"github.com/joho/godotenv"
)

func GetConfiguration(path string) (*Settings, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	var settingsSecret *m.Secret[string]
	settingsSecretType := reflect.TypeOf(settingsSecret)

	parsingOptions := env.Options{
		FuncMap: map[reflect.Type]env.ParserFunc{
			settingsSecretType: func(v string) (interface{}, error) {
				secret, err := m.NewSecret(v)
				if err != nil {
					return nil, nil
				}

				return secret, nil
			},
		},
	}

	var settings Settings
	if err := env.ParseWithOptions(&settings, parsingOptions); err != nil {
		return nil, err
	}

	return &settings, nil
}

// 530 -> 210 -> ??>
