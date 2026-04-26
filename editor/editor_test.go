package editor

import (
	"path/filepath"
	"testing"

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
	cellsFoot := getCellInfoFromString(fullPath, 0, 23, &th.footerStyle)
	expect = append(expect, cellsFoot...)

	testInit(&testCase{simpleFile, expect}, t)
}

func testInit(tc *testCase, t *testing.T) {
	ed, _, mt := getTestEditor(simpleFile)

	ed.Start()

	mt.Drain()

	for _, v := range tc.expOut {
		cell := mt.GetCell(vt.Coord{X: v.X, Y: v.Y})
		if v.C != string(cell.C) {
			t.Errorf("Mismatch string at %d,%d: Actual: %q; Expected: %q", v.X, v.Y, string(cell.C), v.C)
		}
		if v.Fg != cell.S.Fg() {
			t.Errorf("Mismatch foreground at %d,%d: Actual: %s; Expected: %s", v.X, v.Y, cell.S.Fg().String(), v.Fg.String())
		}
		if v.Bg != cell.S.Bg() {
			t.Errorf("Mismatch background at %d,%d: Actual: %s; Expected: %s", v.X, v.Y, cell.S.Bg().String(), v.Bg.String())
		}
		if v.Attr != cell.S.Attr() {
			t.Errorf("Mismatch attr at %d,%d: Actual: %x; Expected: %x", v.X, v.Y, cell.S.Attr(), v.Attr)
		}
	}
}
