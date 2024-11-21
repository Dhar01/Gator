package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBurl           string `json:"db_url"`
	CurrentUserName string `json:"user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func Read() (Config, error) {
	var config Config

	fullPath, err := getConfigFilePath()
	if err != nil {
		return config, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, err
	}

	return config, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(&cfg); err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(home, configFileName)

	return fullPath, nil
}
