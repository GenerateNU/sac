package mocks

import (
	"github.com/GenerateNU/sac/backend/integrations/email"
)

type ResendMockClient struct{}

func NewResendMockClient() email.EmailClientInterface {
	return &ResendMockClient{}
}

func (c *ResendMockClient) SendPasswordResetEmail(name, email, token string) error {
	return nil
}

func (c *ResendMockClient) SendEmailVerification(email, code string) error {
	return nil
}

func (c *ResendMockClient) SendWelcomeEmail(name, email string) error {
	return nil
}

func (c *ResendMockClient) SendPasswordChangedEmail(name, email string) error {
	return nil
}
