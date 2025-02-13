package main

// #cgo LDFLAGS: -L../assn-1-rsg/build -lrsg -lstdc++
// #include <string.h>
// #include <stdlib.h>
// #include "../assn-1-rsg/main.h"
import "C"
import "unsafe"
import "fmt"

func main() {
	// cs := C.CString("Hello from stdio")
	// C.myprint(cs)
	cStr := C.rsg_main(C.CString("../assn-1-rsg/data/excuse.g"))
	goStr := C.GoString(cStr)
	fmt.Print(goStr)
	C.free(unsafe.Pointer(cStr))
	// C.free(unsafe.Pointer(cs))
}
