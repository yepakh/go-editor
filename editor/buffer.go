package editor

import (
	"errors"
	"os"

	piecetable "github.com/yepakh/go-editor/piece-table"
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
			buff.Render()
		}
	}()

	return &buff, nil
}

func (buff *Buffer) Render() {
	lineOff, charOff := buff.Cursor.GetOffsets()
	RenderBuffer(buff.Data, lineOff, charOff)
	RenderFooter(buff.filepath)
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
