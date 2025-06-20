package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port            string
	OsuClientId     int
	OsuClientSecret string
	OsuCallbackURL  string
	OsuApiKey       string
	SessionSecret   string
	IsProduction    bool
}

func LoadConfig() *AppConfig {}
