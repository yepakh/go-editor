package buffer

import (
	"errors"
	"os"
	"unicode/utf8"

	"github.com/yepakh/go-editor/src/cursor"
	"github.com/yepakh/go-editor/src/render"
	"github.com/yepakh/go-editor/src/utils"
)

type Buffer struct {
	filepath          string
	hasUnsavedChanges bool
	Lines             [][]rune
	Cursor            *cursor.Cursor
}

func Init(filePath string) (*Buffer, error) {
	if err := utils.IsValidPathOrEmpty(filePath); err != nil && !errors.Is(err, utils.ErrEmptyPath) {
		return nil, err
	}

	buff := Buffer{filePath, false, make([][]rune, 0), nil}
	buff.load()

	renderChan := make(chan struct{})
	buff.Cursor = cursor.InitCursor(&buff.Lines, renderChan)

	go func() {
		for range renderChan {
			buff.Render()
		}
	}()

	return &buff, nil
}

func (buff *Buffer) Render() {
	lineOff, charOff := buff.Cursor.GetOffsets()
	render.RenderBuffer(&buff.Lines, lineOff, charOff)
	render.RenderFooter(buff.filepath)
}

func (buff *Buffer) GetFilepath() string {
	return buff.filepath
}

func (buff *Buffer) GetMaxLines() int {
	return len(buff.Lines)
}

func (buff *Buffer) GetMaxCharsInLine(line int) int {
	return len(buff.Lines[line])
}

func (buff *Buffer) Close(force bool, create bool) error {
	// if not forced and unsaved - error

	// if file does not exist and not create - error
	// else if create - create dir and save

	return nil
}

func (buff *Buffer) load() {
	buff.Lines = append(buff.Lines, make([]rune, 0))

	data, err := os.ReadFile(buff.filepath)
	if err != nil {
		return
	}

	linePos := 0
	charPos := 0

	for i := 0; i < len(data); {
		r, size := utf8.DecodeRune(data[i:])

		if r == '\n' {
			buff.Lines = append(buff.Lines, make([]rune, 0))
			linePos++
		} else {
			buff.Lines[linePos] = append(buff.Lines[linePos], r)
			charPos++
		}

		i += size
	}
}
