package config

import m "github.com/garrettladley/mattress"

type SearchSettings struct {
	OpenAIApiKey *m.Secret[string]
}

type intermediateSearchSettings struct {
	OpenAIApiKey string `env:"OPENAI_API_KEY`
}

func (i *intermediateSearchSettings) into() (*SearchSettings, error) {
	openAiApiKey, err := m.NewSecret(i.OpenAIApiKey)
	if err != nil {
		return nil, err
	}

	return &SearchSettings{
		OpenAIApiKey: openAiApiKey,
	}, nil
}
