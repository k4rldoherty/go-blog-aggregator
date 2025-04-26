package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const fileName = "/.gatorconfig.json"

func Read() (Config, error) {
	filePath, err := getFilePath(fileName)
	if err != nil {
		return Config{}, fmt.Errorf(" err: %v", err)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error opening file: %v ", filePath)
	}
	var appConfig Config
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&appConfig); err != nil {
		return Config{}, fmt.Errorf("error decoding json")
	}
	return appConfig, nil
}

func getFilePath(fileName string) (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory")
	}
	filepath := homePath + fileName
	return filepath, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	filePath, err := getFilePath(fileName)
	if err != nil {
		return fmt.Errorf("error gettting file")
	}
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v ", err)
	}
	jsonParser := json.NewEncoder(file)
	err = jsonParser.Encode(cfg)
	if err != nil {
		return fmt.Errorf("error encoding json: %v ", err)
	}
	return nil
}
