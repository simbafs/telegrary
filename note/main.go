package note

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// WritePost overwrites the note
func Write(path string, content string) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, []byte(content), 0644)
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
	cmd := exec.Command("tree", "-I", "LICENSE|Makefile|new.sh|README.md", prefix)
	tree, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(tree), nil
}
