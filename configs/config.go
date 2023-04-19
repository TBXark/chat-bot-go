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
	ChatID           []int64       `json:"chat_id"`
	MaxHistoryLength int           `json:"max_history_length"`
	MaxHistoryTokens int           `json:"max_history_tokens"`
	Params           ChatGPTParams `json:"params"`
}

type ChatGPTParams struct {
	InitMessage string         `json:"init_message"`
	ExtraParams map[string]any `json:"extra_params"`
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
	for i, chat := range config.Telegram.AvailableChat {
		if chat.MaxHistoryLength == 0 {
			config.Telegram.AvailableChat[i].MaxHistoryLength = 4
		}
		if chat.MaxHistoryTokens == 0 {
			config.Telegram.AvailableChat[i].MaxHistoryTokens = 1024
		}
	}
	return config, nil
}
