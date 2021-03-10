package syncg

import (
	"log"
	"sync"
)

type MutexG struct {
	num  int
	lock sync.Mutex
}

func (self *MutexG) Add(n int) (int, int) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.num += n
	return self.num - n, self.num
}

type CondG struct {
	num  int
	c    *sync.Cond
	lock sync.Mutex
}

func NewCondG(n int) *CondG {
	cond := &CondG{
		num: n,
	}
	cond.c = sync.NewCond(&cond.lock)
	return cond
}

//Wait Consumer
func (self *CondG) Wait() {
	self.c.L.Lock()
	defer self.c.L.Unlock()
	log.Printf("V in,num = %v\n", self.num)
	for self.num == 0 {
		self.c.Wait()
	}
	self.num--
	log.Printf("V out,num = %v\n", self.num)
}

//Signal Producer
func (self *CondG) Signal(i int) {
	self.c.L.Lock()
	defer self.c.L.Unlock()
	log.Printf("P in,num = %v\n", self.num)
	self.num += i
	for range make([]struct{}, i) {
		self.c.Signal()
	}
	log.Printf("P out,num = %v\n", self.num)
}

func (self *CondG) Broadcast(i int) {
	self.c.L.Lock()
	defer self.c.L.Unlock()
	self.c.Broadcast()
	self.num += i
}
