package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Username string
	Password string
	Host     string
	DBName   string
}

func LoadConfigFromFile(filename string) (*Config, error) {
	// Read JSON data from file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into DBConfig struct
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
