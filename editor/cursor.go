package editor

import (
	"github.com/gdamore/tcell/v3"
)

type Cursor struct {
	charPos          int
	line             int
	lineOffset       int
	charOffset       int
	savedCharPos     int
	buffer           *Buffer
	renderBufChannel chan<- struct{}
}

func InitCursor(buffer *Buffer, renderBufChannel chan<- struct{}) *Cursor {
	return &Cursor{0, 0, 0, 0, 0, buffer, renderBufChannel}
}

var eventHandlers = map[tcell.Key]func(*Cursor){
	tcell.KeyUp:    func(cursor *Cursor) { cursor.MoveCursor(0, -1) },
	tcell.KeyDown:  func(cursor *Cursor) { cursor.MoveCursor(0, 1) },
	tcell.KeyLeft:  func(cursor *Cursor) { cursor.MoveCursor(-1, 0) },
	tcell.KeyRight: func(cursor *Cursor) { cursor.MoveCursor(1, 0) },
}

func (cursor *Cursor) HandleCursorEvent(key tcell.Key) bool {
	handler, ok := eventHandlers[key]

	if !ok {
		return false
	}

	handler(cursor)
	return true
}

func (cursor *Cursor) GetAbsoluteCursorCoords() (charPos, line int) {
	return cursor.charPos, cursor.line
}

func (cursor *Cursor) MoveCursor(x int, y int) {
	currX, currY := cursor.GetAbsoluteCursorCoords()
	cursor.SetCursorTo(currX+x, currY+y)
}

func (cursor *Cursor) SetCursorTo(targetX, targetY int) {
	if !cursor.setPosition(targetX, targetY) {
		return
	}

	cursor.renderCursor(false)
}

func (cursor *Cursor) RefreshCursor() {
	cursor.renderCursor(true)
}

func (cursor *Cursor) GetOffsets() (lineOff, charOff int) {
	return cursor.lineOffset, cursor.charOffset
}

func (cursor *Cursor) renderCursor(renderBuf bool) {
	scrW, scrH := GetContentSceenSize()

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
	SetRenderCursor(scrX, scrY)

	if renderBuf {
		cursor.renderBufChannel <- struct{}{}
	}
}

func (cursor *Cursor) setPosition(targetX, targetY int) bool {
	if cursor.buffer.Data.GetLineNum() == 0 {
		cursor.charPos = 0
		cursor.line = 0
		return true
	}

	absX, absY := cursor.GetAbsoluteCursorCoords()

	charChanged := targetX != absX
	if !charChanged {
		targetX = cursor.savedCharPos
	}

	lineCount := cursor.buffer.Data.GetLineNum()
	if targetY >= lineCount {
		targetY = lineCount - 1
	} else if targetY < 0 {
		targetY = 0
	}

	targetLineLen := cursor.buffer.Data.GetLineLen(targetY)
	if targetX > targetLineLen && targetLineLen > 0 {
		targetX = targetLineLen
	} else if targetX < 0 || targetLineLen == 0 {
		targetX = 0
	}

	if targetX == absX && targetY == absY {
		return false
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
