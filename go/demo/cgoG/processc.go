package cgoG

/*
#include <stdio.h>
#include <stdlib.h>

char ch = 'M';
signed char sch = 'N';
unsigned char uch = 253;
short st = 233;
signed short sst = 233;
unsigned short ust = 233;
int i = 257;
long l = 11112222;
unsigned long ul = 11112222;
long long ll = 1111222211111;
unsigned long long ull = 1111222211111;
float f = 3.14;
double db = 3.15;
void * p;
char *str = "const string";
char str1[64] = "char array";

void printI(void *i)
{
    printf("print i = %d\n", (*(int *)i));
}

struct ImgInfo {
    char *imgPath;
    int format;
    unsigned int width;
    unsigned int height;
};

void printStruct(struct ImgInfo *imgInfo)
{
    if(!imgInfo) {
        fprintf(stderr, "imgInfo is null\n");
        return ;
    }

    fprintf(stdout, "imgPath = %s\n", imgInfo->imgPath);
    fprintf(stdout, "format = %d\n", imgInfo->format);
    fprintf(stdout, "width = %d\n", imgInfo->width);
}

#define FFALIGN(x, a) (((x)+(a)-1)&~((a)-1))

*/
import "C"

import (
	"reflect"
	"unsafe"
	"fmt"
)

/*
wchar_t -->  C.wchar_t  -->  
void * -> unsafe.Pointer

// Go string to C string
func C.CString(string) *C.char

var val []byte
(*C.char)(unsafe.Pointer(&val[0]))
*/

func printType(g interface{}){
	fmt.Println(g, "-", reflect.TypeOf(g))
}

/*
C语言类型	CGO类型	Go语言类型
char	C.char	byte
singed char	C.schar	int8
unsigned char	C.uchar	uint8
short	C.short	int16
unsigned short	C.ushort	uint16
int	C.int	int32
unsigned int	C.uint	uint32
long	C.long	int32
unsigned long	C.ulong	uint32
long long int	C.longlong	int64
unsigned long long int	C.ulonglong	uint64
float	C.float	float32
double	C.double	float64
size_t	C.size_t	uint
*/
func checkType(g interface{}){
	printType(g)
	switch t := g.(type){
	case int8:
		printType(C.schar(g.(int8)))
	case C.char:
		printType(int8(g.(C.char)))
	case *C.char:
		printType(C.GoString(g.(*C.char)))
	case C.schar:
		printType(int8(g.(C.schar)))
	case uint8://byte
		printType(C.uchar(g.(uint8)))
	case C.uchar:
		printType(uint8(g.(C.uchar)))
	case int16:
		printType(C.short(g.(int16)))
	case C.short:
		printType(int16(g.(C.short)))
	case uint16:
		printType(C.ushort(g.(uint16)))
	case C.ushort:
		printType(uint16(g.(C.ushort)))
	case int:
		printType(C.int(g.(int)))
	case C.int:
		printType(int(g.(C.int)))
	case int32:
		printType(C.long(g.(int32)))
	case C.long:
		printType(int32(g.(C.long)))
	case uint32:
		printType(C.uint(g.(uint32)))
	case C.uint:
		printType(uint32(g.(C.uint)))
	case C.ulong:
		printType(uint32(g.(C.ulong)))
	case int64:
		printType(C.longlong(g.(int64)))
	case C.longlong:
		printType(int64(g.(C.longlong)))
	case uint64:
		printType(C.ulonglong(g.(uint64)))
	case C.ulonglong:
		printType(uint64(g.(C.ulonglong)))
	case float32:
		printType(C.float(g.(float32)))
	case C.float:
		printType(float32(g.(C.float)))
	case float64:
		printType(C.double(g.(float64)))
	case C.double:
		printType(float64(g.(C.double)))
	default:
		fmt.Println("undefine type:",t)		
	}
	
	fmt.Println("---------------------------")
}
func Go2C(){
	fmt.Println("----------------Go to C---------------")

	//ret := C.FFALIGN(7,8) cgo: doesn't handle #defines 
	var g1 byte;		g1 = 'z';					checkType(g1)
	var g2 int8;		g2 = 0x3F;					checkType(g2)
	var g3 uint8;		g3 = 0xFF;					checkType(g3)
	var g4 int16;		g4 = 0x0FFF;				checkType(g4)
	var g5 uint16;		g5 = 0xFFFF;				checkType(g5)
	var g6 int;			g6 = 333;					checkType(g6)
	var g7 int32;		g7 = 0x3FFFFFFF;			checkType(g7)
	var g8 uint32;		g8 = 0xFFFFFFFF;			checkType(g8)
	var g9 int64;		g9 = 0x0FFFFFFFFFFFFFFF;	checkType(g9)
	var g10 uint64;		g10 = 0xFFFFFFFFFFFFFFFF;	checkType(g10)
	var g11 float32;	g11 = 133.789;				checkType(g11)
	var g12 float64;	g12 = 99999999.999999999;	checkType(g12)
	C.printI(unsafe.Pointer(&g6))//unsafe.Pointer->void*

	fmt.Println("----------------C to Go---------------")
	checkType(C.ch)
	checkType(C.sch)
	checkType(C.uch)
	checkType(C.st)
	checkType(C.sst)
	checkType(C.ust)
	checkType(C.i)
	checkType(C.l)
	checkType(C.ul)
	checkType(C.ll)
	checkType(C.ull)
	checkType(C.f)
	checkType(C.db)
	checkType(C.str)//*C.char

	// 区别常量字符串和char数组，转换成Go类型不一样,常量字符串是*C.char,数组是[n]C.char
	fmt.Println(reflect.TypeOf(C.str1))//[64]C.char
	var charray []byte
	for i := range C.str1 {
		if C.str1[i] != 0 {
			charray = append(charray, byte(C.str1[i]))
		}
	}

	fmt.Println(charray)
	fmt.Println(string(charray))

	imgInfo := C.struct_ImgInfo{imgPath: C.CString("../images/xx.jpg"), format: 0, width: 500, height: 400}
	defer C.free(unsafe.Pointer(imgInfo.imgPath))
	C.printStruct(&imgInfo)
	
	fmt.Println("----------------C Print----------------")
}