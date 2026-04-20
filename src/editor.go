package editor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v3"
	buffer "github.com/yepakh/notepad/src/buffer"
	"github.com/yepakh/notepad/src/cursor"
	"github.com/yepakh/notepad/src/render"
)

type Editor struct {
	initDirectory   string
	loadedBuffers   []*buffer.Buffer
	displayedBuffer *buffer.Buffer
}

func (ed *Editor) Init() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Failed to get current directory: %v", err)
	}

	return ed.InitFromDir(cwd)
}

func (ed *Editor) Start() {
	eventChan := render.InitScreen()

	ed.displayedBuffer = ed.loadedBuffers[0]
	cursor.InitBuffer(ed.displayedBuffer)

	handleUserInput(eventChan, ed.displayedBuffer)
}

func (ed *Editor) Close() {
	for _, v := range ed.loadedBuffers {
		err := v.Close(false, false)

		if err != nil {
			//handle
		}
	}
}

func (ed *Editor) InitFromDir(dir string) error {
	ed.initDirectory = dir
	return ed.loadBuffer("")
}

func (ed *Editor) InitFromFile(filePath string) error {
	dir := filepath.Dir(filePath)

	ed.initDirectory = dir
	return ed.loadBuffer(filePath)
}

func handleUserInput(evChan <-chan tcell.Event, buf *buffer.Buffer) {
	for {
		event := <-evChan

		switch ev := event.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyRune {
				// Handle character
			} else if cursor.HandleCursorEvent(ev.Key(), buf) {
				continue
			} else if ev.Key() == tcell.KeyCancel {
				return
			}
		case *tcell.EventResize:
			render.Sync()
		}
	}
}

func (ed *Editor) loadBuffer(filePath string) error {
	newBuff, err := buffer.Load(filePath)
	if err != nil {
		return err
	}

	ed.loadedBuffers = append(ed.loadedBuffers, newBuff)
	return nil
}
