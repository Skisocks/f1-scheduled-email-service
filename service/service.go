package service

import (
	"email-service/clients"
	"email-service/repository"
)

type ScheduleEmailService struct {
	F1APIClient *clients.SportsIO
	Repository repository.Repository
}

func NewEmailService(
	F1APIClient *clients.SportsIO,
	Repository repository.Repository,
	) *ScheduleEmailService {
	return &ScheduleEmailService{F1APIClient, Repository}
}

func (es ScheduleEmailService) Run() {

}
