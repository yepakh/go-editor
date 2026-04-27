package editor

import (
	"testing"

	"github.com/gdamore/tcell/v3/vt"
)

func validateTerm(mt vt.MockTerm, expect []cellInfo, t *testing.T) {
	for _, v := range expect {
		cell := mt.GetCell(vt.Coord{X: v.X, Y: v.Y})
		if v.C != string(cell.C) {
			t.Errorf("Mismatch string at %d,%d: Actual: %q; Expected: %q", v.X, v.Y, string(cell.C), v.C)
		}
		if v.Fg != cell.S.Fg() {
			t.Errorf("Mismatch foreground at %d,%d: Actual: %s; Expected: %s", v.X, v.Y, cell.S.Fg().String(), v.Fg.String())
		}
		if v.Bg != cell.S.Bg() {
			t.Errorf("Mismatch background at %d,%d: Actual: %s; Expected: %s", v.X, v.Y, cell.S.Bg().String(), v.Bg.String())
		}
		if v.Attr != cell.S.Attr() {
			t.Errorf("Mismatch attr at %d,%d: Actual: %x; Expected: %x", v.X, v.Y, cell.S.Attr(), v.Attr)
		}
	}
}

func validateCursorPosition(ed *Editor, x, y int, t *testing.T) {
	curX, curY := ed.displayedBuffer.Cursor.getRelativeCursorCoords()
	if curX != x || curY != y {
		t.Errorf("Cursor is at the wrong position. Actual: %d, %d; Expected %d, %d", curX, curY, x, y)
	}
}
