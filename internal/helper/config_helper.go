package helper

import (
	"encoding/json"
	"os"
)

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

type Config struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	TelegramToken string `json:"telegramToken"`
	ChatID        string `json:"chatID"`
}
