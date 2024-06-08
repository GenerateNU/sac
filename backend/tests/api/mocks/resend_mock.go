package mocks

import (
	"github.com/GenerateNU/sac/backend/integrations/email"
)

type ResendMockClient struct{}

func NewResendMockClient() email.Emailer {
	return &ResendMockClient{}
}

func (c *ResendMockClient) SendPasswordReset(name, email, token string) error {
	return nil
}

func (c *ResendMockClient) SendVerification(email, code string) error {
	return nil
}

func (c *ResendMockClient) SendWelcome(name, email string) error {
	return nil
}

func (c *ResendMockClient) SendPasswordChangedEmail(name, email string) error {
	return nil
}
