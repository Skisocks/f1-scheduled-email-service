package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	EmailHandler EmailHandler `yaml:"EMAIL_HANDLER"`
	ErgastClient ErgastClient `yaml:"ERGAST_F1"`
	F1APIClient  F1APIClient  `yaml:"F1_API"`
	Repository   Repository   `yaml:"repository"`
}

type EmailHandler struct {
	SenderAddress  string `yaml:"SENDER_EMAIL_ADDRESS"`
	SenderPassword string `yaml:"SENDER_EMAIL_PASSWORD"`
	SMTPServer     string `yaml:"SMTP_SERVER"`
	SMTPHost       string `yaml:"SMTP_HOST"`
}

type ErgastClient struct {
	BaseURL                      string `yaml:"BASE_URL"`
	DriversEndpoint              string `yaml:"DRIVERS_ENDPOINT"`
	DriverStandingsEndpoint      string `yaml:"DRIVER_STANDINGS_ENDPOINT"`
	ConstructorStandingsEndpoint string `yaml:"CONSTRUCTORS_STANDINGS_ENDPOINT"`
}

type F1APIClient struct {
	Host          string `yaml:"HOST"`
	APIKey        string `yaml:"API_KEY"`
	BaseURL       string `yaml:"BASE_URL"`
	EventEndpoint string `yaml:"CURRENT_EVENT_ENDPOINT"`
	Timezone      string `yaml:"TIMEZONE"`
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
	//Todo: handle error

	var cfg Config
	decoder:= yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		// Todo: handle error
	}

	return &cfg, err
}