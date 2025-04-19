package main

import (
	"os"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUsername string `json:"current_username"`
}

func NewConfig(dbUrl, username string) *Config {
	return &Config{
		DBUrl:           dbUrl,
		CurrentUsername: username,
	}
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/.gatorconfig.json", nil
}
