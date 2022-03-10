package main

import (
	"fmt"
	"strconv"
	_ "embed"

	"strings"

	tgbot "github.com/simba-fs/telegrary/bot"

	"github.com/simba-fs/telegrary/config"
	"github.com/simba-fs/telegrary/note"
	"github.com/simba-fs/telegrary/util"

	log "github.com/sirupsen/logrus"
)

//go:embed help/bot.txt
var help string

func init() {
	tgbot.AddCmd("start", func(ctx *tgbot.Context) {
		ctx.Call("help")
	})
	tgbot.AddCmd("help", func(ctx *tgbot.Context) {
		ctx.Send(help)
	})
	tgbot.AddCmd("read", func(ctx *tgbot.Context) {
		year, month, day := util.GetDate(strings.Split(ctx.Update.Message.Text, " ")[1:])
		diary, err := note.Read(util.Path(year, month, day))
		if err != nil {
			ctx.Send("No diary found")
			return
		}
		ctx.Send(fmt.Sprintf("===== %s/%s/%s.md =====\n%s", year, month, day, diary))
	})
	tgbot.AddCmd("write", func(ctx *tgbot.Context) {
		year, month, day := util.GetDate(strings.Split(ctx.Update.Message.Text, " ")[1:])
		log.Debugln(ctx.Update.Message.Text)

		// get content
		a := strings.Split(ctx.Update.Message.Text, " ")
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
		err := note.Write(util.Path(year, month, day), content, false)
		if err != nil {
			ctx.Send("Write failed")
			log.Fatal(err)
			return
		}
		ctx.Send("write successfully, use /read to read it")
	})
	tgbot.AddCmd("tree", func(ctx *tgbot.Context) {
		tree, err := note.Tree(config.Config.Root)
		if err != nil {
			ctx.Send("Tree failed")
			log.Error(err)
			return
		}
		ctx.Send(tree)
	})
}

// startBot starts the bot
func startBot() {
	tgbot.Run(config.Config.Token)
}
