package piecetable

import "testing"

func TestDeleteCharBefore_InvalidPos(t *testing.T) {
	pt := InitPieceTable("Hello")
	res := pt.DeleteCharBefore(0, 10)
	if res {
		t.Error("Expected DeleteCharBefore to return false for out of bounds charPos")
	}
	assertLine(pt.lines[0], "Hello", t)
}

func TestDeleteCharBefore_FirstLineStart(t *testing.T) {
	pt := InitPieceTable("Hello")
	res := pt.DeleteCharBefore(0, 0)
	if res {
		t.Error("Expected DeleteCharBefore to return false for first line at charPos 0")
	}
	assertLine(pt.lines[0], "Hello", t)
}

func TestDeleteCharBefore_MergeLines(t *testing.T) {
	pt := InitPieceTable("Hello\nWorld")
	res := pt.DeleteCharBefore(1, 0)
	if !res {
		t.Error("Expected DeleteCharBefore to return true for merging lines")
	}
	if pt.GetLineNum() != 1 {
		t.Errorf("Expected 1 line after merge, got %d", pt.GetLineNum())
	}
	assertLine(pt.lines[0], "HelloWorld", t)
}

func TestDeleteCharBefore_EndLine(t *testing.T) {
	pt := InitPieceTable("Hello")
	res := pt.DeleteCharBefore(0, 5)
	if res {
		t.Error("Expected DeleteCharBefore to return false (no line deleted)")
	}
	assertLine(pt.lines[0], "Hell", t)
}

func TestDeleteCharBefore_Middle(t *testing.T) {
	pt := InitPieceTable("Hello")
	// Delete 'e' at pos 2 (before pos 2 is 'e')
	res := pt.DeleteCharBefore(0, 2)
	if res {
		t.Error("Expected DeleteCharBefore to return false")
	}
	assertLine(pt.lines[0], "Hllo", t)
}

func TestDeleteCharBefore_EmptyPiece(t *testing.T) {
	pt := InitPieceTable("A")
	res := pt.DeleteCharBefore(0, 1)
	if res {
		t.Error("Expected DeleteCharBefore to return false")
	}
	assertLine(pt.lines[0], "", t)
	if len(pt.lines[0].pieces) != 0 {
		t.Errorf("Expected 0 pieces, got %d", len(pt.lines[0].pieces))
	}
}

func TestDeleteCharBefore_WithInsertions(t *testing.T) {
	pt := InitPieceTable("Hello")
	pt.InsertChar(0, 5, '!') // "Hello!"
	assertLine(pt.lines[0], "Hello!", t)
	
	res := pt.DeleteCharBefore(0, 6) // delete '!'
	if res {
		t.Error("Expected DeleteCharBefore to return false")
	}
	assertLine(pt.lines[0], "Hello", t)

	pt.InsertChar(0, 2, 'X') // "HeXllo"
	assertLine(pt.lines[0], "HeXllo", t)
	
	res = pt.DeleteCharBefore(0, 3) // delete 'X'
	if res {
		t.Error("Expected DeleteCharBefore to return false")
	}
	assertLine(pt.lines[0], "Hello", t)
}

func TestDeleteCharBefore_SequentialMerge(t *testing.T) {
	pt := InitPieceTable("A\nB\nC")
	// Merge B into A
	res := pt.DeleteCharBefore(1, 0)
	if !res {
		t.Error("Expected true")
	}
	if pt.GetLineNum() != 2 {
		t.Errorf("Expected 2 lines, got %d", pt.GetLineNum())
	}
	assertLine(pt.lines[0], "AB", t)
	assertLine(pt.lines[1], "C", t)

	// Merge C into AB
	res = pt.DeleteCharBefore(1, 0)
	if !res {
		t.Error("Expected true")
	}
	if pt.GetLineNum() != 1 {
		t.Errorf("Expected 1 line, got %d", pt.GetLineNum())
	}
	assertLine(pt.lines[0], "ABC", t)
}

func TestDeleteCharBefore_DeleteFirstCharOfPiece(t *testing.T) {
	pt := InitPieceTable("ABC")
	// Delete 'A' at pos 1
	res := pt.DeleteCharBefore(0, 1)
	if res {
		t.Error("Expected DeleteCharBefore to return false")
	}
	assertLine(pt.lines[0], "BC", t)
}

func TestDeleteCharBefore_Complex(t *testing.T) {
	pt := InitPieceTable("ABC")
	pt.InsertChar(0, 1, '1') // "A1BC"
	pt.InsertChar(0, 3, '2') // "A1B2C"
	assertLine(pt.lines[0], "A1B2C", t)

	// Pieces should be:
	// 0: Orig "A" (len 1)
	// 1: Add "1" (len 1)
	// 2: Orig "B" (len 1)
	// 3: Add "2" (len 1)
	// 4: Orig "C" (len 1)

	pt.DeleteCharBefore(0, 4) // Delete '2'
	assertLine(pt.lines[0], "A1BC", t)

	pt.DeleteCharBefore(0, 2) // Delete '1'
	assertLine(pt.lines[0], "ABC", t)
}
