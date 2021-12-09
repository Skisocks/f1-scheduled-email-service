package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	EmailHandler EmailHandler `envconfig:"EMAIL_HANDLER"`
	Ergast       Ergast       `envconfig:"ERGAST_F1"`
	SportsIO     SportsIO     `envconfig:"SPORTS_IO"`
	Repository   Repository   `envconfig:"DATABASE"`
}

type EmailHandler struct {
	EmailName      string `envconfig:"EMAIL_NAME"`
	SenderAddress  string `envconfig:"EMAIL_SENDER_EMAIL_ADDRESS"`
	SenderPassword string `envconfig:"EMAIL_SENDER_EMAIL_PASSWORD"`
	SMTPServer     string `envconfig:"EMAIL_SMTP_SERVER"`
	SMTPHost       int    `envconfig:"EMAIL_SMTP_HOST"`
}

type Ergast struct {
	BaseURL                      string `envconfig:"ERGAST_BASE_URL"`
	DriversEndpoint              string `envconfig:"ERGAST_DRIVERS_ENDPOINT"`
	DriverStandingsEndpoint      string `envconfig:"ERGAST_DRIVER_STANDINGS_ENDPOINT"`
	Season                       string `envconfig:"ERGAST_SEASON"`
	ConstructorStandingsEndpoint string `envconfig:"ERGAST_CONSTRUCTORS_STANDINGS_ENDPOINT"`
}

type SportsIO struct {
	Host          string `envconfig:"SPORTS_IO_HOST"`
	APIKey        string `envconfig:"SPORTS_IO_API_KEY"`
	BaseURL       string `envconfig:"SPORTS_IO_BASE_URL"`
	EventEndpoint string `envconfig:"SPORTS_IO_CURRENT_EVENT_ENDPOINT"`
	Timezone      string `envconfig:"SPORTS_IO_TIMEZONE"`
}

type Repository struct {
	Name     string `envconfig:"DB_NAME" `
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Host     string `envconfig:"DB_HOST"`
	Port     int    `envconfig:"DB_PORT"`
	Schema   string `envconfig:"DB_SCHEMA"`
}

// GetConfig loads the variables from config.ini
func GetConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		// Todo: Log here?
		return nil, err
	}
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		// Todo: Log here?
		return nil, err
	}
	return &cfg, nil
}
