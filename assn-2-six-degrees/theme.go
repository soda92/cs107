package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct {
	fyne.Theme
}

func NewMyTheme() fyne.Theme {
	return &myTheme{Theme: theme.DarkTheme()}
}
