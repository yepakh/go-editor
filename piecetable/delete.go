package piecetable

import "slices"

func (pt *PieceTable) DeleteCharBefore(lineNum, charPos int) (lineDel bool) {
	line := pt.lines[lineNum]
	lineLen := line.getLength()

	if charPos > lineLen {
		return false
	}

	if charPos == 0 {
		if lineNum == 0 {
			return false
		}

		prevLine := pt.lines[lineNum-1]
		newPieceStart := len(prevLine.add)

		currLineBuf := line.getLine()
		prevLine.add = append(prevLine.add, currLineBuf...)

		newPiece := PieceTableRecord{false, newPieceStart, len(currLineBuf)}
		prevLine.pieces = append(prevLine.pieces, &newPiece)
		pt.lines = slices.Delete(pt.lines, lineNum, lineNum+1)
		return true
	}

	piece, pi, nextPieceStart := line.getPieceToUpdate(charPos)

	// if cursor before piece end - just reduce the len
	if charPos == nextPieceStart {
		piece.len--
		if piece.len == 0 {
			line.pieces = slices.Delete(line.pieces, pi, pi+1)
		}
		return false
	}

	// if cursor in the middle of a piece - spit and reduce len
	newPieceLen := nextPieceStart - charPos
	newPiece := PieceTableRecord{piece.isOrig, piece.startInd + piece.len - newPieceLen, newPieceLen}
	piece.len = piece.len - newPieceLen - 1
	if piece.len == 0 {
		line.pieces = slices.Delete(line.pieces, pi, pi+1)
	}
	line.pieces = slices.Insert(line.pieces, pi+1, &newPiece)
	return false
}
