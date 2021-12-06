package config

import (
	"github.com/joho/godotenv"
	"path/filepath"
)

const (
	Directory = "config"
)

type Config struct {
	EmailHandler	EmailHandler
	ErgastClient	ErgastClient
	F1APIClient		F1APIClient
	Repository		Repository
}

type EmailHandler struct {
	SenderAddress  string `env:"SENDER_EMAIL_ADDRESS"`
	SenderPassword string `env:"SENDER_EMAIL_PASSWORD"`
	SMTPServer     string `env:"SMTP_SERVER"`
	SMTPHost       string `env:"SMTP_HOST"`
}

type ErgastClient struct {
	BaseURI                      string `env:"BASE_URL"`
	DriversEndpoint              string `env:"DRIVERS_ENDPOINT"`
	DriverStandingsEndpoint      string `env:"DRIVER_STANDINGS_ENDPOINT"`
	ConstructorStandingsEndpoint string `env:"CONSTRUCTORS_STANDINGS_ENDPOINT"`
}

type F1APIClient struct {
	Host          string `env:"HOST"`
	APIKey        string `env:"API_KEY"`
	BaseURI       string `env:"BASE_URI"`
	EventEndpoint string `env:"CURRENT_EVENT_ENDPOINT"`
	Timezone      string `env:"TIMEZONE"`
}

type Repository struct {

}

// GetConfig loads the variables in .env as env vars
func GetConfig() error {
	file := filepath.Join(Directory, ".env")
	if err := godotenv.Load(file); err != nil {
		return err
	}
	return nil
}