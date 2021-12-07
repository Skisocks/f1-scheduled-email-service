package service

import (
	"email-service/clients"
	"email-service/repository"
)

type ScheduleEmailService struct {
	SportsIOClient *clients.SportsIO
	ErgastClient   *clients.Ergast
	Repository     repository.Repository
}

func NewEmailService(
	SportsIOClient *clients.SportsIO,
	ErgastClient *clients.Ergast,
	Repository repository.Repository, ) *ScheduleEmailService {
	return &ScheduleEmailService{
		SportsIOClient,
		ErgastClient,
		Repository}
}

func (es ScheduleEmailService) Run() {

}
