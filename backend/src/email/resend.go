package email

import (
	"fmt"
	"os"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/resend/resend-go/v2"
)

type ResendClientInterface interface {
	SendPasswordResetEmail(name, email, token string) *errors.Error
	SendEmailVerification(email, code string) *errors.Error
	SendWelcomeEmail(name, email string) *errors.Error
	SendPasswordChangedEmail(name, email string) *errors.Error
}

type ResendClient struct {
	Client *resend.Client
	Dev    bool
}

func NewResendClient(settings config.ResendSettings, dev bool) *ResendClient {
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

func send(email, subject, html string, client *resend.Client) *errors.Error {
	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev", // TODO: get domain
		To:      []string{email},
		Subject: subject,
		Html:    html,
	}

	err := hystrix.Do("send-email", func() error {
		_, err := client.Emails.Send(params)
		return err
	}, nil)
	if err != nil {
		return &errors.FailedToSendEmail
	}

	return nil
}

func (r *ResendClient) SendPasswordResetEmail(name, email, token string) *errors.Error {
	if r.Dev {
		return nil
	}

	template, err := getTemplateString("password_reset")
	if err != nil {
		return &errors.FailedToGetTemplate
	}

	return send(email, "Password Reset", fmt.Sprintf(*template, name, token), r.Client)
}
func (r *ResendClient) SendEmailVerification(email, code string) *errors.Error {
	if r.Dev {
		return nil
	}

	template, err := getTemplateString("email_verification")
	if err != nil {
		return &errors.FailedToGetTemplate
	}

	return send(email, "Email Verification", fmt.Sprintf(*template, code), r.Client)
}

func (r *ResendClient) SendWelcomeEmail(name, email string) *errors.Error {
	if r.Dev {
		return nil
	}

	template, err := getTemplateString("welcome")
	if err != nil {
		return &errors.FailedToGetTemplate
	}

	return send(email, "Welcome to Resend", fmt.Sprintf(*template, name), r.Client)
}

func (r *ResendClient) SendPasswordChangedEmail(name, email string) *errors.Error {
	if r.Dev {
		return nil
	}

	template, err := getTemplateString("password_change_complete")
	if err != nil {
		return &errors.FailedToGetTemplate
	}

	return send(email, "Password Changed", fmt.Sprintf(*template, name), r.Client)
}

func getTemplateString(name string) (*string, error) {
	// TODO: use default file location
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	htmlBytes, err := os.ReadFile(fmt.Sprintf("%v/templates/emails/%s.html", cwd, name))
	if err != nil {
		return nil, err
	}

	htmlString := string(htmlBytes)

	return &htmlString, nil
}
