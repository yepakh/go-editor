package piecetable

import (
	"slices"
	"strings"
)

type PieceTableRecord struct {
	isOrig   bool
	startInd int
	len      int
}

type PieceTableLine struct {
	isDeleted bool
	orig      []rune
	add       []rune
	pieces    []*PieceTableRecord
}

type PieceTable struct {
	lines []*PieceTableLine
}

func InitPieceTable(text string) *PieceTable {
	lines := strings.Split(text, "\n")
	pt := PieceTable{make([]*PieceTableLine, len(lines))}
	for i, line := range lines {
		runes := []rune(line)
		ptLine := PieceTableLine{false, runes, make([]rune, 0), make([]*PieceTableRecord, 1)}
		ptLine.pieces[0] = &PieceTableRecord{true, 0, len(runes)}
		pt.lines[i] = &ptLine
	}

	return &pt
}

func (pt *PieceTable) InsertChar(lineNum, pos int, char rune) {
	line := pt.lines[lineNum]

	pieceEnd := 0
	lastAddBufInd := 0
	for i, piece := range line.pieces {
		// search skip until we find a buffer which can be modifed
		// e.g want to insert in pos 3, and buffer is 0-2
		// we can insert at the end of this buffer, if it's add buffer
		pieceEnd += piece.len
		if pieceEnd < pos {
			if !piece.isOrig {
				lastAddBufInd = piece.startInd + piece.len
			}
			continue
		}

		// if an 'add' piece and position is the next char of this piece, add to it
		if pieceEnd == pos && !piece.isOrig {
			// if adding to the end of 'add' - should append. If not - insert
			addBufPos := piece.startInd + piece.len
			piece.len++

			if addBufPos == len(line.add) {
				line.add = append(line.add, char)
				return
			}

			// if inserting - need to adjust next pieces
			line.add = slices.Insert(line.add, addBufPos, char)
			line.incrStartOfAddPieces(i)

			return
		}

		// if char is in the end of 'orig' piece - insert to the beginning of next 'add' piece
		if pieceEnd == pos && piece.isOrig {
			// if no more pieces after this one - create a new piece
			if i == len(line.pieces)-1 {
				line.add = append(line.add, char)
				newStInd := len(line.add) - 1
				line.pieces = append(line.pieces, &PieceTableRecord{false, newStInd, 1})
				return
			}

			nextPiece := line.pieces[i+1]
			// if piece exists and it's add - insert and modify it
			if !nextPiece.isOrig {
				nextPiece.len++
				line.add = slices.Insert(line.add, nextPiece.startInd, char)
				line.incrStartOfAddPieces(i + 1)
				return
			}

			// TODO: merge if next piece is from 'orig' buffer
		}

		// if char is in the middle of 'add' piece - insert and offset
		if !piece.isOrig {
			posInBuffer := piece.startInd + piece.len - (pieceEnd - pos)
			piece.len++
			// if inserting - need to adjust next buffers
			line.add = slices.Insert(line.add, posInBuffer, char)

			// move start index of all next 'add' slices by 1
			line.incrStartOfAddPieces(i)

			return
		}

		// if insert in the middle of 'orig' piece - split it and add 'add' piece
		if piece.isOrig {
			posInBuffer := piece.startInd + piece.len - (pieceEnd - pos)
			// modify orig piece
			piece.len -= (pieceEnd - pos)

			// add 'add' piece

			// if add buffer is finished - add to the end, else - insert
			if lastAddBufInd == len(line.add) {
				line.add = append(line.add, char)
			} else {
				line.add = slices.Insert(line.add, lastAddBufInd, char)
				line.incrStartOfAddPieces(i)
			}

			// is last piece - append, else - insert
			addPiece := PieceTableRecord{false, lastAddBufInd, 1}
			isLastPiece := false
			if i == len(line.pieces)-1 {
				line.pieces = append(line.pieces, &addPiece)
				isLastPiece = true
			} else {
				line.pieces = slices.Insert(line.pieces, i+1, &addPiece)
			}

			// add new 'orig' piece
			newOrigPiece := PieceTableRecord{true, posInBuffer, pieceEnd - pos}
			if isLastPiece {
				line.pieces = append(line.pieces, &newOrigPiece)
			} else {
				line.pieces = slices.Insert(line.pieces, i+2, &newOrigPiece)
			}

			return
		}
	}
}

func (pt *PieceTable) GetLineNum() int {
	return len(pt.lines)
}

func (pt *PieceTable) GetLineLen(lineNum int) int {
	sum := 0
	for _, p := range pt.lines[lineNum].pieces {
		sum += p.len
	}

	return sum
}

func (pt *PieceTable) GetLines(stInd, count int) [][]rune {
	resLn := make([][]rune, 0, count)
	for i := 0; i < count && i < len(pt.lines); i++ {
		resLn = append(resLn, pt.getLine(pt.lines[stInd+i]))
	}

	return resLn
}

func (line *PieceTableLine) incrStartOfAddPieces(stInd int) {
	if stInd == len(line.pieces)-1 {
		return
	}

	// move start index of all next 'add' pieces by 1
	for j := stInd + 1; j < len(line.pieces); j++ {
		pc := line.pieces[j]
		if pc.isOrig {
			continue
		}

		pc.startInd++
	}
}

func (pt *PieceTable) getLine(line *PieceTableLine) []rune {
	outputStr := make([]rune, 0)
	for _, v := range line.pieces {
		if v.isOrig {
			outputStr = append(outputStr, line.orig[v.startInd:v.startInd+v.len]...)
		} else {
			outputStr = append(outputStr, line.add[v.startInd:v.startInd+v.len]...)
		}
	}

	return outputStr
}
