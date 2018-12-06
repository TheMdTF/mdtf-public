package main


// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -L../../../../lib -lsample_matching_algorithm
// #include <stdlib.h>
// #include "../../../../c/sample_matching_algorithm.h"
import "C"
import (
	"fmt"
)

func main() {

	r := C.cpp_create_template()

	fmt.Printf("A random number from a c library: %d\n", r)
}
