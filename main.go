package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v3"
	"github.com/yepakh/go-editor/editor"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	var ed *editor.Editor
	path := getPath()

	if path == "" {
		ed, err = editor.Init(&screen)
	} else if info, err := os.Stat(path); err == nil && info.IsDir() {
		ed, err = editor.InitFromDir(&screen, path)
	} else {
		ed, err = editor.InitFromFile(&screen, path)
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
