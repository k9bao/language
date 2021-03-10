package syncg

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func mutexDemo(m *MutexG, i int) {
	old, new := m.Add(i)
	log.Println(i, "+", old, "=", new)
}
func TestMutexG(t *testing.T) {
	mutex := MutexG{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutexDemo(&mutex, i)
		}(i)
	}
	wg.Wait()

}

func condWait(c *CondG, i int) {
	log.Println(i, ": consumer in")
	c.Wait()
	log.Println(i, ": consumer ok")
}

func condSignal(c *CondG, i int) {
	log.Println(i, ": producer in")
	c.Signal(1)
	log.Println(i, ": producer ok")
}

func TestCondG(t *testing.T) {
	init := 1
	custum := 2
	producer := custum - init
	c := NewCondG(init)
	var wg sync.WaitGroup
	wg.Add(custum)
	for i := range make([]struct{}, custum) {
		go func(i int) {
			defer wg.Done()
			condWait(c, i)
		}(i)
	}
	wg.Add(producer)
	for i := range make([]struct{}, producer) {
		go func(i int) {
			defer wg.Done()
			condSignal(c, i)
		}(i + 1)
	}
	//log.Println("---broadcast in")
	//c.Broadcast(1)
	//log.Println("---broadcast out")
	wg.Wait()
}

//func AddInt32(addr *int32, delta int32) (new int32)
//func LoadInt32(addr *int32) (val int32)

func atomicDemo() {
	var iTest int32
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("gorount is over")
				return
			default:
				time.Sleep(time.Millisecond)
				log.Println("get-------", atomic.LoadInt32(&iTest))
				log.Println("set-------", atomic.AddInt32(&iTest, -4))
			}
		}
	}(ctx)
	start := time.Now()
	for {
		if time.Since(start) > time.Second*2 {
			log.Println("begin stop")
			break
		} else {
			time.Sleep(time.Millisecond)
			log.Println("get+++++++++", atomic.LoadInt32(&iTest))
			log.Println("set+++++++++", atomic.AddInt32(&iTest, 5))

		}
	}
	time.Sleep(time.Second)
}

func TestSync(t *testing.T) {
	atomicDemo()
}
