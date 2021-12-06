package main

import (
	"email-service/clients"
	"email-service/config"
	"email-service/service"
	"fmt"
	"time"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	// Get config struct
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get config: %s", err)
	}

	F1APIClient := clients.NewF1APIClient(&cfg.F1APIClient, time.Second * 5)

	EmailService := service.NewEmailService(F1APIClient)

	err := EmailService.Run()
	if err != nil {
		return err
	}

	return nil
}
