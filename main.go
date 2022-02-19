package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/simba-fs/telegrary/bot"
	"github.com/simba-fs/telegrary/note"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigtoml"
)

// Config is the type of config
type Config struct {
	Token string
	Root  string
}

var config Config

func init() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	loader := aconfig.LoaderFor(&config, aconfig.Config{
		SkipFlags: true,
		SkipEnv:   true,
		Files:     []string{path.Join(configDir, "telegrary.toml"), "telegrary.toml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".toml": aconfigtoml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}

	if config.Root == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		config.Root = path.Join(home, ".local", "share", "telegrary")
	}
}

func main() {
	log.Println(os.Args[1:], config)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "bot":
			log.Println("start bot")
			bot.Run(config.Token)
			return
		case "-h", "--help":
			fmt.Println("Usage: telegrary [bot | year month day]\n  configfile: ~/.config/telegrary.toml, ./telegrary.toml")
			return
		}
	}

	// convert date from string to int
	var date []int
	for _, v := range os.Args[1:] {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		date = append(date, i)
	}
	log.Println(date)

	year, month, day := time.Now().Date()
	switch len(date) {
	case 3:
		year, month, day = date[0], time.Month(date[1]), date[2]
	case 2:
		month, day = time.Month(date[0]), date[1]
	case 1:
		day = date[0]
	}

	note.Open(fmt.Sprintf("%s/%d/%d/%d.md", config.Root, year, month, day))
	log.Println(year, month, day)

}
