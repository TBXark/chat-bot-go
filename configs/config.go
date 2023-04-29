package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	OpenAI   OpenAI   `json:"openai"`
	Telegram Telegram `json:"telegram"`
	Database Database `json:"database"`
}

func (c Config) init() {
	for i, chat := range c.Telegram.AvailableChat {
		if chat.MaxHistoryLength == 0 {
			c.Telegram.AvailableChat[i].MaxHistoryLength = 4
		}
		if chat.MaxHistoryTokens == 0 {
			c.Telegram.AvailableChat[i].MaxHistoryTokens = 1024
		}
	}
}

type Database struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type OpenAI struct {
	Key   string `json:"key"`
	Model string `json:"model"`
}

type Telegram struct {
	Token         string       `json:"token"`
	Admin         []int64      `json:"admin"`
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

func NewConfig(path string) ([]*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	list := make([]*Config, 0)
	if e := json.Unmarshal(file, &list); e == nil {
		if len(list) > 0 {
			for _, c := range list {
				c.init()
			}
			return list, nil
		}
	}

	cfg := Config{}
	if e := json.Unmarshal(file, &cfg); e == nil {
		cfg.init()
		return []*Config{&cfg}, nil
	}
	return nil, fmt.Errorf("invalid config file: %s", path)
}
