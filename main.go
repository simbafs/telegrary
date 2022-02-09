package main

import (
	"fmt"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigtoml"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	Token string
}

var config Config

func init() {
	loader := aconfig.LoaderFor(&config, aconfig.Config{
		EnvPrefix:  "gitdiary",
		FlagPrefix: "gitdiary",
		Files:      []string{"~/.config/gitdiary.toml", "gitdiary.toml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".toml": aconfigtoml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(config)

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
}
