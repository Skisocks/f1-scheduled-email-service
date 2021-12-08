package handlers

import (
	"email-service/config"
	"email-service/models"
)

type emailHandler interface {
	SendEmail()
}

type EmailHandler struct {
	config	*config.EmailHandler
	users	[]models.User
}

func NewEmailHandler(cfg *config.EmailHandler, users []models.User) *EmailHandler {
	return &EmailHandler{
		config: cfg,
		users:  users,
	}
}

func (eh *EmailHandler) SendEmail() {

}