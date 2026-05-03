package editor

import "github.com/gdamore/tcell/v3"

var eventHandlers = map[tcell.Key]func(*Buffer, *tcell.EventKey){
	tcell.KeyUp:    func(buffer *Buffer, key *tcell.EventKey) { buffer.Cursor.MoveCursor(0, -1) },
	tcell.KeyDown:  func(buffer *Buffer, key *tcell.EventKey) { buffer.Cursor.MoveCursor(0, 1) },
	tcell.KeyLeft:  func(buffer *Buffer, key *tcell.EventKey) { buffer.Cursor.MoveCursor(-1, 0) },
	tcell.KeyRight: func(buffer *Buffer, key *tcell.EventKey) { buffer.Cursor.MoveCursor(1, 0) },

	tcell.KeyRune:  func(buffer *Buffer, key *tcell.EventKey) { buffer.InsertChar([]rune(key.Str())[0]) },
	tcell.KeyEnter: func(buffer *Buffer, key *tcell.EventKey) { buffer.InserNewLine() },
}

func (buf *Buffer) HandleBufferEvent(ev *tcell.EventKey) bool {
	key := ev.Key()
	handler, ok := eventHandlers[key]

	if !ok {
		return false
	}

	handler(buf, ev)
	return true
}
