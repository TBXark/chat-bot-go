package chat

import (
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"strings"
)

type Command interface {
	Info() (name string, description string)
	Handle(update *bot.Update, dependency *HandleContext) error
}

type CommandHandler struct {
	Command map[string]Handler
}

func NewCommandHandler(api *bot.BotAPI) *CommandHandler {
	h := &CommandHandler{
		Command: make(map[string]Handler),
	}
	h.init(api)
	return h
}

func (h *CommandHandler) init(api *bot.BotAPI) {
	h.AddCommand(NewStartCommand())
	h.AddCommand(NewTokenCommand())
	_ = h.Bind(api)
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

func (h *CommandHandler) Handle(update *bot.Update, ctx *HandleContext) error {
	if update.Message == nil {
		return nil
	}
	if update.Message.IsCommand() {
		command := update.Message.Command()
		if handler, ok := h.Command[command]; ok {
			err := handler.Handle(update, ctx)
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

func (s *StartCommand) Handle(update *bot.Update, ctx *HandleContext) error {
	ctx.session.ClearHistory()
	_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, "New conversation started"))
	return nil
}

type TokenCommand struct {
}

func NewTokenCommand() *TokenCommand {
	return &TokenCommand{}
}

func (s *TokenCommand) Info() (name string, description string) {
	return "token", "Manage tokens"
}

func (s *TokenCommand) Handle(update *bot.Update, ctx *HandleContext) error {
	if !ctx.isAdmin {
		_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, "You are not admin"))
		return nil
	}
	if update.Message.Chat.ID < 0 {
		_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, "This command is not allowed in group"))
		return nil
	}
	cmp := strings.Split(update.Message.CommandArguments(), " ")
	switch cmp[0] {
	case "add":
		if len(cmp) < 2 {
			_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, "Missing token"))
			return nil
		}
		if err := ctx.openai.AddClient(cmp[1]); err != nil {
			_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, err.Error()))
			return nil
		}
		_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, "Token added"))
		break
	case "remove":
		if len(cmp) < 2 {
			_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, "Missing token"))
			return nil
		}
		if err := ctx.openai.RemoveClient(cmp[1]); err != nil {
			_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, err.Error()))
			return nil
		}
		_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, "Token removed"))
		break
	case "list":
		allClient := ctx.openai.AllClient()
		message := "Tokens:\n"
		for _, client := range allClient {
			message += "\t" + client + "\n"
		}
		_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, message))
	default:
		message := "Usage:\n"
		message += "\t/token add <token>\n"
		message += "\t/token remove <token>\n"
		message += "\t/token list\n"
		_, _ = ctx.api.Send(bot.NewMessage(update.Message.Chat.ID, message))
	}
	return nil
}
