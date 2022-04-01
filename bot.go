package main

import (
	_ "embed"
	"fmt"
	"path"
	"strconv"
	"strings"


	"github.com/simba-fs/telegrary/config"
	"github.com/simba-fs/telegrary/note"
	"github.com/simba-fs/telegrary/util"
	"github.com/simba-fs/telegrary/bot"

	log "github.com/sirupsen/logrus"
)

//go:embed help/bot.txt
var help string

var loginedUsers = make(map[string]bool, 0)

func auth(ctx *bot.Context) bool {
	if _, ok := loginedUsers[ctx.Update.Message.From.UserName]; !ok {
		ctx.Send("please login first")
		return false
	}
	return true
}

func init() {
	bot.AddCmd("start", func(ctx *bot.Context) bool {
		if len(ctx.Args) == 1 {
			ctx.Send("please enter secret key")
			return true
		}

		if util.Hash(ctx.Args[1]) != config.Config.Secret {
			ctx.Send("wrong secret key")
			return true
		}
		username := ctx.Update.Message.From.UserName
		loginedUsers[username] = true

		ctx.Send("Welcome to Telegrary! " + username)
		return true
	})
	bot.AddCmd("help", auth, func(ctx *bot.Context) bool {
		ctx.Send(help)
		return true
	})
	bot.AddCmd("read", auth, func(ctx *bot.Context) bool {
		year, month, day := util.GetDate(strings.Split(ctx.Update.Message.Text, " ")[1:])
		diary, err := note.Read(util.Path(year, month, day))
		if err != nil {
			ctx.Send("No diary found")
			return true
		}
		ctx.Send(fmt.Sprintf("===== %s/%s/%s.md =====\n%s", year, month, day, diary))
		return true
	})
	bot.AddCmd("write", auth, func(ctx *bot.Context) bool {
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
			return true
		}
		ctx.Send("write successfully, use /read to read it")

		return true
	})
	bot.AddCmd("tree", auth, func(ctx *bot.Context) bool {
		prefix := ""
		if len(strings.Split(ctx.Update.Message.Text, " ")) > 1 {
			prefix = strings.Split(ctx.Update.Message.Text, " ")[1]
		}
		tree, err := note.Tree(path.Join(config.Config.Root, prefix))
		if err != nil {
			ctx.Send("Tree failed")
			log.Error(err)
			return true
		}
		ctx.Send(tree)

		return true
	})
}

// startBot starts the bot
func startBot() {
	bot.Run(config.Config.Token)
}
