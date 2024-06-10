package email

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/templates/emails"

	"github.com/a-h/templ"

	"github.com/resend/resend-go/v2"
)

type Emailer interface {
	SendWelcome(ctx context.Context, email string, name string) error
}

type ResendClient struct {
	client *resend.Client
	dev    bool
}

func NewResendClient(settings config.ResendSettings, dev bool) Emailer {
	return &ResendClient{
		client: resend.NewClient(settings.APIKey.Expose()),
		dev:    dev,
	}
}

func (r *ResendClient) SendWelcome(ctx context.Context, name string, email string) error {
	if r.dev {
		return nil
	}

	return r.send(ctx, constants.ONBOARDING_EMAIL, email, "Welcome to Hippo", emails.Welcome(name))
}

func (r *ResendClient) send(ctx context.Context, fromEmail string, toEmail string, subject string, template templ.Component) error {
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

	if _, err := r.client.Emails.Send(params); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
