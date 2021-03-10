package main

import (
	"context"
	"fmt"
	"sync"

	body "github.com/knowledgebao/gotest/tags/src/imply/body"

	"github.com/knowledgebao/gotest/tags/src/imply/driver"
)

func main() {
	fmt.Println(driver.List())
	config := body.Config{
		Type: "body2",
		IP:   "120.0.0.1",
	}
	ctx := context.Background()
	imp1, err := driver.Create(ctx, config)
	if err == nil {
		var wg sync.WaitGroup
		wg.Add(1)

	LOOP_TASKS:
		for {
			select {
			case <-ctx.Done():
				break LOOP_TASKS
			case t, ok := <-imp1.Output():
				if !ok {
					break LOOP_TASKS
				}
				fmt.Printf("%v\n", t)
			}
		}
		wg.Wait()
	} else {
		fmt.Printf("create %v is fail", config.Type)
	}
}
