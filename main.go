package main

import (
	"database/sql"
	"email-service/clients"
	"email-service/config"
	"email-service/handlers"
	"email-service/repositories"
	"email-service/services"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var TimeOut time.Duration = time.Second * 5

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

	// Initialise API Clients
	SportsIOClient := clients.NewSportsIOClient(&cfg.SportsIO, TimeOut)
	ErgastClient := clients.NewErgastClient(&cfg.Ergast, TimeOut)

	// Create new database connection pool
	DB, err := setupDB(&cfg.Repository)
	if err != nil {
		log.Fatal(err)
	}

	// Initialise repo
	Repository := repositories.NewRepository(DB)

	// Initialise email handler
	EmailHandler := handlers.NewEmailHandler(&cfg.EmailHandler)

	EmailService := services.NewEmailService(SportsIOClient, ErgastClient, Repository, EmailHandler)

	EmailService.Run()
	if err != nil {
		return err
	}

	return nil
}

func setupDB(cfg *config.Repository) (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Schema,
	)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	return db, err
}

func closeDB(DB *sql.DB) error {
	err := DB.Close()
	if err != nil {
		return err
	}
	return err
}
