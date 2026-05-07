
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
