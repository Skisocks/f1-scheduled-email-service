package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	EmailHandler EmailHandler `yaml:"EMAIL_HANDLER"`
	Ergast       Ergast       `yaml:"ERGAST_F1"`
	SportsIO     SportsIO     `yaml:"SPORTS_IO"`
	Repository   Repository   `yaml:"DATABASE"`
}

type EmailHandler struct {
	EmailName      string `yaml:"EMAIL_NAME"`
	SenderAddress  string `yaml:"SENDER_EMAIL_ADDRESS"`
	SenderPassword string `yaml:"SENDER_EMAIL_PASSWORD"`
	SMTPServer     string `yaml:"SMTP_SERVER"`
	SMTPHost       int    `yaml:"SMTP_HOST"`
}

type Ergast struct {
	BaseURL                      string        `yaml:"BASE_URL"`
	DriversEndpoint              string        `yaml:"DRIVERS_ENDPOINT"`
	DriverStandingsEndpoint      string        `yaml:"DRIVER_STANDINGS_ENDPOINT"`
	Season                       string        `yaml:"SEASON"`
	ConstructorStandingsEndpoint string        `yaml:"CONSTRUCTORS_STANDINGS_ENDPOINT"`
	Timeout                      time.Duration `yaml:"TIMEOUT"`
}

type SportsIO struct {
	Host          string        `yaml:"HOST"`
	APIKey        string        `yaml:"API_KEY"`
	BaseURL       string        `yaml:"BASE_URL"`
	EventEndpoint string        `yaml:"CURRENT_EVENT_ENDPOINT"`
	Timezone      string        `yaml:"TIMEZONE"`
	Timeout       time.Duration `yaml:"TIMEOUT"`
}

type Repository struct {
	Name     string `yaml:"DB_NAME"`
	User     string `yaml:"DB_USER"`
	Password string `yaml:"DB_PASSWORD"`
	Host     string `yaml:"DB_HOST"`
	Port     int    `yaml:"DB_PORT"`
	Schema   string `yaml:"DB_SCHEMA"`
}

// GetConfig loads the variables from config.ini
func GetConfig() (*Config, error) {
	f, err := os.Open("./config.yml")
	if err != nil {
		// Todo: handle error
		return nil, err
	}
	defer f.Close()
	// Todo: handle error

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		// Todo: handle error
		return nil, err
	}

	return &cfg, nil
}
