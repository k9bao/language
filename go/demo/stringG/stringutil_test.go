package stringG

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	s := "123456789"
	fmt.Println("内容", s)
	strh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println("数据内存地址", strh.Data, strh.Len)
	fmt.Println("-----------------------")

	b := String2bytes(s)
	fmt.Println("内容", b)
	sliceh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Println("数据内存地址", sliceh.Data, sliceh.Len, sliceh.Cap)
	fmt.Println("-----------------------")

	b = String2bytesV2(s)
	fmt.Println("内容", b)
	sliceh = (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Println("数据内存地址", sliceh.Data, sliceh.Len, sliceh.Cap)
	fmt.Println("-----------------------")

	s = Bytes2String(b)
	fmt.Println("内容", s)
	strh = (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println("数据内存地址", strh.Data, strh.Len)
}

func TestBytes2Uint16(t *testing.T) {
	s8 := make([]byte, 0, 8)
	s8 = append(s8, 'a')
	s8 = append(s8, 'b')
	s8 = append(s8, 'b')
	s8 = append(s8, 'd')
	s8 = append(s8, 'e')
	s8 = append(s8, 'f')
	fmt.Println("内容", s8)
	sliceh := (*reflect.SliceHeader)(unsafe.Pointer(&s8))
	fmt.Println("数据内存地址", sliceh.Data, sliceh.Len, sliceh.Cap)
	fmt.Println("-----------------------")
	s16 := Bytes2Uint16(s8)
	fmt.Println("内容", s16)
	sliceh = (*reflect.SliceHeader)(unsafe.Pointer(&s16))
	fmt.Println("数据内存地址", sliceh.Data, sliceh.Len, sliceh.Cap)
}

func TestBuilder(t *testing.T) {
	var b1 strings.Builder
	fmt.Println(b1.Len())
	b1.Grow(100)
	b1.WriteString("ABC")
	b1.WriteByte('D')
	b1.Write(String2bytesV2("EFG"))
	//b1.WriteRune("HIJ")
	fmt.Println(b1.String())
}
