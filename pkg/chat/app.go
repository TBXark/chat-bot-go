package chat

import (
	"errors"
	"fmt"
	"github.com/TBXark/chat-bot-go/configs"
	"github.com/TBXark/chat-bot-go/pkg/dao"
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"strings"
)

type App struct {
	bot      *bot.BotAPI
	openai   *OpenAI
	sessions map[int64]*Session
	handler  []Handler
	config   *configs.Config
}

func NewApp(cfg *configs.Config, dao *dao.Dao) *App {
	api, err := bot.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Fatal(err)
	}
	sessions := make(map[int64]*Session)
	for _, chat := range cfg.Telegram.AvailableChat {
		for _, id := range chat.ChatID {
			sessions[id] = NewSession(id, &chat, dao)
		}
	}
	app := &App{
		bot:      api,
		openai:   NewOpenAI(&cfg.OpenAI, dao),
		sessions: sessions,
		config:   cfg,
	}
	app.init(cfg)
	return app
}

func (a *App) init(cfg *configs.Config) {
	a.AddHandler(NewCommandHandler(a.bot))
	a.AddHandler(NewChatHandler())
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
	dep := &HandleContext{
		isAdmin: false,
		api:     a.bot,
		openai:  a.openai,
		session: session,
	}
	for _, admin := range a.config.Telegram.Admin {
		if update.Message.From.ID == admin {
			dep.isAdmin = true
			break
		}
	}
	for _, handler := range a.handler {
		if err := handler.Handle(update, dep); err != nil {
			if !errors.Is(err, io.EOF) {
				_, _ = a.bot.Send(bot.NewMessage(update.Message.Chat.ID, err.Error()))
			}
			break
		}
	}
}

func (a *App) canHandleUpdate(update *bot.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}
	if update.Message.Chat.ID < 0 {
		for _, e := range update.Message.Entities {
			if e.Type == "mention" {
				text := update.Message.Text[e.Offset : e.Offset+e.Length]
				if text[1:] == a.bot.Self.UserName {
					update.Message.Text = strings.Replace(update.Message.Text, text, "", 1)
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
