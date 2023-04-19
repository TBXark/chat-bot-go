package chat

import (
	"context"
	"errors"
	"fmt"
	"github.com/TBXark/chat-bot-go/configs"
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"io"
)

type HandleContext struct {
	api     *bot.BotAPI
	session *Session
}

type Handler interface {
	Handle(update *bot.Update, ctx *HandleContext) error
}

type GPTHandler struct {
	ai *openai.Client
}

func NewChatHandler(cfg *configs.Config) *GPTHandler {
	return &GPTHandler{
		ai: openai.NewClient(cfg.Openai.Key),
	}
}

func (h *GPTHandler) Handle(update *bot.Update, ctx *HandleContext) error {
	question := &openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: update.Message.Text,
	}
	return ctx.session.Chat(question, func(history []*openai.ChatCompletionMessage) (*openai.ChatCompletionMessage, error) {
		var messages []openai.ChatCompletionMessage
		for _, m := range history {
			messages = append(messages, *m)
		}
		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 2048,
			Messages:  messages,
			Stream:    true,
		}
		stream, err := h.ai.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return nil, err
		}
		defer stream.Close()

		answer := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "",
		}
		send, err := ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, "Thinking..."))
		if err != nil {
			return nil, err
		}
		lengthDelta := 0
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return nil, err
			}
			lengthDelta += len(response.Choices[0].Delta.Content)
			answer.Content += response.Choices[0].Delta.Content
			if lengthDelta > 100 {
				lengthDelta = 0
				_, _ = ctx.api.Send(bot.NewEditMessageText(update.Message.Chat.ID, send.MessageID, answer.Content))
			}
		}
		answerMsg := bot.NewEditMessageText(update.Message.Chat.ID, send.MessageID, answer.Content)
		answerMsg.ParseMode = "MarkdownV2"
		if _, err = ctx.api.Send(answerMsg); err != nil {
			answerMsg.ParseMode = ""
			_, _ = ctx.api.Send(answerMsg)
		}
		return &answer, nil
	})
}
