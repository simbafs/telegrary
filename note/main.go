package note

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/simba-fs/telegrary/util"

	log "github.com/sirupsen/logrus"
)

func mkdir(path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	return nil
}

// Open use the default editor to open the note
func Open(path string) error {
	if err := mkdir(path); err != nil {
		return err
	}
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	log.Debugln(editor, path)
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// WritePost overwrites the note
func Write(path string, content string, overwrite bool) error {
	if err := mkdir(path); err != nil {
		return err
	}

	flag := os.O_CREATE | os.O_WRONLY
	if !overwrite {
		flag = flag | os.O_APPEND
	}

	f, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(content))
	return err
}

// Read returns the note
func Read(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

// List returns a list of notes with the given prefix
func List(prefix string) ([]string, error) {
	var notes []string
	err := filepath.Walk(prefix,
		func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() {
				notes = append(notes, path)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return notes, nil
}

// Delete deletes the note
func Delete(path string) error {
	err := os.Remove(path)
	return err
}

// Tree returns a list of notes in tree form with the given prefix
func Tree(prefix string) (string, error) {
	return util.Tree(prefix)
}
