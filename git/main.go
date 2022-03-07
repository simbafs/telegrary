package git

import (
	"github.com/simba-fs/telegrary/config"
	log "github.com/sirupsen/logrus"

	"os/exec"
	"time"
)

func Commit() error {
	path := config.Config.Root

	// init
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	cmd.Run()

	// add
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = path
	err := cmd.Run()
	if err != nil {
		return err
	}

	// commit
	flag := "-m"
	if config.Config.GitSign {
		flag = "-sm"
	}
	cmd = exec.Command("git", "commit", flag, time.Now().Format("2006-01-02"))
	cmd.Dir = path
	return cmd.Run()
}

func Push() error {
	path, repo := config.Config.Root, config.Config.GitRepo

	// add remote
	cmd := exec.Command("git", "remote", "add", "origin", repo)
	cmd.Dir = path
	err := cmd.Run()
	log.Debug("git remote add origin", err)

	// push
	cmd = exec.Command("git", "push", "origin", "main")
	cmd.Dir = path
	err = cmd.Run()
	log.Debug("git push origin master", err)

	return err 
}
