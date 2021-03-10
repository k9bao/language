package test

import (
	"context"
	"log"
	"testing"
	"time"
)

func CancelContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("gorount is over")
				return
			default:
				time.Sleep(time.Second)
				log.Println("gorount is run")
			}
		}
	}(ctx)
	start := time.Now()
	for {
		if time.Since(start) > time.Second*5 {
			log.Println("begin stop")
			break
		} else {
			time.Sleep(time.Second)
		}
	}
	cancel()
	time.Sleep(time.Second)
}

func TestContext(t *testing.T) {
	CancelContext()
}
