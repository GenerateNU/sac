package store

type Stores struct {
	Limiter Storer
}

func ConfigureStores(limiter RedisSettings) *Stores {
	return &Stores{
		Limiter: NewRedisClient(limiter),
	}
}
