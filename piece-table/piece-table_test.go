package piecetable

import (
	"reflect"
	"testing"
)

func TestInitPieceTable(t *testing.T) {
	text := "Hello\nWorld"
	pt := InitPieceTable(text)

	if pt.GetLineNum() != 2 {
		t.Errorf("Expected 2 lines, got %d", pt.GetLineNum())
	}

	if pt.GetLineLen(0) != 5 {
		t.Errorf("Expected line 0 length 5, got %d", pt.GetLineLen(0))
	}

	if pt.GetLineLen(1) != 5 {
		t.Errorf("Expected line 1 length 5, got %d", pt.GetLineLen(1))
	}
}

func TestGetLines(t *testing.T) {
	text := "Line1\nLine2\nLine3"
	pt := InitPieceTable(text)

	lines := pt.GetLines(0, 3)
	expected := [][]rune{
		[]rune("Line1"),
		[]rune("Line2"),
		[]rune("Line3"),
	}

	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("Expected %v, got %v", expected, lines)
	}
}
