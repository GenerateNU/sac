package store

import "github.com/GenerateNU/sac/backend/config"

type Stores struct {
	Session Storer
	Limiter Storer
}

func ConfigureStores(session config.RedisSettings, limiter config.RedisSettings) *Stores {
	return &Stores{
		Session: NewRedisClient(session),
		Limiter: NewRedisClient(limiter),
	}
}
