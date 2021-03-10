package timeutil

import (
	"context"
	"log"
	"testing"
	"time"
)

func StatDemo() {
	now := time.Now()
	var stat = NewState(1000, time.Second, 10*time.Millisecond)
	defer stat.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		tick := time.NewTicker(100 * time.Millisecond)
		defer tick.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Println("ctx Done")
				return
			case <-tick.C:
				stat.Add()
			}
		}
	}(ctx)

	for {
		log.Println(stat.Count(), stat.AllCount())
		time.Sleep(time.Second)
		if time.Since(now) > time.Second*10 {
			log.Println("begin Stop")
			break
		}
	}
}
func TestStat(t *testing.T) {
	StatDemo()
}
