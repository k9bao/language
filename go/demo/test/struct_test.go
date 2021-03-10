package test

import (
	"log"
	"testing"
)

type Name1 struct {
	i1 int
	str  string
}

type Name2 struct {
	Name1
}

func (n2 *Name2) String() {
	log.Println(n2.i1,"+",n2.str)
}


func TestStruct(t *testing.T) {
	n1 := Name2{}
	n1.i1 = 3
	n1.str = "test"
	n1.String()
	log.Println(n1)
}
