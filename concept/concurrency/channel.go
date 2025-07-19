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

func TryUnbufferedChannel1() {
	ch := make(chan string)
	go func() {
		ch <- "yo"
	}()

	go func() {
		ch <- "hey"
	}()

	msg1 := <-ch
	msg2 := <-ch
	fmt.Println(msg1)
	fmt.Println(msg2)
}

func TryUnbufferedChannel2() {
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

func TryBufferedChannel1() {
	ch := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, char := range chars {
		//select {
		//case ch <- char:
		//}
		ch <- char
	}
	close(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
