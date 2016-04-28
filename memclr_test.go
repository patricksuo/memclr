package memclr

import (
	"reflect"
	"testing"
	"unsafe"
)

func BenchmarkMemclr(b *testing.B) {
	buf := make([]byte, 1000, 1000)
	bufVal := reflect.ValueOf(buf)
	bufCap := cap(buf)
	bufPtr := unsafe.Pointer(bufVal.Pointer())
	for i := 0; i < b.N; i++ {
		Memclr(bufPtr, uintptr(bufCap))
	}
}
