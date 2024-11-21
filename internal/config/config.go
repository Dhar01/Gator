package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

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

	filePath := path + configFileName

	file, err := os.Open(filePath)
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
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return home, nil
}

func write(cfg Config) error {
	return nil
}

func (cfg *Config) SetUser() {

}
