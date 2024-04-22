package helpers

import (
	"github.com/GenerateNU/sac/backend/email"
	"github.com/GenerateNU/sac/backend/file"
	"github.com/GenerateNU/sac/backend/search"
	"github.com/GenerateNU/sac/backend/types"
)

func NewMockDependencies() *types.Dependencies {
	return &types.Dependencies{
		Search: search.NewPineconeMockClient(),
		AI:     search.NewOpenAIMockClient(),
		Email:  email.NewResendMockClient(),
		File:   file.NewAWSMockClient(),
	}
}
