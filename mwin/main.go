package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewWindow(a fyne.App, text string) *container.InnerWindow {
	w := widget.NewLabel(text)
	w2 := container.NewInnerWindow(text, container.NewVBox(w))
	w2.Resize(fyne.NewSize(200, 200))
	return w2
}

func MoveWindow(w *container.InnerWindow, offsetx int) {
	time.Sleep(1 * time.Second)
	// w.Content().Move(w.Content().Position().AddXY(float32(offsetx), 0))
	w.Move(w.Position().AddXY(float32(offsetx), 0))
}

func main() {
	a := app.New()
	w1 := NewWindow(a, "w1")
	w2 := NewWindow(a, "w2")
	w := container.NewMultipleWindows(w1, w2)
	w0 := a.NewWindow("Window")
	w0.SetContent(w)
	w0.Resize(fyne.NewSize(400, 400))
	w0.CenterOnScreen()
	w0.Show()
	go MoveWindow(w1, 100)
	go MoveWindow(w2, 300)
	a.Run()
}
