package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	PackSizes []int `json:"pack_sizes"`
}

// LoadConfig reads the pack sizes configuration from a JSON file
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
