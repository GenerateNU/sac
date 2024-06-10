package helpers

import (
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/GenerateNU/sac/backend/tests/api/mocks"
)

type StoresMock struct {
	Limiter store.Storer
}

func NewStoresMock(session store.Storer, limiter store.Storer) *store.Stores {
	return &store.Stores{
		Limiter: limiter,
	}
}

func ConfigureRedisMock() *store.Stores {
	return NewStoresMock(mocks.NewRedisClient(), mocks.NewRedisClient())
}
