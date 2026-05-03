package piecetable

import "testing"

func TestInsertChar(t *testing.T) {
	t.Run("Insert at the end of a line (appending)", func(t *testing.T) {
		pt := InitPieceTable("Hello")

		// Manually create an add piece to test append
		pt.InsertChar(0, 5, '!')
		assertLine(pt.lines[0], "Hello!", t)
		assertPiece(pt.lines[0], 1, false, 0, 1, t)
		assertAddBuf(pt.lines[0], "!", t)

		pt.InsertChar(0, 6, 'X')
		assertLine(pt.lines[0], "Hello!X", t)
		assertPiece(pt.lines[0], 1, false, 0, 2, t)
		assertAddBuf(pt.lines[0], "!X", t)
	})

	t.Run("Insert in the middle of an add piece", func(t *testing.T) {
		pt := InitPieceTable("Hello")
		pt.lines[0].add = []rune("ab")
		pt.lines[0].pieces = append(pt.lines[0].pieces, &PieceTableRecord{false, 0, 2})

		assertLine(pt.lines[0], "Helloab", t)

		pt.InsertChar(0, 6, 'X')

		assertLine(pt.lines[0], "HelloaXb", t)
		assertAddBuf(pt.lines[0], "abX", t)

		assertPiece(pt.lines[0], 0, true, 0, 5, t)
		assertPiece(pt.lines[0], 1, false, 0, 1, t)
		assertPiece(pt.lines[0], 2, false, 2, 1, t)
		assertPiece(pt.lines[0], 3, false, 1, 1, t)
	})

	t.Run("Insert in the middle of an orig piece", func(t *testing.T) {
		pt := InitPieceTable("Hello")
		pt.InsertChar(0, 2, 'X')

		assertLine(pt.lines[0], "HeXllo", t)
		assertAddBuf(pt.lines[0], "X", t)

		assertPiece(pt.lines[0], 0, true, 0, 2, t)
		assertPiece(pt.lines[0], 1, false, 0, 1, t)
		assertPiece(pt.lines[0], 2, true, 2, 3, t)
	})
}

func TestInsertNewLine(t *testing.T) {
	t.Run("Inser new line to the end of line", func(t *testing.T) {
		pt := InitPieceTable("Hello\nWorld!")
		assertLine(pt.lines[0], "Hello", t)
		assertLine(pt.lines[1], "World!", t)

		pt.InsertNewLine(0, 5)

		assertLine(pt.lines[0], "Hello", t)
		assertLine(pt.lines[1], "", t)
		assertLine(pt.lines[2], "World!", t)
	})

	t.Run("Inser new line in the middle of a piece", func(t *testing.T) {
		pt := InitPieceTable("HelloWorld!")
		assertLine(pt.lines[0], "HelloWorld!", t)

		pt.InsertNewLine(0, 5)

		assertLine(pt.lines[0], "Hello", t)
		assertLine(pt.lines[1], "World!", t)
	})
}
