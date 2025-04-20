package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUsername string `json:"current_username"`
}

const configFileName = ".gatorconfig.json"

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
	return filepath.Join(homeDir, configFileName), nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	var cfg Config

	configPath, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(fileData, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUsername = username
	return write(*c)
}
