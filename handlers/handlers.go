package handlers

import (
	"email-service/config"
	"fmt"
	"gopkg.in/gomail.v2"
)

type EmailSender interface {
	SendEmail(subject string, body string, UserEmails []string) error
}

type emailHandler struct {
	config *config.EmailHandler
}

func NewEmailHandler(cfg *config.EmailHandler) *emailHandler {
	return &emailHandler{
		config: cfg,
	}
}

func (eh *emailHandler) SendEmail(subject string, body string, UserEmails []string) error {
	m := gomail.NewMessage()

	// Set message headers
	m.SetAddressHeader("From", eh.config.SenderAddress, eh.config.EmailName)
	m.SetHeader("To", UserEmails...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// SMTP settings
	d := gomail.NewDialer(
		eh.config.SMTPServer,
		eh.config.SMTPHost,
		eh.config.SenderAddress,
		eh.config.SenderPassword,
	)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return nil
}
