package main

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

func init() {
	tgbot.AddCmd("start", func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
		tgbot.Reply(bot, update, update.Message.Text)
	})
	tgbot.AddCmd("help", func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
		tgbot.Reply(bot, update, fmt.Sprintf("telegrary = telegram + diary\ncommands: %s", tgbot.CommandsList))
	})
	tgbot.AddCmd("read", func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
		year, month, day := getDate(strings.Split(update.Message.Text, " ")[1:])
		diary, err := note.Read(fmt.Sprintf("%s/%d/%d/%d.md", config.Root, year, month, day))
		if err != nil {
			tgbot.Reply(bot, update, "No diary found")
			return
		}
		tgbot.Reply(bot, update, fmt.Sprintf("===== %d/%d/%d.md =====\n%s", year, month, day, diary))
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
			tgbot.Reply(bot, update, "Write failed")
			log.Fatal(err)
			return
		}
		tgbot.Reply(bot, update, "write successfully, use /read to read it")
	})
	tgbot.AddCmd("tree", func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
		tree, err := note.Tree(config.Root)
		if err != nil {
			tgbot.Reply(bot, update, "Tree failed")
			log.Error(err)
			return
		}
		tgbot.Reply(bot, update, tree)
	})
}

func startBot(token string) {
	tgbot.Run(token)
}
