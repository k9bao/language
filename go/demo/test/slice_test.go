package test

import (
	"fmt"
	"testing"
)

func TestSliceFor(t *testing.T) {
	c1 := make(chan int, 6)
	for i := 0; i < 6; i++ {
		c1 <- i
	}
	for temp := range c1 {
		fmt.Println(temp)
		c1 <- temp
	}
	fmt.Println(len(c1))
}

func fun1(s []int) {
	fmt.Println(len(s))
}
func TestSliceNormal(t *testing.T) {
	fun1(nil)
}
