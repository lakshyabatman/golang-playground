package main

import (
	"fmt"
	"sync"
)

type SafeInteger struct {
	value int
	lock  sync.Mutex
}

func (s *SafeInteger) Update(num int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.value = num
}

func main() {
	w := sync.WaitGroup{}
	safeInteger := SafeInteger{}
	for i := 0; i < 100; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			safeInteger.Update(i)
		}()
	}
	w.Wait()
	fmt.Println(safeInteger.value)
}
