package chat

import (
	"encoding/json"
	"errors"
	"github.com/TBXark/chat-bot-go/configs"
	"github.com/TBXark/chat-bot-go/pkg/dao"
	tokenizer "github.com/samber/go-gpt-3-encoder"
	"github.com/sashabaranov/go-openai"
	"log"
	"sync"
)

type Session struct {
	id      int64
	lock    sync.Mutex
	dao     *dao.Dao
	config  *configs.ChatConfig
	history []*openai.ChatCompletionMessage
}

func NewSession(id int64, cfg *configs.ChatConfig, dao *dao.Dao) *Session {
	s := &Session{
		id:      id,
		dao:     dao,
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
		tokens += countToken(s.history[i].Content)
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
	_ = s.dao.UpdateChatHistory(s.id, "")
}

func (s *Session) SaveHistory() error {
	bytes, err := json.Marshal(s.history)
	if err != nil {
		return err
	}
	err = s.dao.UpdateChatHistory(s.id, string(bytes))
	return err
}

func (s *Session) RestoreHistory() error {
	history, err := s.dao.FindChatHistoryByChatId(s.id)
	if err != nil {
		return err
	}
	if history.Content == "" {
		s.history = make([]*openai.ChatCompletionMessage, 0)
		return nil
	}
	var messages []*openai.ChatCompletionMessage
	if jErr := json.Unmarshal([]byte(history.Content), &messages); jErr != nil {
		return jErr
	}
	s.history = messages
	return nil
}

func (s *Session) Chat(question *openai.ChatCompletionMessage, with func([]*openai.ChatCompletionMessage) (*openai.ChatCompletionMessage, error)) error {
	if !s.lock.TryLock() {
		return errors.New("session is busy")
	}
	defer s.lock.Unlock()
	s.trimHistory()
	history := append(s.history, question)
	answer, err := with(history)
	if err == nil {
		s.history = append(history, answer)
	}
	go func() {
		sErr := s.SaveHistory()
		if sErr != nil {
			log.Printf("save history error: %v", sErr)
		}
	}()
	return err
}

var encoder *tokenizer.Encoder

func init() {
	e, err := tokenizer.NewEncoder()
	if err == nil {
		encoder = e
	}
}

func countToken(text string) int {
	if encoder != nil {
		encode, err := encoder.Encode(text)
		if err != nil {
			return len(text)
		} else {
			return len(encode)
		}
	} else {
		return len(text)
	}
}
