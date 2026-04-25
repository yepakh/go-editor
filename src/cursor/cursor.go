package cursor

import (
	"github.com/gdamore/tcell/v3"
	"github.com/yepakh/go-editor/src/buffer"
	"github.com/yepakh/go-editor/src/render"
)

type Cursor struct {
	charPos      int
	line         int
	lineOffset   int
	charOffset   int
	savedCharPos int
}

var ActiveCursor = Cursor{0, 0, 0, 0, 0}

func (cursor *Cursor) GetAbsoluteCursorCoords() (charPos, line int) {
	return cursor.charPos, cursor.line
}

func (cursor *Cursor) MoveCursor(x int, y int, buf *buffer.Buffer) {
	currX, currY := cursor.GetAbsoluteCursorCoords()
	cursor.SetCursorTo(currX+x, currY+y, buf)
}

func (cursor *Cursor) SetCursorTo(targetX, targetY int, buf *buffer.Buffer) {
	if !cursor.setPosition(targetX, targetY, buf) {
		return
	}

	cursor.renderCursor(buf, false)
}

func (cursor *Cursor) RefreshCursor(buf *buffer.Buffer) {
	cursor.renderCursor(buf, true)
}

func (cursor *Cursor) renderCursor(buf *buffer.Buffer, renderBuf bool) {
	scrW, scrH := render.GetBufferSceenSize()

	scrMinY, scrMaxY := cursor.lineOffset, cursor.lineOffset+scrH-1
	if cursor.line < scrMinY {
		cursor.lineOffset = cursor.line
		renderBuf = true
	} else if cursor.line > scrMaxY {
		cursor.lineOffset += cursor.line - scrMaxY
		renderBuf = true
	}

	scrMinX, scrMaxX := cursor.charOffset, cursor.charOffset+scrW-1
	if cursor.charPos < scrMinX {
		cursor.charOffset = cursor.charPos
		renderBuf = true
	} else if cursor.charPos > scrMaxX {
		cursor.charOffset += cursor.charPos - scrMaxX
		renderBuf = true
	}

	scrX, scrY := cursor.getRelativeCursorCoords()
	render.SetCursor(scrX, scrY)

	if renderBuf {
		render.RenderBuffer(buf, cursor.lineOffset, cursor.charOffset)
	}
}

func (cursor *Cursor) setPosition(targetX, targetY int, buf *buffer.Buffer) bool {
	if len(buf.Lines) == 0 {
		cursor.charPos = 0
		cursor.line = 0
		return true
	}

	if targetY >= len(buf.Lines) {
		targetY = len(buf.Lines) - 1
	} else if targetY < 0 {
		targetY = 0
	}

	if targetX >= len(buf.Lines[targetY]) && len(buf.Lines[targetY]) > 0 {
		targetX = len(buf.Lines[targetY]) - 1
	} else if targetX < 0 || len(buf.Lines[targetY]) == 0 {
		targetX = 0
	}

	absX, absY := cursor.GetAbsoluteCursorCoords()
	if targetX == absX && targetY == absY {
		return false
	}

	charChanged := targetX == absX
	if !charChanged {
		targetX = cursor.savedCharPos
	}

	cursor.charPos = targetX
	if charChanged {
		cursor.savedCharPos = targetX
	}
	cursor.line = targetY
	return true
}

func (cursor *Cursor) getRelativeCursorCoords() (x, y int) {
	return cursor.charPos - cursor.charOffset, cursor.line - cursor.lineOffset
}

var eventHandlers = map[tcell.Key]func(*buffer.Buffer){
	tcell.KeyUp:    func(buf *buffer.Buffer) { ActiveCursor.MoveCursor(0, -1, buf) },
	tcell.KeyDown:  func(buf *buffer.Buffer) { ActiveCursor.MoveCursor(0, 1, buf) },
	tcell.KeyLeft:  func(buf *buffer.Buffer) { ActiveCursor.MoveCursor(-1, 0, buf) },
	tcell.KeyRight: func(buf *buffer.Buffer) { ActiveCursor.MoveCursor(1, 0, buf) },
}

func HandleCursorEvent(key tcell.Key, buf *buffer.Buffer) bool {
	handler, ok := eventHandlers[key]

	if !ok {
		return false
	}

	handler(buf)
	return true
}

func InitCursor(buf *buffer.Buffer) {
	ActiveCursor.MoveCursor(0, 0, buf)
}
