package test

import (
	"fmt"
	"strconv"
	"testing"
)

//Lazy Evaluation
func Interge() <-chan int {
	yield := make(chan int, 12)
	count := 0
	go func() {
		for {
			yield <- count
			fmt.Println("gorount:" + strconv.Itoa(count))
			count++
		}
	}()
	return yield
}

var resume <-chan int

func getInetge() int {
	return <-resume
}

func Laza() {
	resume = Interge()
	fmt.Println(getInetge())
	fmt.Println(getInetge())
	fmt.Println(getInetge())
	fmt.Println(getInetge())
	fmt.Println(getInetge())
	fmt.Println(getInetge())
	fmt.Println(getInetge())
}

func TestLaza(t *testing.T) {
	Laza()
}
