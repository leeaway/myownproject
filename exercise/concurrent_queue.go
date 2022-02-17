package exercise

import (
	"fmt"
	"sync"
)

/*
	使用Mutex实现一个并发安全的队列
*/

type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(n int) *SliceQueue {
	return &SliceQueue{
		data: make([]interface{}, n),
	}
}

func (s *SliceQueue) Offer(d interface{}) {
	s.mu.Lock()
	s.data = append(s.data, d)
	s.mu.Unlock()
}

func (s *SliceQueue) Poll() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	tmp := s.data[0]
	s.data = s.data[1:]
	return tmp
}

func (s *SliceQueue) Display() {
	total := 0
	for _, q := range s.data {
		if q == nil {
			continue
		}
		total++
		fmt.Printf("%v ", q)
	}
	fmt.Println()
	fmt.Printf("total: %v\n", total)
}
