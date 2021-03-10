package main

import (
	"fmt"
	"time"

	"../../os"
)

func CPUDemo() {
	os.InitCPU()
	defer os.UninitCPU()
	for {
		fmt.Println(os.CPUPrecent())
		time.Sleep(time.Second)
	}
}

func main() {
	CPUDemo()
}
