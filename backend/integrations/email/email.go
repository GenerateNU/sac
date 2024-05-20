package email

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/templates/emails"

	"github.com/a-h/templ"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/resend/resend-go/v2"
)

type EmailClient interface {
	SendPasswordResetEmail(ctx context.Context, name string, email string, token string) error
	SendEmailVerification(ctx context.Context, email string, code string) error
	SendWelcomeEmail(ctx context.Context, name string, email string) error
	SendPasswordChangedEmail(ctx context.Context, name string, email string) error
}

type ResendClient struct {
	Client *resend.Client
	Dev    bool
}

func NewResendClient(settings config.ResendSettings, dev bool) EmailClient {
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

func send(ctx context.Context, fromEmail string, toEmail string, subject string, template templ.Component, client *resend.Client) error {
	buffer := new(bytes.Buffer)
	err := template.Render(context.Background(), buffer)
	if err != nil {
		return fmt.Errorf("failed to render email template: %w", err)
	}

	params := &resend.SendEmailRequest{
		From:    fromEmail,
		To:      []string{toEmail},
		Subject: subject,
		Html:    buffer.String(),
	}

	err = hystrix.DoC(ctx, "send-email", func(context.Context) error {
		_, err := client.Emails.Send(params)
		return err
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (r *ResendClient) SendPasswordResetEmail(ctx context.Context, name string, email string, token string) error {
	if r.Dev {
		return nil
	}

	return send(ctx, constants.DEFAULT_FROM_EMAIL, email, "Password Reset", emails.PasswordReset(name, fmt.Sprintf("https://hipponeu.com/reset/%s", token)), r.Client)
}

func (r *ResendClient) SendEmailVerification(ctx context.Context, email string, code string) error {
	if r.Dev {
		return nil
	}

	return send(ctx, constants.DEFAULT_FROM_EMAIL, email, "Email Verification", emails.Verification(code), r.Client)
}

func (r *ResendClient) SendWelcomeEmail(ctx context.Context, name string, email string) error {
	if r.Dev {
		return nil
	}

	return send(ctx, constants.ONBOARDING_EMAIL, email, "Welcome to Hippo", emails.Welcome(name), r.Client)
}

func (r *ResendClient) SendPasswordChangedEmail(ctx context.Context, name string, email string) error {
	if r.Dev {
		return nil
	}

	return send(ctx, constants.DEFAULT_FROM_EMAIL, email, "Password Changed", emails.PasswordChangeComplete(name), r.Client)
}
