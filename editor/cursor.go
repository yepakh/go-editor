package editor

import "github.com/yepakh/go-editor/piecetable"

type Cursor struct {
	charPos          int
	line             int
	lineOffset       int
	charOffset       int
	savedCharPos     int
	data             *piecetable.PieceTable
	renderBufChannel chan<- struct{}
}

func InitCursor(buffer *Buffer, renderBufChannel chan<- struct{}) *Cursor {
	return &Cursor{0, 0, 0, 0, 0, buffer.Data, renderBufChannel}
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
	if cursor.data.GetLineNum() == 0 {
		cursor.charPos = 0
		cursor.line = 0
		return true
	}

	absX, absY := cursor.GetAbsoluteCursorCoords()

	charChanged := targetX != absX
	if !charChanged {
		targetX = cursor.savedCharPos
	}

	lineCount := cursor.data.GetLineNum()
	if targetY >= lineCount {
		targetY = lineCount - 1
	} else if targetY < 0 {
		targetY = 0
	}

	targetLineLen := cursor.data.GetLineLen(targetY)
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
