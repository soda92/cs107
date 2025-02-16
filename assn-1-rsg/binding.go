package main

// #cgo LDFLAGS: -Lbuild -lrsg
// #include <string.h>
// #include <stdlib.h>
// #include "main.h"
import "C"
import (
	"unsafe"
)

func get_output(file string) string {
	c_str := C.rsg_main(C.CString(file))
	go_str := C.GoString(c_str)
	C.free(unsafe.Pointer(c_str))

	return go_str
}
