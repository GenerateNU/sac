package helpers

import (
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/GenerateNU/sac/backend/tests/api/mocks"
)

type StoresMock struct {
	Session store.Storer
	Limiter store.Storer
}

func NewStoresMock(session store.Storer, limiter store.Storer) *store.Stores {
	return &store.Stores{
		Session: session,
		Limiter: limiter,
	}
}

func ConfigureRedisMock() *store.Stores {
	return NewStoresMock(mocks.NewRedisClient(), mocks.NewRedisClient())
}
