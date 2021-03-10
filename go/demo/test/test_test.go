package test

import (
	"fmt"
	"strings"
	"testing"
)

//go test [build/test flags] [packages] [build/test flags & test binary flags]

//go test 会运行当前目录下所有Test***测试例
//go test -v Test_test.go 会运行指定目录Test_test.go的所有Test***测试例
//go test -v -bench="." -run="Test_test.go" 会运行指定目录Test_test.go下的所有Test+Benchmark测试例
//go test -v -bench="." -benchtime="3s" -run="Test_test.go" 会运行指定目录Test_test.go下的所有Benchmark测试例，如果不指定运行时间，go test觉得稳定了就会结束执行

//go test -v Test_test.go "-test.run" TestHello
func TestHello(t *testing.T) {
	fmt.Println("TestHello")
}

//命令：go test -run="Test_test.go" -bench="." -benchtime="3s"
//结果：BenchmarkStringJoin1-4 300000 4351 ns/op 32 B/op 2 allocs/op
//其中：-4表示4个CPU线程执行；300000表示总共执行了30万次；4531ns/op，表示每次执行耗时4531纳秒；
//      32B/op表示每次执行分配了32字节内存；2 allocs/op表示每次执行分配了2次对象。
func BenchmarkStringJoin1(b *testing.B) {
	b.ReportAllocs()
	input := []string{"Hello", "World"}
	for i := 0; i < b.N; i++ {
		result := strings.Join(input, " ")
		if result != "Hello World" {
			b.Error("Unexpected result: " + result)
		}
	}
}

func GetReturn() (int, int) {
	a := 1
	b := 2
	return a, b
}

//isOverlap rc1 include rc2 return true
func isOverlap(rc1 Rect, rc2 Rect) bool {
	if rc1.Left+rc1.Width > rc2.Left &&
		rc2.Left+rc2.Width > rc1.Left &&
		rc1.Top+rc1.Height > rc2.Top &&
		rc2.Top+rc2.Height > rc1.Top {
		return true
	}
	return false
}

//go test -v Test_test.go "-test.run" TestWorld
func TestWorld(t *testing.T) {
	ts := 1585210514852
	pts := 93972791
	//pkt.TS = pkt.PTS + *flagWriteFileTime2/1000000 //ms
	writeTime := (ts - pts) * 1000000
	println(writeTime)
}
