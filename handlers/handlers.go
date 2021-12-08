package handlers

import (
	"email-service/config"
	"email-service/repository"
	"fmt"
	"gopkg.in/gomail.v2"
	// "gopkg.in/gomail.v2"
)

type EmailHandler interface {
	SendEmail() error
}

type emailHandler struct {
	config *config.EmailHandler
	repo   repository.Repository
}

func NewEmailHandler(cfg *config.EmailHandler, repo repository.Repository) *emailHandler {
	return &emailHandler{
		config: cfg,
		repo:   repo,
	}
}

func (eh *emailHandler) SendEmail() error {
	m := gomail.NewMessage()

	// Set message headers
	m.SetAddressHeader("From", eh.config.SenderAddress, "F1 Info")
	m.SetHeader("To", eh.repo.GetUserEmails()...)
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
