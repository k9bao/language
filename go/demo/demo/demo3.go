package demo

type TestStruct struct {
	num1 int32
	num2 int32
}

func test1() {
	structArr1 := []TestStruct{}
	for i := 0; i < 100000000; i++ {
		structArr1 = append(structArr1, TestStruct{})
	}
}

func test2() {
	structArr2 := []*TestStruct{}
	for i := 0; i < 100000000; i++ {
		structArr2 = append(structArr2, &TestStruct{})
	}
}
func main() {

}
