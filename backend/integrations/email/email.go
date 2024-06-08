package email

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/templates/emails"
	"github.com/afex/hystrix-go/hystrix"

	"github.com/a-h/templ"

	"github.com/resend/resend-go/v2"
)

type Emailer interface {
	SendVerification(ctx context.Context, email string, code string) error
	SendWelcome(ctx context.Context, email string, name string) error
}

type ResendClient struct {
	Client *resend.Client
	Dev    bool
}

func NewResendClient(settings config.ResendSettings, dev bool) Emailer {
	return &ResendClient{
		Client: resend.NewClient(settings.APIKey.Expose()),
		Dev:    dev,
	}
}

func send(ctx context.Context, fromEmail string, toEmail string, subject string, template templ.Component, client *resend.Client) error {
	buffer := new(bytes.Buffer)
	err := template.Render(ctx, buffer)
	if err != nil {
		return fmt.Errorf("failed to render email template: %w", err)
	}

	params := &resend.SendEmailRequest{
		From:    fromEmail,
		To:      []string{toEmail},
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

func (r *ResendClient) SendVerification(ctx context.Context, email string, code string) error {
	if r.Dev {
		return nil
	}

	return send(ctx, constants.DEFAULT_FROM_EMAIL, email, "Email Verification", emails.Verification(code), r.Client)
}

func (r *ResendClient) SendWelcome(ctx context.Context, name string, email string) error {
	if r.Dev {
		return nil
	}

	return send(ctx, constants.ONBOARDING_EMAIL, email, "Welcome to Hippo", emails.Welcome(name), r.Client)
}
