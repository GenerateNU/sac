package helpers

import (
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/GenerateNU/sac/backend/tests/api/mocks"
)

type StoresMock struct {
	store.Limiter
	store.Storer
}

func NewStoresMock(session store.Storer, limiter store.Limiter) *store.Stores {
	return &store.Stores{
		Session: session,
		Limiter: limiter,
	}
}

func ConfigureRedisMock() *store.Stores {
	return NewStoresMock(mocks.NewRedisClient(), mocks.NewLimiter(mocks.NewRedisClient()))
}
