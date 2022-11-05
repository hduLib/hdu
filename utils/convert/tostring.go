package convert

import (
	"reflect"
	"unsafe"
)

// ToString is used to zero copy converT `string`
func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ToBytes is used to zero copy converT `[]byte`.
// Attention! Since `string` is immutable in Go, theconversion
// from `string` to `[]byte` can easily casuse panic as change
// the converted `[]byte` is actually change the immutable
// `string` stored in .RODATA(maybe) section.
func ToBytes(s string) []byte {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	bh.Cap = len(s)
	return *(*[]byte)(unsafe.Pointer(bh))
}
