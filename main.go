package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/simba-fs/telegrary/config"
	"github.com/simba-fs/telegrary/git"
	"github.com/simba-fs/telegrary/note"
	"github.com/simba-fs/telegrary/util"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigtoml"

	log "github.com/sirupsen/logrus"
)

var (
	//go:embed COMMIT
	commit  string
	//go:embed VERSION
	version string
	date    string = "unknown"
	builtBy string = "selfbuild"
)

var configPath []string

//go:embed help/cli.txt
var helpText string

func init() {
	// get config path, e.g. ~/.config
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath = []string{
		path.Join(".", "telegrary.toml"),
		path.Join(configDir, "telegrary.toml"),
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
	if config.Config.Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}
	
	// formet buildinfo
	version = strings.Trim(version, "\n")
	commit = strings.Trim(commit, "\n")
}

func main() {
	log.Debugln(os.Args[1:], config.Config)

	if len(os.Args) > 1 {
		// parse flag
		switch os.Args[1] {
		case "hash":
			if len(os.Args) < 2 {
				return
			}
			fmt.Println(util.Hash(os.Args[2]))
			return
		case "bot":
			git.Pull()
			if config.Config.Token == "" {
				log.Fatal("token is required")
			}
			log.Debugln("start bot")
			startBot()
			return
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
			git.Push()
			break
		case "-v", "--version", "version":
			fmt.Printf("Builder: %s\n", builtBy)
			fmt.Printf("BuildTime: %s\n", date)
			fmt.Printf("GitCommit: %s\n", commit)
			fmt.Printf("Version: %s\n", version)
			return
		case "-h", "--help", "help":
			fmt.Println(helpText)
			return // do not trigger git commit
		}
	}
	git.Pull()
	year, month, day := util.GetDate(os.Args[1:])

	note.Open(util.Path(year, month, day))
	log.Debugln("open", year, month, day)

	git.Save()
}
