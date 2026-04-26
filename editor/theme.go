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
	contStyle := tcell.Style{}
	contStyle = contStyle.Foreground(tcell.NewRGBColor(192, 202, 245))
	contStyle = contStyle.Background(tcell.NewRGBColor(26, 27, 38))

	footStyle := tcell.Style{}
	footStyle = footStyle.Foreground(tcell.NewRGBColor(22, 22, 30))
	footStyle = footStyle.Background(tcell.NewRGBColor(122, 162, 247))

	rPanStyle := tcell.Style{}
	rPanStyle = rPanStyle.Foreground(tcell.NewRGBColor(68, 75, 107))
	rPanStyle = rPanStyle.Background(tcell.NewRGBColor(26, 27, 38))
	return Theme{contStyle, footStyle, rPanStyle}
}
