package editor

import (
	"path/filepath"

	"github.com/gdamore/tcell/v3"
	"github.com/yepakh/go-editor/render"
)

type Editor struct {
	screen          tcell.Screen
	render          *render.Render
	screenEventCh   <-chan tcell.Event
	initDirectory   string
	loadedBuffers   []*Buffer
	displayedBuffer *Buffer
}

type Option func(*Editor) error

func InitEditor(options ...Option) (*Editor, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	ed := Editor{}
	ed.screen = screen

	r, eventChan := render.InitRenderScreen(ed.screen)
	ed.render = r
	ed.screenEventCh = eventChan

	for _, opt := range options {
		err := opt(&ed)
		if err != nil {
			return nil, err
		}
	}

	return &ed, nil
}

func WithDirectory(dir string) Option {
	return func(ed *Editor) error {
		ed.initDirectory = dir

		err := ed.loadBuffer("")
		if err != nil {
			return err
		}

		return nil
	}
}
func WithFile(filePath string) Option {
	return func(ed *Editor) error {
		dir := filepath.Dir(filePath)
		absPath, _ := filepath.Abs(filePath)

		ed.initDirectory = dir
		err := ed.loadBuffer(absPath)
		if err != nil {
			return err
		}

		return nil
	}
}

func (ed *Editor) Start() chan struct{} {
	ed.displayedBuffer = ed.loadedBuffers[0]
	ed.displayedBuffer.FullRender()

	quit := make(chan struct{})
	go ed.handleUserInput(ed.screenEventCh, quit)
	return quit
}

func (ed *Editor) Close() {
	ed.render.CloseRenderScreen()
}

func (ed *Editor) handleUserInput(evChan <-chan tcell.Event, quitCh chan struct{}) {
	for {
		event := <-evChan

		switch ev := event.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyCtrlQ:
				if ed.handleCloseEvent() {
					close(quitCh)
				}
			default:
				ed.displayedBuffer.HandleBufferEvent(ev)
			}
		case *tcell.EventResize:
			ed.handleResizeEvent()
		}
	}
}

func (ed *Editor) handleResizeEvent() {
	ed.render.RenderSync()
	ed.displayedBuffer.Refresh()
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
	newBuff, err := InitBuffer(filePath, ed.render)
	if err != nil {
		return err
	}

	ed.loadedBuffers = append(ed.loadedBuffers, newBuff)
	return nil
}
