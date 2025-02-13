package main

// #cgo LDFLAGS: -Lbuild -lrsg -lstdc++
// #include <string.h>
// #include <stdlib.h>
// #include "main.h"
import "C"
import (
	"unsafe"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func get_output() string {
	cStr := C.rsg_main(C.CString("data/excuse.g"))
	goStr := C.GoString(cStr)
	C.free(unsafe.Pointer(cStr))
	return goStr
}

func main() {
	a := app.New()
	w := a.NewWindow("Assignment 1")
	w.CenterOnScreen()

	w.SetContent(container.NewVBox(widget.NewLabel(get_output())))
	w.ShowAndRun()
}
