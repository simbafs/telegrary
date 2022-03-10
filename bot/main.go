package bot

import (
	"github.com/simba-fs/telegrary/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type Context struct {
	Bot *tgbotapi.BotAPI
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
func (ctx *Context) Call(name string){
	exec, ok := Commands[name]
	if !ok {
		return
	}
	exec(ctx)
}

type Command func(ctx *Context)

var Commands map[string]Command = make(map[string]Command)

func AddCmd(name string, command Command) {
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
		exec, ok := Commands[update.Message.Command()]
		if !ok {
			continue
		}
		exec(&Context{
			Bot: bot, 
			Update: &update,
		})
	}
}
