package diary

import (
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	err := WriteNote("./diary/2022/01/15.md", "01/15")
	if err != nil {
		t.Error(err)
	}

	err = WriteNote("./diary/2022/02/15.md", "02/15")
	if err != nil {
		t.Error(err)
	}
}

func TestList(t *testing.T) {
	list, err := ListNotes("./diary/")
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}

func TestTree(t *testing.T) {
	tree, err := Tree("./diary/")
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestRead(t *testing.T) {
	note, err := ReadNote("./diary/2022/01/15.md")
	if err != nil {
		t.Error(err)
	}
	t.Log(note)

	if note != "01/15" {
		t.Error("content note error")
	}
}

func TestClean(t *testing.T) {
	os.RemoveAll("./diary")
}
