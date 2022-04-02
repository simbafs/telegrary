package git

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/simba-fs/telegrary/config"

	log "github.com/sirupsen/logrus"
)

// Commit call git commit with some checking
func Commit() error {
	// init
	cmd := exec.Command("git", "init")
	cmd.Dir = config.Config.Root
	cmd.Run()

	// add
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = config.Config.Root
	err := cmd.Run()
	if err != nil {
		return err
	}

	// status
	cmd = exec.Command("git", "status")
	cmd.Dir = config.Config.Root
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return err
	}

	// commit
	flag := "-m"
	if config.Config.GitSign {
		flag = "-sm"
	}
	cmd = exec.Command("git", "commit", flag, time.Now().Format("2006-01-02"))
	cmd.Dir = config.Config.Root
	return cmd.Run()
}

// Push push the notes to remote
func Push() error {
	// add remote
	cmd := exec.Command("git", "remote", "add", "origin", config.Config.GitRepo)
	cmd.Dir = config.Config.Root
	err := cmd.Run()
	log.Debug("git remote add origin", err)

	// push
	cmd = exec.Command("git", "push", "origin", "main")
	cmd.Dir = config.Config.Root
	err = cmd.Run()
	log.Debug("git push origin master", err)

	return err
}

// Save call Commit and Push in a single function with some checking
func Save() {
	// git save
	if Commit() == nil {
		fmt.Println("commit notes")
	}

	if config.Config.GitRepo != "" {
		log.Debug("git push")
		if Push() == nil {
			fmt.Println("push notes")
		}
	}
}

// Pull pull/clone the notes from remote
func Pull() error {
	if config.Config.GitRepo == "" {
		return nil
	}

	if _, err := os.Stat(config.Config.Root); os.IsNotExist(err) {
		// clone
		cmd := exec.Command("git", "clone", config.Config.GitRepo, config.Config.Root)
		cmd.Dir = path.Join(config.Config.Root, "..")
		return cmd.Run()
	} else {
		// pull
		cmd := exec.Command("git", "pull")
		cmd.Dir = config.Config.Root
		return cmd.Run()
	}
}
