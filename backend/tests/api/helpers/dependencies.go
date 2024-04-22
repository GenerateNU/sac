package helpers

import (
	"github.com/GenerateNU/sac/backend/tests/api/mocks"
	"github.com/GenerateNU/sac/backend/types"
)

func NewMockDependencies() *types.Dependencies {
	return &types.Dependencies{
		Search: mocks.NewPineconeMockClient(),
		AI:     mocks.NewOpenAIMockClient(),
		Email:  mocks.NewResendMockClient(),
		File:   mocks.NewAWSMockClient(),
	}
}
