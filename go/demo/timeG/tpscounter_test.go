package timeutil

import (
	"context"
	"log"
	"testing"
	"time"
)

func TPSCounterDemo() {
	now := time.Now()

	tps := NewTPSCounter()
	defer tps.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		tick := time.NewTicker(10 * time.Millisecond)
		defer tick.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Println("ctx Done")
				return
			case <-tick.C:
				//tps.AddCount(rand.Int31n(1))
				tps.AddCount(10)
			}
		}
	}(ctx)

	for {
		log.Println(tps.GetTPS(1), tps.GetTPS(2), tps.GetTPS(3), tps.GetTPS(4), tps.GetTPS(5), tps.GetTPS(6))
		time.Sleep(900 * time.Millisecond)
		if time.Since(now) > time.Second*20 {
			log.Println("begin Stop")
			break
		}
	}
}
func TestTPS(t *testing.T) {
	TPSCounterDemo()
}
