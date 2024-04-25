package helpers

import (
	"github.com/GenerateNU/sac/backend/integrations"
	"github.com/GenerateNU/sac/backend/tests/api/mocks"
)

func NewMockDependencies() *integrations.Integrations {
	return &integrations.Integrations{
		Search: mocks.NewPineconeMockClient(),
		AI:     mocks.NewOpenAIMockClient(),
		Email:  mocks.NewResendMockClient(),
		File:   mocks.NewAWSMockClient(),
	}
}
