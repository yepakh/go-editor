package render

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v3"
	piecetable "github.com/yepakh/go-editor/piecetable"
)

type Render struct {
	screen tcell.Screen
	theme  *Theme
}

var rightSidePadding int = 6

func InitRenderScreen(sc tcell.Screen) (*Render, <-chan tcell.Event) {
	screen := sc
	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}

	r := Render{screen, nil}

	r.setTheme()
	r.SetRenderCursor(0, 0)

	return &r, screen.EventQ()
}

func (r *Render) RefreshLine(data *piecetable.PieceTable, lineNum, lineOff, charOff int) {
	width, _ := r.GetContentSceenSize()
	line := data.GetLines(lineNum, 1)[0]

	for j := range width {
		var char rune = 0
		if j+charOff < len(line) {
			char = line[j+charOff]
		}
		r.screen.SetContent(j+rightSidePadding, lineNum-lineOff, char, nil, r.theme.ContentStyle)
	}

	r.screen.Show()
}

func (r *Render) RenderFromLine(data *piecetable.PieceTable, lineNum, lineOff, charOff int) {
	width, height := r.GetContentSceenSize()
	linesToSkip := lineNum - lineOff
	lines := data.GetLines(lineNum, height-linesToSkip)

	for i := range height {
		if i >= len(lines) {
			r.screen.PutStr(0, 0, "")
			for j := range width + rightSidePadding {
				r.screen.SetContent(j, i+linesToSkip, 0, nil, r.theme.ContentStyle)
			}
			continue
		}
		line := lines[i]
		r.RenderLineNumber(linesToSkip+lineOff+i, i+linesToSkip)
		for j := range width {
			var char rune = 0
			if j+charOff < len(line) {
				char = line[j+charOff]
			}
			r.screen.SetContent(j+rightSidePadding, i+linesToSkip, char, nil, r.theme.ContentStyle)
		}
	}

	r.screen.Show()
}

func (r *Render) RenderBuffer(data *piecetable.PieceTable, lineOff, charOff int) {
	width, height := r.GetContentSceenSize()
	lines := data.GetLines(lineOff, height)
	for i := range height {
		if i >= len(lines) {
			r.screen.PutStr(0, 0, "")
			for j := range width + rightSidePadding {
				r.screen.SetContent(j, i, 0, nil, r.theme.ContentStyle)
			}
			continue
		}
		line := lines[i]
		r.RenderLineNumber(lineOff+i, i)
		for j := range width {
			var char rune = 0
			if j+charOff < len(line) {
				char = line[j+charOff]
			}
			r.screen.SetContent(j+rightSidePadding, i, char, nil, r.theme.ContentStyle)
		}
	}

	r.screen.Show()
}

func (r *Render) ClearScreen() { r.screen.Clear() }

func (r *Render) RenderSync() { r.screen.Sync() }

func (r *Render) CloseRenderScreen() { r.screen.Fini() }

func (r *Render) Reset() {
	r.screen.Clear()
}

func (r *Render) RenderLineNumber(lineNum, linePos int) {
	padding := rightSidePadding
	lineStr := fmt.Sprint(lineNum + 1)
	lineNumLen := len(lineStr)

	r.screen.SetContent(padding-1, linePos, 0, nil, r.theme.RightPanelStyle)
	padding--

	for i := range padding {
		if i < lineNumLen {
			rn := rune(lineStr[lineNumLen-i-1])
			r.screen.SetContent(padding-i-1, linePos, rn, nil, r.theme.RightPanelStyle)
		} else {
			r.screen.SetContent(padding-i-1, linePos, 0, nil, r.theme.RightPanelStyle)
		}
	}
}

func (r *Render) RenderFooter(filepath string) {
	chars := []rune(filepath)
	w, h := r.screen.Size()

	for j := range w {
		if j < len(chars) {
			r.screen.SetContent(j, h-1, chars[j], nil, r.theme.FooterStyle)
		} else {
			r.screen.SetContent(j, h-1, 0, nil, r.theme.FooterStyle)
		}
	}

	r.screen.Show()
}

func (r *Render) GetContentSceenSize() (width, height int) {
	w, h := r.screen.Size()
	return w - rightSidePadding, h - 1
}

func (r *Render) SetRenderCursor(x, y int) {
	r.screen.ShowCursor(x+rightSidePadding, y)
	r.screen.Show()
}

func (r *Render) setTheme() {
	theme := InitTheme()
	r.theme = &theme

	r.screen.SetCursorStyle(tcell.CursorStyleSteadyBlock, r.theme.ContentStyle.GetForeground())
	r.screen.SetStyle(r.theme.ContentStyle)
}
