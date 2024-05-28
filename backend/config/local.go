package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func readLocal(v *viper.Viper, path string, useDevDotEnv bool) (*Settings, error) {
	var intermediateSettings intermediateSettings

	env := string(EnvironmentLocal)

	v.SetConfigName(env)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read %s configuration: %w", env, err)
	}

	if err := v.Unmarshal(&intermediateSettings); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	settings, err := intermediateSettings.into()
	if err != nil {
		return nil, fmt.Errorf("failed to convert intermediate settings into final settings: %w", err)
	}

	if useDevDotEnv {
		err = godotenv.Load(fmt.Sprintf("%s/.env.dev", path))
	} else {
		err = godotenv.Load(fmt.Sprintf("%s/.env.template", path))
	}

	if err != nil {
		return nil, fmt.Errorf("failed to load %s/.env.template: %w", path, err)
	}

	awsSettings, err := readAWSSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to read AWS settings: %w", err)
	}

	settings.AWS = *awsSettings

	resendSettings, err := readResendSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to read Resend settings: %w", err)
	}

	settings.Resend = *resendSettings

	googleSettings, err := readGoogleOAuthSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to read Google OAuth settings: %w", err)
	}

	settings.GoogleSettings = *googleSettings

	outlookSettings, err := readOutlookOAuthSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to read Outlook OAuth settings: %w", err)
	}

	settings.OutlookSettings = *outlookSettings

	return settings, nil
}
