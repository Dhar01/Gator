package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	Db_url    string `json:"db_url"`
	User_name string `json:"user_name"`
}

func Read() (Config, error) {
	var config Config

	path, err := getConfigFilePath()
	if err != nil {
		return config, fmt.Errorf("failed to get config file path: %w", err)
	}

	file, err := os.Open(path)
	if err != nil {
		return config, fmt.Errorf("cannot open the file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("failed to decode config: %w", err)
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home directory not found: %w", err)
	}

	filePath := path + configFileName

	return filePath, nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("failed to get the config file path: %w", err)
	}

	data, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("file couldn't found: %w", err)
	}

	encoder := json.NewEncoder(data)
	if err = encoder.Encode(&cfg); err != nil {
		return fmt.Errorf("cannot encode the file")
	}

	return nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.User_name = userName

	if err := write(*cfg); err != nil {
		return fmt.Errorf("failed to update user in config: %w", err)
	}

	fmt.Println("User updated successfully")
	return nil
}
