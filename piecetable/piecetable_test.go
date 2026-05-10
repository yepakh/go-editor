package piecetable

import (
	"testing"
)

func TestInitPieceTable(t *testing.T) {
	text := "Hello\nWorld"
	pt := InitPieceTable(text)

	if pt.GetLineNum() != 2 {
		t.Errorf("Expected 2 lines, got %d", pt.GetLineNum())
	}

	assertLine(pt.lines[0], "Hello", t)
	assertLine(pt.lines[1], "World", t)
}

func TestGetLines(t *testing.T) {
	text := "Line1\nLine2\nLine3"
	pt := InitPieceTable(text)

	assertLine(pt.lines[0], "Line1", t)
	assertLine(pt.lines[1], "Line2", t)
	assertLine(pt.lines[2], "Line3", t)
}
