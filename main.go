package main

import (
	"os"

	editor "github.com/yepakh/notepad/src"
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

	editor.Start()
	editor.Close()
}

func getPath() string {
	args := os.Args
	if len(args) < 1 {
		return ""
	}

	return args[1]
}
