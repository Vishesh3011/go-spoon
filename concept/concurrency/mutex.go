package concurrency

import (
	"fmt"
	"sync"
)

type someStruct struct {
	mu     sync.Mutex
	numMap map[string]int
}

func (s *someStruct) Add(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.numMap["someKey"] = i
}

func TryMutex() {
	s := &someStruct{numMap: make(map[string]int)}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(i)
	}
	wg.Wait()
	fmt.Println(s.numMap["someKey"])
}
