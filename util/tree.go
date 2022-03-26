package util

import (
	"github.com/xlab/treeprint"
	"io/ioutil"
	"path"
)

// list adds nodes and branches to tree recursively
func list(current string, tree treeprint.Tree) error {
	files, err := ioutil.ReadDir(current)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.Name()[0] == '.' {
			continue
		}
		if file.IsDir() {
			list(path.Join(current, file.Name()), tree.AddBranch(file.Name()))
		} else {
			tree.AddNode(file.Name())
		}
	}
	return nil
}

// Tree is the implementation of linux command tree
func Tree(path string) (string, error) {
	tree := treeprint.New()
	err := list(path, tree)
	if err != nil {
		return "", err
	}
	return tree.String(), nil
}
