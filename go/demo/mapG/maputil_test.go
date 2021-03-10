package mapg

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	fmt.Println("hello")
	m := Create()
	m.Insert("123", "234")
	fmt.Println(m)
}
func BenchmarkMemcopy2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Println("hello world")
	}
}
