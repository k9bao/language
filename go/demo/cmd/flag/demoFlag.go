package main

import (
	"flag"
	"fmt"
)

//定义一个全局变量的命令行接收参数
var testFlag = flag.String("test", "default value", "help message.")

//打印值的函数
func print(f *flag.Flag) {
	if f != nil {
		fmt.Println(f.Value)
	} else {
		fmt.Println(nil)
	}
}

func main() {
	//没有用flag.Parse()解析前
	fmt.Print("test:")
	print(flag.Lookup("test"))
	fmt.Print("test1:")
	print(flag.Lookup("test1"))

	//用flag.Parse()解析后
	flag.Parse()

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())

	fmt.Print("test:")
	print(flag.Lookup("test"))
	fmt.Print("test1:")
	print(flag.Lookup("test1"))
}
