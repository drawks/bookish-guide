package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds the bot configuration
type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

// LoadConfig loads configuration from a JSON file
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error decoding config: %w", err)
	}

	return &config, nil
}

// LoadConfigFromEnv loads configuration from environment variables
func LoadConfigFromEnv() (*Config, error) {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("DISCORD_TOKEN environment variable not set")
	}

	prefix := os.Getenv("DISCORD_PREFIX")
	if prefix == "" {
		prefix = "!"
	}

	return &Config{
		Token:  token,
		Prefix: prefix,
	}, nil
}
