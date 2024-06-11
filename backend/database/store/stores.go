package store

import "github.com/GenerateNU/sac/backend/config"

type Stores struct {
	Limiter Storer
}

func ConfigureStores(limiter config.RedisSettings) *Stores {
	return &Stores{
		Limiter: NewRedisClient(limiter),
	}
}
