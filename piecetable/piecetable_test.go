package piecetable

import (
	"reflect"
	"testing"
)

func TestInsertChar(t *testing.T) {
	t.Run("Insert at the end of an add piece (appending)", func(t *testing.T) {
		pt := InitPieceTable("Hello")
		// Manually create an add piece to test append
		pt.lines[0].add = []rune("!")
		pt.lines[0].pieces = append(pt.lines[0].pieces, &PieceTableRecord{false, 0, 1})

		pt.InsertChar(0, 6, 'X')
		expected := "Hello!X"
		actual := string(pt.GetLines(0, 1)[0])
		if actual != expected {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	})

	t.Run("Insert in the middle of an add piece", func(t *testing.T) {
		pt := InitPieceTable("Hello")
		// Manually create an add piece
		pt.lines[0].add = []rune("ab")
		pt.lines[0].pieces = append(pt.lines[0].pieces, &PieceTableRecord{false, 0, 2})

		pt.InsertChar(0, 6, 'X')

		expected := "HelloaXb"
		actual := string(pt.GetLines(0, 1)[0])
		if actual != expected {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	})

	t.Run("Insert in the middle of an orig piece", func(t *testing.T) {
		pt := InitPieceTable("Hello")
		pt.InsertChar(0, 2, 'X')
		expected := "HeXllo"
		actual := string(pt.GetLines(0, 1)[0])
		if actual != expected {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	})

	t.Run("Insert in the middle of orig piece - append char to add", func(t *testing.T) {
		pt := InitPieceTable("Hello World!")
		// Simulate state: "Hello nice World!"
		pt.lines[0].add = []rune(" nice")
		pt.lines[0].pieces = []*PieceTableRecord{
			{true, 0, 5},
			{false, 0, 5},
			{true, 5, 7},
		}

		pt.InsertChar(0, 16, 'X')
		expected := "Hello nice WorldX!"
		actual := string(pt.GetLines(0, 1)[0])
		if actual != expected {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	})

	t.Run("Insert in the middle of orig piece - next add pieces should move", func(t *testing.T) {
		pt := InitPieceTable("Hello World!")
		// Simulate state: "Hello nice World! Bye!"
		pt.lines[0].add = []rune(" nice Bye!")
		pt.lines[0].pieces = []*PieceTableRecord{
			{true, 0, 5},
			{false, 0, 5},
			{true, 5, 7},
			{false, 5, 5},
		}

		pt.InsertChar(0, 13, 'X')
		expected := "Hello nice WoXrld! Bye!"
		actual := string(pt.GetLines(0, 1)[0])
		if actual != expected {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	})

	t.Run("Insert in the middle of add piece - next add pieces should move", func(t *testing.T) {
		pt := InitPieceTable("Hello World!")
		// Simulate state: "Hello nice World! Bye!"
		pt.lines[0].add = []rune(" nice Bye!")
		pt.lines[0].pieces = []*PieceTableRecord{
			{true, 0, 5},
			{false, 0, 5},
			{true, 5, 7},
			{false, 5, 5},
		}

		pt.InsertChar(0, 8, 'X')
		expected := "Hello niXce World! Bye!"
		actual := string(pt.GetLines(0, 1)[0])
		if actual != expected {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	})

	t.Run("Append to the end of add piece - next add pieces should move", func(t *testing.T) {
		pt := InitPieceTable("Hello World!")
		// Simulate state: "Hello nice World! Bye!"
		pt.lines[0].add = []rune(" nice Bye!")
		pt.lines[0].pieces = []*PieceTableRecord{
			{true, 0, 5},
			{false, 0, 5},
			{true, 5, 7},
			{false, 5, 5},
		}

		pt.InsertChar(0, 10, 'X')
		expected := "Hello niceX World! Bye!"
		actual := string(pt.GetLines(0, 1)[0])
		if actual != expected {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	})
}

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
