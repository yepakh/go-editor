package render

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v3"
)

var screen tcell.Screen
var theme Theme
var rightSidePadding int = 6

func InitScreen() <-chan tcell.Event {
	var err error
	screen, err = tcell.NewScreen()

	if err != nil {
		log.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}

	setTheme()
	SetCursor(0, 0)

	return screen.EventQ()
}

func RenderBuffer(lines *[][]rune, lineOff, charOff int) {
	screen.Clear()

	width, height := GetBufferSceenSize()
	for i := 0; i < height && lineOff+i < len(*lines); i++ {
		line := (*lines)[lineOff+i]
		SetLineNumber(lineOff+i, i)
		for j := 0; j < width && j+charOff < len(line); j++ {
			screen.SetContent(j+rightSidePadding, i, line[j], nil, tcell.StyleDefault)
		}
	}

	screen.Show()
}

func Sync() { screen.Sync() }

func CloseScreen() { screen.Fini() }

func Reset() {
	screen.Clear()
}

func SetLineNumber(lineNum, linePos int) {
	padding := rightSidePadding
	lineStr := fmt.Sprint(lineNum + 1)
	lineNumLen := len(lineStr)

	screen.SetContent(padding-1, linePos, 0, nil, theme.rightPanelStyle)
	padding--

	for i := range padding - 1 {
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
			screen.SetContent(j, h-1, chars[j], nil, theme.GetFooterStyle())
		} else {
			screen.SetContent(j, h-1, 0, nil, theme.GetFooterStyle())
		}
	}

	screen.Show()
}

func GetBufferSceenSize() (width, height int) {
	w, h := screen.Size()
	return w - rightSidePadding, h - 1
}

func SetCursor(x, y int) {
	screen.ShowCursor(x+rightSidePadding, y)
	screen.Show()
}

func setTheme() {
	theme = InitTheme()

	screen.SetCursorStyle(tcell.CursorStyleSteadyBlock, theme.contentStyle.GetForeground())
	screen.SetStyle(theme.contentStyle)

}
