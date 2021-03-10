package osG

import (
	"fmt"
	"testing"
	"time"
)

func CPUDemo() {
	InitCPU()
	defer UninitCPU()
	for {
		fmt.Println(CPUPrecent())
		time.Sleep(time.Second)
	}
}

func TestContext(t *testing.T) {
	CPUDemo()
}
