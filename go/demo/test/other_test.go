package test

import (
	"log"
	"testing"
)

func demo1() {
	num := 6
	for i := 0; i < 10; i++ {
		log.Println(num % 6)
		num--
	}
}

func TestOtherDemo(t *testing.T) {
	demo1()
}
