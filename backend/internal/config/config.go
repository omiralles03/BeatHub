package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
	PORT              string
	OSU_CLIENT_ID     string
	OSU_CLIENT_SECRET string
	OSU_CALLBACK_URL  string
	OSU_API_KEY       string
	SESSION_SECRET    string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	VALKEY_HOST string
	VALKEY_PORT string
}

func LoadConfig() (*AppConfig, error) {

	appCfg := &AppConfig{
		PORT:              os.Getenv("PORT"),
		OSU_CLIENT_ID:     os.Getenv("OSU_CLIENT_ID"),
		OSU_CLIENT_SECRET: os.Getenv("OSU_CLIENT_SECRET"),
		OSU_CALLBACK_URL:  os.Getenv("OSU_CALLBACK_URL"),
		OSU_API_KEY:       os.Getenv("OSU_API_KEY"),
		SESSION_SECRET:    os.Getenv("SESSION_SECRET"),

		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		VALKEY_HOST: os.Getenv("VALKEY_HOST"),
		VALKEY_PORT: os.Getenv("VALKEY_PORT"),
	}

	// Default PORT
	if appCfg.PORT == "" {
		appCfg.PORT = "3000"
	}

	// NOTE: maybe remove API_KEY
	// Critical variables should be set in the .env
	if appCfg.OSU_CLIENT_ID == "" || appCfg.OSU_CLIENT_SECRET == "" || appCfg.OSU_API_KEY == "" {
		return nil, fmt.Errorf("missing one or more required OSU API environment variables (OSU_CLIENT_ID, OSU_CLIENT_SECRET, OSU_API_KEY)")
	}
	if appCfg.DB_USER == "" || appCfg.DB_PASSWORD == "" || appCfg.DB_NAME == "" {
		return nil, fmt.Errorf("missing one or more required database environment variables (DB_USER, DB_PASSWORD, DB_NAME)")
	}
	if appCfg.SESSION_SECRET == "" {
		return nil, fmt.Errorf("missing SESSION_SECRET environment variable")
	}

	return appCfg, nil
}
