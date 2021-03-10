package test

//https://blog.csdn.net/chenbaoke/article/details/39932845
import (
	"fmt"
	"testing"
)

type Sample struct {
	a   int
	str string
}

func FmtDemo() {
	s := new(Sample)
	s.a = 1
	s.str = "hello"
	fmt.Printf("%v\n", *s)  //{1 hello}
	fmt.Printf("%+v\n", *s) // {a:1 str:hello}
	fmt.Printf("%#v\n", *s) // main.Sample{a:1, str:"hello"}
	fmt.Printf("%T\n", *s)  // main.Sample
}

func TestFmt(t *testing.T) {
	out := fmt.Sprintf("test_%v_%v.jpg", int(100), int(100))
	fmt.Println(out)
}
