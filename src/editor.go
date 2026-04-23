package editor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v3"
	buffer "github.com/yepakh/go-editor/src/buffer"
	"github.com/yepakh/go-editor/src/cursor"
	"github.com/yepakh/go-editor/src/render"
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

func (ed *Editor) Start() chan struct{} {
	eventChan := render.InitScreen()

	ed.displayedBuffer = ed.loadedBuffers[0]
	render.RenderBuffer(ed.displayedBuffer, 0, 0)
	cursor.InitCursor(ed.displayedBuffer)

	quit := make(chan struct{})
	go ed.handleUserInput(eventChan, ed.displayedBuffer, quit)
	return quit
}

func (ed *Editor) Close() {
	render.CloseScreen()
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

func (ed *Editor) handleUserInput(evChan <-chan tcell.Event, buf *buffer.Buffer, quitCh chan struct{}) {
	for {
		event := <-evChan

		switch ev := event.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyRune {
				// Handle character
			} else if cursor.HandleCursorEvent(ev.Key(), buf) {
				continue
			} else if ev.Key() == tcell.KeyCtrlQ {
				if ed.handleCloseEvent() {
					close(quitCh)
				}
			}
		case *tcell.EventResize:
			render.Sync()
		}
	}
}

func (ed *Editor) handleCloseEvent() bool {
	for _, v := range ed.loadedBuffers {
		err := v.Close(false, false)

		if err != nil {
			//handle
		}
	}

	return true
}

func (ed *Editor) loadBuffer(filePath string) error {
	newBuff, err := buffer.Load(filePath)
	if err != nil {
		return err
	}

	ed.loadedBuffers = append(ed.loadedBuffers, newBuff)
	return nil
}
