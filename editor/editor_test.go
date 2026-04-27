package editor

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/gdamore/tcell/v3/vt"
)

type testCase struct {
	fileName string
	expOut   []cellInfo
}

func TestInitialization(t *testing.T) {
	th := InitTheme()

	expect := make([]cellInfo, 0)
	expect = append(expect, getCellInfoFromString("    1 ", 0, 0, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line one", 6, 0, &th.contentStyle)...)
	expect = append(expect, getCellInfoFromString("    2 ", 0, 1, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line two", 6, 1, &th.contentStyle)...)
	expect = append(expect, getCellInfoFromString("   10 ", 0, 9, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line ten", 6, 9, &th.contentStyle)...)

	fullPath, _ := filepath.Abs(simpleFile)
	expect = append(expect, getCellInfoFromString(fullPath, 0, 23, &th.footerStyle)...)

	ed, _, mt := getTestEditor(simpleFile)
	ed.Start()
	mt.Drain()

	validateTerm(mt, expect, t)
}

func TestResizeAndScroll(t *testing.T) {
	th := InitTheme()
	ed, sc, mt := getTestEditor(simpleFile)
	ed.Start()

	currW, currH := sc.Size()
	mtSize := mt.Backend().GetSize()
	if currW != int(mtSize.X) || currH != int(mtSize.Y) {
		t.Errorf("Incorrect init screen size. Actual: %d, %d; Expected %v, %v", currW, currH, mtSize.X, mtSize.Y)
	}

	newW, newH := 20, 5
	mt.SetSize(vt.Coord{X: vt.Col(newW), Y: vt.Row(newH)})
	validateCursorPosition(ed, 0, 0, t)

	time.Sleep(defaultSleep)
	mt.Drain()

	currW, currH = sc.Size()
	if currW != newW || currH != newH {
		t.Errorf("Incorrect screen size after resize. Actual: %d, %d; Expected %v, %v", currW, currH, newW, newH)
	}

	expect := make([]cellInfo, 0)
	expect = append(expect, getCellInfoFromString("    1 ", 0, 0, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line one", 6, 0, &th.contentStyle)...)
	expect = append(expect, getCellInfoFromString("    2 ", 0, 1, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line two", 6, 1, &th.contentStyle)...)
	expect = append(expect, getCellInfoFromString("    4 ", 0, 3, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line four", 6, 3, &th.contentStyle)...)

	fullPath, _ := filepath.Abs(simpleFile)
	expect = append(expect, getCellInfoFromString(string(fullPath[:20]), 0, 4, &th.footerStyle)...)

	validateTerm(mt, expect, t)
}
