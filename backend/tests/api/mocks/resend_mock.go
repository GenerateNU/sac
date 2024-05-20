package mocks

import (
	"context"

	"github.com/GenerateNU/sac/backend/integrations/email"
)

type ResendMockClient struct{}

func NewResendMockClient() email.EmailClient {
	return &ResendMockClient{}
}

func (c *ResendMockClient) SendPasswordResetEmail(ctx context.Context, name string, email string, token string) error {
	return nil
}

func (c *ResendMockClient) SendEmailVerification(ctx context.Context, name string, code string) error {
	return nil
}

func (c *ResendMockClient) SendWelcomeEmail(ctx context.Context, name string, email string) error {
	return nil
}

func (c *ResendMockClient) SendPasswordChangedEmail(ctx context.Context, name string, email string) error {
	return nil
}
