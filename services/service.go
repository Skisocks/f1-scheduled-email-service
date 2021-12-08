package services

import (
	"email-service/clients"
	"email-service/handlers"
	"email-service/models"
	"email-service/repositories"
	"fmt"
)

type ScheduleEmailService struct {
	sportsIOClient clients.CurrentEventGetter
	ergastClient clients.StandingsGetter
	repository   repositories.UserEmailGetter
	emailHandler handlers.EmailSender
}

func NewEmailService(
	SportsIOClient clients.CurrentEventGetter,
	ErgastClient clients.StandingsGetter,
	Repository repositories.UserEmailGetter,
	EmailHandler handlers.EmailSender,
) *ScheduleEmailService {
	return &ScheduleEmailService{
		SportsIOClient,
		ErgastClient,
		Repository,
		EmailHandler,
	}
}

func (ses ScheduleEmailService) Run() {
	eventsResponse, _ := ses.sportsIOClient.GetEventsResponse()
	if isEvent(eventsResponse) == true {
		emailSubject := fmt.Sprintf("The %s session of the %s today starts at %s", ) //
		emailBody := fmt.Sprintf("")


		// Send email to the current users
		err := ses.emailHandler.SendEmail(emailSubject, emailBody, ses.repository.GetUserEmails())
		if err != nil {
			return
		}


	}
}

// isEvent takes a models.EventsResponse and returns a boolean
// depending on whether there is a qualifying session or a race today.
func isEvent(EventsResponse *models.EventsResponse) bool {
	isEvent := false
	for i := range EventsResponse.Events {
		switch EventsResponse.Events[i].Type {
		case "Race", "1st Qualifying":
			isEvent = true
		}
	}
	return isEvent
}

func GetEvent(EventsResponse *models.EventsResponse) models.Event {
	for i := range EventsResponse.Events {
		}
	}
	return
}