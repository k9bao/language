package stringG

import (
	"reflect"
	"unsafe"
)

// $GOROOT/src/reflect/value.go

// type StringHeader struct {
//     Data uintptr
//     Len  int
// }

func String2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func String2bytesV2(s string) []byte {
	strh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceh := reflect.SliceHeader{Data: strh.Data, Len: strh.Len, Cap: strh.Len}
	return *(*[]byte)(unsafe.Pointer(&sliceh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Bytes2Uint16(u8 []byte) []uint16 {
	sh := (*reflect.SliceHeader)((unsafe.Pointer(&u8)))
	sh.Len /= 2
	sh.Cap /= 2
	return *(*[]uint16)(unsafe.Pointer(sh))
}
