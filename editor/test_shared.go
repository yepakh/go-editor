package editor

import (
	"log"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/gdamore/tcell/v3/vt"
)

const simpleFile = "testdata/simple.txt"
const longFile = "testdata/long.txt"
const emptyFile = "testdata/empty.txt"
const noFile = ""

type cellInfo struct {
	X    vt.Col
	Y    vt.Row
	C    string
	Fg   color.Color
	Bg   color.Color
	Attr vt.Attr
}

func getCellInfoFromString(s string, x, y int, st *tcell.Style) []cellInfo {
	cells := make([]cellInfo, 0, len(s))
	for i := 0; i < len(s); i++ {
		char := s[i]
		ci := cellInfo{
			X:    vt.Col(i + x),
			Y:    vt.Row(y),
			C:    string(char),
			Fg:   st.GetForeground(),
			Bg:   st.GetBackground(),
			Attr: vt.Attr(st.GetAttributes())}
		cells = append(cells, ci)
	}

	return cells
}

func getTestEditor(filePath string) (*Editor, tcell.Screen, vt.MockTerm) {
	mockTty := vt.NewMockTerm(vt.MockOptColors(256*256*256), vt.MockOptSize{})
	mockScreen, err := tcell.NewTerminfoScreenFromTty(mockTty)
	tcell.ShimScreen(mockScreen)

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	ed, err := InitFromFile(&screen, filePath)
	if err != nil {
		log.Fatal(err)
	}
	return ed, screen, mockTty
}
