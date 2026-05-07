package editor

import "github.com/gdamore/tcell/v3"

var eventHandlers = map[tcell.Key]func(*Buffer, *tcell.EventKey){
	tcell.KeyUp:    func(b *Buffer, ek *tcell.EventKey) { b.Cursor.MoveCursor(0, -1) },
	tcell.KeyDown:  func(b *Buffer, ek *tcell.EventKey) { b.Cursor.MoveCursor(0, 1) },
	tcell.KeyLeft:  func(b *Buffer, ek *tcell.EventKey) { b.Cursor.MoveCursor(-1, 0) },
	tcell.KeyRight: func(b *Buffer, ek *tcell.EventKey) { b.Cursor.MoveCursor(1, 0) },

	tcell.KeyRune:      func(b *Buffer, ek *tcell.EventKey) { b.InsertChar([]rune(ek.Str())[0]) },
	tcell.KeyEnter:     func(b *Buffer, ek *tcell.EventKey) { b.InserNewLine() },
	tcell.KeyBackspace: func(b *Buffer, ek *tcell.EventKey) { b.DeleteChar() },
}

func (buf *Buffer) HandleBufferEvent(ev *tcell.EventKey) bool {
	ek := ev.Key()
	handler, ok := eventHandlers[ek]

	if !ok {
		return false
	}

	handler(buf, ev)
	return true
}
