package chat

import (
	"errors"
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"io"
)

type HandleContext struct {
	isAdmin bool
	api     *bot.BotAPI
	openai  *OpenAI
	session *Session
}

type Handler interface {
	Handle(update *bot.Update, ctx *HandleContext) error
}

type GPTHandler struct {
}

func NewChatHandler() *GPTHandler {
	return &GPTHandler{}
}

func (h *GPTHandler) Handle(update *bot.Update, ctx *HandleContext) error {
	question := &openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: update.Message.Text,
	}
	return ctx.session.Chat(question, func(history []*openai.ChatCompletionMessage) (*openai.ChatCompletionMessage, error) {
		answer := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "",
		}

		stream, err := ctx.openai.CreateChatCompletionStream(history)
		if err != nil {
			return nil, err
		}
		defer stream.Close()

		send, err := ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, "Thinking..."))
		if err != nil {
			return nil, err
		}

		lengthDelta := 0
		for {
			response, rErr := stream.Recv()
			if errors.Is(rErr, io.EOF) {
				break
			}
			if rErr != nil {
				return nil, rErr
			}
			lengthDelta += len(response.Choices[0].Delta.Content)
			answer.Content += response.Choices[0].Delta.Content
			if lengthDelta > 100 {
				lengthDelta = 0
				_, _ = ctx.api.Send(bot.NewEditMessageText(update.Message.Chat.ID, send.MessageID, answer.Content))
			}
		}
		answerMsg := bot.NewEditMessageText(update.Message.Chat.ID, send.MessageID, answer.Content)
		answerMsg.ParseMode = "Markdown"
		if _, err = ctx.api.Send(answerMsg); err != nil {
			answerMsg.ParseMode = ""
			_, _ = ctx.api.Send(answerMsg)
		}
		return &answer, nil
	})
}
