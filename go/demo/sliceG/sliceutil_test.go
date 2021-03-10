package sliceG

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/tmthrgd/go-memset"
)

func TestSlice(t *testing.T) {
	list1 := make([]int, 0, 8)
	arrar := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	list2 := arrar[:5:5]
	fmt.Println(list2, len(list2), cap(list2))
	return
	listappend(list1)
	listappend(list2)
	fmt.Println(list1, len(list1), cap(list1))
	fmt.Println(list2, len(list2), cap(list2))
	fmt.Println(arrar, len(arrar), cap(arrar))
	temp := make([]int, 0)
	fmt.Println("***************", len(temp))
}

func TestList(t *testing.T) {
	test1 := make([]int, 0, 8)
	test1 = append(test1, 1)
	h := (*[3]uintptr)(unsafe.Pointer(&test1))
	fmt.Println("----", h[0], h[1], h[2], &h[0], &h[1], &h[2], test1)
	listappend(test1)
	fmt.Println("----", h[0], h[1], h[2], &h[0], &h[1], &h[2], test1)
}

func TestByteConv(t *testing.T) {
	uv8 := []byte{1, 2, 1, 2, 1, 2}
	sliceHeader8 := (*reflect.SliceHeader)((unsafe.Pointer(&uv8)))
	uv16 := make([]uint16, 0)
	sliceHeader16 := (*reflect.SliceHeader)((unsafe.Pointer(&uv16)))
	sliceHeader16.Len = len(uv8) / 2
	sliceHeader16.Cap = sliceHeader16.Len
	sliceHeader16.Data = sliceHeader8.Data
	fmt.Println(uv8, uv16)
	var b16 uint16
	b16 = binary.LittleEndian.Uint16(uv8)
	// b16 = uint16(uv8[0])
	// b16 = b16<<8 | uint16(uv8[1])
	fmt.Println(b16)
}

//BenchmarkMemcopy-8                  1000           1580994 ns/op
func BenchmarkMemcopy(b *testing.B) {
	len := 1920 * 1080 * 3 / 2
	dst := make([]byte, len, len)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len; j++ {
			dst[i] = 100
		}
	}
}

//BenchmarkMemcopy2-8                10000            109099 ns/op
func BenchmarkMemcopy2(b *testing.B) {
	len := 1920 * 1080 * 3 / 2
	dst := make([]byte, len, len)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Memset(dst, 100)
	}
}

//BenchmarkMemcopy3-8                20000             74700 ns/op
func BenchmarkMemcopy3(b *testing.B) {
	len := 1920 * 1080 * 3 / 2
	dst := make([]byte, len, len)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memset.Memset(dst, 100, uint64(len))
	}
}

func TestSlice(t *testing.T) {
	list1 := make([]int, 0, 8)
	arrar := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	list2 := arrar[:5:5]
	fmt.Println(list2, len(list2), cap(list2))
	return
	listappend(list1)
	listappend(list2)
	fmt.Println(list1, len(list1), cap(list1))
	fmt.Println(list2, len(list2), cap(list2))
	fmt.Println(arrar, len(arrar), cap(arrar))
	temp := make([]int, 0)
	fmt.Println("***************", len(temp))
}
