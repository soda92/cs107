package main

import (
	// "fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewWindow(c1 chan fyne.Size, text string) *container.InnerWindow {
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
	c := make(chan fyne.Size)
	w1 := NewWindow(c, "w1")
	w2 := NewWindow(c, "w2")
	f := func() {
		// var v fyne.Size
		for {
			v := <-c
			w2.Resize(v)
			// w1.Resize(v)
		}
	}

	f2 := func() {
		// var v fyne.Size
		for {
			time.Sleep(100 * time.Millisecond)
			s := w1.Size()
			c <- s
		}
	}
	go f()
	go f2()
	w := container.NewMultipleWindows(w1, w2)
	w0 := a.NewWindow("Window")
	w0.SetContent(w)
	w0.Resize(fyne.NewSize(800, 400))
	w0.CenterOnScreen()
	w0.Show()
	MoveWindow(w1, 100)
	MoveWindow(w2, 300)
	c2 := make(chan fyne.Position)
	g1 := func() {
		time.Sleep(1 * time.Second)
		baseLoc := w1.Position()
		for {
			time.Sleep(10 * time.Millisecond)
			loc := w1.Position()

			delta := loc.SubtractXY(baseLoc.X, baseLoc.Y)
			c2 <- delta

		}
	}
	g2 := func() {
		time.Sleep(1 * time.Second)
		for {
			v := <-c2
			w2.Move(w.Position().AddXY(v.X+400, v.Y))
			// w1.Resize(v)
		}
	}
	go g1()
	go g2()
	a.Run()
}
