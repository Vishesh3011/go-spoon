package concurrency

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Work done")
			return
		default:
			fmt.Println("Doing some work")
		}
	}
}

func TryDoneChannel() {
	myChan := make(chan bool)
	go doWork(myChan)

	time.Sleep(5 * time.Second)
	close(myChan)
}
