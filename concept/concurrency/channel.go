package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func readerForChannel(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if val, ok := <-ch; !ok {
			fmt.Println("Channel closed")
			return
		} else {
			fmt.Println(fmt.Sprintf("Reader %d received %d in its channel", id, val))
		}
	}
}

func TryChannel() {
	stTime := time.Now()

	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(4)

	go readerForChannel(1, ch, wg)
	go readerForChannel(2, ch, wg)
	go readerForChannel(3, ch, wg)
	go readerForChannel(4, ch, wg)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()

	//for i := 0; i < 100; i++ {
	//	fmt.Println(i)
	//}

	fmt.Println(time.Since(stTime))
}
