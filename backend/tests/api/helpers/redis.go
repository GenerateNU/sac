package helpers

import (
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/GenerateNU/sac/backend/tests/api/mocks"
)

type StoresMock struct {
	store.Limiter
	store.Refresher
}

func NewStoresMock(limiter store.Limiter, refresher store.Refresher) *store.Stores {
	return &store.Stores{
		Limiter:   limiter,
		Refresher: refresher,
	}
}

func ConfigureRedisMock() *store.Stores {
	return NewStoresMock(mocks.NewLimiterMock(mocks.NewRedisMockClient()), mocks.NewRefresherMock(mocks.NewRedisMockClient()))
}
