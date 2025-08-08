package models

import (
	"fmt"
	"net/smtp"
)

type EmailService struct {
	SMTPHost string
	SMTPPort int
	Username string
	Password string  // Use a generate app password for the account with gmail
	From     string
}

func (e *EmailService) SendPasswordResetEmail(toEmail string, resetLink string) error {
	subject := "Reset Your Password"
	body := fmt.Sprintf("Click the link to reset your password:\n\n%s", resetLink)

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s",
        e.From, toEmail, subject, body)

    addr := fmt.Sprintf("%s:%d", e.SMTPHost, e.SMTPPort)
    auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPHost)

    return smtp.SendMail(addr, auth, e.From, []string{toEmail}, []byte(msg))
}