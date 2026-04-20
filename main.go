package main

import (
	"os"

	editor "github.com/yepakh/go-editor/src"
)

func main() {
	path := getPath()
	editor := editor.Editor{}

	if path == "" {
		editor.Init()
	} else if info, err := os.Stat(path); err == nil && info.IsDir() {
		editor.InitFromDir(path)
	} else {
		editor.InitFromFile(path)
	}

	quitCh := editor.Start()

	<-quitCh
	editor.Close()
}

func getPath() string {
	args := os.Args
	if len(args) < 2 {
		return ""
	}

	return args[1]
}
