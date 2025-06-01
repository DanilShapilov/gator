package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return write(*c)
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	cfgData, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	cfg := Config{}
	err = json.Unmarshal(cfgData, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0600)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(homeDir, configFileName)
	return configPath, nil
}
