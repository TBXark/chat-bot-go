package chat

import (
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
	handler  []Handler
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
	app := &App{
		bot:      api,
		ai:       ai,
		sessions: sessions,
	}
	app.init(cfg)
	return app
}

func (a *App) init(cfg *configs.Config) {
	cmd := NewCommandHandler()
	cmd.AddCommand(NewStartCommand())
	_ = cmd.Bind(a.bot)

	a.AddHandler(cmd)
	a.AddHandler(NewChatHandler(cfg))
}

func (a *App) AddHandler(handler Handler) {
	a.handler = append(a.handler, handler)
}

func (a *App) Run() {
	u := bot.NewUpdate(0)
	u.Timeout = 60
	updates := a.bot.GetUpdatesChan(u)
	for update := range updates {
		go a.handleUpdate(&update)
	}
}

func (a *App) handleUpdate(update *bot.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if !a.canHandleUpdate(update) {
		return
	}
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	session, ok := a.sessions[update.Message.Chat.ID]
	if !ok {
		text := fmt.Sprintf("No session for chat %d", update.Message.Chat.ID)
		_, _ = a.bot.Send(bot.NewMessage(update.Message.Chat.ID, text))
		return
	}
	for _, handler := range a.handler {
		if err := handler.Handle(a.bot, session, update); err != nil {
			if !errors.Is(err, io.EOF) {
				_, _ = a.bot.Send(bot.NewMessage(update.Message.Chat.ID, err.Error()))
			}
			break
		}
	}
}

func (a *App) canHandleUpdate(update *bot.Update) bool {
	if update.Message == nil {
		return false
	}
	if update.Message.Chat.ID < 0 {
		for _, e := range update.Message.Entities {
			if e.Type == "mention" {
				text := strings.TrimLeft(update.Message.Text[e.Offset:e.Offset+e.Length], "@")
				if text == a.bot.Self.UserName {
					return true
				}
			}
		}
		if update.Message.ReplyToMessage != nil {
			if update.Message.ReplyToMessage.From.UserName == a.bot.Self.UserName {
				return true
			}
		}
		return false
	}
	return true
}
