package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Assignment 1")

	files := []string{"excuse.g",
		"bond.g",
		"trek.g",
		"poem.g"}

	tabs := container.NewAppTabs()
	for _, f := range files {
		t := container.NewVBox()
		for range 5 {
			t1 := widget.NewLabel(get_output2("data/" + f))
			t.Add(t1)
		}

		tabs.Append(
			container.NewTabItem(f, container.NewVScroll(t)))
	}

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(500, 500))
	w.CenterOnScreen()
	w.ShowAndRun()
}
