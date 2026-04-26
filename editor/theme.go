package editor

import (
	"github.com/gdamore/tcell/v3"
)

type Theme struct {
	contentStyle    tcell.Style
	footerStyle     tcell.Style
	rightPanelStyle tcell.Style
}

func InitTheme() Theme {
	return tokyoNightTheme()
}

func tokyoNightTheme() Theme {
	contStyle := tcell.StyleDefault.
		Foreground(tcell.NewRGBColor(192, 202, 245)).
		Background(tcell.NewRGBColor(26, 27, 38))

	footStyle := tcell.StyleDefault.
		Foreground(tcell.NewRGBColor(22, 22, 30)).
		Background(tcell.NewRGBColor(122, 162, 247))

	rPanStyle := tcell.StyleDefault.
		Foreground(tcell.NewRGBColor(68, 75, 107)).
		Background(tcell.NewRGBColor(26, 27, 38))

	return Theme{contStyle, footStyle, rPanStyle}
}
