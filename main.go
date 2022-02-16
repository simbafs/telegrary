package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	// "github.com/simba-fs/gitdiary/note"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigtoml"
)

// Config is the type of config
type Config struct {
	Token string
}

var config Config

func init() {
	loader := aconfig.LoaderFor(&config, aconfig.Config{
		SkipFlags: true,
		SkipEnv:   true,
		Files:     []string{"~/.config/gitdiary.toml", "gitdiary.toml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".toml": aconfigtoml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(os.Args[1:], config)

	// convert date from string to int
	var date []int
	for _, v := range os.Args[1:] {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		date = append(date, i)
	}
	fmt.Println(date)

	year, month, day := time.Now().Date()
	switch len(date) {
	case 3:
		year, month, day = date[0], time.Month(date[1]), date[2]
	case 2:
		month, day = time.Month(date[0]), date[1]
	case 1:
		day = date[0]
	}
	fmt.Println(year, month, day)
}
