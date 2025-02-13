package main

// #cgo LDFLAGS: -Lbuild -lrsg -lstdc++
// #include <string.h>
// #include <stdlib.h>
// #include "main.h"
import "C"
import (
	"fyne.io/fyne/v2/app"
	"unsafe"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func get_output() string {
	c_str := C.rsg_main(C.CString("data/excuse.g"))
	go_str := C.GoString(c_str)
	C.free(unsafe.Pointer(c_str))
	return go_str
}

func main() {
	a := app.New()
	w := a.NewWindow("Assignment 1")
	w.CenterOnScreen()

	w.SetContent(container.NewVBox(widget.NewLabel(get_output())))
	w.ShowAndRun()
}
