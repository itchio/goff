package goff

//#include <stdlib.h>
import "C"
import "unsafe"

func CString(s string) *C.char {
	if s == "" {
		return (*C.char)(nil)
	}
	return C.CString(s)
}

func FreeString(c *C.char) {
	if c != nil {
		C.free(unsafe.Pointer(c))
	}
}

func BoolToInt(b bool) C.int {
	if b {
		return 1
	}
	return 0
}
