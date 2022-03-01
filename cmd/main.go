package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/simba-fs/telegrary/note"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigtoml"

	log "github.com/sirupsen/logrus"
)

// Config is the type of config
type Config struct {
	Token string
	Root  string
}

var config Config

var configPath []string

// go:embed help.txt
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
	loader := aconfig.LoaderFor(&config, aconfig.Config{
		SkipFlags: true,
		SkipEnv:   true,
		Files:     configPath,
		FileDecoders: map[string]aconfig.FileDecoder{
			".toml": aconfigtoml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}

	// set default diary root
	if config.Root == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		config.Root = path.Join(home, ".local", "share", "telegrary")
	}

	// set log level
	log.SetLevel(log.ErrorLevel)
}

func getDate(raw []string) (int, int, int) {
	// convert date from string to int
	var date []int
	for _, v := range raw {
		i, err := strconv.Atoi(v)
		if err != nil {
			break
		}
		date = append(date, i)
	}
	log.Debugln(date)

	year, month, day := time.Now().Date()
	switch len(date) {
	case 3:
		year, month, day = date[0], time.Month(date[1]), date[2]
	case 2:
		month, day = time.Month(date[0]), date[1]
	case 1:
		day = date[0]
	}

	return year, int(month), day
}

func main() {
	log.Debugln(os.Args[1:], config)

	if len(os.Args) > 1 {
		// parse flag

		switch os.Args[1] {
		case "bot":
			if config.Token == "" {
				log.Fatal("token is required")
			}
			log.Debugln("start bot")
			startBot(config.Token)
		case "config":
			for _, v := range configPath {
				if _, err := os.Stat(v); err == nil {
					note.Open(v)
					return
				}
			}
		case "-h", "--help", "help":
			fmt.Println(helpText)
		}
		return
	}

	year, month, day := getDate(os.Args[1:])

	note.Open(fmt.Sprintf("%s/%d/%d/%d.md", config.Root, year, month, day))
	log.Debugln(year, month, day)

}
