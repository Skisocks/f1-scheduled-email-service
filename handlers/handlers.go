package handlers

import (
	"email-service/config"
	"email-service/models"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

type EmailSender interface {
	SendEmail(currentEvent *models.CurrentEvent, userEmails []string)
}

type emailHandler struct {
	logger *zap.Logger
	config *config.EmailHandler
}

func NewEmailHandler(logger *zap.Logger, cfg *config.EmailHandler) *emailHandler {
	return &emailHandler{
		logger: logger,
		config: cfg,
	}
}

func (eh *emailHandler) SendEmail(currentEvent *models.CurrentEvent, userEmails []string) {
	m := gomail.NewMessage()

	// Create email contents
	subject := fmt.Sprintf(
		"The %s session of the %s today starts at %s",
		currentEvent.Type,
		currentEvent.Name,
		currentEvent.Datetime.Format("15:04:05 MST"),
	)
	body := fmt.Sprintf("")

	// Set message headers
	m.SetAddressHeader("From", eh.config.SenderAddress, eh.config.EmailName)
	m.SetHeader("To", userEmails...)
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
		eh.logger.Panic(fmt.Sprintf("failed to send email: %s", err))
		return
	}
	eh.logger.Info("Email sent")
}
