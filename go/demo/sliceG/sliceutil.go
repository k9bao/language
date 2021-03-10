package sliceG

import (
	"fmt"
	"unsafe"
)

//风险：不安全，s赋值给x后，x[0]指向的内容随时可能被GC回收
func String2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2String(sh []byte) string {
	return *(*string)(unsafe.Pointer(&sh))
}

func listappend(list []int) {
	h := (*[3]uintptr)(unsafe.Pointer(&list))
	fmt.Println("++++", h[0], h[1], h[2], &h[0], &h[1], &h[2], list)
	list = append(list, 1)
	fmt.Println("++++", h[0], h[1], h[2], &h[0], &h[1], &h[2], list)
}

func Memset(data []byte, value byte) {
	if value == 0 {
		for i := range data {
			data[i] = 0
		}
	} else if len(data) != 0 {
		data[0] = value

		for i := 1; i < len(data); i *= 2 {
			copy(data[i:], data[:i])
		}
	}
}
