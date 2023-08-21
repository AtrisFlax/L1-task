package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна
выводить итоговое значение счетчика.
*/

func main() {
	ac := NewAtomicCounter()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 100; c++ {
				ac.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ac value:", ac.Get())
}

type AtomicCounter struct {
	value uint64
}

func (c *AtomicCounter) Inc() {
	atomic.AddUint64(&c.value, 1)
}

func (c *AtomicCounter) Get() uint64 {
	return atomic.LoadUint64(&c.value)
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}
