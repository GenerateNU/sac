package mocks

import (
	"context"

	"github.com/GenerateNU/sac/backend/integrations/email"
)

type ResendClient struct{}

func NewResendClient() email.Emailer {
	return &ResendClient{}
}

func (c *ResendClient) SendWelcome(ctx context.Context, email string, name string) error {
	return nil
}
