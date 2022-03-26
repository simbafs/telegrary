package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/simba-fs/telegrary/config"
	log "github.com/sirupsen/logrus"
)

type Context struct {
	Bot    *tgbotapi.BotAPI
	Update *tgbotapi.Update
}

// Send sends a message to the user
func (ctx *Context) Send(text string) {
	msg := tgbotapi.NewMessage(ctx.Update.Message.Chat.ID, text)
	// msg.ReplyToMessageID = ctx.Update.Message.MessageID
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Error(err)
	}
}

// Call calls another command with the same context
func (ctx *Context) Call(name string) {
	cmds, ok := Commands[name]
	if !ok {
		return
	}

	for _, cmd := range cmds {
		if !cmd(ctx) {
			break
		}
	}
}

// return false to prevent the next handler from being called
type CmdHandler func(ctx *Context) bool

var Commands map[string][]CmdHandler = make(map[string][]CmdHandler)

func AddCmd(name string, command ...CmdHandler) {
	Commands[name] = command
}

// Run starts the bot
func Run(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = config.Config.Debug

	log.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Extract the command from the Message.
		handlers, ok := Commands[update.Message.Command()]
		if !ok {
			continue
		}
		for _, handler := range handlers {
			if !handler(&Context{Bot: bot, Update: &update}) {
				break
			}
		}
	}
}
