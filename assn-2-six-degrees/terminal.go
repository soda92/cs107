package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"log"

	"fyne.io/fyne/v2/layout"
	"github.com/fyne-io/terminal"
)

func RunCommand(t *terminal.Terminal, command string) {
	t.Write([]byte(command))
	t.Write([]byte{'\r'})
}

func main() {
	a := app.New()
	a.Settings().SetTheme(NewMyTheme())
	w := a.NewWindow("Six Degrees")

	// 	w.SetContent(widget.NewLabel("main"))

	w.Resize(fyne.NewSize(800, 500))
	w.CenterOnScreen()
	// 	// w.ShowAndRun()
	// // run new terminal and close app on terminal exit.
	t := terminal.New()
	t.SetStartDir(".")
	t.Resize(fyne.NewSize(800, 400))
	go func() {
		_ = t.RunLocalShell()
		log.Printf("Terminal's shell exited with exit code: %d", t.ExitCode())
		a.Quit()
	}()

	button := widget.NewButton("run1", func() {
		RunCommand(t, "./build/imdb-test.exe")
	})

	button2 := widget.NewButton("run2", func() {
		RunCommand(t, "./build/six-degrees.exe")
	})

	button3 := widget.NewButton("build", func() {
		RunCommand(t, "make")
	})
	c2 := container.New(layout.NewHBoxLayout(), button, button2, button3)

	c := container.NewBorder(nil, c2, nil, nil, t)
	w.SetContent(c)
	w.ShowAndRun()
}
