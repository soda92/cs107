package main

// #cgo LDFLAGS: -Lbuild -lrsg -lstdc++
// #include <string.h>
// #include <stdlib.h>
// #include "main.h"
import "C"
import "unsafe"
import "fmt"

func main() {
	cStr := C.rsg_main(C.CString("data/excuse.g"))
	goStr := C.GoString(cStr)
	fmt.Print(goStr)
	C.free(unsafe.Pointer(cStr))
}
