package helpers

import (
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/GenerateNU/sac/backend/tests/api/mocks"
)

type StoresMock struct {
	Limiter     store.LimiterInterface
	Blacklist   store.BlacklistInterface
	ActiveToken store.ActiveTokenInterface
}

func NewStoresMock(limiter store.LimiterInterface, blacklist store.BlacklistInterface, activeToken store.ActiveTokenInterface) *store.Stores {
	return &store.Stores{
		Limiter:     limiter,
		Blacklist:   blacklist,
		ActiveToken: activeToken,
	}
}

func ConfigureRedisMock() *store.Stores {
	limiter := mocks.NewLimiterMock(mocks.NewRedisMockClient())
	blacklist := mocks.NewBlacklistMock(mocks.NewRedisMockClient())
	activeToken := mocks.NewActiveTokenMock(mocks.NewRedisMockClient())

	stores := NewStoresMock(limiter, blacklist, activeToken)

	return stores
}
