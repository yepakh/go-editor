package piecetable

import "slices"

func (pt *PieceTable) InsertChar(lineNum, pos int, char rune) {
	line := pt.lines[lineNum]

	// if inserting to the end of line + last buf is add - just append char and increase piece len
	// if insering to the end of line + last buf is orig - append char and create a new piece
	if pt.GetLineLen(lineNum) == pos {
		lastPiece := line.pieces[len(line.pieces)-1]
		line.add = append(line.add, char)

		if lastPiece.isOrig {
			newPiece := PieceTableRecord{false, len(line.add) - 1, 1}
			line.pieces = append(line.pieces, &newPiece)
		} else {
			lastPiece.len++
		}

		return
	}

	// if insering in the middle of any piece - append char, split the piece, in the middle put a new add piece
	piece, pieceInd, lastPCharPos := line.getPieceToUpdate(pos)

	posInBuffer := piece.startInd + piece.len - (lastPCharPos - pos)
	// modify found piece
	piece.len -= (lastPCharPos - pos)

	// add 'add' piece
	line.add = append(line.add, char)

	// is last piece - append, else - insert
	addPiece := PieceTableRecord{false, len(line.add) - 1, 1}
	newOrigPiece := PieceTableRecord{piece.isOrig, posInBuffer, lastPCharPos - pos}

	if pieceInd == len(line.pieces)-1 {
		line.pieces = append(line.pieces, &addPiece)
		line.pieces = append(line.pieces, &newOrigPiece)
	} else {
		line.pieces = slices.Insert(line.pieces, pieceInd+1, &addPiece)
		line.pieces = slices.Insert(line.pieces, pieceInd+2, &newOrigPiece)
	}
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
