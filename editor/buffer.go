package editor

import (
	"errors"
	"os"

	piecetable "github.com/yepakh/go-editor/piecetable"
)

var ChangesNotSaved = errors.New("cannot close file, changes not saved")

type Buffer struct {
	filepath          string
	hasUnsavedChanges bool
	Data              *piecetable.PieceTable
	Cursor            *Cursor
}

func InitBuffer(filePath string) (*Buffer, error) {
	if err := IsValidPathOrEmpty(filePath); err != nil && !errors.Is(err, ErrEmptyPath) {
		return nil, err
	}

	buff := Buffer{filePath, false, nil, nil}
	buff.load()

	renderChan := make(chan struct{})
	buff.Cursor = InitCursor(&buff, renderChan)

	go func() {
		for range renderChan {
			buff.RenderUpdates()
		}
	}()

	return &buff, nil
}

func (buff *Buffer) InserNewLine() {
	cursCurX, cursCurY := buff.Cursor.GetAbsoluteCursorCoords()
	buff.Data.InsertNewLine(cursCurY, cursCurX)

	buff.Cursor.SetCursorTo(0, cursCurY+1)
	lineOff, charOff := buff.Cursor.GetOffsets()
	RenderFromLine(buff.Data, cursCurY, lineOff, charOff)
}

func (buff *Buffer) InsertChar(char rune) {
	cursCurX, cursCurY := buff.Cursor.GetAbsoluteCursorCoords()
	buff.Data.InsertChar(cursCurY, cursCurX, char)

	buff.Cursor.MoveCursor(1, 0)
	lineOff, charOff := buff.Cursor.GetOffsets()
	ReRenderLine(buff.Data, cursCurY, lineOff, charOff)
}

func (buff *Buffer) RenderUpdates() {
	lineOff, charOff := buff.Cursor.GetOffsets()
	RenderBuffer(buff.Data, lineOff, charOff)
}

func (buff *Buffer) FullRender() {
	ClearScreen()
	lineOff, charOff := buff.Cursor.GetOffsets()
	RenderBuffer(buff.Data, lineOff, charOff)
	RenderFooter(buff.filepath)
}

func (buff *Buffer) Refresh() {
	// force refreshing cursor position
	curAbsX, curAbxY := buff.Cursor.GetAbsoluteCursorCoords()
	buff.Cursor.SetCursorTo(curAbsX, curAbxY)

	buff.FullRender()
}

func (buff *Buffer) GetFilepath() string {
	return buff.filepath
}

func (buff *Buffer) Close(force bool, create bool) error {
	// if not forced and unsaved - error

	// if file does not exist and not create - error
	// else if create - create dir and save

	return nil
}

func (buff *Buffer) load() {
	data, err := os.ReadFile(buff.filepath)
	if err != nil {
		return
	}

	buff.Data = piecetable.InitPieceTable(string(data))
}
