package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/TBXark/chat-bot-go/configs"
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"strings"
)

type App struct {
	bot      *bot.BotAPI
	ai       *openai.Client
	sessions map[int64]*Session
}

func NewApp(cfg *configs.Config) *App {
	api, err := bot.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Fatal(err)
	}
	ai := openai.NewClient(cfg.Openai.Key)
	sessions := make(map[int64]*Session)
	for _, chat := range cfg.Telegram.AvailableChat {
		sessions[chat.ChatID] = NewSession(chat.ChatID)
	}
	return &App{
		bot:      api,
		ai:       ai,
		sessions: sessions,
	}
}

func (a *App) Run() {
	u := bot.NewUpdate(0)
	u.Timeout = 60
	updates := a.bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Chat.ID < 0 {
			canHandleThisMessage := false
			for _, e := range update.Message.Entities {
				if e.Type == "mention" {
					text := strings.TrimLeft(update.Message.Text[e.Offset:e.Offset+e.Length], "@")
					if text == a.bot.Self.UserName {
						canHandleThisMessage = true
						break
					}
				}
			}
			if update.Message.ReplyToMessage != nil {
				if update.Message.ReplyToMessage.From.UserName == a.bot.Self.UserName {
					canHandleThisMessage = true
				}
			}
			if !canHandleThisMessage {
				continue
			}
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		session, ok := a.sessions[update.Message.Chat.ID]
		if !ok {
			_, _ = a.bot.Send(bot.NewMessage(update.Message.Chat.ID, "Sorry, I don't know you"))
			continue
		}
		go a.handleUpdate(session, update)
	}
}

func (a *App) handleUpdate(session *Session, update bot.Update) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in f: %v", r)
		}
	}()
	sErr := session.Chat(func(history []*openai.ChatCompletionMessage) (*openai.ChatCompletionMessage, error) {
		ctx := context.Background()
		var messages []openai.ChatCompletionMessage
		for _, m := range history {
			messages = append(messages, *m)
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: update.Message.Text,
		})
		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 2048,
			Messages:  messages,
			Stream:    true,
		}
		stream, err := a.ai.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return nil, err
		}
		defer stream.Close()

		answer := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "",
		}
		send, err := a.bot.Send(bot.NewMessage(update.Message.Chat.ID, "Thinking..."))
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
				_, _ = a.bot.Send(bot.NewEditMessageText(update.Message.Chat.ID, send.MessageID, answer.Content))
			}
		}
		answerMsg := bot.NewEditMessageText(update.Message.Chat.ID, send.MessageID, answer.Content)
		answerMsg.ParseMode = "MarkdownV2"
		if _, err = a.bot.Send(answerMsg); err != nil {
			answerMsg.ParseMode = ""
			_, _ = a.bot.Send(answerMsg)
		}
		return &answer, nil
	})
	if sErr != nil {
		_, _ = a.bot.Send(bot.NewMessage(update.Message.Chat.ID, sErr.Error()))
	}
}
