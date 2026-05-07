package editor

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v3"
	piecetable "github.com/yepakh/go-editor/piecetable"
)

var screen tcell.Screen
var theme Theme
var rightSidePadding int = 6

func InitRenderScreen(sc tcell.Screen) <-chan tcell.Event {
	screen = sc
	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}

	setTheme()
	SetRenderCursor(0, 0)

	return screen.EventQ()
}

func RefreshLine(data *piecetable.PieceTable, lineNum, lineOff, charOff int) {
	width, _ := GetContentSceenSize()
	line := data.GetLines(lineNum, 1)[0]

	for j := range width {
		var char rune = 0
		if j+charOff < len(line) {
			char = line[j+charOff]
		}
		screen.SetContent(j+rightSidePadding, lineNum-lineOff, char, nil, theme.contentStyle)
	}

	screen.Show()
}

func RenderFromLine(data *piecetable.PieceTable, lineNum, lineOff, charOff int) {
	width, height := GetContentSceenSize()
	linesToSkip := lineNum - lineOff
	lines := data.GetLines(lineNum, height-linesToSkip)

	for i := range height {
		if i >= len(lines) {
			screen.PutStr(0, 0, "")
			for j := range width + rightSidePadding {
				screen.SetContent(j, i+linesToSkip, 0, nil, theme.contentStyle)
			}
			continue
		}
		line := lines[i]
		RenderLineNumber(linesToSkip+lineOff+i, i+linesToSkip)
		for j := range width {
			var char rune = 0
			if j+charOff < len(line) {
				char = line[j+charOff]
			}
			screen.SetContent(j+rightSidePadding, i+linesToSkip, char, nil, theme.contentStyle)
		}
	}

	screen.Show()
}

func RenderBuffer(data *piecetable.PieceTable, lineOff, charOff int) {
	width, height := GetContentSceenSize()
	lines := data.GetLines(lineOff, height)
	for i := range height {
		if i >= len(lines) {
			screen.PutStr(0, 0, "")
			for j := range width + rightSidePadding {
				screen.SetContent(j, i, 0, nil, theme.contentStyle)
			}
			continue
		}
		line := lines[i]
		RenderLineNumber(lineOff+i, i)
		for j := range width {
			var char rune = 0
			if j+charOff < len(line) {
				char = line[j+charOff]
			}
			screen.SetContent(j+rightSidePadding, i, char, nil, theme.contentStyle)
		}
	}

	screen.Show()
}

func ClearScreen() { screen.Clear() }

func RenderSync() { screen.Sync() }

func CloseRenderScreen() { screen.Fini() }

func Reset() {
	screen.Clear()
}

func RenderLineNumber(lineNum, linePos int) {
	padding := rightSidePadding
	lineStr := fmt.Sprint(lineNum + 1)
	lineNumLen := len(lineStr)

	screen.SetContent(padding-1, linePos, 0, nil, theme.rightPanelStyle)
	padding--

	for i := range padding {
		if i < lineNumLen {
			r := rune(lineStr[lineNumLen-i-1])
			screen.SetContent(padding-i-1, linePos, r, nil, theme.rightPanelStyle)
		} else {
			screen.SetContent(padding-i-1, linePos, 0, nil, theme.rightPanelStyle)
		}
	}
}

func RenderFooter(filepath string) {
	chars := []rune(filepath)
	w, h := screen.Size()

	for j := range w {
		if j < len(chars) {
			screen.SetContent(j, h-1, chars[j], nil, theme.footerStyle)
		} else {
			screen.SetContent(j, h-1, 0, nil, theme.footerStyle)
		}
	}

	screen.Show()
}

func GetContentSceenSize() (width, height int) {
	w, h := screen.Size()
	return w - rightSidePadding, h - 1
}

func SetRenderCursor(x, y int) {
	screen.ShowCursor(x+rightSidePadding, y)
	screen.Show()
}

func setTheme() {
	theme = InitTheme()

	screen.SetCursorStyle(tcell.CursorStyleSteadyBlock, theme.contentStyle.GetForeground())
	screen.SetStyle(theme.contentStyle)
}
