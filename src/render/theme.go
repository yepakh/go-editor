package render

import "github.com/gdamore/tcell/v3/color"

type theme struct {
	fgColor color.Color
	bgColor color.Color
}

var Theme theme

func InitTheme() {
	Theme = theme{color.Orange, color.Coral}
}
