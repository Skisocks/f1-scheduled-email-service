package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

const (
	Directory = "config"
)

type Config struct {
	EmailHandler EmailHandler `yaml:"email_handler"`
	ErgastClient ErgastClient `yaml:"ERGAST_F1"`
	F1APIClient  F1APIClient  `yaml:"f_1_api_client"`
	Repository   Repository   `yaml:"repository"`
}

type EmailHandler struct {
	SenderAddress  string `env:"SENDER_EMAIL_ADDRESS"`
	SenderPassword string `env:"SENDER_EMAIL_PASSWORD"`
	SMTPServer     string `env:"SMTP_SERVER"`
	SMTPHost       string `env:"SMTP_HOST"`
}

type ErgastClient struct {
	BaseURI                      string `yaml:"BASE_URL"`
	DriversEndpoint              string `yaml:"DRIVERS_ENDPOINT"`
	DriverStandingsEndpoint      string `yaml:"DRIVER_STANDINGS_ENDPOINT"`
	ConstructorStandingsEndpoint string `yaml:"CONSTRUCTORS_STANDINGS_ENDPOINT"`
}

type F1APIClient struct {
	Host          string `env:"HOST"`
	APIKey        string `env:"API_KEY"`
	BaseURI       string `env:"F1_API_BASE_URI"`
	EventEndpoint string `env:"CURRENT_EVENT_ENDPOINT"`
	Timezone      string `env:"TIMEZONE"`
}

type Repository struct {
	// Todo: add Repository variables
}

// GetConfig loads the variables in .env as env vars
func GetConfig() (*Config, error) {
	f, err := os.Open("config/config.yml")
	if err != nil {
		//Todo: handle error
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder:= yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		// Todo: handle error
	}

	return &cfg, err
}