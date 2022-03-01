package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type Command func(bot *tgbotapi.BotAPI, update *tgbotapi.Update)

var Commands map[string]Command = make(map[string]Command)
var CommandsList []string = make([]string, 3)

func AddCmd(name string, command Command) {
	Commands[name] = command
	CommandsList = append(CommandsList, name)
}

// Run starts the bot
func Run(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

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
		exec(bot, &update)
	}
}

// Reply sends a message to the chat where the command was received
func Reply(bot *tgbotapi.BotAPI, update *tgbotapi.Update, text string) {
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text))
}

