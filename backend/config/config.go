package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Settings struct {
	Application     ApplicationSettings
	Database        DatabaseSettings
	SuperUser       SuperUserSettings
	Auth            AuthSettings
	AWS             AWSSettings
	Resend          ResendSettings
	Calendar        CalendarSettings
	GoogleSettings  OAuthSettings
	OutlookSettings OAuthSettings
}

type intermediateSettings struct {
	Application ApplicationSettings           `yaml:"application"`
	Database    intermediateDatabaseSettings  `yaml:"database"`
	SuperUser   intermediateSuperUserSettings `yaml:"superuser"`
	Auth        intermediateAuthSettings      `yaml:"authsecret"`
	Calendar    intermediateCalendarSettings  `yaml:"calendar"`
}

func (int *intermediateSettings) into() (*Settings, error) {
	databaseSettings, err := int.Database.into()
	if err != nil {
		return nil, err
	}

	var intermediateSettings IntermediateSettings
	if err := env.Parse(&intermediateSettings); err != nil {
		return nil, err
	}

	settings, err := intermediateSettings.Into()
	return settings, err
}
