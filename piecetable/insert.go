package piecetable

import (
	"slices"
)

func (pt *PieceTable) InsertNewLine(lineNum, pos int) {
	if pos == pt.GetLineLen(lineNum) {
		newLine := &PieceTableLine{make([]rune, 0), make([]rune, 0), make([]*PieceTableRecord, 0)}
		newLine.pieces = append(newLine.pieces, &PieceTableRecord{false, 0, 0})
		pt.lines = slices.Insert(pt.lines, lineNum+1, newLine)

		return
	}

	line := pt.lines[lineNum]
	piece, pieceInd, lastPieceCharPos := line.getPieceToUpdate(pos)

	// get all strings after pos for a new line
	newAddBuf := make([]rune, 0)
	if piece.isOrig {
		newAddBuf = append(newAddBuf, line.orig[pos:piece.startInd+piece.len]...)
	} else {
		newAddBuf = append(newAddBuf, line.add[pos:piece.startInd+piece.len]...)
	}
	piece.len -= lastPieceCharPos - pos

	if pieceInd < len(line.pieces)-1 {
		for _, p := range line.pieces[pieceInd+1:] {
			if p.isOrig {
				newAddBuf = append(newAddBuf, line.orig[p.startInd:p.startInd+p.len]...)
			} else {
				newAddBuf = append(newAddBuf, line.add[p.startInd:p.startInd+p.len]...)
			}

		}

		line.pieces = line.pieces[:pieceInd+1]
	}

	newLine := PieceTableLine{[]rune{}, []rune(newAddBuf), make([]*PieceTableRecord, 1)}
	newLine.pieces[0] = &PieceTableRecord{false, 0, len(newAddBuf)}
	pt.lines = slices.Insert(pt.lines, lineNum+1, &newLine)
}

func (pt *PieceTable) InsertChar(lineNum, pos int, char rune) {
	line := pt.lines[lineNum]

	// if inserting to the end of line + last buf is add - just append char and increase piece len
	// if insering to the end of line + last buf is orig - append char and create a new piece
	if pt.GetLineLen(lineNum) == pos {
		lastPiece := line.pieces[len(line.pieces)-1]
		line.add = append(line.add, char)

		if !lastPiece.isOrig {
			lastPiece.len++
			return
		}

		newPiece := PieceTableRecord{false, len(line.add) - 1, 1}
		line.pieces = append(line.pieces, &newPiece)
		return
	}

	// if insering in the middle of any piece - append char, split the piece, in the middle put a new add piece
	piece, pieceInd, lastPCharPos := line.getPieceToUpdate(pos)

	posInBuffer := piece.startInd + piece.len - (lastPCharPos - pos)
	piece.len -= (lastPCharPos - pos)
	line.add = append(line.add, char)

	addPiece := PieceTableRecord{false, len(line.add) - 1, 1}
	line.pieces = slices.Insert(line.pieces, pieceInd+1, &addPiece)

	newOrigPiece := PieceTableRecord{piece.isOrig, posInBuffer, lastPCharPos - pos}
	line.pieces = slices.Insert(line.pieces, pieceInd+2, &newOrigPiece)
}

func (ln *PieceTableLine) getPieceToUpdate(pos int) (piece *PieceTableRecord, pieceInd, lastPieceCharPos int) {
	lastPieceCharPos = 0

	for i, piece := range ln.pieces {
		lastPieceCharPos += piece.len
		if lastPieceCharPos <= pos {
			continue
		}

		return piece, i, lastPieceCharPos
	}

	return nil, 0, 0
}
