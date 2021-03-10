package demo

import (
	"fmt"
)

func init() {
	demofun = func(para int) error {
		fmt.Println("implate DemooImpl", para)
		return nil
	}
}
