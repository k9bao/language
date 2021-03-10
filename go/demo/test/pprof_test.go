package test

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

func ProfileWrite(dur time.Duration) {
	f, err := os.Create(time.Now().Format("20060102_150405.cpu.data"))
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer f.Close()
	defer pprof.StopCPUProfile()

	f1, err := os.Create(time.Now().Format("20060102_150405.mem.data"))
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(f1)
	defer f1.Close()

	start := time.Now()
	for {
		if time.Since(start) > dur {
			break
		}
		fmt.Println("test")
		time.Sleep(time.Millisecond)
	}
}

func TestProfile(t *testing.T) {
	for {
		ProfileWrite(30 * time.Second)
		time.Sleep(30 * time.Second)
	}
}
