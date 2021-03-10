package demo

import (
	"fmt"
)

var demofun func(para int) error

func DemoRun() {
	if demofun == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}
