package main

import (
	"log"
	"os"

	"github.com/yepakh/go-editor/editor"
)

func main() {
	var ed *editor.Editor
	var err error
	path := getPath()

	if path == "" {
		ed, err = editor.InitEditor()
	} else if info, err := os.Stat(path); err == nil && info.IsDir() {
		ed, err = editor.InitEditor(editor.WithDirectory(path))
	} else {
		ed, err = editor.InitEditor(editor.WithFile(path))
	}

	if err != nil {
		log.Fatal(err)
	}

	quitCh := ed.Start()

	<-quitCh
	ed.Close()
}

func getPath() string {
	args := os.Args
	if len(args) < 2 {
		return ""
	}

	return args[1]
}
