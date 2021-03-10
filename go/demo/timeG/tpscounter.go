package timeutil

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

//TPSCounter statistic in 1s/2s/3s... counts
type TPSCounter struct {
	startT     time.Time //modify bkIndex per 1s
	allCount   uint64    //all user counts
	bucket     []uint64  //bucket slice,1 bucket have 1 second counts
	bucketSize int       //bucket size
	bucketOut  []uint64  //bucket slice sum,1 ... n second counts
	bkIndex    int       //in 1 second to time.Now() bucket index
	cancel     context.CancelFunc
	bucketNew  uint64 //if c is full,save to this para,avoid discard user data
	mutex      sync.Mutex
}

//NewState statistic count/interval,if frequency==0, sync execute
func NewTPSCounter() *TPSCounter {
	bucketSize := 5
	ctx, cancel := context.WithCancel(context.Background())
	s := &TPSCounter{
		bucket:     make([]uint64, bucketSize, bucketSize),
		bucketOut:  make([]uint64, bucketSize, bucketSize),
		allCount:   0,
		bucketSize: bucketSize,
		cancel:     cancel,
	}
	s.startT = time.Now()
	go s.run(ctx)
	return s
}

func (s *TPSCounter) Close() error {
	s.cancel()
	return nil
}

func (s *TPSCounter) AddCount(count uint64) error {
	s.allCount += uint64(count)
	atomic.AddUint64(&s.bucketNew, count)
	return nil
}

func (s *TPSCounter) GetAllCount() uint64 {
	return s.allCount
}

func (s *TPSCounter) GetTPS(second int) uint64 {
	if second > s.bucketSize {
		second = s.bucketSize
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.bucketOut[second-1]
}

func (s *TPSCounter) run(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.bkIndex = (s.bkIndex + 1) % s.bucketSize
			s.bucket[s.bkIndex] = atomic.SwapUint64(&s.bucketNew, 0)

			s.mutex.Lock()
			index := s.bkIndex
			for i := range s.bucketOut {
				s.bucketOut[i] = 0
			}
			for i := 0; i < s.bucketSize; i++ {
				if index < 0 {
					index += s.bucketSize
				}
				for j := i; j < s.bucketSize; j++ {
					s.bucketOut[j] += s.bucket[index]
				}
				index--
			}
			s.mutex.Unlock()
		}
	}
}
