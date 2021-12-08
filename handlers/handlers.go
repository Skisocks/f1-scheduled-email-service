package handlers

import (
	"email-service/config"
	"fmt"
	"gopkg.in/gomail.v2"
)

type EmailSender interface {
	SendEmail(UserEmails []string) error
}

type emailHandler struct {
	config *config.EmailHandler
}

func NewEmailHandler(cfg *config.EmailHandler) *emailHandler {
	return &emailHandler{
		config: cfg,
	}
}

func (eh *emailHandler) SendEmail(UserEmails []string) error {
	m := gomail.NewMessage()

	// Set message headers
	m.SetAddressHeader("From", eh.config.SenderAddress, "F1 Info")
	m.SetHeader("To", UserEmails...)
	m.SetHeader("Subject", "Test subject") // Todo: Make subject
	m.SetBody("text/plain", "Test body")   // Todo: Make body

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
