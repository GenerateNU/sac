package mocks

import (
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/integrations/email"
)

type ResendMockClient struct{}

func NewResendMockClient() email.EmailClientInterface {
	return &ResendMockClient{}
}

func (c *ResendMockClient) SendPasswordResetEmail(name, email, token string) *errors.Error {
	return nil
}

func (c *ResendMockClient) SendEmailVerification(email, code string) *errors.Error {
	return nil
}

func (c *ResendMockClient) SendWelcomeEmail(name, email string) *errors.Error {
	return nil
}

func (c *ResendMockClient) SendPasswordChangedEmail(name, email string) *errors.Error {
	return nil
}
