package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) runCommand() {

}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	if err := write(cfg); err != nil {
		errMsg := fmt.Errorf("failed writing CurrentUserName to config: %s", err)
		return errMsg
	}
	return nil
}

// Read returns the local config files content
func Read() (*Config, error) {
	// Get the config filepath
	fullPath, err := getConfigFilePath()
	if err != nil {
		errMsg := fmt.Errorf("config file not found: %s", err)
		return nil, errMsg
	}

	// Read the file
	data, err := os.ReadFile(fullPath)
	if err != nil {
		errMsg := fmt.Errorf("config file failed to read: %s", err)
		return nil, errMsg
	}

	// Unmarshal JSON bytes to struct
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		errMsg := fmt.Errorf("unmarshal json failed: %s", err)
		return nil, errMsg
	}

	return &cfg, nil
}

// getConfigFilePath gets local $HOME directory
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		errMsg := fmt.Errorf("could not finde home directory: %s", err)
		return "", errMsg
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

// write to local config file
func write(cfg *Config) error {
	// Get the config filepath
	dir, err := getConfigFilePath()
	if err != nil {
		errMsg := fmt.Errorf("config file not found: %s", err)
		return errMsg

	}

	// Marshal struct to JSON bytes
	data, err := json.Marshal(cfg)
	if err != nil {
		errMsg := fmt.Errorf("marshal struct failed: %s", err)
		return errMsg
	}

	// Write to file
	err = os.WriteFile(dir, data, 0644)
	if err != nil {
		errMsg := fmt.Errorf("writing to json file failed: %s", err)
		return errMsg
	}

	return nil
}
