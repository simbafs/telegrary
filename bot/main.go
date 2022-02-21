package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type Command func(bot *tgbotapi.BotAPI, update *tgbotapi.Update)

var Commands map[string]Command = make(map[string]Command)

func AddCmd(name string, command Command) {
	Commands[name] = command
}

// Run starts the bot
func Run(token string) {
	log.SetLevel(log.DebugLevel)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Debugf("Authorized on account %s", bot.Self.UserName)

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
