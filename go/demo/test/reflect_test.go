package test

import (
	"fmt"
	"reflect"
	"testing"
)

const tagName = "Testing"

type Info struct {
	Name string `Testing:"-"`
	Age  int    `Testing:"age,min=17,max=60"`
	Sex  string `Testing:"sex,required"`
}

func ReflectDemo() {
	info := Info{
		Name: "benben",
		Age:  23,
		Sex:  "male",
	}

	//通过反射，我们获取变量的动态类型
	t := reflect.TypeOf(info)
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) //获取结构体的每一个字段
		tag := field.Tag.Get(tagName)
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}

func CheckType(args ...interface{}) {
	if true {
		for _, arg := range args {
			switch arg.(type) {
			case int:
				fmt.Println(arg, "is an int value.")
			case string:
				fmt.Println(arg, "is a string value.")
			case int64:
				fmt.Println(arg, "is an int64 value.")
			case uintptr:
				fmt.Println(arg, "is an uintptr value.")
			case *uintptr:
				fmt.Println(arg, "is an *uintptr value.")
			default:
				fmt.Println(arg, "is an unknown type.")
			}
		}
	}
}

func TestReflect(t *testing.T) {
	ReflectDemo()
}
