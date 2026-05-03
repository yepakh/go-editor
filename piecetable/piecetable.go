package piecetable

import (
	"strings"
)

type PieceTableRecord struct {
	isOrig   bool
	startInd int
	len      int
}

type PieceTableLine struct {
	orig   []rune
	add    []rune
	pieces []*PieceTableRecord
}

type PieceTable struct {
	lines []*PieceTableLine
}

func InitPieceTable(text string) *PieceTable {
	lines := strings.Split(text, "\n")
	pt := PieceTable{make([]*PieceTableLine, len(lines))}
	for i, line := range lines {
		runes := []rune(line)
		ptLine := PieceTableLine{runes, make([]rune, 0), make([]*PieceTableRecord, 1)}
		ptLine.pieces[0] = &PieceTableRecord{true, 0, len(runes)}
		pt.lines[i] = &ptLine
	}

	return &pt
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
