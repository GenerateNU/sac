package email

import (
	"fmt"
	"os"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/resend/resend-go/v2"
)

type EmailServiceInterface interface {
	SendPasswordResetEmail(name, email, token string) *errors.Error
	SendEmailVerification(email, code string) *errors.Error
	SendWelcomeEmail(name, email string) *errors.Error
	SendPasswordChangedEmail(name, email string) *errors.Error
}

type EmailService struct {
	Client *resend.Client
}

func NewEmailClient(settings config.ResendSettings) *EmailService {
	return &EmailService{
		Client: resend.NewClient(settings.APIKey.Expose()),
	}
}

func (e *EmailService) SendPasswordResetEmail(name, email, token string) *errors.Error {
	template, err := getTemplateString("password_reset")
	if err != nil {
		return &errors.FailedToGetTemplate
	}

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{email},
		Subject: "Password Reset",
		Html:    fmt.Sprintf(*template, name, token),
	}

	_, err = e.Client.Emails.Send(params)
	if err != nil {
		return &errors.FailedToSendEmail
	}

	return nil
}

func (e *EmailService) SendEmailVerification(email, code string) *errors.Error {
	template, err := getTemplateString("email_verification")
	if err != nil {
		return &errors.FailedToGetTemplate
	}

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{email},
		Subject: "Email Verification",
		Html:    fmt.Sprintf(*template, code),
	}

	_, err = e.Client.Emails.Send(params)
	if err != nil {
		return &errors.FailedToSendEmail
	}

	return nil
}

func (e *EmailService) SendWelcomeEmail(name, email string) *errors.Error {
	template, err := getTemplateString("welcome")
	if err != nil {
		return &errors.FailedToGetTemplate
	}

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{email},
		Subject: "Welcome to Resend",
		Html:    fmt.Sprintf(*template, name),
	}

	_, err = e.Client.Emails.Send(params)
	if err != nil {
		return &errors.FailedToSendEmail
	}

	return nil
}

func (e *EmailService) SendPasswordChangedEmail(name, email string) *errors.Error {
	template, err := getTemplateString("password_change_complete")
	if err != nil {
		return &errors.FailedToGetTemplate
	}

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{email},
		Subject: "Password Changed",
		Html:    fmt.Sprintf(*template, name),
	}

	_, err = e.Client.Emails.Send(params)
	if err != nil {
		return &errors.FailedToSendEmail
	}

	return nil
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
