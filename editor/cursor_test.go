package editor

import (
	"testing"
	"time"

	"github.com/gdamore/tcell/v3/vt"
)

func TestPreserveCursor(t *testing.T) {
	ed, _, mt := getTestEditor(longFile)
	ed.Start()
	validateCursorPosition(ed, 0, 0, t)

	for range 20 {
		mt.KeyPress(vt.KeyRight)
	}
	time.Sleep(defaultSleep)
	validateCursorPosition(ed, 20, 0, t)

	mt.KeyPress(vt.KeyDown)
	time.Sleep(defaultSleep)
	validateCursorPosition(ed, 13, 1, t)

	mt.KeyPress(vt.KeyDown)
	time.Sleep(defaultSleep)
	validateCursorPosition(ed, 20, 2, t)
}

func TestScroll(t *testing.T) {
	th := InitTheme()
	ed, _, mt := getTestEditor(simpleFile)
	mt.SetSize(vt.Coord{X: vt.Col(20), Y: vt.Row(5)})
	ed.Start()

	expect := make([]cellInfo, 0)
	expect = append(expect, getCellInfoFromString("    1 ", 0, 0, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line one", 6, 0, &th.contentStyle)...)
	expect = append(expect, getCellInfoFromString("    4 ", 0, 3, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line four", 6, 3, &th.contentStyle)...)

	validateTerm(mt, expect, t)
	validateCursorPosition(ed, 0, 0, t)

	for range 8 {
		mt.KeyPress(vt.KeyDown)
	}
	time.Sleep(defaultSleep)

	validateCursorPosition(ed, 0, 3, t)

	expect = make([]cellInfo, 0)
	expect = append(expect, getCellInfoFromString("    6 ", 0, 0, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line six", 6, 0, &th.contentStyle)...)
	expect = append(expect, getCellInfoFromString("    9 ", 0, 3, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line nine", 6, 3, &th.contentStyle)...)

	validateTerm(mt, expect, t)

	for range 4 {
		mt.KeyPress(vt.KeyUp)
	}
	time.Sleep(defaultSleep)

	expect = make([]cellInfo, 0)
	expect = append(expect, getCellInfoFromString("    5 ", 0, 0, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line five", 6, 0, &th.contentStyle)...)
	expect = append(expect, getCellInfoFromString("    8 ", 0, 3, &th.rightPanelStyle)...)
	expect = append(expect, getCellInfoFromString("line eight", 6, 3, &th.contentStyle)...)

	validateTerm(mt, expect, t)
}
