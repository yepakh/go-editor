package editor

import (
	"testing"
	"time"

	"github.com/gdamore/tcell/v3/vt"
)

func TestBufferModifications(t *testing.T) {
	th := InitTheme()

	t.Run("InsertChar rendering", func(t *testing.T) {
		ed, _, mt := getTestEditor(simpleFile)
		ed.Start()
		time.Sleep(defaultSleep)

		// simple.txt starts with "line one"
		// Insert 'x' at 0,0
		mt.KeyPress(vt.KeyX)
		time.Sleep(defaultSleep)

		validateCursorPosition(ed, 1, 0, t)

		expect := make([]cellInfo, 0)
		expect = append(expect, getCellInfoFromString("xline one", 6, 0, &th.contentStyle)...)
		validateTerm(mt, expect, t)
	})

	t.Run("InsertNewLine rendering", func(t *testing.T) {
		ed, _, mt := getTestEditor(simpleFile)
		ed.Start()
		time.Sleep(defaultSleep)

		// Move to end of "line one" (length 8)
		for range 8 {
			mt.KeyPress(vt.KeyRight)
		}
		time.Sleep(defaultSleep)
		validateCursorPosition(ed, 8, 0, t)

		mt.KeyPress(vt.KeyEnter)
		time.Sleep(defaultSleep)

		validateCursorPosition(ed, 0, 1, t)

		expect := make([]cellInfo, 0)
		expect = append(expect, getCellInfoFromString("line one", 6, 0, &th.contentStyle)...)
		expect = append(expect, getCellInfoFromString("    2 ", 0, 1, &th.rightPanelStyle)...)
		expect = append(expect, getCellInfoFromString("", 6, 1, &th.contentStyle)...)
		expect = append(expect, getCellInfoFromString("    3 ", 0, 2, &th.rightPanelStyle)...)
		expect = append(expect, getCellInfoFromString("line two", 6, 2, &th.contentStyle)...)
		validateTerm(mt, expect, t)
	})

	t.Run("DeleteChar rendering", func(t *testing.T) {
		ed, _, mt := getTestEditor(simpleFile)
		ed.Start()
		time.Sleep(defaultSleep)

		// "line one" -> Delete 'e' at end
		for range 8 {
			mt.KeyPress(vt.KeyRight)
		}
		time.Sleep(defaultSleep)

		mt.KeyPress(vt.KeyBackspace)
		time.Sleep(defaultSleep)

		validateCursorPosition(ed, 7, 0, t)

		expect := make([]cellInfo, 0)
		expect = append(expect, getCellInfoFromString("line on ", 6, 0, &th.contentStyle)...) // space because of how it might be rendered or cleared
		validateTerm(mt, expect, t)
	})

	t.Run("DeleteChar at start of line (Merge)", func(t *testing.T) {
		ed, _, mt := getTestEditor(simpleFile)
		ed.Start()
		time.Sleep(defaultSleep)

		// Move to start of "line two"
		mt.KeyPress(vt.KeyDown)
		time.Sleep(defaultSleep)
		validateCursorPosition(ed, 0, 1, t)

		mt.KeyPress(vt.KeyBackspace)
		time.Sleep(defaultSleep)

		// "line one" + "line two" = "line oneline two"
		validateCursorPosition(ed, 8, 0, t)

		expect := make([]cellInfo, 0)
		expect = append(expect, getCellInfoFromString("line oneline two", 6, 0, &th.contentStyle)...)
		validateTerm(mt, expect, t)
	})
}
