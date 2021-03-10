package timeutil

import (
	"errors"
	"time"
)

//Stat Statistics of the number within 1 second from the current time, this algorithm can be implemented more efficiently through the bucket mechanism
type Stat struct {
	c        chan time.Time
	interval time.Duration
	count    int
	t        time.Time
	sync     bool
	startT   time.Time
	allCount uint64
}

//NewState statistic count/interval,if frequency==0, sync execute
func NewState(cap int64, interval, frequency time.Duration) *Stat {
	s := &Stat{
		c:        make(chan time.Time, cap),
		interval: interval,
	}
	if frequency == 0 {
		s.sync = true
	} else {
		go s.run(frequency)
	}
	return s
}

func (s *Stat) run(frequency time.Duration) error {
	ticker := time.NewTicker(frequency)
EXIT:
	for {
		select {
		case <-ticker.C:
			if err := s.calc(); err != nil {
				break EXIT
			}
		}
	}
	return nil
}

func (s *Stat) calc() error {
	if !s.t.IsZero() {
		if time.Since(s.t) < s.interval {
			s.count = len(s.c)
			return nil
		}
	}
	for {
		if v, ok := <-s.c; ok {
			if time.Since(v) < s.interval {
				s.count = len(s.c)
				s.t = v
				return nil
			}
		} else {
			return errors.New("chan is close")
		}
	}
}

func (s *Stat) Add() int {
	if s.startT.IsZero() {
		s.startT = time.Now()
	}
	s.allCount++
	if len(s.c) != cap(s.c) {
		s.c <- time.Now()
	} else {
		panic("chan not have enough space")
	}
	if s.sync == true {
		s.calc()
	}
	return s.count
}

func (s *Stat) Count() int {
	return s.count
}

func (s *Stat) AllCount() uint64 {
	if time.Since(s.startT) > time.Second {
		return s.allCount / uint64(time.Since(s.startT)/time.Second)
	}
	return 0
}

func (s *Stat) Close() error {
	close(s.c)
	return nil
}
