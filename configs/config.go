package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	OpenaiKey     string `json:"openai_key"`
	TelegramToken string `json:"telegram_token"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
