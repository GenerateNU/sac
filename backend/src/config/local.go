package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func readLocal(v *viper.Viper) (*Settings, error) {
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

	return settings, nil
}
