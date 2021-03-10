package test

import (
	"fmt"
	"testing"
	"time"
)

/*
1,chan 由发送者(生产者)关闭，如果是双向chan，则有其中一方关闭
2,对于值为nil的channel或者对同一个channel重复close, 都会panic, 关闭只读channel会报编译错误
*/

func TestChan(t *testing.T) {
	c1 := make(chan int, 6)
	c2 := make(chan int, 6)
	var defNum int64
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
			time.Sleep(time.Second)
			fmt.Println("in", defNum)
		}
		//close(c1)
	}()
EXIT:
	for {
		select {
		case v, ok := <-c1:
			if !ok {
				fmt.Println("close")
				break EXIT
			} else {
				c2 <- v
				fmt.Println(v, defNum, cap(c2), len(c2))
			}
		default:
			defNum++
		}
	}
	//close(c1) //panic
	fmt.Println("end")
}

func TestReadFromCloesChan(t *testing.T) {
	c1 := make(chan int, 6)
	for i := 0; i < 6; i++ {
		c1 <- i
	}
	close(c1)
EXIT:
	for {
		select {
		case v, ok := <-c1:
			if !ok {
				fmt.Println("close")
				break EXIT
			} else {
				fmt.Println(v)
			}
		}
	}
	fmt.Println("end")
}
