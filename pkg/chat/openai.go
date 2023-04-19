package chat

import (
	"context"
	"github.com/TBXark/chat-bot-go/configs"
	"github.com/TBXark/chat-bot-go/pkg/dao"
	"github.com/sashabaranov/go-openai"
	"math/rand"
)

type OpenAI struct {
	dao           *dao.Dao
	defaultClient *openaiClient
	clientQueue   []*openaiClient
}

type openaiClient struct {
	key string
	*openai.Client
}

func newClientItem(key string) *openaiClient {
	return &openaiClient{
		key:    key,
		Client: openai.NewClient(key),
	}
}

func (i *openaiClient) check() error {
	_, err := i.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Ping",
			},
		},
	})
	return err
}

func NewOpenAI(cfg *configs.OpenAI, db *dao.Dao) *OpenAI {
	ai := &OpenAI{
		dao:           db,
		defaultClient: newClientItem(cfg.Key),
		clientQueue:   make([]*openaiClient, 0),
	}
	_ = ai.RestoreClients()
	return ai
}

func (o *OpenAI) GetRandClient() *openai.Client {
	if len(o.clientQueue) == 0 {
		return o.defaultClient.Client
	}
	index := rand.Intn(len(o.clientQueue)+1) - 1
	if index < 0 {
		return o.defaultClient.Client
	}
	return o.clientQueue[index].Client
}

func (o *OpenAI) AllClient() []string {
	list := []string{
		o.defaultClient.key,
	}
	for _, c := range o.clientQueue {
		list = append(list, c.key)
	}

	return list
}

func (o *OpenAI) AddClient(key string) error {
	client := newClientItem(key)
	if err := client.check(); err != nil {
		return err
	}
	err := o.dao.CreateOpenAIToken(key)
	if err != nil {
		return err
	}
	o.clientQueue = append(o.clientQueue, &openaiClient{
		key:    key,
		Client: openai.NewClient(key),
	})
	return nil
}

func (o *OpenAI) RemoveClient(key string) error {
	err := o.dao.DeleteOpenAIToken(key)
	if err != nil {
		return err
	}
	for i, c := range o.clientQueue {
		if c.key == key {
			o.clientQueue = append(o.clientQueue[:i], o.clientQueue[i+1:]...)
			break
		}
	}
	return nil
}

func (o *OpenAI) RestoreClients() error {
	tokens, err := o.dao.FinaAllActiveOpenAIToken()
	if err != nil {
		return err
	}
	for _, t := range tokens {
		item := &openaiClient{
			key:    t.Token,
			Client: openai.NewClient(t.Token),
		}
		if item.check() != nil {
			continue
		}
		o.clientQueue = append(o.clientQueue, item)
	}
	return nil
}
