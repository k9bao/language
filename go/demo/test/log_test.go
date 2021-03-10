package test

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	fmt.Println("begin TestLog ...")
	file, err := os.Create("test.log") // 创建日志文件
	if err != nil {
		// 打印日志 并退出程序
		log.Fatalln("fail to create test.log file!")
	}

	// 创建logger对象　这种方式会显示触发日志文件行数
	logger := log.New(io.MultiWriter(file, os.Stdout), "", log.LstdFlags|log.Llongfile)
	// Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	// Ltime                         // the time in the local time zone: 01:23:23
	// Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	// Llongfile                     // full file name and line number: /a/b/c/d.go:23
	// Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	// LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	// LstdFlags     = Ldate | Ltime // initial values for the standard logger
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //
	logger.SetOutput(io.MultiWriter(file, os.Stdout))       //ioutil.Discard/os.Stdout/os.Stdout/os.Stderr/
	logger.SetPrefix("prefix:")
	logger.Output(1, "outputs") //其他输出函数最后都是调用此函数。一般无需直接调用，第一个参数表示跳过的栈帧数，内部一般为2，直接调用写1即可。

	logger.Println(logger.Flags())
	logger.Println(logger.Prefix())

	logger.Print("Print")
	logger.Printf("Printf=%020v", 5555)
	logger.Println("Println")

	logger.Fatal("Fatal")
	logger.Fatalf("%v", "Fatalf")
	logger.Fatalln("Fatalln")

	logger.Panic("Panic")
	logger.Panicf("%v", "Panicf")
	logger.Panicln("Panicfln")
}
