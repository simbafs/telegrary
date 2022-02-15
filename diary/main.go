package diary

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// WritePost overwrites the note
func WriteNote(path string, content string) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, []byte(content), 0644)
}

// ReadNote returns the note
func ReadNote(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

// ListNotes returns a list of notes with the given prefix
func ListNotes(prefix string) ([]string, error) {
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

// DeleteNote deletes the note
func DeleteNote(path string) error {
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
