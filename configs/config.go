package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	Openai   OpenAI   `json:"openai"`
	Telegram Telegram `json:"telegram"`
}

type OpenAI struct {
	Key string `json:"key"`
}

type Telegram struct {
	Token         string       `json:"token"`
	AvailableChat []ChatConfig `json:"available_chat"`
}

type ChatConfig struct {
	ChatID int64 `json:"chat_id"`
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
