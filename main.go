package main

import (
	"database/sql"
	"personal/f1/f1-scheduled-email-service/clients"
	"personal/f1/f1-scheduled-email-service/handlers"
	"personal/f1/f1-scheduled-email-service/services"

	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	// Setup logging
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)

	// Get configs
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Panic(fmt.Sprintf("failed to retrive config file: %s", err.Error()))
		return err
	}

	// Create new database connection pool
	DB, err := setupDB(&cfg.Repository)
	if err != nil {
		logger.Panic(fmt.Sprintf("failed to setup database: %s", err.Error()))
	}
	defer func(DB *sql.DB) {
		err := closeDB(DB)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to close database connection: %s", err.Error()))
		}
	}(DB)

	// Initialise API Clients
	SportsIOClient := clients.NewSportsIOClient(logger, &cfg.SportsIO)
	ErgastClient := clients.NewErgastClient(logger, &cfg.Ergast)

	// Initialise repo
	Repository := repositories.NewRepository(logger, DB)

	// Initialise email handler
	EmailHandler := handlers.NewEmailHandler(logger, &cfg.EmailHandler)

	EmailService := services.NewEmailService(SportsIOClient, ErgastClient, Repository, EmailHandler)

	EmailService.Run()
	if err != nil {
		return err
	}

	return nil
}

func setupDB(cfg *config.Repository) (db *sql.DB, err error) {
	// Create connection string
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Schema,
	)
	db, err = sql.Open("postgres", psqlconn)
	err = db.Ping()
	return
}

func closeDB(DB *sql.DB) error {
	return DB.Close()
}
