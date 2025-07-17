package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func reader(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Hello ", id)
}

func TryWaitGroup() {
	//	initial implementation
	stTime := time.Now()

	//for loop time
	//for i := 1; i < 3; i++ {
	//	fmt.Println("Hello ", i)
	//}

	// goroutine time
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go reader(wg, 1)
	go reader(wg, 2)

	wg.Wait()
	fmt.Println(time.Since(stTime))
}
