package chat

import (
	"errors"
	"github.com/TBXark/chat-bot-go/configs"
	"github.com/sashabaranov/go-openai"
	"sync"
)

type Session struct {
	config  *configs.ChatConfig
	history []*openai.ChatCompletionMessage
	lock    sync.Mutex
}

func NewSession(cfg *configs.ChatConfig) *Session {
	s := &Session{
		config:  cfg,
		history: make([]*openai.ChatCompletionMessage, 0),
	}
	_ = s.RestoreHistory()
	return s
}

func (s *Session) trimHistory() {
	if len(s.history) > s.config.MaxHistoryLength {
		s.history = s.history[len(s.history)-s.config.MaxHistoryLength:]
	}
	tokens := 0
	for i := len(s.history) - 1; i >= 0; i-- {
		tokens += len(s.history[i].Content)
		if tokens > s.config.MaxHistoryTokens {
			s.history = s.history[i+1:]
			break
		}
	}
	if len(s.history) > 0 && s.history[0].Role == openai.ChatMessageRoleAssistant {
		s.history = s.history[1:]
	}
	if len(s.history) > 0 && s.history[0].Role != openai.ChatMessageRoleSystem {
		s.history = append([]*openai.ChatCompletionMessage{{
			Role:    openai.ChatMessageRoleSystem,
			Content: s.config.Params.InitMessage,
		}}, s.history...)
	}
}

func (s *Session) ClearHistory() {
	s.history = make([]*openai.ChatCompletionMessage, 0)
}

func (s *Session) SaveHistory() error {
	// TODO: save history to db
	return nil
}

func (s *Session) RestoreHistory() error {
	// TODO: restore history from db
	return nil
}

func (s *Session) Chat(with func([]*openai.ChatCompletionMessage) (*openai.ChatCompletionMessage, error)) error {
	if !s.lock.TryLock() {
		return errors.New("session is busy")
	}
	defer s.lock.Unlock()
	s.trimHistory()
	answer, err := with(s.history)
	if err == nil {
		s.history = append(s.history, answer)
	}
	_ = s.SaveHistory()
	return err
}
