package main
//#include <cuda.h>
//#cgo LDFLAGS: -lcuda
import "C"
import "fmt"
func main() {
	buf := C.CString(string(make([]byte, 256)))
	C.cuDeviceGetName(buf, 256, C.CUdevice(0))
	fmt.Println("Hello, your GPU is:", C.GoString(buf))
}