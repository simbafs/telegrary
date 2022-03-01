package main

import (
	"os"
	"strconv"

	// "github.com/simba-fs/telegrary/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "telegrary",
	Short: "Telegrary is a diary manager with build in telegram bot",
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "bot",
		Short: "Start telegram bot",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("bot start")
		},
	}, &cobra.Command{
		Use:   "config",
		Short: "Edit config file",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("edit config")
		},
	})
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func edit(year, month, day int) error{
	log.Println("edit", year, month, day)
	return nil
}

func main() {
	if len(os.Args) > 1 {
		if _, err := strconv.Atoi(os.Args[1]); err == nil {
			edit(0,0,0)
		}else{
			execute()
		}
	} else {
		edit(0, 0, 0)
	}

}
