package cmd

import (
	"fmt"
	"strconv"

	// "log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbot "github.com/simba-fs/telegrary/bot"
	"github.com/simba-fs/telegrary/note"

	log "github.com/sirupsen/logrus"
)

func reply(bot *tgbotapi.BotAPI, update *tgbotapi.Update, text string) {
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text))
}

func init() {
	tgbot.AddCmd("help", func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
		reply(bot, update, fmt.Sprintf("telegrary = telegram + diary\ncommands: %s", tgbot.CommandsList))
	})
	tgbot.AddCmd("read", func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
		year, month, day := getDate(strings.Split(update.Message.Text, " ")[1:])
		diary, err := note.Read(fmt.Sprintf("%s/%d/%d/%d.md", config.Root, year, month, day))
		if err != nil {
			reply(bot, update, "No diary found")
			return
		}
		reply(bot, update, fmt.Sprintf("===== %d/%d/%d.md =====\n%s", year, month, day, diary))
	})
	tgbot.AddCmd("write", func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
		year, month, day := getDate(strings.Split(update.Message.Text, " ")[1:])
		log.Debugln(update.Message.Text)

		// get content
		a := strings.Split(update.Message.Text, " ")
		for k, v := range a {
			_, err := strconv.Atoi(v)
			if k == 0 || err == nil && k <= 3 {
				a = a[1:]
			} else {
				break
			}
		}
		content := "\n" + strings.Trim(strings.Join(a, " "), " ")

		// write
		err := note.Write(fmt.Sprintf("%s/%d/%d/%d.md", config.Root, year, month, day), content, false)
		if err != nil {
			reply(bot, update, "Write failed")
			log.Fatal(err)
			return
		}
		reply(bot, update, "write successfully, use /read to read it")
	})
}

func startBot(token string) {
	tgbot.Run(token)
}
