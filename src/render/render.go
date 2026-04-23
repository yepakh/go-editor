package render

import (
	"log"

	"github.com/gdamore/tcell/v3"
	"github.com/yepakh/go-editor/src/buffer"
)

var screen tcell.Screen
var theme Theme

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

func RenderBuffer(buf *buffer.Buffer, lineOff, charOff int) {
	screen.Clear()
	SetFooter(buf.GetFilepath())

	width, height := GetBufferSceenSize()
	for i := 0; i < height && lineOff+i < len(buf.Lines); i++ {
		line := buf.Lines[lineOff+i]
		for j := 0; j < width && j+charOff < len(line); j++ {
			screen.SetContent(j, i, line[j], nil, tcell.StyleDefault)
		}
	}

	screen.Show()
}

func Sync() { screen.Sync() }

func CloseScreen() { screen.Fini() }

func Reset() {
	screen.Clear()
}

func SetFooter(filepath string) {
	chars := []rune(filepath)
	w, h := screen.Size()

	for j := range w {
		if j < len(chars) {
			screen.SetContent(j, h-1, chars[j], nil, theme.GetFooterStyle())
		} else {
			screen.SetContent(j, h-1, 0, nil, theme.GetFooterStyle())
		}
	}
}

func GetBufferSceenSize() (width, height int) {
	w, h := screen.Size()
	return w, h - 1
}

func SetCursor(x, y int) {
	screen.ShowCursor(x, y)
	screen.Show()
}

func setTheme() {
	theme = InitTheme()

	screen.SetCursorStyle(tcell.CursorStyleSteadyBlock, theme.contentStyle.GetForeground())
	screen.SetStyle(theme.contentStyle)

}
