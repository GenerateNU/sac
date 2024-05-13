package email

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/templates/emails/password_change_complete"
	"github.com/GenerateNU/sac/backend/templates/emails/password_reset"
	"github.com/GenerateNU/sac/backend/templates/emails/verification"
	"github.com/GenerateNU/sac/backend/templates/emails/welcome"

	"github.com/a-h/templ"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/resend/resend-go/v2"
)

type EmailClientInterface interface {
	SendPasswordResetEmail(name string, email string, token string) error
	SendEmailVerification(email string, code string) error
	SendWelcomeEmail(name string, email string) error
	SendPasswordChangedEmail(name string, email string) error
}

type ResendClient struct {
	Client *resend.Client
	Dev    bool
}

func NewResendClient(settings config.ResendSettings, dev bool) EmailClientInterface {
	hystrix.ConfigureCommand("send-email", hystrix.CommandConfig{
		Timeout:               5000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	return &ResendClient{
		Client: resend.NewClient(settings.APIKey.Expose()),
		Dev:    dev,
	}
}

func send(email string, subject string, template templ.Component, client *resend.Client) error {
	buffer := new(bytes.Buffer)
	err := template.Render(context.Background(), buffer)
	if err != nil {
		return fmt.Errorf("failed to render email template: %w", err)
	}

	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("onboarding%s", constants.DOMAIN),
		To:      []string{email},
		Subject: subject,
		Html:    buffer.String(),
	}

	err = hystrix.Do("send-email", func() error {
		_, err := client.Emails.Send(params)
		return err
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (r *ResendClient) SendPasswordResetEmail(name string, email string, token string) error {
	if r.Dev {
		return nil
	}

	return send(email, "Password Reset", password_reset.Index(name, fmt.Sprintf("https://hipponeu.com/reset/%s", token)), r.Client)
}

func (r *ResendClient) SendEmailVerification(email string, code string) error {
	if r.Dev {
		return nil
	}

	return send(email, "Email Verification", verification.Index(code), r.Client)
}

func (r *ResendClient) SendWelcomeEmail(name string, email string) error {
	if r.Dev {
		return nil
	}

	return send(email, "Welcome to Resend", welcome.Index(name), r.Client)
}

func (r *ResendClient) SendPasswordChangedEmail(name string, email string) error {
	if r.Dev {
		return nil
	}

	return send(email, "Password Changed", password_change_complete.Index(name), r.Client)
}
