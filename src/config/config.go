package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Username string
	Password string
	Host     string
	DBName   string
	DBPath   string
}

func LoadConfigFromFile(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
