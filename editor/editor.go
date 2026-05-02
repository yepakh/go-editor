package editor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v3"
)

type Editor struct {
	screen          *tcell.Screen
	initDirectory   string
	loadedBuffers   []*Buffer
	displayedBuffer *Buffer
}

func Init(screen *tcell.Screen) (*Editor, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Failed to get current directory: %v", err)
	}

	return InitFromDir(screen, cwd)
}

func InitFromDir(screen *tcell.Screen, dir string) (*Editor, error) {
	ed := Editor{}
	ed.screen = screen
	ed.initDirectory = dir

	err := ed.loadBuffer("")
	if err != nil {
		return nil, err
	}

	return &ed, nil
}

func InitFromFile(screen *tcell.Screen, filePath string) (*Editor, error) {
	dir := filepath.Dir(filePath)
	absPath, _ := filepath.Abs(filePath)

	ed := Editor{}
	ed.screen = screen
	ed.initDirectory = dir

	err := ed.loadBuffer(absPath)
	if err != nil {
		return nil, err
	}

	return &ed, nil
}

func (ed *Editor) Start() chan struct{} {
	eventChan := InitRenderScreen(*ed.screen)

	ed.displayedBuffer = ed.loadedBuffers[0]
	ed.displayedBuffer.Render()

	quit := make(chan struct{})
	go ed.handleUserInput(eventChan, ed.displayedBuffer, quit)
	return quit
}

func (ed *Editor) Close() {
	CloseRenderScreen()
}

func (ed *Editor) handleUserInput(evChan <-chan tcell.Event, buf *Buffer, quitCh chan struct{}) {
	for {
		event := <-evChan

		switch ev := event.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyRune {
				ed.displayedBuffer.InsertChar([]rune(ev.Str())[0])
			} else if buf.Cursor.HandleCursorEvent(ev.Key()) {
				continue
			} else if ev.Key() == tcell.KeyCtrlQ {
				if ed.handleCloseEvent() {
					close(quitCh)
				}
			}
		case *tcell.EventResize:
			ed.handleResizeEvent()
		}
	}
}

func (ed *Editor) handleResizeEvent() {
	RenderSync()
	ed.displayedBuffer.Cursor.RefreshCursor()
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
	newBuff, err := InitBuffer(filePath)
	if err != nil {
		return err
	}

	ed.loadedBuffers = append(ed.loadedBuffers, newBuff)
	return nil
}
