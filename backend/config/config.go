package config

import (
	"github.com/joho/godotenv"
)

func GetConfiguration(path string) (*Settings, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	redisSettings := make([]RedisSettings, len(int.Redis))
	for i, r := range int.Redis {
		redisInstance, err := r.into()
		if err != nil {
			return nil, err
		}

		redisSettings[i] = *redisInstance
	}

	superUserSettings, err := int.SuperUser.into()
	if err != nil {
		return nil, err
	}

	settings, err := intermediateSettings.into()
	return settings, err
}
