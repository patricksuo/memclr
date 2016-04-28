package memclr

import (
	"unsafe"
)

//go:linkname Memclr runtime.memclr
func Memclr(ptr unsafe.Pointer, n uintptr)
