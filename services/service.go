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
		SportsIOClient,
		ErgastClient,
		Repository,
		EmailHandler,
	}
}

func (ses ScheduleEmailService) Run() {
	eventsResponse, _ := ses.sportsIOClient.GetEventsResponse()
	todaysEvent := getEvent(eventsResponse)

	if todaysEvent != nil {
		emailSubject := fmt.Sprintf(
			"The %s session of the %s today starts at %s",
			todaysEvent.Type,
			todaysEvent.Name,
			todaysEvent.Datetime.Format("15:04:05 MST"),
		)
		emailBody := fmt.Sprintf("")

		// Send email to the current users
		err := ses.emailHandler.SendEmail(emailSubject, emailBody, ses.repository.GetUserEmails())
		if err != nil {
			return
		}
	}
}

// getEvent takes a models.EventsResponse it returns a models.CurrentEvent
// depending on whether there is a race or 1st qualifying session today.
// If getEvent == nil then there is no event today.
func getEvent(EventsResponse *models.EventsResponse) *models.CurrentEvent {
	var todaysEvent *models.CurrentEvent = nil
	for i := range EventsResponse.Events {
		// Check if there's a race or 1st qualifying today
		switch EventsResponse.Events[i].Type {
		case "Race", "1st Qualifying":
			{
				// Take the required variables from the EventResponse and add them to todaysEvent
				todaysEvent = &models.CurrentEvent{
					Name:     EventsResponse.Events[i].Competition.Name,
					Type:     EventsResponse.Events[i].Type,
					Datetime: EventsResponse.Events[i].Date,
				}
				return todaysEvent
			}
		}
	}
	return todaysEvent
}
