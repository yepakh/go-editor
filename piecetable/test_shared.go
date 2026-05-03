package piecetable

import "testing"

func assertAddBuf(line *PieceTableLine, expected string, t *testing.T) {
	actual := string(line.add)
	if expected != actual {
		t.Errorf("Expected add buffer '%v', got '%v'", expected, actual)
	}
}

func assertLine(line *PieceTableLine, expected string, t *testing.T) {
	actual := string(line.getLine())
	if expected != actual {
		t.Errorf("Expected full line to be '%v', got '%v'", expected, actual)
	}
}

func assertPiece(line *PieceTableLine, pieceInd int, isOrig bool, expStInd, expLen int, t *testing.T) {
	piece := line.pieces[pieceInd]

	if piece.isOrig != isOrig {
		t.Errorf("Expected piece %v isOrig to be '%v', got '%v'", pieceInd, isOrig, piece.isOrig)
	}

	if piece.startInd != expStInd {
		t.Errorf("Expected piece %v start index to be '%v', got '%v'", pieceInd, expStInd, piece.startInd)
	}

	if piece.len != expLen {
		t.Errorf("Expected piece %v length to be '%v', got '%v'", pieceInd, expLen, piece.len)
	}
}
