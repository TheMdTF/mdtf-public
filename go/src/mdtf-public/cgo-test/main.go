package main


// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -L../../../../lib -lsample_matching_algorithm
// #include <stdlib.h>
// #include "../../../../c/sample_matching_algorithm.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	image := C.CString("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg==")
	defer C.free(unsafe.Pointer(image))

	r := C.cpp_create_template(image)

	fmt.Printf("A random number from a c library: %d\n", r)
}
