package helpers

import "github.com/GenerateNU/sac/backend/src/types"

func NewMockDependencies() *types.Dependencies {
	return &types.Dependencies{
		Search: NewPineconeMockClient(),
		AI:     NewOpenAIMockClient(),
		Email:  NewResendMockClient(),
		File:   NewAWSMockClient(),
		Redis: NewRedisMockClient(),
	}
}
