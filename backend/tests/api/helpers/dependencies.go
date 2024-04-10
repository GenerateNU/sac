package helpers

import (
	"github.com/GenerateNU/sac/backend/src/email"
	"github.com/GenerateNU/sac/backend/src/file"
	"github.com/GenerateNU/sac/backend/src/search"
	"github.com/GenerateNU/sac/backend/src/types"
)

func NewMockDependencies() *types.Dependencies {
	return &types.Dependencies{
		Search: search.NewPineconeMockClient(),
		AI:     search.NewOpenAIMockClient(),
		Email:  email.NewResendMockClient(),
		File:   file.NewAWSMockClient(),
	}
}
