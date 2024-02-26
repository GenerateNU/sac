package auth

import (
	"fmt"
	"os"

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

func (e *EmailService) SendPasswordResetEmail(name, email, token string) error {
	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{"generatesac@gmail.com"},
		Subject: "Password Reset",
		Html:    fmt.Sprintf("<p>Hello %s, <br> Forgot your password? you can reset your password by clicking on this link: <a href='https://sac.resend.dev/reset-password/%s'>Reset Password</a> You can also copy and paste the following link into your browser: https://sac.resend.dev/reset-password/%s <br><br> If you did not request a password reset, please ignore this email.</p>", name, token, token),
	}

	_, err := e.Client.Emails.Send(params)
	if err != nil {
		return err
	}

	return nil
}
