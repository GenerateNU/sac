package auth

import (
	"fmt"
	"os"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/resend/resend-go/v2"
)

type EmailService struct {
	Client *resend.Client
}

func NewEmailService() *EmailService {
	client := resend.NewClient(os.Getenv("SAC_RESEND_API_KEY"))
	return &EmailService{
		Client: client,
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
