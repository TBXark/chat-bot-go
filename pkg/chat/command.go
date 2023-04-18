package chat

import (
	"fmt"
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
)

type Command interface {
	Info() (name string, description string)
	Handle(api *bot.BotAPI, session *Session, update *bot.Update) error
}

type CommandHandler struct {
	Command map[string]Handler
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		Command: make(map[string]Handler),
	}
}

func (h *CommandHandler) AddCommand(command Command) {
	name, _ := command.Info()
	h.Command[name] = command
}

func (h *CommandHandler) Bind(api *bot.BotAPI) error {
	var commands []bot.BotCommand
	for _, handler := range h.Command {
		name, description := handler.(Command).Info()
		commands = append(commands, bot.BotCommand{
			Command:     name,
			Description: description,
		})
	}
	setMyCommands := bot.NewSetMyCommands(commands...)
	_, err := api.Request(setMyCommands)
	return err
}

func (h *CommandHandler) Handle(api *bot.BotAPI, session *Session, update *bot.Update) error {
	if update.Message == nil {
		return nil
	}
	if update.Message.IsCommand() {
		command := update.Message.Command()
		if handler, ok := h.Command[command]; ok {
			err := handler.Handle(api, session, update)
			if err != nil {
				return err
			}
			return io.EOF
		}
	}
	return nil
}

type StartCommand struct {
}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (s *StartCommand) Info() (name string, description string) {
	return "start", "Start a new conversation"
}

func (s *StartCommand) Handle(api *bot.BotAPI, session *Session, update *bot.Update) error {
	session.ClearHistory()
	text := fmt.Sprintf("New conversation started with %d", update.Message.Chat.ID)
	_, _ = api.Send(bot.NewMessage(update.Message.Chat.ID, text))
	return nil
}
