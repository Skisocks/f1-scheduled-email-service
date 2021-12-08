package handlers

import (
	"email-service/config"
	"email-service/models"
	"email-service/repository"
)

type emailHandler interface {
	SendEmail()
}

type EmailHandler struct {
	config *config.EmailHandler
	repo  repository.Repository
}

func NewEmailHandler(cfg *config.EmailHandler, repo repository.Repository) *EmailHandler {
	return &EmailHandler{
		config: cfg,
		repo:  users,
	}
}

func (eh *EmailHandler) SendEmail() {

}
