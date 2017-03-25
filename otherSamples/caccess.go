package main
//#include <cuda.h>
//#cgo LDFLAGS: -lcuda
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(100)
	fmt.Println("Hello, your GPU is:", Random())
	buf := C.CString(string(make([]byte, 256)))
	fmt.Println(buf)
	C.cuDeviceGetName(buf, 256, C.CUdevice(0))
	fmt.Println("Hello, your GPU is:", C.GoString(buf))
}