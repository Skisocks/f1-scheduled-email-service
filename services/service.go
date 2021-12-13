package services

import (
	"email-service/clients"
	"email-service/handlers"
	"email-service/repositories"
)

type ScheduleEmailService struct {
	sportsIOClient clients.CurrentEventGetter
	ergastClient   clients.StandingsGetter
	repository     repositories.UserEmailGetter
	emailHandler   handlers.EmailSender
}

func NewEmailService(
	SportsIOClient clients.CurrentEventGetter,
	ErgastClient clients.StandingsGetter,
	Repository repositories.UserEmailGetter,
	EmailHandler handlers.EmailSender,
) *ScheduleEmailService {
	return &ScheduleEmailService{
		sportsIOClient: SportsIOClient,
		ergastClient:   ErgastClient,
		repository:     Repository,
		emailHandler:   EmailHandler,
	}
}

func (ses ScheduleEmailService) Run() {
	currentEvent := ses.sportsIOClient.GetEvent()

	// If there is an event today send email
	if currentEvent != nil {
		ses.emailHandler.SendEmail(currentEvent, ses.repository.GetUserEmails())
	}
}
