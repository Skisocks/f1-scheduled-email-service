package service

import (
	"email-service/clients"
)

type ScheduleEmailService struct {
	F1APIClient	*clients.F1API
}

func NewEmailService(F1APIClient *clients.F1API) *ScheduleEmailService {
	return &ScheduleEmailService{F1APIClient}
}

func (es ScheduleEmailService) Run() error {

}