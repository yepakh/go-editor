package render

import (
	"log"

	"github.com/gdamore/tcell/v3"
	"github.com/yepakh/notepad/src/buffer"
)

var screen tcell.Screen

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

	return screen.EventQ()
}

func RenderBuffer(buff *buffer.Buffer, lineOff, charOff int) {
	screen.Clear()
	width, height := screen.Size()
	for i := 0; i < height && lineOff+i < len(buff.Lines); i++ {
		line := buff.Lines[lineOff+i]
		for j := 0; j < width && j+charOff < len(line); j++ {
			SetCharacter(j, i, line[charOff+j])
		}
	}

	screen.Show()
}

func Sync() { screen.Sync() }

func CloseScreen() { screen.Fini() }

func Reset() { screen.Clear() }

func SetCharacter(x, y int, char rune) { screen.SetContent(x, y, char, nil, tcell.StyleDefault) }

func GetSceenSize() (width, height int) {
	return screen.Size()
}

func SetCursor(x, y int) {
	screen.ShowCursor(x, y)
	screen.Show()
}

func setTheme() {
	screen.SetCursorStyle(tcell.CursorStyleSteadyBlock)
}
