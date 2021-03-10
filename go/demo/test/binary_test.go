package test

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/*
https://blog.csdn.net/ce123_zhouwei/article/details/6971544
一般操作系统都是小端，而通讯协议是大端的。
4.1 常见CPU的字节序
Big Endian : PowerPC、IBM、Sun
Little Endian : x86、DEC
ARM既可以工作在大端模式，也可以工作在小端模式。
*/
func TestBytes(t *testing.T) {
	var i64 int64 = 2323 //(高)0 0 0 0 0 0 9 19(低) Intel/AMD x86 x64都是小端序
	var buf64 = make([]byte, 8)
	//Big-Endian就是高位字节排放在内存的低地址端，低位字节排放在内存的高地址端。
	binary.BigEndian.PutUint64(buf64, uint64(i64))
	fmt.Println(buf64) //0 0 0 0 0 0 9 19
	fmt.Println(int64(binary.BigEndian.Uint64(buf64)))

	//Little-Endian就是低位字节排放在内存的低地址端，高位字节排放在内存的高地址端。
	var buf64S = make([]byte, 8)
	sh := (*reflect.SliceHeader)((unsafe.Pointer(&buf64S)))
	sh.Data = (uintptr)(unsafe.Pointer(&i64))
	sh.Len = 8
	sh.Cap = 8
	fmt.Println(buf64S) //19 9 0 0 0 0 0 0
	fmt.Println(int64(binary.LittleEndian.Uint64(buf64S)))

	var i32 int32 = 800
	var buf32 = make([]byte, 4)
	binary.BigEndian.PutUint32(buf32, uint32(i32))
	fmt.Println(buf32)
	fmt.Println(int32(binary.BigEndian.Uint32(buf32)))

	var i16 int16 = 27
	var buf16 = make([]byte, 2)
	binary.BigEndian.PutUint16(buf16, uint16(i16))
	fmt.Println(buf16)
	fmt.Println(int16(binary.BigEndian.Uint16(buf16)))

	binary.LittleEndian.PutUint16(buf16, uint16(i16))
	fmt.Println(buf16)
	fmt.Println(int16(binary.LittleEndian.Uint16(buf16)))
}

type Rect struct {
	Left   int `json:"left"`
	Top    int `json:"top"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func TestEnlarge(t *testing.T) {
	r := Rect{1074, 546, 56, 56}
	leftRatio := 0.5
	topRatio := 1.0
	heightRatio := 2.5
	widthRatio := 2.0
	r.Left -= int(float64(r.Width) * leftRatio)
	r.Top -= int(float64(r.Height) * topRatio)
	r.Height = int(float64(r.Height) * heightRatio)
	r.Width = int(float64(r.Width) * widthRatio)
	fmt.Println(r)
}
