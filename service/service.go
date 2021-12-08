package service

import (
	"email-service/clients"
	"email-service/handlers"
	"email-service/repository"
)

type ScheduleEmailService struct {
	sportsIOClient *clients.SportsIO
	ergastClient   *clients.Ergast
	repository     repository.Repository
	emailHandler   handlers.EmailHandler
}

func NewEmailService(
	SportsIOClient *clients.SportsIO,
	ErgastClient *clients.Ergast,
	Repository repository.Repository,
	EmailHandler handlers.EmailHandler,
) *ScheduleEmailService {
	return &ScheduleEmailService{
		SportsIOClient,
		ErgastClient,
		Repository,
		EmailHandler,
	}
}

func (es ScheduleEmailService) Run() {

}
