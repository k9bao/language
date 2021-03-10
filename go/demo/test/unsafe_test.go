package test

import (
	"fmt"
	"testing"
	"unsafe"
)

//uintptr和unsafe.Pointer的区别就是：unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象，uintptr类型的目标会被回收。
func pointerDemo() {
	var x struct {
		a bool  //2bytes
		b int16 //6bytes
		c []int //24bityes(8+8+8)
	}

	//unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞.
	//uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b) 指针的运算
	fmt.Println(unsafe.Sizeof(x), unsafe.Sizeof(x.a), unsafe.Sizeof(x.b), unsafe.Sizeof(x.c)) //32 1 2 24
	fmt.Println(unsafe.Alignof(x.a), unsafe.Alignof(x.b), unsafe.Alignof(x.c))                //1 2 8
	fmt.Println(unsafe.Offsetof(x.a), unsafe.Offsetof(x.b), unsafe.Offsetof(x.c))             //0 2 8
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))        //pb := &x.b 等价
	// tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b) //maybe error,uintptr类型的临时变量只是一个普通的数字,此操作完成后，x可能被GC回收,下边的获取指针可能会出问题(小概率)， uintptr 无法持有对象，uintptr类型的目标会被回收
	// pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 42
	fmt.Println(x.b) // "42"
}

func TestUnsafe(t *testing.T) {
	fmt.Println("begin TestLog ...")
	pointerDemo()
}
