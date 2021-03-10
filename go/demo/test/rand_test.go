package test

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func RandDemo() {
	for i := 0; i < 10; i++ {
		log.Println(rand.Int(), rand.Int31(), rand.Int63(), rand.NormFloat64(), rand.Uint32())
	}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		log.Println(rd.Int(), rd.Int31(), rd.Int63(), rd.NormFloat64(), rd.Uint32())
	}
}

func TestRand(t *testing.T) {
	RandDemo()
}
