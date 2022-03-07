package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"

	"github.com/simba-fs/telegrary/config"
	"github.com/simba-fs/telegrary/git"
	"github.com/simba-fs/telegrary/note"
	"github.com/simba-fs/telegrary/util"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigtoml"

	log "github.com/sirupsen/logrus"
)

var configPath []string

//go:embed help.txt
var helpText string

func init() {
	// get config path, e.g. ~/.config
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath = []string{
		path.Join(configDir, "telegrary.toml"),
		path.Join(".", "telegrary.toml"),
	}

	// load config
	loader := aconfig.LoaderFor(&config.Config, aconfig.Config{
		SkipFlags:          true,
		SkipEnv:            true,
		AllowUnknownFields: true,
		Files:              configPath,
		FileDecoders: map[string]aconfig.FileDecoder{
			".toml": aconfigtoml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}

	// set default diary root
	if config.Config.Root == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		config.Config.Root = path.Join(home, ".local", "share", "telegrary")
	}

	// set log level
	log.SetLevel(log.ErrorLevel)
	// log.SetLevel(log.DebugLevel)
}

func main() {
	log.Debugln(os.Args[1:], config.Config)

	if len(os.Args) > 1 {
		// parse flag
		switch os.Args[1] {
		case "bot":
			if config.Config.Token == "" {
				log.Fatal("token is required")
			}
			log.Debugln("start bot")
			startBot(config.Config.Token)
		case "config":
			for _, v := range configPath {
				if _, err := os.Stat(v); err == nil {
					note.Open(v)
					return
				}
			}
			note.Open(configPath[0])
			return // do not trigger git commit
		case "push":
			break
		case "-h", "--help", "help":
			fmt.Println(helpText)
			return // do not trigger git commit
		}
	} else {

		year, month, day := util.GetDate(os.Args[1:])

		note.Open(fmt.Sprintf("%s/%s/%s/%s.md", config.Config.Root, year, month, day))
		log.Debugln("open", year, month, day)
	}
	// git save
	if git.Commit() == nil {
		fmt.Println("commit notes")
	}

	if config.Config.GitRepo != "" {
		log.Debug("git push")
		if git.Push() == nil {
			fmt.Println("push notes")
		}
	}
}
