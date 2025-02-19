package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"log"

	"flag"
	"fyne.io/fyne/v2/layout"
	"github.com/fyne-io/terminal"
)

func RunCommand(t *terminal.Terminal, command string) {
	t.Write([]byte(command))
	t.Write([]byte{'\r'})
}

func main() {
	cmd := flag.Bool("cmd", false, "whether run cmd version")
	test := flag.Bool("test", false, "run test")
	dg_main := flag.Bool("main", false, "run six degrees")
	name := flag.String("name", "nil", "query name")

	flag.Parse()

	if *cmd {
		if *test {
			imdb_test_main(*name)
		}
		if *dg_main {
			six_dg_main(nil)
		}
		return
	}

	a := app.New()
	a.Settings().SetTheme(NewMyTheme())
	w := a.NewWindow("Six Degrees")

	w.Resize(fyne.NewSize(800, 500))
	w.CenterOnScreen()

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

	go_cmd_imdb_test := widget.NewButton("go-imdb-test", func() {
		imdb_test_main("nil")
	})

	go_cmd_six_degrees := widget.NewButton("go-six-degrees", func() {
		six_dg_main(nil)
	})
	c2 := container.New(layout.NewHBoxLayout(), button, button2, button3, go_cmd_imdb_test, go_cmd_six_degrees)

	c := container.NewBorder(nil, c2, nil, nil, t)
	w.SetContent(c)
	w.ShowAndRun()
}
