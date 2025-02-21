package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewWindow(a fyne.App, text string) fyne.Window {
	w := a.NewWindow(text)

	w.Resize(fyne.NewSize(200, 100))
	w.CenterOnScreen()
	w.Content().Move(w.Content().Position().SubtractXY(200, 0))
	w.Show()
	return w
}

func MoveWindow(w fyne.Window, offsetx int) {
	time.Sleep(1 * time.Second)
	w.Content().Move(w.Content().Position().AddXY(float32(offsetx), 0))
}

func main1() {
	a := app.New()
	w1 := NewWindow(a, "w1")
	w2 := NewWindow(a, "w2")
	go MoveWindow(w1, -200)
	go MoveWindow(w2, 200)
	a.Run()
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := container.NewDocTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	iw := container.NewInnerWindow("123", tabs)

	a := container.NewMultipleWindows(iw)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	// tabs.SetTabLocation(container.TabLocationLeading)
	// myWindow.Resize(fyne.NewSize(200, 400))
	a.Resize(fyne.NewSize(400, 400))

	myWindow.SetContent(a)
	myWindow.ShowAndRun()
}
