package service

import (
	"email-service/clients"
	"email-service/handlers"
	"email-service/models"
	"email-service/repository"
)

type ScheduleEmailService struct {
	sportsIOClient clients.CurrentEventGetter
	ergastClient   clients.StandingsGetter
	repository     repository.UserEmailGetter
	emailHandler   handlers.EmailSender
}

func NewEmailService(
	SportsIOClient clients.CurrentEventGetter,
	ErgastClient clients.StandingsGetter,
	Repository repository.UserEmailGetter,
	EmailHandler handlers.EmailSender,
) *ScheduleEmailService {
	return &ScheduleEmailService{
		SportsIOClient,
		ErgastClient,
		Repository,
		EmailHandler,
	}
}

func (es ScheduleEmailService) isEvent(EventsResponse *models.EventsResponse) bool {
	var isEvent bool
	for i := range EventsResponse.Events {
		switch EventsResponse.Events[i].Type {
		case "Race", "1st Qualifying":
			isEvent = true
		default:
			isEvent = false
		}
	}
	return isEvent
}

func (es ScheduleEmailService) Run() {

	eventsResponse, _ := es.sportsIOClient.GetEventsResponse()
	if es.isEvent(eventsResponse) == true {
	}
}
